// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packet.proto

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

type Packet1 struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet1) Reset()         { *m = Packet1{} }
func (m *Packet1) String() string { return proto.CompactTextString(m) }
func (*Packet1) ProtoMessage()    {}
func (*Packet1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
}

func (m *Packet1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet1.Unmarshal(m, b)
}
func (m *Packet1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet1.Marshal(b, m, deterministic)
}
func (m *Packet1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet1.Merge(m, src)
}
func (m *Packet1) XXX_Size() int {
	return xxx_messageInfo_Packet1.Size(m)
}
func (m *Packet1) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet1.DiscardUnknown(m)
}

var xxx_messageInfo_Packet1 proto.InternalMessageInfo

func (m *Packet1) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Packet2 struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet2) Reset()         { *m = Packet2{} }
func (m *Packet2) String() string { return proto.CompactTextString(m) }
func (*Packet2) ProtoMessage()    {}
func (*Packet2) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{1}
}

func (m *Packet2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet2.Unmarshal(m, b)
}
func (m *Packet2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet2.Marshal(b, m, deterministic)
}
func (m *Packet2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet2.Merge(m, src)
}
func (m *Packet2) XXX_Size() int {
	return xxx_messageInfo_Packet2.Size(m)
}
func (m *Packet2) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet2.DiscardUnknown(m)
}

var xxx_messageInfo_Packet2 proto.InternalMessageInfo

func (m *Packet2) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

type Packet3 struct {
	BoolValue            bool     `protobuf:"varint,1,opt,name=boolValue,proto3" json:"boolValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet3) Reset()         { *m = Packet3{} }
func (m *Packet3) String() string { return proto.CompactTextString(m) }
func (*Packet3) ProtoMessage()    {}
func (*Packet3) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{2}
}

func (m *Packet3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet3.Unmarshal(m, b)
}
func (m *Packet3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet3.Marshal(b, m, deterministic)
}
func (m *Packet3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet3.Merge(m, src)
}
func (m *Packet3) XXX_Size() int {
	return xxx_messageInfo_Packet3.Size(m)
}
func (m *Packet3) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet3.DiscardUnknown(m)
}

var xxx_messageInfo_Packet3 proto.InternalMessageInfo

func (m *Packet3) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

type Packet4 struct {
	DoubleValue          float64  `protobuf:"fixed64,1,opt,name=doubleValue,proto3" json:"doubleValue,omitempty"`
	FloatValue           float32  `protobuf:"fixed32,2,opt,name=floatValue,proto3" json:"floatValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet4) Reset()         { *m = Packet4{} }
func (m *Packet4) String() string { return proto.CompactTextString(m) }
func (*Packet4) ProtoMessage()    {}
func (*Packet4) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{3}
}

func (m *Packet4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet4.Unmarshal(m, b)
}
func (m *Packet4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet4.Marshal(b, m, deterministic)
}
func (m *Packet4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet4.Merge(m, src)
}
func (m *Packet4) XXX_Size() int {
	return xxx_messageInfo_Packet4.Size(m)
}
func (m *Packet4) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet4.DiscardUnknown(m)
}

var xxx_messageInfo_Packet4 proto.InternalMessageInfo

func (m *Packet4) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *Packet4) GetFloatValue() float32 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

type ArrayMessage struct {
	StrArr               []string `protobuf:"bytes,1,rep,name=strArr,proto3" json:"strArr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayMessage) Reset()         { *m = ArrayMessage{} }
func (m *ArrayMessage) String() string { return proto.CompactTextString(m) }
func (*ArrayMessage) ProtoMessage()    {}
func (*ArrayMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{4}
}

func (m *ArrayMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayMessage.Unmarshal(m, b)
}
func (m *ArrayMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayMessage.Marshal(b, m, deterministic)
}
func (m *ArrayMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayMessage.Merge(m, src)
}
func (m *ArrayMessage) XXX_Size() int {
	return xxx_messageInfo_ArrayMessage.Size(m)
}
func (m *ArrayMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayMessage proto.InternalMessageInfo

func (m *ArrayMessage) GetStrArr() []string {
	if m != nil {
		return m.StrArr
	}
	return nil
}

type Camera struct {
	Serial               string   `protobuf:"bytes,1,opt,name=serial,proto3" json:"serial,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PublicIP             string   `protobuf:"bytes,3,opt,name=publicIP,proto3" json:"publicIP,omitempty"`
	PublicPort           string   `protobuf:"bytes,4,opt,name=publicPort,proto3" json:"publicPort,omitempty"`
	PrivateIP            string   `protobuf:"bytes,5,opt,name=privateIP,proto3" json:"privateIP,omitempty"`
	PrivatePort          string   `protobuf:"bytes,6,opt,name=privatePort,proto3" json:"privatePort,omitempty"`
	Ping                 uint32   `protobuf:"varint,7,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Camera) Reset()         { *m = Camera{} }
func (m *Camera) String() string { return proto.CompactTextString(m) }
func (*Camera) ProtoMessage()    {}
func (*Camera) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{5}
}

func (m *Camera) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Camera.Unmarshal(m, b)
}
func (m *Camera) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Camera.Marshal(b, m, deterministic)
}
func (m *Camera) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Camera.Merge(m, src)
}
func (m *Camera) XXX_Size() int {
	return xxx_messageInfo_Camera.Size(m)
}
func (m *Camera) XXX_DiscardUnknown() {
	xxx_messageInfo_Camera.DiscardUnknown(m)
}

var xxx_messageInfo_Camera proto.InternalMessageInfo

func (m *Camera) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *Camera) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Camera) GetPublicIP() string {
	if m != nil {
		return m.PublicIP
	}
	return ""
}

func (m *Camera) GetPublicPort() string {
	if m != nil {
		return m.PublicPort
	}
	return ""
}

func (m *Camera) GetPrivateIP() string {
	if m != nil {
		return m.PrivateIP
	}
	return ""
}

func (m *Camera) GetPrivatePort() string {
	if m != nil {
		return m.PrivatePort
	}
	return ""
}

func (m *Camera) GetPing() uint32 {
	if m != nil {
		return m.Ping
	}
	return 0
}

type CameraListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CameraListRequest) Reset()         { *m = CameraListRequest{} }
func (m *CameraListRequest) String() string { return proto.CompactTextString(m) }
func (*CameraListRequest) ProtoMessage()    {}
func (*CameraListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{6}
}

func (m *CameraListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CameraListRequest.Unmarshal(m, b)
}
func (m *CameraListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CameraListRequest.Marshal(b, m, deterministic)
}
func (m *CameraListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CameraListRequest.Merge(m, src)
}
func (m *CameraListRequest) XXX_Size() int {
	return xxx_messageInfo_CameraListRequest.Size(m)
}
func (m *CameraListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CameraListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CameraListRequest proto.InternalMessageInfo

type CameraListResponse struct {
	CameraList           []*Camera `protobuf:"bytes,1,rep,name=cameraList,proto3" json:"cameraList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CameraListResponse) Reset()         { *m = CameraListResponse{} }
func (m *CameraListResponse) String() string { return proto.CompactTextString(m) }
func (*CameraListResponse) ProtoMessage()    {}
func (*CameraListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{7}
}

func (m *CameraListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CameraListResponse.Unmarshal(m, b)
}
func (m *CameraListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CameraListResponse.Marshal(b, m, deterministic)
}
func (m *CameraListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CameraListResponse.Merge(m, src)
}
func (m *CameraListResponse) XXX_Size() int {
	return xxx_messageInfo_CameraListResponse.Size(m)
}
func (m *CameraListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CameraListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CameraListResponse proto.InternalMessageInfo

func (m *CameraListResponse) GetCameraList() []*Camera {
	if m != nil {
		return m.CameraList
	}
	return nil
}

func init() {
	proto.RegisterType((*Packet1)(nil), "example.Packet1")
	proto.RegisterType((*Packet2)(nil), "example.Packet2")
	proto.RegisterType((*Packet3)(nil), "example.Packet3")
	proto.RegisterType((*Packet4)(nil), "example.Packet4")
	proto.RegisterType((*ArrayMessage)(nil), "example.ArrayMessage")
	proto.RegisterType((*Camera)(nil), "example.Camera")
	proto.RegisterType((*CameraListRequest)(nil), "example.CameraListRequest")
	proto.RegisterType((*CameraListResponse)(nil), "example.CameraListResponse")
}

func init() { proto.RegisterFile("packet.proto", fileDescriptor_e9ef1a6541f9f9e7) }

var fileDescriptor_e9ef1a6541f9f9e7 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x5f, 0x4f, 0x83, 0x30,
	0x14, 0xc5, 0xc3, 0xfe, 0xb0, 0x71, 0x37, 0xa3, 0xab, 0x89, 0x21, 0xce, 0x18, 0xec, 0x83, 0xf2,
	0x84, 0x71, 0xf3, 0x0b, 0x6c, 0xc6, 0x87, 0x45, 0x4d, 0x48, 0x1f, 0x7c, 0x2f, 0xf3, 0xba, 0xa0,
	0x40, 0xb1, 0x2d, 0x46, 0xbf, 0x9f, 0x1f, 0xcc, 0xd0, 0xc2, 0xc6, 0xdb, 0x3d, 0xbf, 0x73, 0x1a,
	0x4e, 0x7b, 0x81, 0x69, 0xc9, 0xb7, 0x9f, 0xa8, 0xa3, 0x52, 0x0a, 0x2d, 0xc8, 0x08, 0x7f, 0x78,
	0x5e, 0x66, 0x48, 0xaf, 0x60, 0x14, 0x1b, 0xe3, 0x8e, 0x9c, 0x81, 0x5b, 0x54, 0x79, 0x82, 0xd2,
	0x77, 0x02, 0x27, 0x1c, 0xb2, 0x46, 0xd1, 0x79, 0x1b, 0x59, 0x90, 0x13, 0xe8, 0x2b, 0x6d, 0x7d,
	0x8f, 0xd5, 0x23, 0xbd, 0x69, 0xcd, 0x25, 0xb9, 0x00, 0x2f, 0x11, 0x22, 0x7b, 0xe5, 0x59, 0x85,
	0x26, 0x32, 0x66, 0x07, 0x40, 0x9f, 0xda, 0xe0, 0x3d, 0x09, 0x60, 0xf2, 0x26, 0xaa, 0x24, 0xc3,
	0x43, 0xd4, 0x61, 0x5d, 0x44, 0x2e, 0x01, 0xde, 0x33, 0xc1, 0xb5, 0x0d, 0xf4, 0x02, 0x27, 0xec,
	0xb1, 0x0e, 0xa1, 0xd7, 0x30, 0x5d, 0x49, 0xc9, 0x7f, 0x5f, 0x50, 0x29, 0xbe, 0xc3, 0xba, 0xba,
	0xd2, 0x72, 0x25, 0xeb, 0x6a, 0xfd, 0xd0, 0x63, 0x8d, 0xa2, 0x7f, 0x0e, 0xb8, 0x0f, 0x3c, 0x47,
	0xc9, 0x4d, 0x04, 0x65, 0xca, 0xb3, 0xa6, 0x7d, 0xa3, 0x08, 0x81, 0x41, 0xc1, 0x73, 0xfb, 0x11,
	0x8f, 0x99, 0x99, 0x9c, 0xc3, 0xb8, 0xac, 0x92, 0x2c, 0xdd, 0x6e, 0x62, 0xbf, 0x6f, 0xf8, 0x5e,
	0xd7, 0xd5, 0xec, 0x1c, 0x0b, 0xa9, 0xfd, 0x81, 0x71, 0x3b, 0xa4, 0x7e, 0x85, 0x52, 0xa6, 0xdf,
	0x5c, 0xe3, 0x26, 0xf6, 0x87, 0xc6, 0x3e, 0x80, 0xfa, 0xea, 0x8d, 0x30, 0xc7, 0x5d, 0xe3, 0x77,
	0x51, 0xdd, 0xa7, 0x4c, 0x8b, 0x9d, 0x3f, 0x0a, 0x9c, 0xf0, 0x88, 0x99, 0x99, 0x9e, 0xc2, 0xcc,
	0xde, 0xe2, 0x39, 0x55, 0x9a, 0xe1, 0x57, 0x85, 0x4a, 0xd3, 0x47, 0x20, 0x5d, 0xa8, 0x4a, 0x51,
	0x28, 0x24, 0xb7, 0x00, 0xdb, 0x3d, 0x35, 0xaf, 0x31, 0x59, 0x1c, 0x47, 0xcd, 0xb6, 0x23, 0x7b,
	0x80, 0x75, 0x22, 0xeb, 0x39, 0xcc, 0x0a, 0xd4, 0x51, 0xfa, 0x21, 0x44, 0xd1, 0xe6, 0xd6, 0xae,
	0x5d, 0x55, 0xe2, 0x9a, 0xbf, 0x65, 0xf9, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x94, 0x33, 0x91,
	0x3d, 0x02, 0x00, 0x00,
}