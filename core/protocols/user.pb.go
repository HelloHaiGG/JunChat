// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package core

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type LoginParams struct {
	UName                string   `protobuf:"bytes,1,opt,name=UName,proto3" json:"UName,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginParams) Reset()         { *m = LoginParams{} }
func (m *LoginParams) String() string { return proto.CompactTextString(m) }
func (*LoginParams) ProtoMessage()    {}
func (*LoginParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}
func (m *LoginParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginParams.Merge(m, src)
}
func (m *LoginParams) XXX_Size() int {
	return m.Size()
}
func (m *LoginParams) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginParams.DiscardUnknown(m)
}

var xxx_messageInfo_LoginParams proto.InternalMessageInfo

func (m *LoginParams) GetUName() string {
	if m != nil {
		return m.UName
	}
	return ""
}

func (m *LoginParams) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LoginParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRsp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	ServerPort           string   `protobuf:"bytes,3,opt,name=ServerPort,proto3" json:"ServerPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRsp) Reset()         { *m = LoginRsp{} }
func (m *LoginRsp) String() string { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()    {}
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}
func (m *LoginRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRsp.Merge(m, src)
}
func (m *LoginRsp) XXX_Size() int {
	return m.Size()
}
func (m *LoginRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRsp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRsp proto.InternalMessageInfo

func (m *LoginRsp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginRsp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRsp) GetServerPort() string {
	if m != nil {
		return m.ServerPort
	}
	return ""
}

type RegisterParams struct {
	UName                string   `protobuf:"bytes,1,opt,name=UName,proto3" json:"UName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterParams) Reset()         { *m = RegisterParams{} }
func (m *RegisterParams) String() string { return proto.CompactTextString(m) }
func (*RegisterParams) ProtoMessage()    {}
func (*RegisterParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}
func (m *RegisterParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterParams.Merge(m, src)
}
func (m *RegisterParams) XXX_Size() int {
	return m.Size()
}
func (m *RegisterParams) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterParams.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterParams proto.InternalMessageInfo

func (m *RegisterParams) GetUName() string {
	if m != nil {
		return m.UName
	}
	return ""
}

func (m *RegisterParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterRsp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRsp) Reset()         { *m = RegisterRsp{} }
func (m *RegisterRsp) String() string { return proto.CompactTextString(m) }
func (*RegisterRsp) ProtoMessage()    {}
func (*RegisterRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}
func (m *RegisterRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRsp.Merge(m, src)
}
func (m *RegisterRsp) XXX_Size() int {
	return m.Size()
}
func (m *RegisterRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRsp proto.InternalMessageInfo

func (m *RegisterRsp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterRsp) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginParams)(nil), "core.LoginParams")
	proto.RegisterType((*LoginRsp)(nil), "core.LoginRsp")
	proto.RegisterType((*RegisterParams)(nil), "core.RegisterParams")
	proto.RegisterType((*RegisterRsp)(nil), "core.RegisterRsp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4e, 0xf3, 0x40,
	0x10, 0x85, 0xe3, 0xc4, 0xf9, 0x95, 0x4c, 0x7e, 0x59, 0xc9, 0x28, 0x85, 0xe5, 0xc2, 0x42, 0xae,
	0xa8, 0x2c, 0x44, 0x2a, 0xda, 0xa4, 0x45, 0xc8, 0x2c, 0xf8, 0x00, 0x06, 0xaf, 0x82, 0x45, 0xe2,
	0xb1, 0x66, 0x0d, 0x54, 0xdc, 0x83, 0x23, 0x51, 0x72, 0x04, 0x64, 0x2e, 0x82, 0x76, 0x9c, 0x20,
	0x07, 0x04, 0xdd, 0xcc, 0xa7, 0x9d, 0x79, 0xef, 0xcd, 0x02, 0x3c, 0x18, 0xcd, 0x71, 0xc5, 0x54,
	0x13, 0xba, 0xb7, 0xc4, 0x3a, 0xba, 0x84, 0xc9, 0x39, 0xad, 0x8b, 0x32, 0xc9, 0x38, 0xdb, 0x1a,
	0x9c, 0xc3, 0x30, 0xbd, 0xc8, 0xb6, 0xda, 0x77, 0x8e, 0x9c, 0xe3, 0xb1, 0x6a, 0x1b, 0x9c, 0xc2,
	0x20, 0x2d, 0x72, 0xbf, 0x2f, 0xcc, 0x96, 0x18, 0xc0, 0x28, 0xc9, 0x8c, 0x79, 0x22, 0xce, 0xfd,
	0x81, 0xe0, 0xaf, 0x3e, 0xba, 0x83, 0x91, 0xac, 0x54, 0xa6, 0x42, 0x04, 0x77, 0x45, 0x79, 0xbb,
	0x6e, 0xa8, 0xa4, 0xb6, 0x4c, 0x24, 0x5c, 0x99, 0x93, 0xda, 0xea, 0x5e, 0xd3, 0xbd, 0x2e, 0x77,
	0x1a, 0x6d, 0x83, 0x21, 0xc0, 0x95, 0xe6, 0x47, 0xcd, 0x09, 0x71, 0xbd, 0xd3, 0xe9, 0x90, 0x68,
	0x09, 0x9e, 0xd2, 0xeb, 0xc2, 0xd4, 0x9a, 0xff, 0xf4, 0xdf, 0x75, 0xdb, 0xff, 0xe6, 0x76, 0x01,
	0x93, 0xfd, 0x8e, 0xdf, 0x0c, 0xff, 0x88, 0x7f, 0xfa, 0x0c, 0x5e, 0x6a, 0x34, 0xaf, 0xa8, 0xac,
	0x99, 0x36, 0x1b, 0xcd, 0x78, 0x02, 0x63, 0x4b, 0x24, 0x38, 0xce, 0x62, 0x7b, 0xdb, 0xb8, 0x73,
	0xd8, 0xc0, 0xeb, 0x20, 0x65, 0xaa, 0xa8, 0x87, 0x67, 0xf0, 0x7f, 0x2f, 0x6c, 0x27, 0x71, 0xde,
	0xbe, 0x38, 0x0c, 0x14, 0xcc, 0x0e, 0xa9, 0x8c, 0x2e, 0xa7, 0xaf, 0x4d, 0xe8, 0xbc, 0x35, 0xa1,
	0xf3, 0xde, 0x84, 0xce, 0xcb, 0x47, 0xd8, 0xbb, 0xf9, 0x27, 0x7f, 0xba, 0xf8, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0x94, 0xd5, 0x51, 0xd8, 0xe1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserControllerClient is the client API for UserController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserControllerClient interface {
	UserLogin(ctx context.Context, in *LoginParams, opts ...grpc.CallOption) (*LoginRsp, error)
	RegisterUser(ctx context.Context, in *RegisterParams, opts ...grpc.CallOption) (*RegisterRsp, error)
}

type userControllerClient struct {
	cc *grpc.ClientConn
}

func NewUserControllerClient(cc *grpc.ClientConn) UserControllerClient {
	return &userControllerClient{cc}
}

func (c *userControllerClient) UserLogin(ctx context.Context, in *LoginParams, opts ...grpc.CallOption) (*LoginRsp, error) {
	out := new(LoginRsp)
	err := c.cc.Invoke(ctx, "/core.UserController/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) RegisterUser(ctx context.Context, in *RegisterParams, opts ...grpc.CallOption) (*RegisterRsp, error) {
	out := new(RegisterRsp)
	err := c.cc.Invoke(ctx, "/core.UserController/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserControllerServer is the server API for UserController service.
type UserControllerServer interface {
	UserLogin(context.Context, *LoginParams) (*LoginRsp, error)
	RegisterUser(context.Context, *RegisterParams) (*RegisterRsp, error)
}

// UnimplementedUserControllerServer can be embedded to have forward compatible implementations.
type UnimplementedUserControllerServer struct {
}

func (*UnimplementedUserControllerServer) UserLogin(ctx context.Context, req *LoginParams) (*LoginRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserControllerServer) RegisterUser(ctx context.Context, req *RegisterParams) (*RegisterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}

func RegisterUserControllerServer(s *grpc.Server, srv UserControllerServer) {
	s.RegisterService(&_UserController_serviceDesc, srv)
}

func _UserController_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.UserController/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).UserLogin(ctx, req.(*LoginParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.UserController/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).RegisterUser(ctx, req.(*RegisterParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.UserController",
	HandlerType: (*UserControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserController_UserLogin_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _UserController_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func (m *LoginParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UName) > 0 {
		i -= len(m.UName)
		copy(dAtA[i:], m.UName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.UName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoginRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ServerPort) > 0 {
		i -= len(m.ServerPort)
		copy(dAtA[i:], m.ServerPort)
		i = encodeVarintUser(dAtA, i, uint64(len(m.ServerPort)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RegisterParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UName) > 0 {
		i -= len(m.UName)
		copy(dAtA[i:], m.UName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.UName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoginParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoginRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovUser(uint64(m.Code))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.ServerPort)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegisterParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegisterRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovUser(uint64(m.Code))
	}
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoginRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerPort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
