// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v1/params.proto

package v1

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

type Params struct {
	Deposit          types.Coin                               `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
	InactiveDuration time.Duration                            `protobuf:"bytes,2,opt,name=inactive_duration,json=inactiveDuration,proto3,stdduration" json:"inactive_duration"`
	MaxPrice         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=max_price,json=maxPrice,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_price"`
	MinPrice         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=min_price,json=minPrice,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_price"`
	StakingShare     github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,5,opt,name=staking_share,json=stakingShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking_share"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a408d0240644eb, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "sentinel.node.v1.Params")
}

func init() { proto.RegisterFile("sentinel/node/v1/params.proto", fileDescriptor_56a408d0240644eb) }

var fileDescriptor_56a408d0240644eb = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0x13, 0x5a, 0x4a, 0x1b, 0x40, 0x2a, 0x11, 0x43, 0xa8, 0x84, 0x53, 0x31, 0xa0, 0x2e,
	0xb5, 0x49, 0x61, 0x61, 0x0d, 0x7d, 0x80, 0x2a, 0x95, 0x18, 0x58, 0x2a, 0x27, 0x71, 0x13, 0xab,
	0x8d, 0x1d, 0xc5, 0x4e, 0x54, 0xde, 0x82, 0x91, 0x47, 0x40, 0x3c, 0x04, 0x73, 0xc7, 0x8e, 0x88,
	0xa1, 0xe5, 0xa6, 0x2f, 0x72, 0xe5, 0xc4, 0x91, 0xee, 0x74, 0x75, 0x97, 0x3b, 0xc5, 0xc9, 0x39,
	0xff, 0xf7, 0x1d, 0xc5, 0xc7, 0x7a, 0x2b, 0x08, 0x93, 0x94, 0x91, 0x3d, 0x62, 0x3c, 0x26, 0xa8,
	0xf2, 0x50, 0x8e, 0x0b, 0x9c, 0x09, 0x98, 0x17, 0x5c, 0x72, 0x7b, 0xdc, 0x95, 0xa1, 0x2a, 0xc3,
	0xca, 0x9b, 0x80, 0x88, 0x8b, 0x8c, 0x0b, 0x14, 0x62, 0xa1, 0xda, 0x43, 0x22, 0xb1, 0x87, 0x22,
	0x4e, 0x59, 0x9b, 0x98, 0xbc, 0x4e, 0x78, 0xc2, 0x9b, 0x23, 0x52, 0x27, 0xfd, 0x15, 0x24, 0x9c,
	0x27, 0x7b, 0x82, 0x9a, 0xb7, 0xb0, 0xdc, 0xa2, 0xb8, 0x2c, 0xb0, 0xa4, 0x5c, 0xa7, 0xde, 0xfd,
	0xe9, 0x59, 0x83, 0x55, 0x23, 0xb6, 0x3f, 0x5b, 0xcf, 0x62, 0x92, 0x73, 0x41, 0xa5, 0x63, 0x4e,
	0xcd, 0xd9, 0xf3, 0xc5, 0x1b, 0xd8, 0x2a, 0xa1, 0x52, 0x42, 0xad, 0x84, 0x5f, 0x38, 0x65, 0x7e,
	0xff, 0x78, 0x76, 0x8d, 0xa0, 0xeb, 0xb7, 0x57, 0xd6, 0x2b, 0xca, 0x70, 0x24, 0x69, 0x45, 0x36,
	0x9d, 0xc0, 0x79, 0xa2, 0x21, 0xed, 0x04, 0xb0, 0x9b, 0x00, 0x2e, 0x75, 0x83, 0x3f, 0x54, 0x90,
	0x9f, 0x17, 0xd7, 0x0c, 0xc6, 0x5d, 0xba, 0xab, 0xd9, 0xa9, 0x35, 0xca, 0xf0, 0x61, 0x93, 0x17,
	0x34, 0x22, 0x4e, 0x6f, 0xda, 0xbb, 0x7f, 0x9c, 0x0f, 0x8a, 0xf4, 0xfb, 0xe2, 0xce, 0x12, 0x2a,
	0xd3, 0x32, 0x84, 0x11, 0xcf, 0x90, 0xfe, 0x5d, 0xed, 0x63, 0x2e, 0xe2, 0x1d, 0x92, 0xdf, 0x73,
	0x22, 0x9a, 0x80, 0x08, 0x86, 0x19, 0x3e, 0xac, 0x14, 0xbc, 0x31, 0x51, 0xa6, 0x4d, 0xfd, 0xc7,
	0x30, 0x51, 0xd6, 0x9a, 0xd6, 0xd6, 0x4b, 0x21, 0xf1, 0x8e, 0xb2, 0x64, 0x23, 0x52, 0x5c, 0x10,
	0xe7, 0xe9, 0xd4, 0x9c, 0x8d, 0x7c, 0xa8, 0x90, 0xff, 0xce, 0xee, 0xfb, 0x07, 0x20, 0x97, 0x24,
	0x0a, 0x5e, 0x68, 0xc8, 0x5a, 0x31, 0xfc, 0xaf, 0xc7, 0x1b, 0x60, 0xfc, 0xaa, 0x81, 0x71, 0xac,
	0x81, 0x79, 0xaa, 0x81, 0xf9, 0xbf, 0x06, 0xe6, 0x8f, 0x2b, 0x30, 0x4e, 0x57, 0x60, 0xfc, 0xbd,
	0x02, 0xe3, 0xdb, 0xa7, 0x3b, 0xdc, 0x6e, 0xab, 0xe6, 0x7c, 0xbb, 0xa5, 0x11, 0xc5, 0x7b, 0x94,
	0x96, 0x21, 0xaa, 0xbc, 0x05, 0x3a, 0xb4, 0x7b, 0xd8, 0x68, 0xd4, 0x7a, 0x0d, 0x9a, 0xfb, 0xfa,
	0x78, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xbf, 0x2c, 0xf6, 0xa8, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.StakingShare.Size()
		i -= size
		if _, err := m.StakingShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.MinPrice) > 0 {
		for iNdEx := len(m.MinPrice) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinPrice[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MaxPrice) > 0 {
		for iNdEx := len(m.MaxPrice) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxPrice[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.InactiveDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.InactiveDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.InactiveDuration)
	n += 1 + l + sovParams(uint64(l))
	if len(m.MaxPrice) > 0 {
		for _, e := range m.MaxPrice {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MinPrice) > 0 {
		for _, e := range m.MinPrice {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.StakingShare.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.InactiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxPrice = append(m.MaxPrice, types.Coin{})
			if err := m.MaxPrice[len(m.MaxPrice)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinPrice = append(m.MinPrice, types.Coin{})
			if err := m.MinPrice[len(m.MinPrice)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
