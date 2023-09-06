// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/events.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/sentinel-official/hub/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventStartSession struct {
	From         string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Subscription uint64 `protobuf:"varint,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Node         string `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *EventStartSession) Reset()         { *m = EventStartSession{} }
func (m *EventStartSession) String() string { return proto.CompactTextString(m) }
func (*EventStartSession) ProtoMessage()    {}
func (*EventStartSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{0}
}
func (m *EventStartSession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStartSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStartSession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStartSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStartSession.Merge(m, src)
}
func (m *EventStartSession) XXX_Size() int {
	return m.Size()
}
func (m *EventStartSession) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStartSession.DiscardUnknown(m)
}

var xxx_messageInfo_EventStartSession proto.InternalMessageInfo

type EventUpdateSession struct {
	From         string          `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Id           uint64          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Subscription uint64          `protobuf:"varint,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Node         string          `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	Address      string          `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Duration     time.Duration   `protobuf:"bytes,6,opt,name=duration,proto3,stdduration" json:"duration"`
	Bandwidth    types.Bandwidth `protobuf:"bytes,7,opt,name=bandwidth,proto3" json:"bandwidth"`
}

func (m *EventUpdateSession) Reset()         { *m = EventUpdateSession{} }
func (m *EventUpdateSession) String() string { return proto.CompactTextString(m) }
func (*EventUpdateSession) ProtoMessage()    {}
func (*EventUpdateSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{1}
}
func (m *EventUpdateSession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateSession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateSession.Merge(m, src)
}
func (m *EventUpdateSession) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateSession) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateSession.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateSession proto.InternalMessageInfo

type EventEndSession struct {
	From         string       `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Id           uint64       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Subscription uint64       `protobuf:"varint,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Node         string       `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	Status       types.Status `protobuf:"varint,5,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
}

func (m *EventEndSession) Reset()         { *m = EventEndSession{} }
func (m *EventEndSession) String() string { return proto.CompactTextString(m) }
func (*EventEndSession) ProtoMessage()    {}
func (*EventEndSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{2}
}
func (m *EventEndSession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEndSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEndSession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEndSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEndSession.Merge(m, src)
}
func (m *EventEndSession) XXX_Size() int {
	return m.Size()
}
func (m *EventEndSession) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEndSession.DiscardUnknown(m)
}

var xxx_messageInfo_EventEndSession proto.InternalMessageInfo

type EventPay struct {
	Id           uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Subscription uint64      `protobuf:"varint,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Node         string      `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Payment      types1.Coin `protobuf:"bytes,4,opt,name=payment,proto3" json:"payment"`
}

func (m *EventPay) Reset()         { *m = EventPay{} }
func (m *EventPay) String() string { return proto.CompactTextString(m) }
func (*EventPay) ProtoMessage()    {}
func (*EventPay) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{3}
}
func (m *EventPay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPay.Merge(m, src)
}
func (m *EventPay) XXX_Size() int {
	return m.Size()
}
func (m *EventPay) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPay.DiscardUnknown(m)
}

var xxx_messageInfo_EventPay proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventStartSession)(nil), "sentinel.session.v1.EventStartSession")
	proto.RegisterType((*EventUpdateSession)(nil), "sentinel.session.v1.EventUpdateSession")
	proto.RegisterType((*EventEndSession)(nil), "sentinel.session.v1.EventEndSession")
	proto.RegisterType((*EventPay)(nil), "sentinel.session.v1.EventPay")
}

func init() { proto.RegisterFile("sentinel/session/v1/events.proto", fileDescriptor_e34d43b21999bd7a) }

var fileDescriptor_e34d43b21999bd7a = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x5c, 0x69, 0x7b, 0x06, 0x1d, 0x22, 0x30, 0xe4, 0x4e, 0xc8, 0x57, 0x32, 0x75,
	0xc1, 0x26, 0xc7, 0xc4, 0x04, 0x2a, 0xdc, 0x8e, 0x52, 0xb1, 0xb0, 0x39, 0xb1, 0x9b, 0x5a, 0x6a,
	0xed, 0x28, 0x76, 0x02, 0xfd, 0x10, 0x27, 0x31, 0x32, 0x32, 0xf2, 0x51, 0x3a, 0xde, 0xc8, 0xc4,
	0x9f, 0xf6, 0x8b, 0x20, 0x3b, 0x7f, 0x00, 0x01, 0x37, 0x76, 0x7b, 0x63, 0x3f, 0xef, 0xe3, 0x9f,
	0x9f, 0xf8, 0x85, 0x13, 0xcd, 0xa5, 0x11, 0x92, 0xaf, 0x88, 0xe6, 0x5a, 0x0b, 0x25, 0x49, 0x1d,
	0x13, 0x5e, 0x73, 0x69, 0x34, 0x2e, 0x4a, 0x65, 0x54, 0x70, 0xbf, 0x53, 0xe0, 0x56, 0x81, 0xeb,
	0xf8, 0x0c, 0x65, 0x4a, 0xaf, 0x95, 0x26, 0x29, 0xd5, 0x9c, 0xd4, 0x71, 0xca, 0x0d, 0x8d, 0x49,
	0xa6, 0x84, 0x6c, 0x9a, 0xce, 0x1e, 0xe4, 0x2a, 0x57, 0xae, 0x24, 0xb6, 0x6a, 0x57, 0x51, 0xae,
	0x54, 0xbe, 0xe2, 0xc4, 0x7d, 0xa5, 0xd5, 0x82, 0xb0, 0xaa, 0xa4, 0xc6, 0x5a, 0x36, 0xfb, 0x8f,
	0x7a, 0x18, 0xb3, 0x29, 0xb8, 0xb6, 0x28, 0x29, 0x95, 0xec, 0x9d, 0x60, 0x66, 0xd9, 0x59, 0xfc,
	0x2d, 0xd1, 0x86, 0x9a, 0xaa, 0xa5, 0x8d, 0x14, 0xbc, 0x77, 0x69, 0xe9, 0xe7, 0x86, 0x96, 0x66,
	0xde, 0x00, 0x07, 0x01, 0x1c, 0x2c, 0x4a, 0xb5, 0x0e, 0xc1, 0x04, 0x4c, 0x8f, 0x13, 0x57, 0x07,
	0x27, 0xd0, 0x17, 0x2c, 0xf4, 0x27, 0x60, 0x3a, 0x48, 0x7c, 0xc1, 0x82, 0x08, 0xde, 0xd1, 0x55,
	0xaa, 0xb3, 0x52, 0x14, 0x96, 0x28, 0x3c, 0x72, 0x3b, 0x7f, 0xac, 0x59, 0x1f, 0xa9, 0x18, 0x0f,
	0x07, 0x8d, 0x8f, 0xad, 0xa3, 0x2b, 0x1f, 0x06, 0xee, 0xc4, 0x37, 0x05, 0xa3, 0x86, 0x1f, 0xe0,
	0xc8, 0x20, 0x84, 0x23, 0xca, 0x58, 0xc9, 0xb5, 0x0e, 0x6f, 0xb9, 0xe5, 0xee, 0x33, 0x78, 0x0e,
	0xc7, 0x5d, 0xa4, 0xe1, 0x70, 0x02, 0xa6, 0xb7, 0x2f, 0x4e, 0x71, 0x93, 0x39, 0xee, 0x32, 0xc7,
	0xaf, 0x5a, 0xc1, 0x6c, 0xbc, 0xfd, 0x7a, 0xee, 0x7d, 0xfc, 0x76, 0x0e, 0x92, 0xbe, 0x29, 0x78,
	0x01, 0x8f, 0xfb, 0xc4, 0xc3, 0x91, 0x73, 0x78, 0x88, 0xfb, 0x07, 0xe0, 0x22, 0xc7, 0x75, 0x8c,
	0x67, 0x9d, 0x66, 0x36, 0xb0, 0x26, 0xc9, 0xaf, 0xa6, 0xe8, 0x13, 0x80, 0x77, 0x5d, 0x1e, 0x97,
	0x92, 0x1d, 0x22, 0x8c, 0x18, 0x0e, 0x9b, 0x07, 0xe0, 0xb2, 0x38, 0xb9, 0x38, 0xfd, 0x07, 0xee,
	0xdc, 0x09, 0x92, 0x56, 0x18, 0x5d, 0x01, 0x38, 0x76, 0x88, 0xaf, 0xe9, 0xa6, 0xe5, 0x00, 0xff,
	0xe5, 0xf0, 0x6f, 0xe0, 0x38, 0xfa, 0x8d, 0xe3, 0x19, 0x1c, 0x15, 0x74, 0xb3, 0xe6, 0xd2, 0x38,
	0x3c, 0x9b, 0x7c, 0x33, 0x23, 0xd8, 0xce, 0x08, 0x6e, 0x67, 0x04, 0xbf, 0x54, 0x42, 0xb6, 0xa1,
	0x75, 0xfa, 0x59, 0xb2, 0xfd, 0x81, 0xbc, 0xcf, 0x3b, 0xe4, 0x6d, 0x77, 0x08, 0x5c, 0xef, 0x10,
	0xf8, 0xbe, 0x43, 0xe0, 0xc3, 0x1e, 0x79, 0xd7, 0x7b, 0xe4, 0x7d, 0xd9, 0x23, 0xef, 0xed, 0x93,
	0x5c, 0x98, 0x65, 0x95, 0xe2, 0x4c, 0xad, 0x49, 0x77, 0xbd, 0xc7, 0x6a, 0xb1, 0x10, 0x99, 0xa0,
	0x2b, 0xb2, 0xac, 0x52, 0xf2, 0xbe, 0x9f, 0x5f, 0x77, 0xeb, 0x74, 0xe8, 0xfe, 0xf7, 0xd3, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x3f, 0xa9, 0xb8, 0xe0, 0x03, 0x00, 0x00,
}

func (m *EventStartSession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStartSession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStartSession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x22
	}
	if m.Subscription != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateSession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateSession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateSession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEvents(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x22
	}
	if m.Subscription != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventEndSession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEndSession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEndSession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x22
	}
	if m.Subscription != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventPay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Payment.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Subscription != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventStartSession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	if m.Subscription != 0 {
		n += 1 + sovEvents(uint64(m.Subscription))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateSession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	if m.Subscription != 0 {
		n += 1 + sovEvents(uint64(m.Subscription))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovEvents(uint64(l))
	l = m.Bandwidth.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventEndSession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	if m.Subscription != 0 {
		n += 1 + sovEvents(uint64(m.Subscription))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func (m *EventPay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	if m.Subscription != 0 {
		n += 1 + sovEvents(uint64(m.Subscription))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Payment.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventStartSession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventStartSession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStartSession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventUpdateSession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventUpdateSession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateSession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventEndSession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventEndSession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEndSession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventPay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventPay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Payment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
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
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
