// Code generated by protoc-gen-go. DO NOT EDIT.
// source: traffic_control.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TrafficControlProvider service

type TrafficControlProviderClient interface {
}

type trafficControlProviderClient struct {
	cc *grpc.ClientConn
}

func NewTrafficControlProviderClient(cc *grpc.ClientConn) TrafficControlProviderClient {
	return &trafficControlProviderClient{cc}
}

// Server API for TrafficControlProvider service

type TrafficControlProviderServer interface {
}

func RegisterTrafficControlProviderServer(s *grpc.Server, srv TrafficControlProviderServer) {
	s.RegisterService(&_TrafficControlProvider_serviceDesc, srv)
}

var _TrafficControlProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TrafficControlProvider",
	HandlerType: (*TrafficControlProviderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "traffic_control.proto",
}

func init() { proto1.RegisterFile("traffic_control.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x29, 0x4a, 0x4c,
	0x4b, 0xcb, 0x4c, 0x8e, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x53, 0x46, 0x12, 0x5c, 0x62, 0x21, 0x10, 0x79, 0x67, 0x88, 0x74, 0x40,
	0x51, 0x7e, 0x59, 0x66, 0x4a, 0x6a, 0x51, 0x12, 0x1b, 0x58, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0x6d, 0x79, 0x36, 0x40, 0x00, 0x00, 0x00,
}
