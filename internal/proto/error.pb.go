// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Error struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto1.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto1.RegisterType((*Error)(nil), "proto.Error")
}

func init() { proto1.RegisterFile("error.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 48 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xec, 0x5c, 0xac, 0xae,
	0x20, 0x51, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x45, 0xcf, 0xf5, 0x1d, 0x00, 0x00, 0x00,
}
