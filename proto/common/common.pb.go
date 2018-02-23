// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto

It has these top-level messages:
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Code int32

const (
	Code_CodeOk               Code = 0
	Code_CodeErr              Code = 1
	Code_CodeErrBadRequest    Code = 2
	Code_CodeErrNotFound      Code = 3
	Code_CodeErrAlreadyExists Code = 4
	Code_CodeErrForbidden     Code = 5
	Code_CodeErrsUnauthorized Code = 6
)

var Code_name = map[int32]string{
	0: "CodeOk",
	1: "CodeErr",
	2: "CodeErrBadRequest",
	3: "CodeErrNotFound",
	4: "CodeErrAlreadyExists",
	5: "CodeErrForbidden",
	6: "CodeErrsUnauthorized",
}
var Code_value = map[string]int32{
	"CodeOk":               0,
	"CodeErr":              1,
	"CodeErrBadRequest":    2,
	"CodeErrNotFound":      3,
	"CodeErrAlreadyExists": 4,
	"CodeErrForbidden":     5,
	"CodeErrsUnauthorized": 6,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}
func (Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("common.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x3d, 0x4e, 0x03, 0x31,
	0x10, 0x85, 0x09, 0x04, 0x23, 0x0d, 0x05, 0x83, 0x13, 0x24, 0x8e, 0x80, 0x52, 0xc4, 0x05, 0x27,
	0x20, 0x28, 0x29, 0x41, 0x42, 0xa2, 0xa1, 0xb3, 0xd7, 0x23, 0xd6, 0xda, 0x5d, 0x0f, 0x8c, 0x6d,
	0xf1, 0x73, 0x0f, 0xee, 0x8b, 0x76, 0xd7, 0x45, 0xaa, 0x37, 0xdf, 0x9b, 0xd7, 0x7c, 0xb0, 0x6a,
	0x78, 0x18, 0x38, 0x9a, 0x39, 0xb6, 0x1f, 0xc2, 0x99, 0xb5, 0x9a, 0x69, 0xf3, 0xb7, 0x80, 0xe5,
	0x23, 0x7b, 0xd2, 0x00, 0x6a, 0xcc, 0xe7, 0x0e, 0x4f, 0xf4, 0x25, 0x5c, 0x8c, 0xf7, 0x5e, 0x04,
	0x17, 0xfa, 0x06, 0xae, 0x2b, 0xec, 0xac, 0x7f, 0xa1, 0xcf, 0x42, 0x29, 0xe3, 0xa9, 0x5e, 0xc1,
	0x55, 0xad, 0x9f, 0x38, 0x1f, 0xb8, 0x44, 0x8f, 0x67, 0xfa, 0x16, 0xd6, 0xb5, 0x7c, 0xe8, 0x85,
	0xac, 0xff, 0xd9, 0x7f, 0x87, 0x94, 0x13, 0x2e, 0xf5, 0x1a, 0xb0, 0x7e, 0x0e, 0x2c, 0x2e, 0x78,
	0x4f, 0x11, 0xcf, 0x8f, 0xf6, 0xe9, 0x35, 0xda, 0x92, 0x5b, 0x96, 0xf0, 0x4b, 0x1e, 0xd5, 0x6e,
	0xf3, 0x76, 0xf7, 0x1e, 0x72, 0x5b, 0xdc, 0xb6, 0xe1, 0xc1, 0x24, 0xee, 0xad, 0x7c, 0xd9, 0xbe,
	0x23, 0x31, 0xce, 0x26, 0x32, 0x93, 0x44, 0x35, 0x72, 0x6a, 0xa2, 0xfb, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x86, 0xea, 0xda, 0xe9, 0x00, 0x00, 0x00,
}