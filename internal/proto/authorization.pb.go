// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authorization.proto

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

// Client API for AuthorizationProvider service

type AuthorizationProviderClient interface {
}

type authorizationProviderClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationProviderClient(cc *grpc.ClientConn) AuthorizationProviderClient {
	return &authorizationProviderClient{cc}
}

// Server API for AuthorizationProvider service

type AuthorizationProviderServer interface {
}

func RegisterAuthorizationProviderServer(s *grpc.Server, srv AuthorizationProviderServer) {
	s.RegisterService(&_AuthorizationProvider_serviceDesc, srv)
}

var _AuthorizationProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthorizationProvider",
	HandlerType: (*AuthorizationProviderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "authorization.proto",
}

func init() { proto1.RegisterFile("authorization.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 69 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x53, 0x46, 0xe2, 0x5c, 0xa2, 0x8e, 0xc8, 0xb2, 0x01, 0x45, 0xf9, 0x65, 0x99, 0x29,
	0xa9, 0x45, 0x49, 0x6c, 0x60, 0x79, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xe2, 0xfe,
	0x35, 0x3d, 0x00, 0x00, 0x00,
}
