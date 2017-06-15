// Code generated by protoc-gen-go. DO NOT EDIT.
// source: say.proto

/*
Package say is a generated protocol buffer package.

It is generated from these files:
	say.proto

It has these top-level messages:
	Text
*/
package say

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Text struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Text) Reset()                    { *m = Text{} }
func (m *Text) String() string            { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()               {}
func (*Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Text)(nil), "say.Text")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SayHello service

type SayHelloClient interface {
	Hello(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error)
}

type sayHelloClient struct {
	cc *grpc.ClientConn
}

func NewSayHelloClient(cc *grpc.ClientConn) SayHelloClient {
	return &sayHelloClient{cc}
}

func (c *sayHelloClient) Hello(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := grpc.Invoke(ctx, "/say.SayHello/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SayHello service

type SayHelloServer interface {
	Hello(context.Context, *Text) (*Text, error)
}

func RegisterSayHelloServer(s *grpc.Server, srv SayHelloServer) {
	s.RegisterService(&_SayHello_serviceDesc, srv)
}

func _SayHello_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayHelloServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/say.SayHello/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayHelloServer).Hello(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

var _SayHello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "say.SayHello",
	HandlerType: (*SayHelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _SayHello_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "say.proto",
}

func init() { proto.RegisterFile("say.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4e, 0xac, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4e, 0xac, 0x54, 0x92, 0xe2, 0x62, 0x09, 0x49,
	0xad, 0x28, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x8d, 0xb4, 0xb9, 0x38, 0x82, 0x13, 0x2b, 0x3d, 0x52, 0x73, 0x72, 0xf2, 0x85,
	0xe4, 0xb9, 0x58, 0x21, 0x0c, 0x4e, 0x3d, 0x90, 0x09, 0x20, 0x3d, 0x52, 0x08, 0xa6, 0x12, 0x43,
	0x12, 0x1b, 0xd8, 0x50, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x4e, 0xcb, 0x69, 0x61,
	0x00, 0x00, 0x00,
}
