// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provider_config.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProviderConfig struct {
}

func (m *ProviderConfig) Reset()                    { *m = ProviderConfig{} }
func (m *ProviderConfig) String() string            { return proto1.CompactTextString(m) }
func (*ProviderConfig) ProtoMessage()               {}
func (*ProviderConfig) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func init() {
	proto1.RegisterType((*ProviderConfig)(nil), "proto.ProviderConfig")
}

func init() { proto1.RegisterFile("provider_config.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 65 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xcb, 0x4c, 0x49, 0x2d, 0x8a, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x02, 0x5c, 0x7c, 0x01, 0x50, 0x79, 0x67, 0xb0, 0x74, 0x12,
	0x1b, 0x58, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x05, 0x13, 0x2e, 0xff, 0x38, 0x00, 0x00,
	0x00,
}