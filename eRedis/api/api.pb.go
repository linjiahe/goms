// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Reply struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Count                int64    `protobuf:"zigzag64,2,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Reply) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type UserT struct {
	Uid                  int64    `protobuf:"zigzag64,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Sex                  int64    `protobuf:"zigzag64,3,opt,name=Sex,proto3" json:"Sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserT) Reset()         { *m = UserT{} }
func (m *UserT) String() string { return proto.CompactTextString(m) }
func (*UserT) ProtoMessage()    {}
func (*UserT) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *UserT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserT.Unmarshal(m, b)
}
func (m *UserT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserT.Marshal(b, m, deterministic)
}
func (m *UserT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserT.Merge(m, src)
}
func (m *UserT) XXX_Size() int {
	return xxx_messageInfo_UserT.Size(m)
}
func (m *UserT) XXX_DiscardUnknown() {
	xxx_messageInfo_UserT.DiscardUnknown(m)
}

var xxx_messageInfo_UserT proto.InternalMessageInfo

func (m *UserT) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserT) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserT) GetSex() int64 {
	if m != nil {
		return m.Sex
	}
	return 0
}

type UidT struct {
	Val                  int64    `protobuf:"zigzag64,1,opt,name=Val,proto3" json:"Val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UidT) Reset()         { *m = UidT{} }
func (m *UidT) String() string { return proto.CompactTextString(m) }
func (*UidT) ProtoMessage()    {}
func (*UidT) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *UidT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UidT.Unmarshal(m, b)
}
func (m *UidT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UidT.Marshal(b, m, deterministic)
}
func (m *UidT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UidT.Merge(m, src)
}
func (m *UidT) XXX_Size() int {
	return xxx_messageInfo_UidT.Size(m)
}
func (m *UidT) XXX_DiscardUnknown() {
	xxx_messageInfo_UidT.DiscardUnknown(m)
}

var xxx_messageInfo_UidT proto.InternalMessageInfo

func (m *UidT) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "service.goms.Request")
	proto.RegisterType((*Reply)(nil), "service.goms.Reply")
	proto.RegisterType((*Empty)(nil), "service.goms.Empty")
	proto.RegisterType((*Null)(nil), "service.goms.Null")
	proto.RegisterType((*UserT)(nil), "service.goms.UserT")
	proto.RegisterType((*UidT)(nil), "service.goms.UidT")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdb, 0x26, 0x69, 0xde, 0x0c, 0xef, 0x41, 0x46, 0x85, 0x90, 0x53, 0x59, 0x2f, 0x3d,
	0xe5, 0xe0, 0x1f, 0xea, 0x4d, 0xb0, 0x7a, 0xb4, 0xc8, 0xb6, 0xe9, 0xc1, 0xdb, 0x6a, 0x86, 0xb0,
	0xb0, 0x69, 0xd6, 0x6c, 0x22, 0xcd, 0xa7, 0xf1, 0xab, 0x4a, 0x36, 0x29, 0x48, 0xa8, 0xe2, 0x6d,
	0xe6, 0xe1, 0x79, 0x7e, 0x33, 0xb3, 0x2c, 0x04, 0x42, 0xcb, 0x58, 0x97, 0x45, 0x55, 0xe0, 0x7f,
	0x43, 0xe5, 0x87, 0x7c, 0xa3, 0x38, 0x2b, 0x72, 0xc3, 0x2e, 0xc0, 0xe7, 0xf4, 0x5e, 0x93, 0xa9,
	0x30, 0x04, 0xff, 0x89, 0x8c, 0x11, 0x19, 0x85, 0xe3, 0xd9, 0x78, 0x1e, 0xf0, 0x43, 0xcb, 0x16,
	0xe0, 0x71, 0xd2, 0xaa, 0xf9, 0xd9, 0x82, 0x67, 0xe0, 0x2d, 0x8b, 0x7a, 0x57, 0x85, 0x93, 0xd9,
	0x78, 0x8e, 0xbc, 0x6b, 0x98, 0x0f, 0xde, 0x63, 0xae, 0xab, 0x86, 0x4d, 0xc1, 0x5d, 0xd5, 0x4a,
	0xb1, 0x3b, 0xf0, 0x12, 0x43, 0xe5, 0x06, 0x4f, 0xc0, 0x49, 0x64, 0x6a, 0x29, 0xc8, 0xdb, 0x12,
	0x11, 0xdc, 0x95, 0xc8, 0xc9, 0x02, 0x02, 0x6e, 0xeb, 0xd6, 0xb5, 0xa6, 0x7d, 0xe8, 0x74, 0xae,
	0x35, 0xed, 0x59, 0x08, 0x6e, 0x22, 0x53, 0x9b, 0xdf, 0x0a, 0x75, 0xc8, 0x6f, 0x85, 0xba, 0xfc,
	0x9c, 0x80, 0xdb, 0xb2, 0xf1, 0x1a, 0xdc, 0x67, 0xb9, 0xcb, 0xf0, 0x3c, 0xfe, 0x7e, 0x69, 0xdc,
	0x9f, 0x19, 0x9d, 0x0e, 0x65, 0xad, 0x1a, 0x36, 0xc2, 0x05, 0xc0, 0xb2, 0x24, 0x51, 0x91, 0x65,
	0x0c, 0x4c, 0x76, 0xe7, 0x08, 0x07, 0xa2, 0x4c, 0x37, 0x6c, 0x84, 0xb7, 0x00, 0x89, 0x4e, 0x7f,
	0x0d, 0x0e, 0xc4, 0xee, 0x49, 0x46, 0x78, 0x03, 0xff, 0x38, 0x89, 0xd4, 0xe6, 0x8e, 0xb0, 0xa3,
	0x63, 0xac, 0x6e, 0xd3, 0x07, 0x52, 0xd4, 0x0f, 0xfc, 0x43, 0xb0, 0x9f, 0x77, 0xef, 0xbd, 0x38,
	0x42, 0xcb, 0xd7, 0xa9, 0xfd, 0x07, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x60, 0x3e,
	0xde, 0x14, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	CreateUser(ctx context.Context, in *UserT, opts ...grpc.CallOption) (*UidT, error)
	UpdateUser(ctx context.Context, in *UserT, opts ...grpc.CallOption) (*Empty, error)
	ReadUser(ctx context.Context, in *UidT, opts ...grpc.CallOption) (*UserT, error)
	DeleteUser(ctx context.Context, in *UidT, opts ...grpc.CallOption) (*Empty, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/service.goms.User/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *UserT, opts ...grpc.CallOption) (*UidT, error) {
	out := new(UidT)
	err := c.cc.Invoke(ctx, "/service.goms.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UserT, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.goms.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ReadUser(ctx context.Context, in *UidT, opts ...grpc.CallOption) (*UserT, error) {
	out := new(UserT)
	err := c.cc.Invoke(ctx, "/service.goms.User/ReadUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *UidT, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.goms.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Ping(context.Context, *Request) (*Reply, error)
	CreateUser(context.Context, *UserT) (*UidT, error)
	UpdateUser(context.Context, *UserT) (*Empty, error)
	ReadUser(context.Context, *UidT) (*UserT, error)
	DeleteUser(context.Context, *UidT) (*Empty, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Ping(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *UserT) (*UidT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) UpdateUser(ctx context.Context, req *UserT) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServer) ReadUser(ctx context.Context, req *UidT) (*UserT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUser not implemented")
}
func (*UnimplementedUserServer) DeleteUser(ctx context.Context, req *UidT) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.goms.User/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.goms.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*UserT))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.goms.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UserT))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ReadUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ReadUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.goms.User/ReadUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ReadUser(ctx, req.(*UidT))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.goms.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*UidT))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.goms.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _User_Ping_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "ReadUser",
			Handler:    _User_ReadUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}