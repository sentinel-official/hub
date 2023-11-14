// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/types/v1/bandwidth.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Bandwidth represents information about upload and download bandwidth.
type Bandwidth struct {
	// Upload bandwidth value represented as a string.
	// It uses a custom type "cosmossdk.io/math.Int".
	// The value is not nullable, as indicated by "(gogoproto.nullable) = false".
	Upload cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=upload,proto3,customtype=cosmossdk.io/math.Int" json:"upload"`
	// Download bandwidth value represented as a string.
	// It uses a custom type "cosmossdk.io/math.Int".
	// The value is not nullable, as indicated by "(gogoproto.nullable) = false".
	Download cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=download,proto3,customtype=cosmossdk.io/math.Int" json:"download"`
}

func (m *Bandwidth) Reset()         { *m = Bandwidth{} }
func (m *Bandwidth) String() string { return proto.CompactTextString(m) }
func (*Bandwidth) ProtoMessage()    {}
func (*Bandwidth) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed82d6e988b8939b, []int{0}
}
func (m *Bandwidth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bandwidth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bandwidth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bandwidth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bandwidth.Merge(m, src)
}
func (m *Bandwidth) XXX_Size() int {
	return m.Size()
}
func (m *Bandwidth) XXX_DiscardUnknown() {
	xxx_messageInfo_Bandwidth.DiscardUnknown(m)
}

var xxx_messageInfo_Bandwidth proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Bandwidth)(nil), "sentinel.types.v1.Bandwidth")
}

func init() { proto.RegisterFile("sentinel/types/v1/bandwidth.proto", fileDescriptor_ed82d6e988b8939b) }

var fileDescriptor_ed82d6e988b8939b = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4e, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4b, 0xcd, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0x4a,
	0xcc, 0x4b, 0x29, 0xcf, 0x4c, 0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84,
	0x29, 0xd1, 0x03, 0x2b, 0xd1, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xea,
	0x83, 0x58, 0x10, 0x85, 0x4a, 0xb5, 0x5c, 0x9c, 0x4e, 0x30, 0xbd, 0x42, 0xa6, 0x5c, 0x6c, 0xa5,
	0x05, 0x39, 0xf9, 0x89, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0xb2, 0x27, 0xee, 0xc9,
	0x33, 0xdc, 0xba, 0x27, 0x2f, 0x9a, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x5c, 0x9c, 0x92, 0xad, 0x97,
	0x99, 0xaf, 0x9f, 0x9b, 0x58, 0x92, 0xa1, 0xe7, 0x99, 0x57, 0x12, 0x04, 0x55, 0x2c, 0x64, 0xc9,
	0xc5, 0x91, 0x92, 0x5f, 0x9e, 0x07, 0xd6, 0xc8, 0x44, 0x8c, 0x46, 0xb8, 0x72, 0x27, 0xef, 0x13,
	0x0f, 0xe5, 0x18, 0x56, 0x3c, 0x92, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0xa7,
	0x74, 0xf3, 0xd3, 0xd2, 0x32, 0x93, 0x33, 0x13, 0x73, 0xf4, 0x33, 0x4a, 0x93, 0x40, 0xde, 0x07,
	0x7b, 0x32, 0x89, 0x0d, 0xec, 0x25, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xe9, 0xa8,
	0xd5, 0x20, 0x01, 0x00, 0x00,
}

func (m *Bandwidth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bandwidth) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bandwidth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Download.Size()
		i -= size
		if _, err := m.Download.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBandwidth(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Upload.Size()
		i -= size
		if _, err := m.Upload.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBandwidth(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintBandwidth(dAtA []byte, offset int, v uint64) int {
	offset -= sovBandwidth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bandwidth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Upload.Size()
	n += 1 + l + sovBandwidth(uint64(l))
	l = m.Download.Size()
	n += 1 + l + sovBandwidth(uint64(l))
	return n
}

func sovBandwidth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBandwidth(x uint64) (n int) {
	return sovBandwidth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bandwidth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBandwidth
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
			return fmt.Errorf("proto: Bandwidth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bandwidth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandwidth
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
				return ErrInvalidLengthBandwidth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBandwidth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Upload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Download", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandwidth
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
				return ErrInvalidLengthBandwidth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBandwidth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Download.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBandwidth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBandwidth
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
func skipBandwidth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBandwidth
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
					return 0, ErrIntOverflowBandwidth
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
					return 0, ErrIntOverflowBandwidth
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
				return 0, ErrInvalidLengthBandwidth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBandwidth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBandwidth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBandwidth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBandwidth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBandwidth = fmt.Errorf("proto: unexpected end of group")
)
