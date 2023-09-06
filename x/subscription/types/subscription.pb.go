// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v2/subscription.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types "github.com/sentinel-official/hub/v1/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type SubscriptionType int32

const (
	TypeUnspecified SubscriptionType = 0
	TypeNode        SubscriptionType = 1
	TypePlan        SubscriptionType = 2
)

var SubscriptionType_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_NODE",
	2: "TYPE_PLAN",
}

var SubscriptionType_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"TYPE_NODE":        1,
	"TYPE_PLAN":        2,
}

func (x SubscriptionType) String() string {
	return proto.EnumName(SubscriptionType_name, int32(x))
}

func (SubscriptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{0}
}

type BaseSubscription struct {
	ID         uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address    string       `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	InactiveAt time.Time    `protobuf:"bytes,3,opt,name=inactive_at,json=inactiveAt,proto3,stdtime" json:"inactive_at"`
	Status     types.Status `protobuf:"varint,4,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt   time.Time    `protobuf:"bytes,5,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *BaseSubscription) Reset()         { *m = BaseSubscription{} }
func (m *BaseSubscription) String() string { return proto.CompactTextString(m) }
func (*BaseSubscription) ProtoMessage()    {}
func (*BaseSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{0}
}
func (m *BaseSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseSubscription.Merge(m, src)
}
func (m *BaseSubscription) XXX_Size() int {
	return m.Size()
}
func (m *BaseSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_BaseSubscription proto.InternalMessageInfo

type NodeSubscription struct {
	*BaseSubscription `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base,omitempty"`
	NodeAddress       string      `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	Gigabytes         int64       `protobuf:"varint,3,opt,name=gigabytes,proto3" json:"gigabytes,omitempty"`
	Hours             int64       `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	Deposit           types1.Coin `protobuf:"bytes,5,opt,name=deposit,proto3" json:"deposit"`
}

func (m *NodeSubscription) Reset()         { *m = NodeSubscription{} }
func (m *NodeSubscription) String() string { return proto.CompactTextString(m) }
func (*NodeSubscription) ProtoMessage()    {}
func (*NodeSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{1}
}
func (m *NodeSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSubscription.Merge(m, src)
}
func (m *NodeSubscription) XXX_Size() int {
	return m.Size()
}
func (m *NodeSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSubscription proto.InternalMessageInfo

type PlanSubscription struct {
	*BaseSubscription `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base,omitempty"`
	PlanID            uint64 `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Denom             string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *PlanSubscription) Reset()         { *m = PlanSubscription{} }
func (m *PlanSubscription) String() string { return proto.CompactTextString(m) }
func (*PlanSubscription) ProtoMessage()    {}
func (*PlanSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{2}
}
func (m *PlanSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlanSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlanSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlanSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanSubscription.Merge(m, src)
}
func (m *PlanSubscription) XXX_Size() int {
	return m.Size()
}
func (m *PlanSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_PlanSubscription proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("sentinel.subscription.v2.SubscriptionType", SubscriptionType_name, SubscriptionType_value)
	proto.RegisterType((*BaseSubscription)(nil), "sentinel.subscription.v2.BaseSubscription")
	proto.RegisterType((*NodeSubscription)(nil), "sentinel.subscription.v2.NodeSubscription")
	proto.RegisterType((*PlanSubscription)(nil), "sentinel.subscription.v2.PlanSubscription")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v2/subscription.proto", fileDescriptor_f6350e663da1ca66)
}

var fileDescriptor_f6350e663da1ca66 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xb1, 0x6f, 0xd3, 0x4e,
	0x18, 0xf5, 0xa5, 0x6e, 0xda, 0x5c, 0xaa, 0xdf, 0xcf, 0x32, 0x15, 0x4a, 0x03, 0x72, 0x42, 0x59,
	0x42, 0x11, 0x77, 0x4a, 0x58, 0x60, 0x8c, 0x9b, 0x20, 0x45, 0x42, 0x21, 0x72, 0xdb, 0x81, 0x2e,
	0xd1, 0xd9, 0xbe, 0xb8, 0x27, 0x25, 0x3e, 0x2b, 0x77, 0x89, 0xe8, 0xc0, 0xc0, 0x86, 0x3a, 0x55,
	0x62, 0xee, 0xc4, 0xc2, 0x9f, 0xd2, 0xb1, 0x23, 0x53, 0x80, 0x74, 0xe3, 0xaf, 0x40, 0x77, 0xb6,
	0xdb, 0x06, 0x89, 0x81, 0x81, 0xed, 0xbe, 0x7b, 0xef, 0xcb, 0xf7, 0xde, 0xbb, 0x2f, 0x86, 0x4f,
	0x05, 0x8d, 0x25, 0x8b, 0xe9, 0x18, 0x8b, 0x99, 0x2f, 0x82, 0x29, 0x4b, 0x24, 0xe3, 0x31, 0x9e,
	0xb7, 0x56, 0x6a, 0x94, 0x4c, 0xb9, 0xe4, 0x76, 0x25, 0x27, 0xa3, 0x15, 0x70, 0xde, 0xaa, 0x3a,
	0x01, 0x17, 0x13, 0x2e, 0xb0, 0x4f, 0x04, 0xc5, 0xf3, 0xa6, 0x4f, 0x25, 0x69, 0xe2, 0x80, 0xb3,
	0xac, 0xb3, 0xba, 0x1d, 0xf1, 0x88, 0xeb, 0x23, 0x56, 0xa7, 0xec, 0xb6, 0x16, 0x71, 0x1e, 0x8d,
	0x29, 0xd6, 0x95, 0x3f, 0x1b, 0x61, 0xc9, 0x26, 0x54, 0x48, 0x32, 0x49, 0x32, 0x82, 0x73, 0xa3,
	0x4e, 0x9e, 0x26, 0x54, 0xe0, 0x79, 0x13, 0x0b, 0x49, 0xe4, 0x4c, 0xa4, 0xf8, 0xee, 0x87, 0x02,
	0xb4, 0x5c, 0x22, 0xe8, 0xc1, 0x1d, 0x39, 0xf6, 0x7d, 0x58, 0x60, 0x61, 0x05, 0xd4, 0x41, 0xc3,
	0x74, 0x8b, 0xcb, 0x45, 0xad, 0xd0, 0xeb, 0x78, 0x05, 0x16, 0xda, 0x15, 0xb8, 0x41, 0xc2, 0x70,
	0x4a, 0x85, 0xa8, 0x14, 0xea, 0xa0, 0x51, 0xf2, 0xf2, 0xd2, 0xee, 0xc2, 0x32, 0x8b, 0x49, 0x20,
	0xd9, 0x9c, 0x0e, 0x89, 0xac, 0xac, 0xd5, 0x41, 0xa3, 0xdc, 0xaa, 0xa2, 0x54, 0x1d, 0xca, 0xd5,
	0xa1, 0xc3, 0x5c, 0x9d, 0xbb, 0x79, 0xb9, 0xa8, 0x19, 0xe7, 0xdf, 0x6a, 0xc0, 0x83, 0x79, 0x63,
	0x5b, 0xda, 0x4d, 0x58, 0x4c, 0xd5, 0x55, 0xcc, 0x3a, 0x68, 0xfc, 0xd7, 0xda, 0x41, 0x37, 0x79,
	0x69, 0xf9, 0x68, 0xde, 0x44, 0x07, 0x9a, 0xe0, 0x65, 0x44, 0xbb, 0x0d, 0x4b, 0xe9, 0x49, 0xcd,
	0x5d, 0xff, 0x8b, 0xb9, 0x9b, 0x69, 0x5b, 0x5b, 0xee, 0xfe, 0x04, 0xd0, 0xea, 0xf3, 0x70, 0x35,
	0x83, 0x0e, 0x34, 0xd5, 0x53, 0xe8, 0x14, 0xca, 0xad, 0x3d, 0xf4, 0xa7, 0x87, 0x43, 0xbf, 0xa7,
	0xe7, 0x9a, 0x57, 0x8b, 0x1a, 0xf0, 0x74, 0xb7, 0xfd, 0x08, 0x6e, 0xc5, 0x3c, 0xa4, 0xc3, 0xd5,
	0xd8, 0xca, 0xea, 0xae, 0x9d, 0x45, 0xf7, 0x10, 0x96, 0x22, 0x16, 0x11, 0xff, 0x54, 0x52, 0xa1,
	0x83, 0x5b, 0xf3, 0x6e, 0x2f, 0xec, 0x6d, 0xb8, 0x7e, 0xc2, 0x67, 0xd3, 0x34, 0x90, 0x35, 0x2f,
	0x2d, 0xec, 0x97, 0x70, 0x23, 0xa4, 0x09, 0x17, 0x2c, 0xb7, 0xbc, 0x83, 0xd2, 0xf5, 0x41, 0x6a,
	0x2a, 0xca, 0xd6, 0x07, 0xed, 0x73, 0x16, 0xbb, 0xa6, 0x72, 0xec, 0xe5, 0xfc, 0xdd, 0x4f, 0x00,
	0x5a, 0x83, 0x31, 0x89, 0xff, 0x81, 0xd9, 0xc7, 0x70, 0x23, 0x19, 0x93, 0x78, 0xc8, 0x42, 0xed,
	0xd3, 0x74, 0xe1, 0x72, 0x51, 0x2b, 0xaa, 0x61, 0xbd, 0x8e, 0x57, 0x54, 0x50, 0x2f, 0x54, 0x86,
	0x42, 0x1a, 0xf3, 0x89, 0xb6, 0x5a, 0xf2, 0xd2, 0x62, 0xef, 0x3d, 0xb4, 0xee, 0xfe, 0xec, 0xe1,
	0x69, 0x42, 0xed, 0x27, 0xd0, 0x3a, 0x7c, 0x3b, 0xe8, 0x0e, 0x8f, 0xfa, 0x07, 0x83, 0xee, 0x7e,
	0xef, 0x55, 0xaf, 0xdb, 0xb1, 0x8c, 0xea, 0xbd, 0xb3, 0x8b, 0xfa, 0xff, 0x0a, 0x3f, 0x8a, 0x45,
	0x42, 0x03, 0x36, 0x62, 0x34, 0xb4, 0x1f, 0xc0, 0x92, 0xa6, 0xf6, 0xdf, 0x74, 0xba, 0x16, 0xa8,
	0x6e, 0x9d, 0x5d, 0xd4, 0x37, 0x15, 0x47, 0xbd, 0xea, 0x0d, 0x38, 0x78, 0xdd, 0xee, 0x5b, 0x85,
	0x5b, 0x50, 0x09, 0xab, 0x9a, 0x1f, 0x3f, 0x3b, 0x86, 0x7b, 0x7c, 0xf9, 0xc3, 0x31, 0xbe, 0x2c,
	0x1d, 0xe3, 0x72, 0xe9, 0x80, 0xab, 0xa5, 0x03, 0xbe, 0x2f, 0x1d, 0x70, 0x7e, 0xed, 0x18, 0x57,
	0xd7, 0x8e, 0xf1, 0xf5, 0xda, 0x31, 0x8e, 0x5f, 0x44, 0x4c, 0x9e, 0xcc, 0x7c, 0x14, 0xf0, 0x09,
	0xce, 0xd3, 0x79, 0xc6, 0x47, 0x23, 0x16, 0x30, 0x32, 0xc6, 0x27, 0x33, 0x5f, 0xfd, 0xb3, 0xde,
	0xad, 0x7e, 0x02, 0xf4, 0xca, 0xfa, 0x45, 0xbd, 0x85, 0xcf, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x3b, 0x2b, 0xa2, 0x23, 0x28, 0x04, 0x00, 0x00,
}

func (m *BaseSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSubscription(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.Status != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.InactiveAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSubscription(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NodeSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSubscription(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Hours != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x20
	}
	if m.Gigabytes != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Gigabytes))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseSubscription != nil {
		{
			size, err := m.BaseSubscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubscription(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PlanSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlanSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlanSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PlanID != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PlanID))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseSubscription != nil {
		{
			size, err := m.BaseSubscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubscription(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubscription(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubscription(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovSubscription(uint64(m.ID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt)
	n += 1 + l + sovSubscription(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSubscription(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovSubscription(uint64(l))
	return n
}

func (m *NodeSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseSubscription != nil {
		l = m.BaseSubscription.Size()
		n += 1 + l + sovSubscription(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.Gigabytes != 0 {
		n += 1 + sovSubscription(uint64(m.Gigabytes))
	}
	if m.Hours != 0 {
		n += 1 + sovSubscription(uint64(m.Hours))
	}
	l = m.Deposit.Size()
	n += 1 + l + sovSubscription(uint64(l))
	return n
}

func (m *PlanSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseSubscription != nil {
		l = m.BaseSubscription.Size()
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.PlanID != 0 {
		n += 1 + sovSubscription(uint64(m.PlanID))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	return n
}

func sovSubscription(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubscription(x uint64) (n int) {
	return sovSubscription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: BaseSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.InactiveAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func (m *NodeSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: NodeSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseSubscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseSubscription == nil {
				m.BaseSubscription = &BaseSubscription{}
			}
			if err := m.BaseSubscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gigabytes", wireType)
			}
			m.Gigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func (m *PlanSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: PlanSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlanSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseSubscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseSubscription == nil {
				m.BaseSubscription = &BaseSubscription{}
			}
			if err := m.BaseSubscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanID", wireType)
			}
			m.PlanID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func skipSubscription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
				return 0, ErrInvalidLengthSubscription
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubscription
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubscription
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubscription        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubscription          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubscription = fmt.Errorf("proto: unexpected end of group")
)
