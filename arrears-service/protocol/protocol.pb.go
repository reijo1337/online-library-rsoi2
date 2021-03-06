// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Arrear struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ReaderID             int32    `protobuf:"varint,2,opt,name=ReaderID,proto3" json:"ReaderID,omitempty"`
	BookID               int32    `protobuf:"varint,3,opt,name=BookID,proto3" json:"BookID,omitempty"`
	Start                string   `protobuf:"bytes,4,opt,name=Start,proto3" json:"Start,omitempty"`
	End                  string   `protobuf:"bytes,5,opt,name=End,proto3" json:"End,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Arrear) Reset()         { *m = Arrear{} }
func (m *Arrear) String() string { return proto.CompactTextString(m) }
func (*Arrear) ProtoMessage()    {}
func (*Arrear) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *Arrear) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arrear.Unmarshal(m, b)
}
func (m *Arrear) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arrear.Marshal(b, m, deterministic)
}
func (m *Arrear) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arrear.Merge(m, src)
}
func (m *Arrear) XXX_Size() int {
	return xxx_messageInfo_Arrear.Size(m)
}
func (m *Arrear) XXX_DiscardUnknown() {
	xxx_messageInfo_Arrear.DiscardUnknown(m)
}

var xxx_messageInfo_Arrear proto.InternalMessageInfo

func (m *Arrear) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Arrear) GetReaderID() int32 {
	if m != nil {
		return m.ReaderID
	}
	return 0
}

func (m *Arrear) GetBookID() int32 {
	if m != nil {
		return m.BookID
	}
	return 0
}

func (m *Arrear) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Arrear) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type PagingArrears struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagingArrears) Reset()         { *m = PagingArrears{} }
func (m *PagingArrears) String() string { return proto.CompactTextString(m) }
func (*PagingArrears) ProtoMessage()    {}
func (*PagingArrears) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *PagingArrears) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingArrears.Unmarshal(m, b)
}
func (m *PagingArrears) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingArrears.Marshal(b, m, deterministic)
}
func (m *PagingArrears) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingArrears.Merge(m, src)
}
func (m *PagingArrears) XXX_Size() int {
	return xxx_messageInfo_PagingArrears.Size(m)
}
func (m *PagingArrears) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingArrears.DiscardUnknown(m)
}

var xxx_messageInfo_PagingArrears proto.InternalMessageInfo

func (m *PagingArrears) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PagingArrears) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PagingArrears) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type NewArrear struct {
	ReaderID             int32    `protobuf:"varint,1,opt,name=ReaderID,proto3" json:"ReaderID,omitempty"`
	BookID               int32    `protobuf:"varint,2,opt,name=BookID,proto3" json:"BookID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewArrear) Reset()         { *m = NewArrear{} }
func (m *NewArrear) String() string { return proto.CompactTextString(m) }
func (*NewArrear) ProtoMessage()    {}
func (*NewArrear) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *NewArrear) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewArrear.Unmarshal(m, b)
}
func (m *NewArrear) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewArrear.Marshal(b, m, deterministic)
}
func (m *NewArrear) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewArrear.Merge(m, src)
}
func (m *NewArrear) XXX_Size() int {
	return xxx_messageInfo_NewArrear.Size(m)
}
func (m *NewArrear) XXX_DiscardUnknown() {
	xxx_messageInfo_NewArrear.DiscardUnknown(m)
}

var xxx_messageInfo_NewArrear proto.InternalMessageInfo

func (m *NewArrear) GetReaderID() int32 {
	if m != nil {
		return m.ReaderID
	}
	return 0
}

func (m *NewArrear) GetBookID() int32 {
	if m != nil {
		return m.BookID
	}
	return 0
}

type SomeArrearsID struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SomeArrearsID) Reset()         { *m = SomeArrearsID{} }
func (m *SomeArrearsID) String() string { return proto.CompactTextString(m) }
func (*SomeArrearsID) ProtoMessage()    {}
func (*SomeArrearsID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *SomeArrearsID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SomeArrearsID.Unmarshal(m, b)
}
func (m *SomeArrearsID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SomeArrearsID.Marshal(b, m, deterministic)
}
func (m *SomeArrearsID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SomeArrearsID.Merge(m, src)
}
func (m *SomeArrearsID) XXX_Size() int {
	return xxx_messageInfo_SomeArrearsID.Size(m)
}
func (m *SomeArrearsID) XXX_DiscardUnknown() {
	xxx_messageInfo_SomeArrearsID.DiscardUnknown(m)
}

var xxx_messageInfo_SomeArrearsID proto.InternalMessageInfo

func (m *SomeArrearsID) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type NothingArrear struct {
	Dummy                bool     `protobuf:"varint,1,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NothingArrear) Reset()         { *m = NothingArrear{} }
func (m *NothingArrear) String() string { return proto.CompactTextString(m) }
func (*NothingArrear) ProtoMessage()    {}
func (*NothingArrear) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{4}
}

func (m *NothingArrear) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NothingArrear.Unmarshal(m, b)
}
func (m *NothingArrear) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NothingArrear.Marshal(b, m, deterministic)
}
func (m *NothingArrear) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NothingArrear.Merge(m, src)
}
func (m *NothingArrear) XXX_Size() int {
	return xxx_messageInfo_NothingArrear.Size(m)
}
func (m *NothingArrear) XXX_DiscardUnknown() {
	xxx_messageInfo_NothingArrear.DiscardUnknown(m)
}

var xxx_messageInfo_NothingArrear proto.InternalMessageInfo

func (m *NothingArrear) GetDummy() bool {
	if m != nil {
		return m.Dummy
	}
	return false
}

type AuthRequest struct {
	AppKey               string   `protobuf:"bytes,1,opt,name=AppKey,proto3" json:"AppKey,omitempty"`
	AppSecret            string   `protobuf:"bytes,2,opt,name=AppSecret,proto3" json:"AppSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{5}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *AuthRequest) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

type SomeString struct {
	String_              string   `protobuf:"bytes,1,opt,name=String,proto3" json:"String,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SomeString) Reset()         { *m = SomeString{} }
func (m *SomeString) String() string { return proto.CompactTextString(m) }
func (*SomeString) ProtoMessage()    {}
func (*SomeString) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{6}
}

func (m *SomeString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SomeString.Unmarshal(m, b)
}
func (m *SomeString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SomeString.Marshal(b, m, deterministic)
}
func (m *SomeString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SomeString.Merge(m, src)
}
func (m *SomeString) XXX_Size() int {
	return xxx_messageInfo_SomeString.Size(m)
}
func (m *SomeString) XXX_DiscardUnknown() {
	xxx_messageInfo_SomeString.DiscardUnknown(m)
}

var xxx_messageInfo_SomeString proto.InternalMessageInfo

func (m *SomeString) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

func init() {
	proto.RegisterType((*Arrear)(nil), "protocol.Arrear")
	proto.RegisterType((*PagingArrears)(nil), "protocol.PagingArrears")
	proto.RegisterType((*NewArrear)(nil), "protocol.NewArrear")
	proto.RegisterType((*SomeArrearsID)(nil), "protocol.SomeArrearsID")
	proto.RegisterType((*NothingArrear)(nil), "protocol.NothingArrear")
	proto.RegisterType((*AuthRequest)(nil), "protocol.AuthRequest")
	proto.RegisterType((*SomeString)(nil), "protocol.SomeString")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xef, 0xea, 0xd3, 0x40,
	0x10, 0x6c, 0xf2, 0x6b, 0x6b, 0xb3, 0x92, 0x52, 0xcf, 0x5a, 0x43, 0x11, 0x2c, 0x87, 0x42, 0x3f,
	0x15, 0x51, 0xfc, 0x56, 0x90, 0xd4, 0xd3, 0x10, 0x84, 0x52, 0x2e, 0x4f, 0x10, 0xdb, 0x25, 0x0d,
	0xb6, 0x4d, 0xbc, 0x5c, 0x91, 0xfa, 0x8e, 0xbe, 0x93, 0xdc, 0x5d, 0xfe, 0x6a, 0x05, 0xbf, 0xed,
	0xcc, 0xde, 0xed, 0xcc, 0xce, 0xc2, 0x38, 0x17, 0x99, 0xcc, 0xf6, 0xd9, 0x69, 0xa5, 0x0b, 0x32,
	0xaa, 0x30, 0x95, 0x30, 0xf4, 0x85, 0xc0, 0x58, 0x90, 0x31, 0xd8, 0x21, 0xf3, 0xac, 0x85, 0xb5,
	0x1c, 0x70, 0x3b, 0x64, 0x64, 0x0e, 0x23, 0x8e, 0xf1, 0x01, 0x45, 0xc8, 0x3c, 0x5b, 0xb3, 0x35,
	0x26, 0x33, 0x18, 0x6e, 0xb2, 0xec, 0x5b, 0xc8, 0xbc, 0x07, 0xdd, 0x29, 0x11, 0x99, 0xc2, 0x20,
	0x92, 0xb1, 0x90, 0x5e, 0x7f, 0x61, 0x2d, 0x1d, 0x6e, 0x00, 0x99, 0xc0, 0xc3, 0xa7, 0xcb, 0xc1,
	0x1b, 0x68, 0x4e, 0x95, 0x34, 0x00, 0x77, 0x17, 0x27, 0xe9, 0x25, 0x31, 0xda, 0xc5, 0x5f, 0xe2,
	0x04, 0xfa, 0x51, 0xfa, 0x13, 0x4b, 0x61, 0x5d, 0x2b, 0x6e, 0x17, 0x27, 0x58, 0x4a, 0xea, 0x9a,
	0x7e, 0x00, 0x67, 0x8b, 0x3f, 0xca, 0x0d, 0xda, 0x8e, 0xad, 0x7f, 0x3a, 0xb6, 0xdb, 0x8e, 0xe9,
	0x4b, 0x70, 0xa3, 0xec, 0x8c, 0xa5, 0x8f, 0x90, 0xfd, 0xe9, 0x84, 0xbe, 0x06, 0x77, 0x9b, 0xc9,
	0x63, 0xed, 0x55, 0xed, 0xc8, 0xae, 0xe7, 0xf3, 0x4d, 0xbf, 0x19, 0x71, 0x03, 0xe8, 0x47, 0x78,
	0xec, 0x5f, 0xe5, 0x91, 0xe3, 0xf7, 0x2b, 0x16, 0x52, 0xc9, 0xf9, 0x79, 0xfe, 0x05, 0xcd, 0x2b,
	0x87, 0x97, 0x88, 0xbc, 0x00, 0xc7, 0xcf, 0xf3, 0x08, 0xf7, 0x02, 0xa5, 0x76, 0xe2, 0xf0, 0x86,
	0xa0, 0xaf, 0x00, 0x94, 0x99, 0x48, 0x8a, 0xf4, 0x92, 0xa8, 0x19, 0xa6, 0xaa, 0x66, 0x18, 0xf4,
	0xf6, 0x97, 0x0d, 0x8f, 0xaa, 0xdc, 0x02, 0x98, 0x05, 0x28, 0x55, 0x14, 0x07, 0xb3, 0x6a, 0x51,
	0x75, 0x9e, 0xaf, 0xea, 0x9b, 0x77, 0xa2, 0x9e, 0x4f, 0x9a, 0x86, 0xa1, 0x68, 0xef, 0x8d, 0x45,
	0xd6, 0xf0, 0x84, 0x63, 0x92, 0x16, 0x12, 0x45, 0x13, 0xe8, 0xd3, 0xe6, 0x69, 0x4d, 0xde, 0xfb,
	0x4f, 0xd6, 0xe0, 0x06, 0x28, 0x0d, 0xdc, 0xdc, 0x42, 0xd6, 0x56, 0xef, 0xc4, 0x7b, 0xf7, 0xf7,
	0x67, 0x98, 0x30, 0x3c, 0xa1, 0xc4, 0xff, 0x19, 0xd0, 0x6a, 0x74, 0xee, 0x42, 0x7b, 0xe4, 0x3d,
	0xf4, 0xd5, 0x0d, 0xc8, 0xb3, 0x96, 0x46, 0x73, 0x93, 0xf9, 0xb4, 0x3b, 0xd2, 0xa4, 0x49, 0x7b,
	0x5f, 0x87, 0x9a, 0x7e, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x68, 0x6d, 0x7d, 0x92, 0x25, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArrearsClient is the client API for Arrears service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArrearsClient interface {
	GetPagedReadersArrears(ctx context.Context, in *PagingArrears, opts ...grpc.CallOption) (Arrears_GetPagedReadersArrearsClient, error)
	RegisterNewArrear(ctx context.Context, in *NewArrear, opts ...grpc.CallOption) (*Arrear, error)
	GetArrearByID(ctx context.Context, in *SomeArrearsID, opts ...grpc.CallOption) (*Arrear, error)
	DeleteArrearByID(ctx context.Context, in *SomeArrearsID, opts ...grpc.CallOption) (*NothingArrear, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*SomeString, error)
}

type arrearsClient struct {
	cc *grpc.ClientConn
}

func NewArrearsClient(cc *grpc.ClientConn) ArrearsClient {
	return &arrearsClient{cc}
}

func (c *arrearsClient) GetPagedReadersArrears(ctx context.Context, in *PagingArrears, opts ...grpc.CallOption) (Arrears_GetPagedReadersArrearsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Arrears_serviceDesc.Streams[0], "/protocol.Arrears/GetPagedReadersArrears", opts...)
	if err != nil {
		return nil, err
	}
	x := &arrearsGetPagedReadersArrearsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Arrears_GetPagedReadersArrearsClient interface {
	Recv() (*Arrear, error)
	grpc.ClientStream
}

type arrearsGetPagedReadersArrearsClient struct {
	grpc.ClientStream
}

func (x *arrearsGetPagedReadersArrearsClient) Recv() (*Arrear, error) {
	m := new(Arrear)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arrearsClient) RegisterNewArrear(ctx context.Context, in *NewArrear, opts ...grpc.CallOption) (*Arrear, error) {
	out := new(Arrear)
	err := c.cc.Invoke(ctx, "/protocol.Arrears/RegisterNewArrear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arrearsClient) GetArrearByID(ctx context.Context, in *SomeArrearsID, opts ...grpc.CallOption) (*Arrear, error) {
	out := new(Arrear)
	err := c.cc.Invoke(ctx, "/protocol.Arrears/GetArrearByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arrearsClient) DeleteArrearByID(ctx context.Context, in *SomeArrearsID, opts ...grpc.CallOption) (*NothingArrear, error) {
	out := new(NothingArrear)
	err := c.cc.Invoke(ctx, "/protocol.Arrears/DeleteArrearByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arrearsClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*SomeString, error) {
	out := new(SomeString)
	err := c.cc.Invoke(ctx, "/protocol.Arrears/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArrearsServer is the server API for Arrears service.
type ArrearsServer interface {
	GetPagedReadersArrears(*PagingArrears, Arrears_GetPagedReadersArrearsServer) error
	RegisterNewArrear(context.Context, *NewArrear) (*Arrear, error)
	GetArrearByID(context.Context, *SomeArrearsID) (*Arrear, error)
	DeleteArrearByID(context.Context, *SomeArrearsID) (*NothingArrear, error)
	Auth(context.Context, *AuthRequest) (*SomeString, error)
}

func RegisterArrearsServer(s *grpc.Server, srv ArrearsServer) {
	s.RegisterService(&_Arrears_serviceDesc, srv)
}

func _Arrears_GetPagedReadersArrears_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PagingArrears)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArrearsServer).GetPagedReadersArrears(m, &arrearsGetPagedReadersArrearsServer{stream})
}

type Arrears_GetPagedReadersArrearsServer interface {
	Send(*Arrear) error
	grpc.ServerStream
}

type arrearsGetPagedReadersArrearsServer struct {
	grpc.ServerStream
}

func (x *arrearsGetPagedReadersArrearsServer) Send(m *Arrear) error {
	return x.ServerStream.SendMsg(m)
}

func _Arrears_RegisterNewArrear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewArrear)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArrearsServer).RegisterNewArrear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Arrears/RegisterNewArrear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArrearsServer).RegisterNewArrear(ctx, req.(*NewArrear))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arrears_GetArrearByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SomeArrearsID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArrearsServer).GetArrearByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Arrears/GetArrearByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArrearsServer).GetArrearByID(ctx, req.(*SomeArrearsID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arrears_DeleteArrearByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SomeArrearsID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArrearsServer).DeleteArrearByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Arrears/DeleteArrearByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArrearsServer).DeleteArrearByID(ctx, req.(*SomeArrearsID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arrears_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArrearsServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Arrears/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArrearsServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Arrears_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Arrears",
	HandlerType: (*ArrearsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewArrear",
			Handler:    _Arrears_RegisterNewArrear_Handler,
		},
		{
			MethodName: "GetArrearByID",
			Handler:    _Arrears_GetArrearByID_Handler,
		},
		{
			MethodName: "DeleteArrearByID",
			Handler:    _Arrears_DeleteArrearByID_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Arrears_Auth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPagedReadersArrears",
			Handler:       _Arrears_GetPagedReadersArrears_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol.proto",
}
