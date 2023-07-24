// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v2/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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
	InactivePendingDuration  time.Duration `protobuf:"bytes,1,opt,name=inactive_pending_duration,json=inactivePendingDuration,proto3,stdduration" json:"inactive_pending_duration"`
	ProofVerificationEnabled bool          `protobuf:"varint,2,opt,name=proof_verification_enabled,json=proofVerificationEnabled,proto3" json:"proof_verification_enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_98343ae164d22c30, []int{0}
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
	proto.RegisterType((*Params)(nil), "sentinel.session.v2.Params")
}

func init() { proto.RegisterFile("sentinel/session/v2/params.proto", fileDescriptor_98343ae164d22c30) }

var fileDescriptor_98343ae164d22c30 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3d, 0x4e, 0xc3, 0x30,
	0x00, 0x85, 0x6d, 0x86, 0xaa, 0x0a, 0x5b, 0x40, 0x22, 0xcd, 0xe0, 0x46, 0x4c, 0x5d, 0xb0, 0x51,
	0x58, 0x99, 0x2a, 0xd8, 0xab, 0x0c, 0x0c, 0x2c, 0x91, 0x93, 0x38, 0xae, 0xa5, 0xd4, 0x8e, 0x62,
	0x27, 0x82, 0x5b, 0x30, 0x72, 0x03, 0x38, 0x4a, 0xc6, 0x8e, 0x4c, 0xfc, 0x24, 0x17, 0x41, 0x75,
	0x6a, 0xc4, 0x66, 0xbf, 0xf7, 0xf9, 0x93, 0xfc, 0xbc, 0x48, 0x33, 0x69, 0x84, 0x64, 0x15, 0xd1,
	0x4c, 0x6b, 0xa1, 0x24, 0xe9, 0x62, 0x52, 0xd3, 0x86, 0xee, 0x34, 0xae, 0x1b, 0x65, 0x94, 0x7f,
	0xe6, 0x08, 0x7c, 0x24, 0x70, 0x17, 0x87, 0xe7, 0x5c, 0x71, 0x65, 0x7b, 0x72, 0x38, 0x4d, 0x68,
	0x88, 0xb8, 0x52, 0xbc, 0x62, 0xc4, 0xde, 0xb2, 0xb6, 0x24, 0x45, 0xdb, 0x50, 0x73, 0x78, 0x62,
	0x93, 0xcb, 0x37, 0xe8, 0xcd, 0x36, 0xd6, 0xed, 0xa7, 0xde, 0x42, 0x48, 0x9a, 0x1b, 0xd1, 0xb1,
	0xb4, 0x66, 0xb2, 0x10, 0x92, 0xa7, 0x8e, 0x0e, 0x60, 0x04, 0x57, 0xa7, 0xf1, 0x02, 0x4f, 0x3a,
	0xec, 0x74, 0xf8, 0xee, 0x08, 0xac, 0xe7, 0xfd, 0xe7, 0x12, 0xbc, 0x7e, 0x2d, 0x61, 0x72, 0xe1,
	0x2c, 0x9b, 0x49, 0xe2, 0x10, 0xff, 0xd6, 0x0b, 0xeb, 0x46, 0xa9, 0x32, 0xed, 0x58, 0x23, 0x4a,
	0x91, 0xdb, 0x34, 0x65, 0x92, 0x66, 0x15, 0x2b, 0x82, 0x93, 0x08, 0xae, 0xe6, 0x49, 0x60, 0x89,
	0x87, 0x7f, 0xc0, 0xfd, 0xd4, 0xaf, 0x93, 0xfe, 0x07, 0x81, 0xf7, 0x01, 0x81, 0x7e, 0x40, 0x70,
	0x3f, 0x20, 0xf8, 0x3d, 0x20, 0xf8, 0x32, 0x22, 0xb0, 0x1f, 0x11, 0xf8, 0x18, 0x11, 0x78, 0xbc,
	0xe6, 0xc2, 0x6c, 0xdb, 0x0c, 0xe7, 0x6a, 0x47, 0xdc, 0x42, 0x57, 0xaa, 0x2c, 0x45, 0x2e, 0x68,
	0x45, 0xb6, 0x6d, 0x46, 0x9e, 0xfe, 0x26, 0x35, 0xcf, 0x35, 0xd3, 0xd9, 0xcc, 0xfe, 0xe3, 0xe6,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x85, 0x43, 0xa7, 0x73, 0x01, 0x00, 0x00,
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
	if m.ProofVerificationEnabled {
		i--
		if m.ProofVerificationEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.InactivePendingDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.InactivePendingDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
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
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.InactivePendingDuration)
	n += 1 + l + sovParams(uint64(l))
	if m.ProofVerificationEnabled {
		n += 2
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field InactivePendingDuration", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.InactivePendingDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofVerificationEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProofVerificationEnabled = bool(v != 0)
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
