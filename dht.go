// Package dht implements a distributed hash table that satisfies the dms3fs routing
// interface. This DHT is modeled after kademlia with S/Kademlia modifications.
package dht

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	opts "github.com/dms3-p2p/go-p2p-kad-dht/opts"
	pb "github.com/dms3-p2p/go-p2p-kad-dht/pb"
	providers "github.com/dms3-p2p/go-p2p-kad-dht/providers"

	proto "github.com/gogo/protobuf/proto"
	cid "github.com/dms3-fs/go-cid"
	ds "github.com/dms3-fs/go-datastore"
	logging "github.com/dms3-fs/go-log"
	goprocess "github.com/jbenet/goprocess"
	goprocessctx "github.com/jbenet/goprocess/context"
	ci "github.com/dms3-p2p/go-p2p-crypto"
	host "github.com/dms3-p2p/go-p2p-host"
	kb "github.com/dms3-p2p/go-p2p-kbucket"
	inet "github.com/dms3-p2p/go-p2p-net"
	peer "github.com/dms3-p2p/go-p2p-peer"
	pstore "github.com/dms3-p2p/go-p2p-peerstore"
	protocol "github.com/dms3-p2p/go-p2p-protocol"
	record "github.com/dms3-p2p/go-p2p-record"
	recpb "github.com/dms3-p2p/go-p2p-record/pb"
	routing "github.com/dms3-p2p/go-p2p-routing"
	base32 "github.com/whyrusleeping/base32"
)

var log = logging.Logger("dht")

// NumBootstrapQueries defines the number of random dht queries to do to
// collect members of the routing table.
const NumBootstrapQueries = 5

// Dms3FsDHT is an implementation of Kademlia with S/Kademlia modifications.
// It is used to implement the base Dms3FsRouting module.
type Dms3FsDHT struct {
	host      host.Host        // the network services we need
	self      peer.ID          // Local peer (yourself)
	peerstore pstore.Peerstore // Peer Registry

	datastore ds.Datastore // Local data

	routingTable *kb.RoutingTable // Array of routing tables for differently distanced nodes
	providers    *providers.ProviderManager

	birth time.Time // When this peer started up

	Validator record.Validator

	ctx  context.Context
	proc goprocess.Process

	strmap map[peer.ID]*messageSender
	smlk   sync.Mutex

	plk sync.Mutex

	protocols []protocol.ID // DHT protocols
}

// New creates a new DHT with the specified host and options.
func New(ctx context.Context, h host.Host, options ...opts.Option) (*Dms3FsDHT, error) {
	var cfg opts.Options
	if err := cfg.Apply(append([]opts.Option{opts.Defaults}, options...)...); err != nil {
		return nil, err
	}
	dht := makeDHT(ctx, h, cfg.Datastore, cfg.Protocols)

	// register for network notifs.
	dht.host.Network().Notify((*netNotifiee)(dht))

	dht.proc = goprocessctx.WithContextAndTeardown(ctx, func() error {
		// remove ourselves from network notifs.
		dht.host.Network().StopNotify((*netNotifiee)(dht))
		return nil
	})

	dht.proc.AddChild(dht.providers.Process())
	dht.Validator = cfg.Validator

	if !cfg.Client {
		for _, p := range cfg.Protocols {
			h.SetStreamHandler(p, dht.handleNewStream)
		}
	}
	return dht, nil
}

// NewDHT creates a new DHT object with the given peer as the 'local' host.
// Dms3FsDHT's initialized with this function will respond to DHT requests,
// whereas Dms3FsDHT's initialized with NewDHTClient will not.
func NewDHT(ctx context.Context, h host.Host, dstore ds.Batching) *Dms3FsDHT {
	dht, err := New(ctx, h, opts.Datastore(dstore))
	if err != nil {
		panic(err)
	}
	return dht
}

// NewDHTClient creates a new DHT object with the given peer as the 'local'
// host. Dms3FsDHT clients initialized with this function will not respond to DHT
// requests. If you need a peer to respond to DHT requests, use NewDHT instead.
// NewDHTClient creates a new DHT object with the given peer as the 'local' host
func NewDHTClient(ctx context.Context, h host.Host, dstore ds.Batching) *Dms3FsDHT {
	dht, err := New(ctx, h, opts.Datastore(dstore), opts.Client(true))
	if err != nil {
		panic(err)
	}
	return dht
}

func makeDHT(ctx context.Context, h host.Host, dstore ds.Batching, protocols []protocol.ID) *Dms3FsDHT {
	rt := kb.NewRoutingTable(KValue, kb.ConvertPeerID(h.ID()), time.Minute, h.Peerstore())

	cmgr := h.ConnManager()
	rt.PeerAdded = func(p peer.ID) {
		cmgr.TagPeer(p, "kbucket", 5)
	}
	rt.PeerRemoved = func(p peer.ID) {
		cmgr.UntagPeer(p, "kbucket")
	}

	return &Dms3FsDHT{
		datastore:    dstore,
		self:         h.ID(),
		peerstore:    h.Peerstore(),
		host:         h,
		strmap:       make(map[peer.ID]*messageSender),
		ctx:          ctx,
		providers:    providers.NewProviderManager(ctx, h.ID(), dstore),
		birth:        time.Now(),
		routingTable: rt,
		protocols:    protocols,
	}
}

// putValueToPeer stores the given key/value pair at the peer 'p'
func (dht *Dms3FsDHT) putValueToPeer(ctx context.Context, p peer.ID, rec *recpb.Record) error {

	pmes := pb.NewMessage(pb.Message_PUT_VALUE, rec.Key, 0)
	pmes.Record = rec
	rpmes, err := dht.sendRequest(ctx, p, pmes)
	if err != nil {
		log.Debugf("putValueToPeer: %v. (peer: %s, key: %s)", err, p.Pretty(), loggableKey(string(rec.Key)))
		return err
	}

	if !bytes.Equal(rpmes.GetRecord().Value, pmes.GetRecord().Value) {
		log.Warningf("putValueToPeer: value not put correctly. (%v != %v)", pmes, rpmes)
		return errors.New("value not put correctly")
	}

	return nil
}

var errInvalidRecord = errors.New("received invalid record")

// getValueOrPeers queries a particular peer p for the value for
// key. It returns either the value or a list of closer peers.
// NOTE: It will update the dht's peerstore with any new addresses
// it finds for the given peer.
func (dht *Dms3FsDHT) getValueOrPeers(ctx context.Context, p peer.ID, key string) (*recpb.Record, []*pstore.PeerInfo, error) {

	pmes, err := dht.getValueSingle(ctx, p, key)
	if err != nil {
		return nil, nil, err
	}

	// Perhaps we were given closer peers
	peers := pb.PBPeersToPeerInfos(pmes.GetCloserPeers())

	if record := pmes.GetRecord(); record != nil {
		// Success! We were given the value
		log.Debug("getValueOrPeers: got value")

		// make sure record is valid.
		err = dht.Validator.Validate(string(record.GetKey()), record.GetValue())
		if err != nil {
			log.Info("Received invalid record! (discarded)")
			// return a sentinal to signify an invalid record was received
			err = errInvalidRecord
			record = new(recpb.Record)
		}
		return record, peers, err
	}

	if len(peers) > 0 {
		log.Debug("getValueOrPeers: peers")
		return nil, peers, nil
	}

	log.Warning("getValueOrPeers: routing.ErrNotFound")
	return nil, nil, routing.ErrNotFound
}

// getValueSingle simply performs the get value RPC with the given parameters
func (dht *Dms3FsDHT) getValueSingle(ctx context.Context, p peer.ID, key string) (*pb.Message, error) {
	meta := logging.LoggableMap{
		"key":  key,
		"peer": p,
	}

	eip := log.EventBegin(ctx, "getValueSingle", meta)
	defer eip.Done()

	pmes := pb.NewMessage(pb.Message_GET_VALUE, []byte(key), 0)
	resp, err := dht.sendRequest(ctx, p, pmes)
	switch err {
	case nil:
		return resp, nil
	case ErrReadTimeout:
		log.Warningf("getValueSingle: read timeout %s %s", p.Pretty(), key)
		fallthrough
	default:
		eip.SetError(err)
		return nil, err
	}
}

// getLocal attempts to retrieve the value from the datastore
func (dht *Dms3FsDHT) getLocal(key string) (*recpb.Record, error) {
	log.Debugf("getLocal %s", key)
	rec, err := dht.getRecordFromDatastore(mkDsKey(key))
	if err != nil {
		log.Warningf("getLocal: %s", err)
		return nil, err
	}

	// Double check the key. Can't hurt.
	if rec != nil && string(rec.GetKey()) != key {
		log.Errorf("BUG getLocal: found a DHT record that didn't match it's key: %s != %s", rec.GetKey(), key)
		return nil, nil

	}
	return rec, nil
}

// getOwnPrivateKey attempts to load the local peers private
// key from the peerstore.
func (dht *Dms3FsDHT) getOwnPrivateKey() (ci.PrivKey, error) {
	sk := dht.peerstore.PrivKey(dht.self)
	if sk == nil {
		log.Warningf("%s dht cannot get own private key!", dht.self)
		return nil, fmt.Errorf("cannot get private key to sign record!")
	}
	return sk, nil
}

// putLocal stores the key value pair in the datastore
func (dht *Dms3FsDHT) putLocal(key string, rec *recpb.Record) error {
	log.Debugf("putLocal: %v %v", key, rec)
	data, err := proto.Marshal(rec)
	if err != nil {
		log.Warningf("putLocal: %s", err)
		return err
	}

	return dht.datastore.Put(mkDsKey(key), data)
}

// Update signals the routingTable to Update its last-seen status
// on the given peer.
func (dht *Dms3FsDHT) Update(ctx context.Context, p peer.ID) {
	log.Event(ctx, "updatePeer", p)
	dht.routingTable.Update(p)
}

// FindLocal looks for a peer with a given ID connected to this dht and returns the peer and the table it was found in.
func (dht *Dms3FsDHT) FindLocal(id peer.ID) pstore.PeerInfo {
	switch dht.host.Network().Connectedness(id) {
	case inet.Connected, inet.CanConnect:
		return dht.peerstore.PeerInfo(id)
	default:
		return pstore.PeerInfo{}
	}
}

// findPeerSingle asks peer 'p' if they know where the peer with id 'id' is
func (dht *Dms3FsDHT) findPeerSingle(ctx context.Context, p peer.ID, id peer.ID) (*pb.Message, error) {
	eip := log.EventBegin(ctx, "findPeerSingle",
		logging.LoggableMap{
			"peer":   p,
			"target": id,
		})
	defer eip.Done()

	pmes := pb.NewMessage(pb.Message_FIND_NODE, []byte(id), 0)
	resp, err := dht.sendRequest(ctx, p, pmes)
	switch err {
	case nil:
		return resp, nil
	case ErrReadTimeout:
		log.Warningf("read timeout: %s %s", p.Pretty(), id)
		fallthrough
	default:
		eip.SetError(err)
		return nil, err
	}
}

func (dht *Dms3FsDHT) findProvidersSingle(ctx context.Context, p peer.ID, key *cid.Cid) (*pb.Message, error) {
	eip := log.EventBegin(ctx, "findProvidersSingle", p, key)
	defer eip.Done()

	pmes := pb.NewMessage(pb.Message_GET_PROVIDERS, key.Bytes(), 0)
	resp, err := dht.sendRequest(ctx, p, pmes)
	switch err {
	case nil:
		return resp, nil
	case ErrReadTimeout:
		log.Warningf("read timeout: %s %s", p.Pretty(), key)
		fallthrough
	default:
		eip.SetError(err)
		return nil, err
	}
}

// nearestPeersToQuery returns the routing tables closest peers.
func (dht *Dms3FsDHT) nearestPeersToQuery(pmes *pb.Message, count int) []peer.ID {
	closer := dht.routingTable.NearestPeers(kb.ConvertKey(string(pmes.GetKey())), count)
	return closer
}

// betterPeerToQuery returns nearestPeersToQuery, but iff closer than self.
func (dht *Dms3FsDHT) betterPeersToQuery(pmes *pb.Message, p peer.ID, count int) []peer.ID {
	closer := dht.nearestPeersToQuery(pmes, count)

	// no node? nil
	if closer == nil {
		log.Warning("betterPeersToQuery: no closer peers to send:", p)
		return nil
	}

	filtered := make([]peer.ID, 0, len(closer))
	for _, clp := range closer {

		// == to self? thats bad
		if clp == dht.self {
			log.Error("BUG betterPeersToQuery: attempted to return self! this shouldn't happen...")
			return nil
		}
		// Dont send a peer back themselves
		if clp == p {
			continue
		}

		filtered = append(filtered, clp)
	}

	// ok seems like closer nodes
	return filtered
}

// Context return dht's context
func (dht *Dms3FsDHT) Context() context.Context {
	return dht.ctx
}

// Process return dht's process
func (dht *Dms3FsDHT) Process() goprocess.Process {
	return dht.proc
}

// Close calls Process Close
func (dht *Dms3FsDHT) Close() error {
	return dht.proc.Close()
}

func (dht *Dms3FsDHT) protocolStrs() []string {
	pstrs := make([]string, len(dht.protocols))
	for idx, proto := range dht.protocols {
		pstrs[idx] = string(proto)
	}

	return pstrs
}

func mkDsKey(s string) ds.Key {
	return ds.NewKey(base32.RawStdEncoding.EncodeToString([]byte(s)))
}
