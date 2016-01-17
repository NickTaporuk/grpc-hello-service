// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	Request
	Response
	LoginRequest
	LoginResponse
	User
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// The request message containing the user's name.
type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greeting.
type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The login request message containing the username and password.
type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// The login response message containing the JWT token.
type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type User struct {
	Username     string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	PasswordHash string `protobuf:"bytes,2,opt,name=passwordHash" json:"passwordHash,omitempty"`
	IsAdmin      bool   `protobuf:"varint,3,opt,name=isAdmin" json:"isAdmin,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Request)(nil), "hello.Request")
	proto.RegisterType((*Response)(nil), "hello.Response")
	proto.RegisterType((*LoginRequest)(nil), "hello.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "hello.LoginResponse")
	proto.RegisterType((*User)(nil), "hello.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Hello service

type HelloClient interface {
	// Sends a greeting.
	Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/hello.Hello/Say", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	// Sends a greeting.
	Say(context.Context, *Request) (*Response, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HelloServer).Say(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Hello_Say_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Auth service

type AuthClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/hello.Auth/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthServer).Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x40, 0xa9, 0x4d, 0x6c, 0x1c, 0xab, 0xc2, 0xd8, 0x43, 0x08, 0x08, 0xb2, 0xa8, 0x78, 0xaa,
	0x10, 0x6f, 0xde, 0x7a, 0x91, 0x1e, 0x3c, 0xad, 0x78, 0x77, 0xa5, 0x43, 0x13, 0x6c, 0x76, 0x63,
	0x66, 0x83, 0xf8, 0xef, 0x6d, 0xf6, 0xa3, 0xb4, 0x97, 0xde, 0xf6, 0xcd, 0xcc, 0xbe, 0xd9, 0x99,
	0x85, 0xf3, 0x8a, 0x36, 0x1b, 0x33, 0x6f, 0x3b, 0x63, 0x0d, 0xa6, 0x0e, 0xc4, 0x0d, 0x4c, 0x24,
	0xfd, 0xf4, 0xc4, 0x16, 0x11, 0x12, 0xad, 0x1a, 0xca, 0x47, 0xb7, 0xa3, 0xc7, 0x33, 0xe9, 0xce,
	0xe2, 0x0e, 0x32, 0x49, 0xdc, 0x1a, 0xcd, 0x84, 0x39, 0x4c, 0x1a, 0x62, 0x56, 0xeb, 0x58, 0x12,
	0x51, 0xbc, 0xc2, 0xf4, 0xcd, 0xac, 0x6b, 0x1d, 0x4d, 0x05, 0x64, 0x3d, 0x53, 0xb7, 0x67, 0xdb,
	0xf1, 0x90, 0x6b, 0x15, 0xf3, 0xaf, 0xe9, 0x56, 0xf9, 0x89, 0xcf, 0x45, 0x16, 0xf7, 0x70, 0x11,
	0x3c, 0xa1, 0xe5, 0x0c, 0x52, 0x6b, 0xbe, 0x49, 0x07, 0x8b, 0x07, 0xf1, 0x09, 0xc9, 0xc7, 0x56,
	0x77, 0xb4, 0x8d, 0x80, 0x69, 0xd4, 0x2e, 0x15, 0x57, 0xa1, 0xd5, 0x41, 0x6c, 0x18, 0xa8, 0xe6,
	0xc5, 0xaa, 0xa9, 0x75, 0x3e, 0xde, 0xa6, 0x33, 0x19, 0xb1, 0x7c, 0x82, 0x74, 0x39, 0xac, 0x07,
	0x1f, 0x60, 0xfc, 0xae, 0xfe, 0xf0, 0x72, 0xee, 0x57, 0x17, 0x06, 0x2c, 0xae, 0x76, 0xec, 0x1f,
	0x5a, 0xbe, 0x40, 0xb2, 0xe8, 0x6d, 0x85, 0x25, 0xa4, 0x6e, 0x02, 0xbc, 0x0e, 0x15, 0xfb, 0x7b,
	0x29, 0x66, 0x87, 0x41, 0x7f, 0xf7, 0xeb, 0xd4, 0x7d, 0xc8, 0xf3, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe7, 0xd1, 0xb7, 0x4f, 0x9f, 0x01, 0x00, 0x00,
}
