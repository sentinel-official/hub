// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/session.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	types "github.com/sentinel-official/hub/types"
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

type Session struct {
	Id           uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Subscription uint64          `protobuf:"varint,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Node         string          `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Address      string          `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Duration     time.Duration   `protobuf:"bytes,5,opt,name=duration,proto3,stdduration" json:"duration"`
	Bandwidth    types.Bandwidth `protobuf:"bytes,6,opt,name=bandwidth,proto3" json:"bandwidth"`
	Status       types.Status    `protobuf:"varint,7,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt     time.Time       `protobuf:"bytes,8,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Session) Reset()      { *m = Session{} }
func (*Session) ProtoMessage() {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b6692acc1d7206a, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Session.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Session)(nil), "sentinel.session.v1.Session")
}

func init() { proto.RegisterFile("sentinel/session/v1/session.proto", fileDescriptor_8b6692acc1d7206a) }

var fileDescriptor_8b6692acc1d7206a = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x73, 0xd3, 0x30,
	0x18, 0xc6, 0xad, 0x10, 0xf2, 0x47, 0xf4, 0x3a, 0x08, 0x06, 0x35, 0xc7, 0xc9, 0xa1, 0x53, 0x16,
	0x24, 0x5c, 0x36, 0x16, 0x68, 0x8e, 0x89, 0xd1, 0x65, 0x62, 0xe1, 0x24, 0x4b, 0x71, 0x74, 0x97,
	0x58, 0xb9, 0x48, 0x36, 0xf0, 0x0d, 0x18, 0x3b, 0x76, 0xcc, 0xc7, 0xe9, 0xd8, 0x91, 0x09, 0xb8,
	0x64, 0x61, 0xe6, 0x13, 0x70, 0x96, 0x2d, 0xf7, 0x20, 0x6c, 0xaf, 0xde, 0xf7, 0x79, 0xf4, 0xfe,
	0x1e, 0xcb, 0xf0, 0x99, 0x55, 0x85, 0xd3, 0x85, 0x5a, 0x31, 0xab, 0xac, 0xd5, 0xa6, 0x60, 0x55,
	0x12, 0x4a, 0xba, 0xd9, 0x1a, 0x67, 0xd0, 0xe3, 0x20, 0xa1, 0xa1, 0x5f, 0x25, 0x13, 0x92, 0x19,
	0xbb, 0x36, 0x96, 0x09, 0x6e, 0x15, 0xab, 0x12, 0xa1, 0x1c, 0x4f, 0x58, 0x66, 0x74, 0x6b, 0x9a,
	0x3c, 0xc9, 0x4d, 0x6e, 0x7c, 0xc9, 0xea, 0xaa, 0xed, 0x92, 0xdc, 0x98, 0x7c, 0xa5, 0x98, 0x3f,
	0x89, 0x72, 0xc1, 0x64, 0xb9, 0xe5, 0xae, 0x5b, 0x35, 0x89, 0xff, 0x9d, 0x3b, 0xbd, 0x56, 0xd6,
	0xf1, 0xf5, 0xa6, 0x15, 0xdc, 0xe3, 0xba, 0x2f, 0x1b, 0x65, 0x6b, 0x58, 0xc1, 0x0b, 0xf9, 0x49,
	0x4b, 0xb7, 0x0c, 0x3b, 0x8e, 0x25, 0xd6, 0x71, 0x57, 0xda, 0x66, 0x7e, 0xfe, 0xbb, 0x07, 0x87,
	0x57, 0x4d, 0x10, 0x74, 0x0a, 0x7b, 0x5a, 0x62, 0x30, 0x05, 0xb3, 0x7e, 0xda, 0xd3, 0x12, 0x9d,
	0xc3, 0x13, 0x5b, 0x0a, 0x9b, 0x6d, 0xf5, 0xa6, 0xa6, 0xc2, 0x3d, 0x3f, 0xf9, 0xab, 0x87, 0x10,
	0xec, 0x17, 0x46, 0x2a, 0xfc, 0x60, 0x0a, 0x66, 0xe3, 0xd4, 0xd7, 0x08, 0xc3, 0x21, 0x97, 0x72,
	0xab, 0xac, 0xc5, 0x7d, 0xdf, 0x0e, 0x47, 0xf4, 0x1a, 0x8e, 0x42, 0x46, 0xfc, 0x70, 0x0a, 0x66,
	0x8f, 0x2e, 0xce, 0x68, 0x13, 0x92, 0x86, 0x90, 0xf4, 0x6d, 0x2b, 0x98, 0x8f, 0x6e, 0xbf, 0xc7,
	0xd1, 0xcd, 0x8f, 0x18, 0xa4, 0x9d, 0x09, 0xbd, 0x81, 0xe3, 0x2e, 0x21, 0x1e, 0xf8, 0x1b, 0x9e,
	0xd2, 0xee, 0x45, 0x7c, 0x44, 0x5a, 0x25, 0x74, 0x1e, 0x34, 0xf3, 0x7e, 0x7d, 0x49, 0x7a, 0x6f,
	0x42, 0x09, 0x1c, 0x34, 0x1f, 0x00, 0x0f, 0xa7, 0x60, 0x76, 0x7a, 0x71, 0xf6, 0x1f, 0xfb, 0x95,
	0x17, 0xa4, 0xad, 0x10, 0x5d, 0xc2, 0x71, 0x53, 0x7d, 0xe4, 0x0e, 0x8f, 0xfc, 0xd2, 0xc9, 0x11,
	0xf6, 0xfb, 0xf0, 0x36, 0x0d, 0xf7, 0xb5, 0xe7, 0x6e, 0x6c, 0x97, 0xee, 0xd5, 0xc9, 0xd7, 0x5d,
	0x1c, 0xdd, 0xec, 0xe2, 0xe8, 0xd7, 0x2e, 0x8e, 0xe6, 0xef, 0x6e, 0xf7, 0x04, 0xdc, 0xed, 0x09,
	0xf8, 0xb9, 0x27, 0xe0, 0xfa, 0x40, 0xa2, 0xbb, 0x03, 0x89, 0xbe, 0x1d, 0x48, 0xf4, 0xe1, 0x45,
	0xae, 0xdd, 0xb2, 0x14, 0x34, 0x33, 0x6b, 0x16, 0xb8, 0x9e, 0x9b, 0xc5, 0x42, 0x67, 0x9a, 0xaf,
	0xd8, 0xb2, 0x14, 0xec, 0x73, 0xf7, 0x6b, 0x7a, 0x5c, 0x31, 0xf0, 0x04, 0x2f, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xb9, 0x69, 0xaa, 0x7c, 0xbb, 0x02, 0x00, 0x00,
}

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSession(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	if m.Status != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSession(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintSession(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x2a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSession(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintSession(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Subscription != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSession(dAtA []byte, offset int, v uint64) int {
	offset -= sovSession(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSession(uint64(m.Id))
	}
	if m.Subscription != 0 {
		n += 1 + sovSession(uint64(m.Subscription))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovSession(uint64(l))
	l = m.Bandwidth.Size()
	n += 1 + l + sovSession(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSession(uint64(m.Status))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovSession(uint64(l))
	return n
}

func sovSession(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
					return ErrIntOverflowSession
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
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
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
func skipSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
				return 0, ErrInvalidLengthSession
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSession
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSession
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSession        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSession = fmt.Errorf("proto: unexpected end of group")
)