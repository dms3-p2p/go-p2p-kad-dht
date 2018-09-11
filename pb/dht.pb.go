// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dht.proto

package dht_pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pb "github.com/dms3-p2p/go-p2p-record/pb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message_MessageType int32

const (
	Message_PUT_VALUE     Message_MessageType = 0
	Message_GET_VALUE     Message_MessageType = 1
	Message_ADD_PROVIDER  Message_MessageType = 2
	Message_GET_PROVIDERS Message_MessageType = 3
	Message_FIND_NODE     Message_MessageType = 4
	Message_PING          Message_MessageType = 5
)

var Message_MessageType_name = map[int32]string{
	0: "PUT_VALUE",
	1: "GET_VALUE",
	2: "ADD_PROVIDER",
	3: "GET_PROVIDERS",
	4: "FIND_NODE",
	5: "PING",
}
var Message_MessageType_value = map[string]int32{
	"PUT_VALUE":     0,
	"GET_VALUE":     1,
	"ADD_PROVIDER":  2,
	"GET_PROVIDERS": 3,
	"FIND_NODE":     4,
	"PING":          5,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dht_a516883840d6c148, []int{0, 0}
}

type Message_ConnectionType int32

const (
	// sender does not have a connection to peer, and no extra information (default)
	Message_NOT_CONNECTED Message_ConnectionType = 0
	// sender has a live connection to peer
	Message_CONNECTED Message_ConnectionType = 1
	// sender recently connected to peer
	Message_CAN_CONNECT Message_ConnectionType = 2
	// sender recently tried to connect to peer repeatedly but failed to connect
	// ("try" here is loose, but this should signal "made strong effort, failed")
	Message_CANNOT_CONNECT Message_ConnectionType = 3
)

var Message_ConnectionType_name = map[int32]string{
	0: "NOT_CONNECTED",
	1: "CONNECTED",
	2: "CAN_CONNECT",
	3: "CANNOT_CONNECT",
}
var Message_ConnectionType_value = map[string]int32{
	"NOT_CONNECTED":  0,
	"CONNECTED":      1,
	"CAN_CONNECT":    2,
	"CANNOT_CONNECT": 3,
}

func (x Message_ConnectionType) String() string {
	return proto.EnumName(Message_ConnectionType_name, int32(x))
}
func (Message_ConnectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dht_a516883840d6c148, []int{0, 1}
}

type Message struct {
	// defines what type of message it is.
	Type Message_MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=dht.pb.Message_MessageType" json:"type,omitempty"`
	// defines what coral cluster level this query/response belongs to.
	// in case we want to implement coral's cluster rings in the future.
	ClusterLevelRaw int32 `protobuf:"varint,10,opt,name=clusterLevelRaw,proto3" json:"clusterLevelRaw,omitempty"`
	// Used to specify the key associated with this message.
	// PUT_VALUE, GET_VALUE, ADD_PROVIDER, GET_PROVIDERS
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Used to return a value
	// PUT_VALUE, GET_VALUE
	Record *pb.Record `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
	// Used to return peers closer to a key in a query
	// GET_VALUE, GET_PROVIDERS, FIND_NODE
	CloserPeers []*Message_Peer `protobuf:"bytes,8,rep,name=closerPeers" json:"closerPeers,omitempty"`
	// Used to return Providers
	// GET_VALUE, ADD_PROVIDER, GET_PROVIDERS
	ProviderPeers        []*Message_Peer `protobuf:"bytes,9,rep,name=providerPeers" json:"providerPeers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_dht_a516883840d6c148, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_MessageType {
	if m != nil {
		return m.Type
	}
	return Message_PUT_VALUE
}

func (m *Message) GetClusterLevelRaw() int32 {
	if m != nil {
		return m.ClusterLevelRaw
	}
	return 0
}

func (m *Message) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetRecord() *pb.Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *Message) GetCloserPeers() []*Message_Peer {
	if m != nil {
		return m.CloserPeers
	}
	return nil
}

func (m *Message) GetProviderPeers() []*Message_Peer {
	if m != nil {
		return m.ProviderPeers
	}
	return nil
}

type Message_Peer struct {
	// ID of a given peer.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// multiaddrs for a given peer
	Addrs [][]byte `protobuf:"bytes,2,rep,name=addrs" json:"addrs,omitempty"`
	// used to signal the sender's connection capabilities to the peer
	Connection           Message_ConnectionType `protobuf:"varint,3,opt,name=connection,proto3,enum=dht.pb.Message_ConnectionType" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Message_Peer) Reset()         { *m = Message_Peer{} }
func (m *Message_Peer) String() string { return proto.CompactTextString(m) }
func (*Message_Peer) ProtoMessage()    {}
func (*Message_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_dht_a516883840d6c148, []int{0, 0}
}
func (m *Message_Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Peer.Merge(dst, src)
}
func (m *Message_Peer) XXX_Size() int {
	return m.Size()
}
func (m *Message_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Peer proto.InternalMessageInfo

func (m *Message_Peer) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Message_Peer) GetAddrs() [][]byte {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *Message_Peer) GetConnection() Message_ConnectionType {
	if m != nil {
		return m.Connection
	}
	return Message_NOT_CONNECTED
}

func init() {
	proto.RegisterType((*Message)(nil), "dht.pb.Message")
	proto.RegisterType((*Message_Peer)(nil), "dht.pb.Message.Peer")
	proto.RegisterEnum("dht.pb.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("dht.pb.Message_ConnectionType", Message_ConnectionType_name, Message_ConnectionType_value)
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDht(dAtA, i, uint64(m.Type))
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDht(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Record != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDht(dAtA, i, uint64(m.Record.Size()))
		n1, err := m.Record.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.CloserPeers) > 0 {
		for _, msg := range m.CloserPeers {
			dAtA[i] = 0x42
			i++
			i = encodeVarintDht(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ProviderPeers) > 0 {
		for _, msg := range m.ProviderPeers {
			dAtA[i] = 0x4a
			i++
			i = encodeVarintDht(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.ClusterLevelRaw != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintDht(dAtA, i, uint64(m.ClusterLevelRaw))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Message_Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Peer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDht(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Addrs) > 0 {
		for _, b := range m.Addrs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDht(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.Connection != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDht(dAtA, i, uint64(m.Connection))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDht(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Message) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovDht(uint64(m.Type))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovDht(uint64(l))
	}
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovDht(uint64(l))
	}
	if len(m.CloserPeers) > 0 {
		for _, e := range m.CloserPeers {
			l = e.Size()
			n += 1 + l + sovDht(uint64(l))
		}
	}
	if len(m.ProviderPeers) > 0 {
		for _, e := range m.ProviderPeers {
			l = e.Size()
			n += 1 + l + sovDht(uint64(l))
		}
	}
	if m.ClusterLevelRaw != 0 {
		n += 1 + sovDht(uint64(m.ClusterLevelRaw))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Message_Peer) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDht(uint64(l))
	}
	if len(m.Addrs) > 0 {
		for _, b := range m.Addrs {
			l = len(b)
			n += 1 + l + sovDht(uint64(l))
		}
	}
	if m.Connection != 0 {
		n += 1 + sovDht(uint64(m.Connection))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDht(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDht(x uint64) (n int) {
	return sovDht(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDht
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Message_MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDht
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDht
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Record == nil {
				m.Record = &pb.Record{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CloserPeers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDht
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CloserPeers = append(m.CloserPeers, &Message_Peer{})
			if err := m.CloserPeers[len(m.CloserPeers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderPeers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDht
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderPeers = append(m.ProviderPeers, &Message_Peer{})
			if err := m.ProviderPeers[len(m.ProviderPeers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterLevelRaw", wireType)
			}
			m.ClusterLevelRaw = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClusterLevelRaw |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDht(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDht
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDht
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDht
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDht
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addrs = append(m.Addrs, make([]byte, postIndex-iNdEx))
			copy(m.Addrs[len(m.Addrs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connection", wireType)
			}
			m.Connection = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDht
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Connection |= (Message_ConnectionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDht(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDht
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDht(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDht
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDht
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDht
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDht
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDht
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDht(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDht = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDht   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("dht.proto", fileDescriptor_dht_a516883840d6c148) }

var fileDescriptor_dht_a516883840d6c148 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0x9b, 0x40,
	0x10, 0x86, 0xbb, 0x80, 0xdd, 0x78, 0xc0, 0x64, 0x33, 0xca, 0x01, 0xa5, 0x92, 0x85, 0x7c, 0xa2,
	0x87, 0x80, 0x4a, 0xa4, 0x1e, 0x7a, 0xa8, 0xe4, 0x02, 0x8d, 0x2c, 0xa5, 0xd8, 0xda, 0x3a, 0xe9,
	0xd1, 0x32, 0xb0, 0x72, 0x50, 0x1d, 0x2f, 0x02, 0x92, 0xca, 0x4f, 0xd8, 0x1e, 0xfb, 0x08, 0x95,
	0x9f, 0xa4, 0x02, 0x42, 0x8b, 0x7d, 0xc8, 0x69, 0xff, 0x7f, 0xf7, 0xff, 0x66, 0x86, 0x11, 0x30,
	0x48, 0xee, 0x4b, 0x3b, 0xcb, 0x45, 0x29, 0xb0, 0x5f, 0xcb, 0xe8, 0xe2, 0xdd, 0x3a, 0x2d, 0xef,
	0x1f, 0x23, 0x3b, 0x16, 0x0f, 0x4e, 0xf2, 0x50, 0x5c, 0x5d, 0x66, 0x6e, 0xe6, 0xac, 0x45, 0x75,
	0x5c, 0xe6, 0x3c, 0x16, 0x79, 0xe2, 0x64, 0x91, 0xd3, 0xa8, 0x06, 0x1d, 0xff, 0x54, 0xe0, 0xf5,
	0x17, 0x5e, 0x14, 0xab, 0x35, 0x47, 0x07, 0x94, 0x72, 0x97, 0x71, 0x83, 0x98, 0xc4, 0xd2, 0xdd,
	0x37, 0x76, 0x53, 0xd5, 0x7e, 0x7e, 0x6e, 0xcf, 0xc5, 0x2e, 0xe3, 0xac, 0x0e, 0x22, 0x05, 0xf9,
	0x3b, 0xdf, 0x19, 0x92, 0x49, 0x2c, 0x8d, 0x55, 0x12, 0xdf, 0x42, 0xbf, 0x29, 0x6f, 0xc8, 0x26,
	0xb1, 0x54, 0xf7, 0xcc, 0x6e, 0xbb, 0x45, 0x36, 0xab, 0x15, 0x7b, 0x0e, 0xe0, 0x7b, 0x50, 0xe3,
	0x8d, 0x28, 0x78, 0x3e, 0xe7, 0x3c, 0x2f, 0x8c, 0x13, 0x53, 0xb6, 0x54, 0xf7, 0xfc, 0xb8, 0x69,
	0xf5, 0xc8, 0xba, 0x41, 0xfc, 0x00, 0xc3, 0x2c, 0x17, 0x4f, 0x69, 0xd2, 0x92, 0x83, 0x17, 0xc8,
	0xc3, 0x28, 0x5a, 0x70, 0x1a, 0x6f, 0x1e, 0x8b, 0x92, 0xe7, 0x37, 0xfc, 0x89, 0x6f, 0xd8, 0xea,
	0x87, 0x01, 0x26, 0xb1, 0x7a, 0xec, 0xf8, 0xfa, 0x62, 0x03, 0x4a, 0x85, 0xa0, 0x0e, 0x52, 0x9a,
	0xd4, 0x1b, 0xd1, 0x98, 0x94, 0x26, 0x78, 0x0e, 0xbd, 0x55, 0x92, 0xe4, 0x85, 0x21, 0x99, 0xb2,
	0xa5, 0xb1, 0xc6, 0xe0, 0x47, 0x80, 0x58, 0x6c, 0xb7, 0x3c, 0x2e, 0x53, 0xb1, 0xad, 0x3f, 0x5d,
	0x77, 0x47, 0xc7, 0x03, 0x79, 0xff, 0x12, 0xf5, 0x0a, 0x3b, 0xc4, 0x38, 0x05, 0xb5, 0xb3, 0x5d,
	0x1c, 0xc2, 0x60, 0x7e, 0xbb, 0x58, 0xde, 0x4d, 0x6e, 0x6e, 0x03, 0xfa, 0xaa, 0xb2, 0xd7, 0x41,
	0x6b, 0x09, 0x52, 0xd0, 0x26, 0xbe, 0xbf, 0x9c, 0xb3, 0xd9, 0xdd, 0xd4, 0x0f, 0x18, 0x95, 0xf0,
	0x0c, 0x86, 0x55, 0xa0, 0xbd, 0xf9, 0x4a, 0xe5, 0x8a, 0xf9, 0x3c, 0x0d, 0xfd, 0x65, 0x38, 0xf3,
	0x03, 0xaa, 0xe0, 0x09, 0x28, 0xf3, 0x69, 0x78, 0x4d, 0x7b, 0xe3, 0x6f, 0xa0, 0x1f, 0x0e, 0x52,
	0xd1, 0xe1, 0x6c, 0xb1, 0xf4, 0x66, 0x61, 0x18, 0x78, 0x8b, 0xc0, 0x6f, 0x3a, 0xfe, 0xb7, 0x04,
	0x4f, 0x41, 0xf5, 0x26, 0x61, 0x9b, 0xa0, 0x12, 0x22, 0xe8, 0xde, 0x24, 0xec, 0x50, 0x54, 0xfe,
	0xa4, 0xfd, 0xda, 0x8f, 0xc8, 0xef, 0xfd, 0x88, 0xfc, 0xd9, 0x8f, 0x48, 0xd4, 0xaf, 0x7f, 0xaf,
	0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x92, 0x33, 0x1f, 0xa6, 0x02, 0x00, 0x00,
}
