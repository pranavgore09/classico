// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	Text
*/
package echo

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
	proto.RegisterType((*Text)(nil), "echo.Text")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Echo_serviceDesc.Streams[0], c.cc, "/echo.echo/echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoEchoClient{stream}
	return x, nil
}

type Echo_EchoClient interface {
	Send(*Text) error
	Recv() (*Text, error)
	grpc.ClientStream
}

type echoEchoClient struct {
	grpc.ClientStream
}

func (x *echoEchoClient) Send(m *Text) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoEchoClient) Recv() (*Text, error) {
	m := new(Text)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Echo service

type EchoServer interface {
	Echo(Echo_EchoServer) error
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).Echo(&echoEchoServer{stream})
}

type Echo_EchoServer interface {
	Send(*Text) error
	Recv() (*Text, error)
	grpc.ServerStream
}

type echoEchoServer struct {
	grpc.ServerStream
}

func (x *echoEchoServer) Send(m *Text) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoEchoServer) Recv() (*Text, error) {
	m := new(Text)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.echo",
	HandlerType: (*EchoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "echo",
			Handler:       _Echo_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xa4, 0xb8, 0x58, 0x42, 0x52,
	0x2b, 0x4a, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x23, 0x1d, 0x2e, 0xb0, 0x1a, 0x21, 0x15, 0x28, 0xcd, 0xa5, 0x07, 0xd6, 0x0e,
	0x52, 0x2f, 0x85, 0xc4, 0x56, 0x62, 0xd0, 0x60, 0x34, 0x60, 0x4c, 0x62, 0x03, 0x1b, 0x6b, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xb4, 0xa1, 0x25, 0x64, 0x00, 0x00, 0x00,
}