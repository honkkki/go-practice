// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/message.proto

package message

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 包含人名的一个请求消息
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 包含问候语的响应消息
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type NumRequest struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumRequest) Reset()         { *m = NumRequest{} }
func (m *NumRequest) String() string { return proto.CompactTextString(m) }
func (*NumRequest) ProtoMessage()    {}
func (*NumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{2}
}

func (m *NumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumRequest.Unmarshal(m, b)
}
func (m *NumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumRequest.Marshal(b, m, deterministic)
}
func (m *NumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumRequest.Merge(m, src)
}
func (m *NumRequest) XXX_Size() int {
	return xxx_messageInfo_NumRequest.Size(m)
}
func (m *NumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumRequest proto.InternalMessageInfo

func (m *NumRequest) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type NumRes struct {
	Res                  int64    `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumRes) Reset()         { *m = NumRes{} }
func (m *NumRes) String() string { return proto.CompactTextString(m) }
func (*NumRes) ProtoMessage()    {}
func (*NumRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{3}
}

func (m *NumRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumRes.Unmarshal(m, b)
}
func (m *NumRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumRes.Marshal(b, m, deterministic)
}
func (m *NumRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumRes.Merge(m, src)
}
func (m *NumRes) XXX_Size() int {
	return xxx_messageInfo_NumRes.Size(m)
}
func (m *NumRes) XXX_DiscardUnknown() {
	xxx_messageInfo_NumRes.DiscardUnknown(m)
}

var xxx_messageInfo_NumRes proto.InternalMessageInfo

func (m *NumRes) GetRes() int64 {
	if m != nil {
		return m.Res
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "message.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "message.HelloReply")
	proto.RegisterType((*NumRequest)(nil), "message.NumRequest")
	proto.RegisterType((*NumRes)(nil), "message.NumRes")
}

func init() { proto.RegisterFile("pb/message.proto", fileDescriptor_8447775385e7eb85) }

var fileDescriptor_8447775385e7eb85 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72,
	0x95, 0x94, 0xb8, 0x78, 0x3c, 0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x25, 0x35, 0x2e, 0x2e, 0xa8, 0x9a, 0x82, 0x9c, 0x4a, 0x21, 0x09, 0x2e, 0x98, 0x66, 0xa8,
	0x22, 0xb8, 0x59, 0x72, 0x5c, 0x5c, 0x7e, 0xa5, 0xb9, 0x30, 0x93, 0x04, 0xb8, 0x98, 0xf3, 0x4a,
	0x73, 0xc1, 0x6a, 0x98, 0x83, 0x40, 0x4c, 0x25, 0x29, 0x2e, 0x36, 0xb0, 0x7c, 0x31, 0x48, 0xae,
	0x28, 0xb5, 0x18, 0x26, 0x57, 0x94, 0x5a, 0x6c, 0xe4, 0xcc, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a,
	0x92, 0x5a, 0x24, 0x64, 0xc1, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xb6, 0x51, 0x48, 0x54, 0x0f, 0xe6,
	0x6e, 0x64, 0x57, 0x4a, 0x09, 0xa3, 0x0b, 0x17, 0xe4, 0x54, 0x2a, 0x31, 0x18, 0x59, 0x72, 0xb1,
	0x04, 0xe4, 0x94, 0x16, 0x0b, 0x19, 0x72, 0xb1, 0x83, 0x68, 0xff, 0xbc, 0x54, 0x21, 0x84, 0x4a,
	0x84, 0xd3, 0xa4, 0xf8, 0x51, 0x05, 0x8b, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xe1, 0x62, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x81, 0x0f, 0x51, 0x2b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// SayHello 方法
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/message.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// SayHello 方法
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/message.proto",
}

// PlusClient is the client API for Plus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlusClient interface {
	PlusOne(ctx context.Context, in *NumRequest, opts ...grpc.CallOption) (*NumRes, error)
}

type plusClient struct {
	cc *grpc.ClientConn
}

func NewPlusClient(cc *grpc.ClientConn) PlusClient {
	return &plusClient{cc}
}

func (c *plusClient) PlusOne(ctx context.Context, in *NumRequest, opts ...grpc.CallOption) (*NumRes, error) {
	out := new(NumRes)
	err := c.cc.Invoke(ctx, "/message.Plus/PlusOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlusServer is the server API for Plus service.
type PlusServer interface {
	PlusOne(context.Context, *NumRequest) (*NumRes, error)
}

// UnimplementedPlusServer can be embedded to have forward compatible implementations.
type UnimplementedPlusServer struct {
}

func (*UnimplementedPlusServer) PlusOne(ctx context.Context, req *NumRequest) (*NumRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlusOne not implemented")
}

func RegisterPlusServer(s *grpc.Server, srv PlusServer) {
	s.RegisterService(&_Plus_serviceDesc, srv)
}

func _Plus_PlusOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlusServer).PlusOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Plus/PlusOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlusServer).PlusOne(ctx, req.(*NumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.Plus",
	HandlerType: (*PlusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlusOne",
			Handler:    _Plus_PlusOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/message.proto",
}
