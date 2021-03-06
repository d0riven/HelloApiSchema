// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserInput struct {
	EmailAddress         string   `protobuf:"bytes,1,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Birthday             string   `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserInput) Reset()         { *m = CreateUserInput{} }
func (m *CreateUserInput) String() string { return proto.CompactTextString(m) }
func (*CreateUserInput) ProtoMessage()    {}
func (*CreateUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *CreateUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserInput.Unmarshal(m, b)
}
func (m *CreateUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserInput.Marshal(b, m, deterministic)
}
func (m *CreateUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserInput.Merge(m, src)
}
func (m *CreateUserInput) XXX_Size() int {
	return xxx_messageInfo_CreateUserInput.Size(m)
}
func (m *CreateUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserInput proto.InternalMessageInfo

func (m *CreateUserInput) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *CreateUserInput) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CreateUserInput) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateUserInput) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *CreateUserInput) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CreateUserOutput struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailAddress         string   `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName            string   `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Birthday             string   `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserOutput) Reset()         { *m = CreateUserOutput{} }
func (m *CreateUserOutput) String() string { return proto.CompactTextString(m) }
func (*CreateUserOutput) ProtoMessage()    {}
func (*CreateUserOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CreateUserOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserOutput.Unmarshal(m, b)
}
func (m *CreateUserOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserOutput.Marshal(b, m, deterministic)
}
func (m *CreateUserOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserOutput.Merge(m, src)
}
func (m *CreateUserOutput) XXX_Size() int {
	return xxx_messageInfo_CreateUserOutput.Size(m)
}
func (m *CreateUserOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserOutput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserOutput proto.InternalMessageInfo

func (m *CreateUserOutput) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateUserOutput) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *CreateUserOutput) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CreateUserOutput) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateUserOutput) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *CreateUserOutput) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UpdateUserInput struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailAddress         string   `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName            string   `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Birthday             string   `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInput) Reset()         { *m = UpdateUserInput{} }
func (m *UpdateUserInput) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInput) ProtoMessage()    {}
func (*UpdateUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UpdateUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInput.Unmarshal(m, b)
}
func (m *UpdateUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInput.Marshal(b, m, deterministic)
}
func (m *UpdateUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInput.Merge(m, src)
}
func (m *UpdateUserInput) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInput.Size(m)
}
func (m *UpdateUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInput proto.InternalMessageInfo

func (m *UpdateUserInput) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserInput) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *UpdateUserInput) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UpdateUserInput) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UpdateUserInput) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *UpdateUserInput) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UpdateUserOutput struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailAddress         string   `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName            string   `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Birthday             string   `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserOutput) Reset()         { *m = UpdateUserOutput{} }
func (m *UpdateUserOutput) String() string { return proto.CompactTextString(m) }
func (*UpdateUserOutput) ProtoMessage()    {}
func (*UpdateUserOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UpdateUserOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserOutput.Unmarshal(m, b)
}
func (m *UpdateUserOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserOutput.Marshal(b, m, deterministic)
}
func (m *UpdateUserOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserOutput.Merge(m, src)
}
func (m *UpdateUserOutput) XXX_Size() int {
	return xxx_messageInfo_UpdateUserOutput.Size(m)
}
func (m *UpdateUserOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserOutput.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserOutput proto.InternalMessageInfo

func (m *UpdateUserOutput) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserOutput) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *UpdateUserOutput) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UpdateUserOutput) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UpdateUserOutput) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *UpdateUserOutput) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type GetUserInput struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInput) Reset()         { *m = GetUserInput{} }
func (m *GetUserInput) String() string { return proto.CompactTextString(m) }
func (*GetUserInput) ProtoMessage()    {}
func (*GetUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *GetUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInput.Unmarshal(m, b)
}
func (m *GetUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInput.Marshal(b, m, deterministic)
}
func (m *GetUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInput.Merge(m, src)
}
func (m *GetUserInput) XXX_Size() int {
	return xxx_messageInfo_GetUserInput.Size(m)
}
func (m *GetUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInput proto.InternalMessageInfo

func (m *GetUserInput) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserOutput struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailAddress         string   `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName            string   `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Birthday             string   `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserOutput) Reset()         { *m = GetUserOutput{} }
func (m *GetUserOutput) String() string { return proto.CompactTextString(m) }
func (*GetUserOutput) ProtoMessage()    {}
func (*GetUserOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *GetUserOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserOutput.Unmarshal(m, b)
}
func (m *GetUserOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserOutput.Marshal(b, m, deterministic)
}
func (m *GetUserOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserOutput.Merge(m, src)
}
func (m *GetUserOutput) XXX_Size() int {
	return xxx_messageInfo_GetUserOutput.Size(m)
}
func (m *GetUserOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserOutput proto.InternalMessageInfo

func (m *GetUserOutput) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUserOutput) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *GetUserOutput) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *GetUserOutput) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *GetUserOutput) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *GetUserOutput) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserInput)(nil), "proto.CreateUserInput")
	proto.RegisterType((*CreateUserOutput)(nil), "proto.CreateUserOutput")
	proto.RegisterType((*UpdateUserInput)(nil), "proto.UpdateUserInput")
	proto.RegisterType((*UpdateUserOutput)(nil), "proto.UpdateUserOutput")
	proto.RegisterType((*GetUserInput)(nil), "proto.GetUserInput")
	proto.RegisterType((*GetUserOutput)(nil), "proto.GetUserOutput")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0xc7, 0x99, 0xa4, 0xa9, 0xcd, 0x6b, 0x6b, 0xcb, 0x28, 0x1a, 0xe2, 0x07, 0x12, 0x37, 0xe2,
	0xa2, 0x81, 0xba, 0x73, 0xe7, 0x07, 0x14, 0x11, 0x14, 0x2a, 0xdd, 0xb8, 0x29, 0x53, 0x67, 0xac,
	0x83, 0x6d, 0x12, 0x26, 0x53, 0x41, 0xc4, 0x8d, 0x57, 0xf0, 0x12, 0x9e, 0x40, 0xf4, 0x1c, 0x5e,
	0xc1, 0xad, 0x77, 0x90, 0xcc, 0x34, 0xad, 0x6d, 0x69, 0x5d, 0x67, 0x35, 0xcc, 0xff, 0xbd, 0xf9,
	0xe7, 0x37, 0x6f, 0xde, 0x0b, 0xc0, 0x20, 0x66, 0xa2, 0x16, 0x89, 0x50, 0x86, 0xd8, 0x52, 0x8b,
	0xbb, 0xd9, 0x0d, 0xc3, 0x6e, 0x8f, 0xf9, 0x24, 0xe2, 0x3e, 0x09, 0x82, 0x50, 0x12, 0xc9, 0xc3,
	0x20, 0xd6, 0x49, 0xde, 0x1b, 0x82, 0xca, 0x89, 0x60, 0x44, 0xb2, 0x56, 0xcc, 0xc4, 0x59, 0x10,
	0x0d, 0x24, 0xde, 0x85, 0x32, 0xeb, 0x13, 0xde, 0x6b, 0x13, 0x4a, 0x05, 0x8b, 0x63, 0x07, 0xed,
	0xa0, 0x3d, 0xbb, 0x59, 0x52, 0xe2, 0x91, 0xd6, 0xf0, 0x06, 0xd8, 0x3d, 0x12, 0xcb, 0x76, 0x40,
	0xfa, 0xcc, 0x31, 0x54, 0x42, 0x21, 0x11, 0x2e, 0x48, 0x9f, 0xe1, 0x2d, 0x80, 0x5b, 0x2e, 0xd2,
	0xa8, 0xa9, 0xa2, 0xb6, 0x52, 0x54, 0xd8, 0x85, 0x42, 0x87, 0x0b, 0x79, 0x47, 0xc9, 0xa3, 0x93,
	0xd3, 0x47, 0xd3, 0x3d, 0x76, 0x60, 0x29, 0xfd, 0xac, 0xa5, 0x42, 0xe9, 0xd6, 0xfb, 0x44, 0x50,
	0x1d, 0xa3, 0x5e, 0x0e, 0x64, 0xc2, 0xba, 0x0c, 0x06, 0xa7, 0x0a, 0xd0, 0x6c, 0x1a, 0x9c, 0xce,
	0xb2, 0x1b, 0xff, 0xb1, 0x9b, 0x0b, 0xd9, 0x73, 0x8b, 0xd8, 0xad, 0xf9, 0xec, 0xf9, 0x49, 0xf6,
	0x0f, 0x04, 0x95, 0x56, 0x44, 0x27, 0xca, 0x9c, 0x11, 0xf4, 0xa4, 0xec, 0x63, 0xf4, 0x6c, 0x95,
	0x7d, 0x1b, 0x4a, 0x0d, 0x26, 0xe7, 0x96, 0xdc, 0x7b, 0x47, 0x50, 0x1e, 0x26, 0x64, 0xea, 0x62,
	0xf5, 0x1f, 0x04, 0xc5, 0x84, 0xfa, 0x8a, 0x89, 0x07, 0x7e, 0xc3, 0xf0, 0x39, 0xe4, 0xf5, 0x68,
	0xe0, 0x35, 0x3d, 0xd8, 0xb5, 0xa9, 0xa1, 0x76, 0xd7, 0x67, 0x74, 0x7d, 0x63, 0xaf, 0xfa, 0xf2,
	0xf5, 0xfd, 0x6a, 0x80, 0x67, 0xf9, 0xc9, 0xbf, 0xe3, 0x10, 0xed, 0x27, 0x66, 0xfa, 0xc1, 0x47,
	0x66, 0x53, 0xad, 0x3b, 0x32, 0x9b, 0xee, 0x8b, 0xd4, 0xac, 0x3e, 0x36, 0x3b, 0x05, 0xb3, 0xc1,
	0x24, 0x5e, 0x19, 0x9e, 0xf8, 0xfb, 0x1c, 0xee, 0xea, 0xa4, 0x38, 0xf4, 0xc0, 0xca, 0xa3, 0x84,
	0x41, 0x79, 0xf8, 0x4f, 0x9c, 0x3e, 0x1f, 0x17, 0xaf, 0xed, 0xe8, 0x9e, 0xfa, 0x2a, 0xbd, 0x93,
	0x57, 0xcb, 0xc1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xc1, 0x7a, 0xee, 0xed, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUserInput, opts ...grpc.CallOption) (*CreateUserOutput, error)
	Update(ctx context.Context, in *UpdateUserInput, opts ...grpc.CallOption) (*UpdateUserOutput, error)
	Get(ctx context.Context, in *GetUserInput, opts ...grpc.CallOption) (*GetUserOutput, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserInput, opts ...grpc.CallOption) (*CreateUserOutput, error) {
	out := new(CreateUserOutput)
	err := c.cc.Invoke(ctx, "/proto.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserInput, opts ...grpc.CallOption) (*UpdateUserOutput, error) {
	out := new(UpdateUserOutput)
	err := c.cc.Invoke(ctx, "/proto.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *GetUserInput, opts ...grpc.CallOption) (*GetUserOutput, error) {
	out := new(GetUserOutput)
	err := c.cc.Invoke(ctx, "/proto.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *CreateUserInput) (*CreateUserOutput, error)
	Update(context.Context, *UpdateUserInput) (*UpdateUserOutput, error)
	Get(context.Context, *GetUserInput) (*GetUserOutput, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *CreateUserInput) (*CreateUserOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Update(ctx context.Context, req *UpdateUserInput) (*UpdateUserOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServiceServer) Get(ctx context.Context, req *GetUserInput) (*GetUserOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetUserInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
