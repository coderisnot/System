// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/message.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// *
// Generic message header for all messages passing through CCTP
// The message body is dynamically-sized to support custom message body
// formats. Other fields must be fixed-size to avoid hash collisions.
//
// Padding: uintNN fields are left-padded, and bytesNN fields are right-padded.
//
// @param version the version of the message format
// @param source_domain domain of home chain
// @param destination_domain domain of destination chain
// @param nonce destination-specific nonce
// @param sender address of sender on source chain as bytes32
// @param recipient address of recipient on destination chain as bytes32
// @param destination_caller address of caller on destination chain as bytes32
// @param message_body raw bytes of message body
type Message struct {
	Version           uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	SourceDomain      uint32 `protobuf:"varint,2,opt,name=source_domain,json=sourceDomain,proto3" json:"source_domain,omitempty"`
	DestinationDomain uint32 `protobuf:"varint,3,opt,name=destination_domain,json=destinationDomain,proto3" json:"destination_domain,omitempty"`
	Nonce             uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Sender            []byte `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient         []byte `protobuf:"bytes,6,opt,name=recipient,proto3" json:"recipient,omitempty"`
	DestinationCaller []byte `protobuf:"bytes,7,opt,name=destination_caller,json=destinationCaller,proto3" json:"destination_caller,omitempty"`
	MessageBody       []byte `protobuf:"bytes,8,opt,name=message_body,json=messageBody,proto3" json:"message_body,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1466d44ee11803b, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Message) GetSourceDomain() uint32 {
	if m != nil {
		return m.SourceDomain
	}
	return 0
}

func (m *Message) GetDestinationDomain() uint32 {
	if m != nil {
		return m.DestinationDomain
	}
	return 0
}

func (m *Message) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Message) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Message) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Message) GetDestinationCaller() []byte {
	if m != nil {
		return m.DestinationCaller
	}
	return nil
}

func (m *Message) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "circle.cctp.v1.Message")
}

func init() { proto.RegisterFile("circle/cctp/v1/message.proto", fileDescriptor_c1466d44ee11803b) }

var fileDescriptor_c1466d44ee11803b = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x46, 0xeb, 0xd2, 0x1f, 0x30, 0x2d, 0x12, 0x16, 0x42, 0x1e, 0x2a, 0xab, 0xc0, 0x92, 0x81,
	0x26, 0xaa, 0x78, 0x83, 0x82, 0xd8, 0x58, 0x32, 0xb2, 0x54, 0x89, 0x73, 0x29, 0x96, 0x12, 0x3b,
	0xb2, 0xdd, 0x88, 0xbc, 0x05, 0x0b, 0xef, 0xc4, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0xb0, 0x53,
	0x51, 0x31, 0xde, 0xef, 0x1c, 0xe9, 0x4a, 0x07, 0xcf, 0xb8, 0xd0, 0x3c, 0x87, 0x88, 0x73, 0x5b,
	0x46, 0xd5, 0x32, 0x2a, 0xc0, 0x98, 0x64, 0x03, 0x61, 0xa9, 0x95, 0x55, 0xe4, 0xcc, 0xd3, 0xf0,
	0x97, 0x86, 0xd5, 0xf2, 0xfa, 0xa3, 0x8f, 0xc7, 0x4f, 0xde, 0x20, 0x14, 0x8f, 0x2b, 0xd0, 0x46,
	0x28, 0x49, 0xd1, 0x1c, 0x05, 0xd3, 0x78, 0x7f, 0x92, 0x1b, 0x3c, 0x35, 0x6a, 0xab, 0x39, 0xac,
	0x33, 0x55, 0x24, 0x42, 0xd2, 0xbe, 0xe3, 0x13, 0x3f, 0x3e, 0xb8, 0x8d, 0x2c, 0x30, 0xc9, 0xc0,
	0x58, 0x21, 0x13, 0x2b, 0x94, 0xdc, 0x9b, 0x47, 0xce, 0x3c, 0x3f, 0x20, 0x9d, 0x7e, 0x81, 0x87,
	0x52, 0x49, 0x0e, 0x74, 0x30, 0x47, 0xc1, 0x20, 0xf6, 0x07, 0xb9, 0xc4, 0x23, 0x03, 0x32, 0x03,
	0x4d, 0x87, 0x73, 0x14, 0x4c, 0xe2, 0xee, 0x22, 0x33, 0x7c, 0xa2, 0x81, 0x8b, 0x52, 0x80, 0xb4,
	0x74, 0xe4, 0xd0, 0xdf, 0xf0, 0xff, 0x35, 0x4f, 0xf2, 0x1c, 0x34, 0x1d, 0x3b, 0xed, 0xf0, 0xf5,
	0xbd, 0x03, 0xe4, 0x0a, 0x4f, 0xba, 0x2a, 0xeb, 0x54, 0x65, 0x35, 0x3d, 0x76, 0xe2, 0x69, 0xb7,
	0xad, 0x54, 0x56, 0xaf, 0x1e, 0x3f, 0x1b, 0x86, 0x76, 0x0d, 0x43, 0xdf, 0x0d, 0x43, 0xef, 0x2d,
	0xeb, 0xed, 0x5a, 0xd6, 0xfb, 0x6a, 0x59, 0xef, 0xf9, 0x76, 0x23, 0xec, 0xeb, 0x36, 0x0d, 0xb9,
	0x2a, 0x22, 0x1f, 0xf3, 0x45, 0xc8, 0x48, 0xaa, 0x34, 0x87, 0x85, 0x6b, 0xfe, 0xe6, 0xd3, 0xdb,
	0xba, 0x04, 0x93, 0x8e, 0x5c, 0xf6, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xf8, 0x06,
	0xdd, 0x96, 0x01, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageBody) > 0 {
		i -= len(m.MessageBody)
		copy(dAtA[i:], m.MessageBody)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.MessageBody)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DestinationCaller) > 0 {
		i -= len(m.DestinationCaller)
		copy(dAtA[i:], m.DestinationCaller)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.DestinationCaller)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Nonce != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x20
	}
	if m.DestinationDomain != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.DestinationDomain))
		i--
		dAtA[i] = 0x18
	}
	if m.SourceDomain != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.SourceDomain))
		i--
		dAtA[i] = 0x10
	}
	if m.Version != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovMessage(uint64(m.Version))
	}
	if m.SourceDomain != 0 {
		n += 1 + sovMessage(uint64(m.SourceDomain))
	}
	if m.DestinationDomain != 0 {
		n += 1 + sovMessage(uint64(m.DestinationDomain))
	}
	if m.Nonce != 0 {
		n += 1 + sovMessage(uint64(m.Nonce))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.DestinationCaller)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.MessageBody)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceDomain", wireType)
			}
			m.SourceDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationDomain", wireType)
			}
			m.DestinationDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestinationDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationCaller", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationCaller = append(m.DestinationCaller[:0], dAtA[iNdEx:postIndex]...)
			if m.DestinationCaller == nil {
				m.DestinationCaller = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageBody", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageBody = append(m.MessageBody[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageBody == nil {
				m.MessageBody = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
