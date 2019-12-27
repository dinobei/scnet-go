// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packet_type.proto

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

type PacketType int32

const (
	PacketType_packetType1        PacketType = 0
	PacketType_packetType2        PacketType = 1
	PacketType_packetType3        PacketType = 2
	PacketType_packetType4        PacketType = 3
	PacketType_arrayMessageType   PacketType = 4
	PacketType_imageRequest       PacketType = 5
	PacketType_imageResponse      PacketType = 6
	PacketType_cameraListRequest  PacketType = 7
	PacketType_cameraListResponse PacketType = 8
)

var PacketType_name = map[int32]string{
	0: "packetType1",
	1: "packetType2",
	2: "packetType3",
	3: "packetType4",
	4: "arrayMessageType",
	5: "imageRequest",
	6: "imageResponse",
	7: "cameraListRequest",
	8: "cameraListResponse",
}

var PacketType_value = map[string]int32{
	"packetType1":        0,
	"packetType2":        1,
	"packetType3":        2,
	"packetType4":        3,
	"arrayMessageType":   4,
	"imageRequest":       5,
	"imageResponse":      6,
	"cameraListRequest":  7,
	"cameraListResponse": 8,
}

func (x PacketType) String() string {
	return proto.EnumName(PacketType_name, int32(x))
}

func (PacketType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d20de2ff11256102, []int{0}
}

func init() {
	proto.RegisterEnum("example.PacketType", PacketType_name, PacketType_value)
}

func init() { proto.RegisterFile("packet_type.proto", fileDescriptor_d20de2ff11256102) }

var fileDescriptor_d20de2ff11256102 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcf, 0xcf, 0x6e, 0x82, 0x40,
	0x10, 0xc7, 0xf1, 0xd2, 0x3f, 0xd0, 0x0c, 0x6d, 0x3a, 0x3b, 0x69, 0x7b, 0xf6, 0xec, 0x81, 0x44,
	0xf0, 0x09, 0x38, 0x6b, 0x62, 0x8c, 0x77, 0x33, 0x92, 0x09, 0x41, 0x85, 0x5d, 0x77, 0xd7, 0x44,
	0x5e, 0xcd, 0xa7, 0x33, 0x12, 0xfc, 0xc3, 0x71, 0x3e, 0xf3, 0xbb, 0x7c, 0x41, 0x19, 0x2e, 0x76,
	0xe2, 0xd7, 0xbe, 0x35, 0x92, 0x18, 0xab, 0xbd, 0xa6, 0x48, 0x4e, 0x5c, 0x9b, 0xbd, 0x8c, 0xcf,
	0x01, 0xc0, 0xa2, 0x7b, 0xaf, 0x5a, 0x23, 0xf4, 0x03, 0xb1, 0xb9, 0x5f, 0x13, 0x7c, 0x19, 0x42,
	0x8a, 0xc1, 0x10, 0x32, 0x7c, 0x1d, 0xc2, 0x14, 0xdf, 0xe8, 0x17, 0x90, 0xad, 0xe5, 0x76, 0x2e,
	0xce, 0x71, 0x29, 0x57, 0xc6, 0x77, 0x42, 0xf8, 0xaa, 0x6a, 0x2e, 0x65, 0x29, 0x87, 0xa3, 0x38,
	0x8f, 0x1f, 0xa4, 0xe0, 0xbb, 0x17, 0x67, 0x74, 0xe3, 0x04, 0x43, 0xfa, 0x03, 0x55, 0x70, 0x2d,
	0x96, 0x67, 0x95, 0xf3, 0xb7, 0x65, 0x44, 0xff, 0x40, 0xcf, 0xdc, 0xcf, 0x3f, 0xf3, 0x11, 0xa8,
	0x46, 0x7c, 0x52, 0x6d, 0xb5, 0x6e, 0x92, 0xbe, 0x28, 0x8f, 0x1f, 0x39, 0xe9, 0x26, 0xec, 0x72,
	0xb3, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x48, 0xe5, 0x2e, 0x03, 0x01, 0x00, 0x00,
}
