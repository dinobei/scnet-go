// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_image.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ImageHeader struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Size                 int32    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageHeader) Reset()         { *m = ImageHeader{} }
func (m *ImageHeader) String() string { return proto.CompactTextString(m) }
func (*ImageHeader) ProtoMessage()    {}
func (*ImageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4cbcb06a3b852c3, []int{0}
}

func (m *ImageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageHeader.Unmarshal(m, b)
}
func (m *ImageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageHeader.Marshal(b, m, deterministic)
}
func (m *ImageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageHeader.Merge(m, src)
}
func (m *ImageHeader) XXX_Size() int {
	return xxx_messageInfo_ImageHeader.Size(m)
}
func (m *ImageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ImageHeader proto.InternalMessageInfo

func (m *ImageHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageHeader) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ImageHeader) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ImageHeader) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ImageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageRequest) Reset()         { *m = ImageRequest{} }
func (m *ImageRequest) String() string { return proto.CompactTextString(m) }
func (*ImageRequest) ProtoMessage()    {}
func (*ImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4cbcb06a3b852c3, []int{1}
}

func (m *ImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageRequest.Unmarshal(m, b)
}
func (m *ImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageRequest.Marshal(b, m, deterministic)
}
func (m *ImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageRequest.Merge(m, src)
}
func (m *ImageRequest) XXX_Size() int {
	return xxx_messageInfo_ImageRequest.Size(m)
}
func (m *ImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImageRequest proto.InternalMessageInfo

func (m *ImageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ImageResponse struct {
	Header               *ImageHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Imagebuffer          []byte       `protobuf:"bytes,2,opt,name=imagebuffer,proto3" json:"imagebuffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ImageResponse) Reset()         { *m = ImageResponse{} }
func (m *ImageResponse) String() string { return proto.CompactTextString(m) }
func (*ImageResponse) ProtoMessage()    {}
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4cbcb06a3b852c3, []int{2}
}

func (m *ImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageResponse.Unmarshal(m, b)
}
func (m *ImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageResponse.Marshal(b, m, deterministic)
}
func (m *ImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageResponse.Merge(m, src)
}
func (m *ImageResponse) XXX_Size() int {
	return xxx_messageInfo_ImageResponse.Size(m)
}
func (m *ImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageResponse proto.InternalMessageInfo

func (m *ImageResponse) GetHeader() *ImageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ImageResponse) GetImagebuffer() []byte {
	if m != nil {
		return m.Imagebuffer
	}
	return nil
}

func init() {
	proto.RegisterType((*ImageHeader)(nil), "example.ImageHeader")
	proto.RegisterType((*ImageRequest)(nil), "example.ImageRequest")
	proto.RegisterType((*ImageResponse)(nil), "example.ImageResponse")
}

func init() { proto.RegisterFile("get_image.proto", fileDescriptor_f4cbcb06a3b852c3) }

var fileDescriptor_f4cbcb06a3b852c3 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x14, 0x84, 0xa9, 0xee, 0xae, 0xfa, 0xba, 0x22, 0x86, 0x22, 0xb9, 0x08, 0x25, 0xa7, 0x1e, 0x24,
	0x07, 0xfd, 0x07, 0xbd, 0xa8, 0xd7, 0xfc, 0x81, 0x92, 0xda, 0xd7, 0x36, 0x62, 0x93, 0xda, 0xa4,
	0x28, 0xfe, 0x7a, 0xc9, 0x33, 0x42, 0x0f, 0x7b, 0x7b, 0x33, 0x0c, 0x33, 0x5f, 0x02, 0x37, 0x03,
	0x86, 0xc6, 0x4c, 0x7a, 0x40, 0x39, 0x2f, 0x2e, 0x38, 0x76, 0x81, 0xdf, 0x7a, 0x9a, 0x3f, 0x50,
	0xbc, 0x41, 0xfe, 0x1a, 0xfd, 0x17, 0xd4, 0x1d, 0x2e, 0x8c, 0xc1, 0xce, 0xea, 0x09, 0x79, 0x56,
	0x66, 0xd5, 0x95, 0xa2, 0x9b, 0x15, 0xb0, 0xff, 0x32, 0x5d, 0x18, 0xf9, 0x59, 0x99, 0x55, 0x7b,
	0xf5, 0x27, 0xd8, 0x1d, 0x1c, 0x46, 0x34, 0xc3, 0x18, 0xf8, 0x39, 0xd9, 0x49, 0xc5, 0x06, 0x6f,
	0x7e, 0x90, 0xef, 0xc8, 0xa5, 0x5b, 0x08, 0x38, 0xd2, 0x88, 0xc2, 0xcf, 0x15, 0x7d, 0x38, 0xb5,
	0x22, 0x1a, 0xb8, 0x4e, 0x19, 0x3f, 0x3b, 0xeb, 0x91, 0x3d, 0xc4, 0x81, 0x08, 0x45, 0xb1, 0xfc,
	0xb1, 0x90, 0x89, 0x59, 0x6e, 0x80, 0x55, 0xca, 0xb0, 0x12, 0x72, 0x7a, 0x5f, 0xbb, 0xf6, 0x3d,
	0x2e, 0x84, 0x7a, 0x54, 0x5b, 0xab, 0xbe, 0x87, 0x5b, 0x8b, 0x41, 0x9a, 0x77, 0xe7, 0xec, 0x7f,
	0x55, 0x7d, 0xf9, 0x8c, 0x81, 0xea, 0xda, 0x03, 0x7d, 0xcc, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x32, 0xaf, 0x99, 0x2b, 0x01, 0x00, 0x00,
}
