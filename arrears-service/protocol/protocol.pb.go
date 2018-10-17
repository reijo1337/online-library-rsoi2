// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

func init() {
	proto.RegisterType((*Arrear)(nil), "protocol.Arrear")
	proto.RegisterType((*PagingArrears)(nil), "protocol.PagingArrears")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x6a, 0x84, 0x30,
	0x14, 0x85, 0x1b, 0xff, 0x6a, 0x2f, 0x54, 0x24, 0x14, 0x1b, 0x5c, 0x89, 0x2b, 0x57, 0x52, 0xda,
	0x27, 0x68, 0x49, 0x91, 0xec, 0x24, 0x3e, 0x41, 0x66, 0x0c, 0x22, 0x33, 0x98, 0x21, 0x66, 0x35,
	0x4f, 0x3f, 0x98, 0xa8, 0xc3, 0x30, 0xbb, 0x73, 0xbe, 0x03, 0xf9, 0x72, 0x21, 0xb9, 0x68, 0x65,
	0xd4, 0x51, 0x9d, 0x6b, 0x1b, 0x70, 0xbc, 0xf5, 0xd2, 0x40, 0xf4, 0xab, 0xb5, 0x14, 0x1a, 0x27,
	0xe0, 0x31, 0x4a, 0x50, 0x81, 0xaa, 0x90, 0x7b, 0x8c, 0xe2, 0x1c, 0x62, 0x2e, 0x45, 0x2f, 0x35,
	0xa3, 0xc4, 0xb3, 0x74, 0xef, 0x38, 0x83, 0xe8, 0x4f, 0xa9, 0x13, 0xa3, 0xc4, 0xb7, 0xcb, 0xda,
	0xf0, 0x07, 0x84, 0x9d, 0x11, 0xda, 0x90, 0xa0, 0x40, 0xd5, 0x1b, 0x77, 0x05, 0xa7, 0xe0, 0xff,
	0x4f, 0x3d, 0x09, 0x2d, 0x5b, 0x62, 0xd9, 0xc0, 0x7b, 0x2b, 0x86, 0x71, 0x1a, 0x9c, 0x7b, 0x7e,
	0x92, 0x63, 0x08, 0xba, 0xf1, 0x2a, 0x57, 0xb1, 0xcd, 0x0b, 0x6b, 0xc5, 0x20, 0x57, 0xa5, 0xcd,
	0xdf, 0x1c, 0x5e, 0xb7, 0x27, 0x1a, 0xc8, 0x1a, 0x69, 0x16, 0xda, 0xbb, 0x7f, 0xce, 0xdb, 0xf2,
	0x59, 0xef, 0xe7, 0x3f, 0x58, 0xf3, 0xf4, 0x3e, 0x38, 0x54, 0xbe, 0x7c, 0xa1, 0x43, 0x64, 0xe1,
	0xcf, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x95, 0x1f, 0xcc, 0xd5, 0x35, 0x01, 0x00, 0x00,
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

// ArrearsServer is the server API for Arrears service.
type ArrearsServer interface {
	GetPagedReadersArrears(*PagingArrears, Arrears_GetPagedReadersArrearsServer) error
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

var _Arrears_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Arrears",
	HandlerType: (*ArrearsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPagedReadersArrears",
			Handler:       _Arrears_GetPagedReadersArrears_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol.proto",
}
