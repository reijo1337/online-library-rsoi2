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

type Reader struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reader) Reset()         { *m = Reader{} }
func (m *Reader) String() string { return proto.CompactTextString(m) }
func (*Reader) ProtoMessage()    {}
func (*Reader) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *Reader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reader.Unmarshal(m, b)
}
func (m *Reader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reader.Marshal(b, m, deterministic)
}
func (m *Reader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reader.Merge(m, src)
}
func (m *Reader) XXX_Size() int {
	return xxx_messageInfo_Reader.Size(m)
}
func (m *Reader) XXX_DiscardUnknown() {
	xxx_messageInfo_Reader.DiscardUnknown(m)
}

var xxx_messageInfo_Reader proto.InternalMessageInfo

func (m *Reader) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Reader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReaderName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReaderName) Reset()         { *m = ReaderName{} }
func (m *ReaderName) String() string { return proto.CompactTextString(m) }
func (*ReaderName) ProtoMessage()    {}
func (*ReaderName) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *ReaderName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReaderName.Unmarshal(m, b)
}
func (m *ReaderName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReaderName.Marshal(b, m, deterministic)
}
func (m *ReaderName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReaderName.Merge(m, src)
}
func (m *ReaderName) XXX_Size() int {
	return xxx_messageInfo_ReaderName.Size(m)
}
func (m *ReaderName) XXX_DiscardUnknown() {
	xxx_messageInfo_ReaderName.DiscardUnknown(m)
}

var xxx_messageInfo_ReaderName proto.InternalMessageInfo

func (m *ReaderName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Nothing struct {
	Dummy                bool     `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func (m *Nothing) GetDummy() bool {
	if m != nil {
		return m.Dummy
	}
	return false
}

func init() {
	proto.RegisterType((*Reader)(nil), "protocol.Reader")
	proto.RegisterType((*ReaderName)(nil), "protocol.ReaderName")
	proto.RegisterType((*Nothing)(nil), "protocol.Nothing")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x03, 0x33, 0x84, 0x38, 0x60, 0x7c, 0x25, 0x1d, 0x2e, 0xb6, 0xa0,
	0xd4, 0xc4, 0x94, 0xd4, 0x22, 0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0x26, 0x4f, 0x17, 0x21, 0x21, 0x2e, 0x16, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x81, 0x8b, 0x0b, 0xa2, 0x1a, 0xc4, 0x83, 0xab, 0x60,
	0x44, 0x52, 0x21, 0xcf, 0xc5, 0xee, 0x97, 0x5f, 0x92, 0x91, 0x99, 0x97, 0x2e, 0x24, 0xc2, 0xc5,
	0x9a, 0x52, 0x9a, 0x9b, 0x5b, 0x09, 0x96, 0xe7, 0x08, 0x82, 0x70, 0x8c, 0x1a, 0x18, 0xb9, 0xd8,
	0x21, 0x66, 0x14, 0x0b, 0x59, 0x71, 0xf1, 0x05, 0xa5, 0xa6, 0x67, 0x16, 0x97, 0xa4, 0x16, 0x41,
	0x1d, 0x21, 0xa2, 0x07, 0x77, 0x29, 0xc2, 0x22, 0x29, 0x01, 0x74, 0x51, 0x25, 0x06, 0x21, 0x4b,
	0x2e, 0x3e, 0xf7, 0xd4, 0x12, 0xa8, 0x49, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0x82, 0x08, 0x55, 0x50,
	0x27, 0x60, 0xd3, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x16, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xa4, 0xb8, 0x67, 0xb5, 0x16, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReadersClient is the client API for Readers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReadersClient interface {
	RegisterReader(ctx context.Context, in *ReaderName, opts ...grpc.CallOption) (*Reader, error)
	GetReadersList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (Readers_GetReadersListClient, error)
}

type readersClient struct {
	cc *grpc.ClientConn
}

func NewReadersClient(cc *grpc.ClientConn) ReadersClient {
	return &readersClient{cc}
}

func (c *readersClient) RegisterReader(ctx context.Context, in *ReaderName, opts ...grpc.CallOption) (*Reader, error) {
	out := new(Reader)
	err := c.cc.Invoke(ctx, "/protocol.Readers/RegisterReader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readersClient) GetReadersList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (Readers_GetReadersListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Readers_serviceDesc.Streams[0], "/protocol.Readers/GetReadersList", opts...)
	if err != nil {
		return nil, err
	}
	x := &readersGetReadersListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Readers_GetReadersListClient interface {
	Recv() (*Reader, error)
	grpc.ClientStream
}

type readersGetReadersListClient struct {
	grpc.ClientStream
}

func (x *readersGetReadersListClient) Recv() (*Reader, error) {
	m := new(Reader)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReadersServer is the server API for Readers service.
type ReadersServer interface {
	RegisterReader(context.Context, *ReaderName) (*Reader, error)
	GetReadersList(*Nothing, Readers_GetReadersListServer) error
}

func RegisterReadersServer(s *grpc.Server, srv ReadersServer) {
	s.RegisterService(&_Readers_serviceDesc, srv)
}

func _Readers_RegisterReader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReaderName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadersServer).RegisterReader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Readers/RegisterReader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadersServer).RegisterReader(ctx, req.(*ReaderName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Readers_GetReadersList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Nothing)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReadersServer).GetReadersList(m, &readersGetReadersListServer{stream})
}

type Readers_GetReadersListServer interface {
	Send(*Reader) error
	grpc.ServerStream
}

type readersGetReadersListServer struct {
	grpc.ServerStream
}

func (x *readersGetReadersListServer) Send(m *Reader) error {
	return x.ServerStream.SendMsg(m)
}

var _Readers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Readers",
	HandlerType: (*ReadersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterReader",
			Handler:    _Readers_RegisterReader_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetReadersList",
			Handler:       _Readers_GetReadersList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol.proto",
}
