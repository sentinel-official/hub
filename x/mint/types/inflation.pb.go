// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/mint/v1/inflation.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// Inflation represents a message for handling inflation parameters.
type Inflation struct {
	// Field 1: Maximum inflation rate.
	// - (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec":
	//   Custom type definition for the field.
	// - (gogoproto.moretags) = "yaml:\"max\"": YAML tag for better representation.
	// - (gogoproto.nullable) = false: Field is not nullable.
	Max cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=max,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max" yaml:"max"`
	// Field 2: Minimum inflation rate.
	// - (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec":
	//   Custom type definition for the field.
	// - (gogoproto.moretags) = "yaml:\"min\"": YAML tag for better representation.
	// - (gogoproto.nullable) = false: Field is not nullable.
	Min cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=min,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min" yaml:"min"`
	// Field 3: Rate of change of inflation.
	// - (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec":
	//   Custom type definition for the field.
	// - (gogoproto.moretags) = "yaml:\"rate_change\"": YAML tag for better representation.
	// - (gogoproto.nullable) = false: Field is not nullable.
	RateChange cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=rate_change,json=rateChange,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"rate_change" yaml:"rate_change"`
	// Field 4: Timestamp indicating when the inflation parameters were set.
	// - (gogoproto.moretags) = "yaml:\"timestamp\"": YAML tag for better representation.
	// - (gogoproto.nullable) = false: Field is not nullable.
	// - (gogoproto.stdtime) = true: Use standard time representation for Go.
	Timestamp time.Time `protobuf:"bytes,4,opt,name=timestamp,proto3,stdtime" json:"timestamp" yaml:"timestamp"`
}

func (m *Inflation) Reset()         { *m = Inflation{} }
func (m *Inflation) String() string { return proto.CompactTextString(m) }
func (*Inflation) ProtoMessage()    {}
func (*Inflation) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e0e2355d878ebb, []int{0}
}
func (m *Inflation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Inflation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Inflation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Inflation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inflation.Merge(m, src)
}
func (m *Inflation) XXX_Size() int {
	return m.Size()
}
func (m *Inflation) XXX_DiscardUnknown() {
	xxx_messageInfo_Inflation.DiscardUnknown(m)
}

var xxx_messageInfo_Inflation proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Inflation)(nil), "sentinel.mint.v1.Inflation")
}

func init() { proto.RegisterFile("sentinel/mint/v1/inflation.proto", fileDescriptor_00e0e2355d878ebb) }

var fileDescriptor_00e0e2355d878ebb = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe3, 0x16, 0x21, 0xd5, 0x5d, 0xaa, 0x88, 0x21, 0x2a, 0xc8, 0xa9, 0x32, 0x75, 0xc1,
	0x56, 0xca, 0x04, 0x63, 0xcb, 0x82, 0xc4, 0x54, 0x10, 0x43, 0x17, 0xe4, 0x04, 0xc7, 0xb1, 0x88,
	0xed, 0xaa, 0x71, 0xab, 0xf6, 0x16, 0x3d, 0x06, 0x47, 0x29, 0x5b, 0x47, 0xc4, 0x10, 0x20, 0xbd,
	0x41, 0x4f, 0x80, 0x92, 0x28, 0x85, 0x11, 0xb1, 0xf9, 0xc9, 0xdf, 0xff, 0xd9, 0xfa, 0x1f, 0xec,
	0xa5, 0x4c, 0x19, 0xa1, 0x58, 0x42, 0xa4, 0x50, 0x86, 0x2c, 0x7c, 0x22, 0x54, 0x94, 0x50, 0x23,
	0xb4, 0xc2, 0xd3, 0x99, 0x36, 0xda, 0xee, 0xd4, 0x04, 0x2e, 0x08, 0xbc, 0xf0, 0xbb, 0x27, 0x5c,
	0x73, 0x5d, 0x5e, 0x92, 0xe2, 0x54, 0x71, 0x5d, 0x97, 0x6b, 0xcd, 0x13, 0x46, 0xca, 0x29, 0x98,
	0x47, 0xc4, 0x08, 0xc9, 0x52, 0x43, 0xe5, 0xb4, 0x02, 0xbc, 0xd7, 0x06, 0x6c, 0xdd, 0xd4, 0x72,
	0x7b, 0x04, 0x9b, 0x92, 0x2e, 0x1d, 0xd0, 0x03, 0xfd, 0xd6, 0xd0, 0xdf, 0x64, 0xae, 0xf5, 0x9e,
	0xb9, 0xa7, 0xa1, 0x4e, 0xa5, 0x4e, 0xd3, 0xa7, 0x67, 0x2c, 0x34, 0x91, 0xd4, 0xc4, 0xf8, 0x96,
	0x71, 0x1a, 0xae, 0xae, 0x59, 0xb8, 0xcf, 0x5c, 0xb8, 0xa2, 0x32, 0xb9, 0xf2, 0x24, 0x5d, 0x7a,
	0xe3, 0x22, 0x5d, 0x4a, 0x84, 0x72, 0x1a, 0xff, 0x91, 0x08, 0x55, 0x48, 0x84, 0xb2, 0x27, 0xb0,
	0x3d, 0xa3, 0x86, 0x3d, 0x86, 0x31, 0x55, 0x9c, 0x39, 0xcd, 0x52, 0x76, 0xf9, 0x37, 0x99, 0x5d,
	0xc9, 0x7e, 0xe5, 0xbd, 0x31, 0x2c, 0xa6, 0x51, 0x39, 0xd8, 0x0f, 0xb0, 0x75, 0xa8, 0xc1, 0x39,
	0xea, 0x81, 0x7e, 0x7b, 0xd0, 0xc5, 0x55, 0x51, 0xb8, 0x2e, 0x0a, 0xdf, 0xd7, 0xc4, 0xf0, 0xac,
	0x78, 0x75, 0x9f, 0xb9, 0x9d, 0x4a, 0x7b, 0x88, 0x7a, 0xeb, 0x0f, 0x17, 0x8c, 0x7f, 0x54, 0xc3,
	0xbb, 0xcd, 0x17, 0xb2, 0x5e, 0x72, 0x64, 0x6d, 0x72, 0x04, 0xb6, 0x39, 0x02, 0x9f, 0x39, 0x02,
	0xeb, 0x1d, 0xb2, 0xb6, 0x3b, 0x64, 0xbd, 0xed, 0x90, 0x35, 0xf1, 0xb9, 0x30, 0xf1, 0x3c, 0xc0,
	0xa1, 0x96, 0xa4, 0xde, 0xe0, 0xb9, 0x8e, 0x22, 0x11, 0x0a, 0x9a, 0x90, 0x78, 0x1e, 0x90, 0x85,
	0x3f, 0x20, 0xcb, 0x6a, 0xed, 0x66, 0x35, 0x65, 0x69, 0x70, 0x5c, 0xfe, 0xe8, 0xe2, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xc1, 0xa0, 0x4f, 0x93, 0x14, 0x02, 0x00, 0x00,
}

func (m *Inflation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Inflation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Inflation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintInflation(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.RateChange.Size()
		i -= size
		if _, err := m.RateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Min.Size()
		i -= size
		if _, err := m.Min.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Max.Size()
		i -= size
		if _, err := m.Max.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintInflation(dAtA []byte, offset int, v uint64) int {
	offset -= sovInflation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Inflation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Max.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.Min.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.RateChange.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func sovInflation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInflation(x uint64) (n int) {
	return sovInflation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Inflation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: Inflation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Inflation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Max.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Min.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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
func skipInflation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
				return 0, ErrInvalidLengthInflation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInflation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInflation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInflation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInflation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInflation = fmt.Errorf("proto: unexpected end of group")
)
