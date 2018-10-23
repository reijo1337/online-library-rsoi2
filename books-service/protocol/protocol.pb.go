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

type Writer struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Writer) Reset()         { *m = Writer{} }
func (m *Writer) String() string { return proto.CompactTextString(m) }
func (*Writer) ProtoMessage()    {}
func (*Writer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *Writer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Writer.Unmarshal(m, b)
}
func (m *Writer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Writer.Marshal(b, m, deterministic)
}
func (m *Writer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Writer.Merge(m, src)
}
func (m *Writer) XXX_Size() int {
	return xxx_messageInfo_Writer.Size(m)
}
func (m *Writer) XXX_DiscardUnknown() {
	xxx_messageInfo_Writer.DiscardUnknown(m)
}

var xxx_messageInfo_Writer proto.InternalMessageInfo

func (m *Writer) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Writer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Book struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Author               *Writer  `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetAuthor() *Writer {
	if m != nil {
		return m.Author
	}
	return nil
}

type NothingBooks struct {
	Dummy                bool     `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NothingBooks) Reset()         { *m = NothingBooks{} }
func (m *NothingBooks) String() string { return proto.CompactTextString(m) }
func (*NothingBooks) ProtoMessage()    {}
func (*NothingBooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *NothingBooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NothingBooks.Unmarshal(m, b)
}
func (m *NothingBooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NothingBooks.Marshal(b, m, deterministic)
}
func (m *NothingBooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NothingBooks.Merge(m, src)
}
func (m *NothingBooks) XXX_Size() int {
	return xxx_messageInfo_NothingBooks.Size(m)
}
func (m *NothingBooks) XXX_DiscardUnknown() {
	xxx_messageInfo_NothingBooks.DiscardUnknown(m)
}

var xxx_messageInfo_NothingBooks proto.InternalMessageInfo

func (m *NothingBooks) GetDummy() bool {
	if m != nil {
		return m.Dummy
	}
	return false
}

type WriterBookName struct {
	Writer               string   `protobuf:"bytes,1,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriterBookName) Reset()         { *m = WriterBookName{} }
func (m *WriterBookName) String() string { return proto.CompactTextString(m) }
func (*WriterBookName) ProtoMessage()    {}
func (*WriterBookName) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *WriterBookName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriterBookName.Unmarshal(m, b)
}
func (m *WriterBookName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriterBookName.Marshal(b, m, deterministic)
}
func (m *WriterBookName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriterBookName.Merge(m, src)
}
func (m *WriterBookName) XXX_Size() int {
	return xxx_messageInfo_WriterBookName.Size(m)
}
func (m *WriterBookName) XXX_DiscardUnknown() {
	xxx_messageInfo_WriterBookName.DiscardUnknown(m)
}

var xxx_messageInfo_WriterBookName proto.InternalMessageInfo

func (m *WriterBookName) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

func (m *WriterBookName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BookInsert struct {
	BookName             string   `protobuf:"bytes,1,opt,name=BookName,proto3" json:"BookName,omitempty"`
	AuthorName           string   `protobuf:"bytes,2,opt,name=AuthorName,proto3" json:"AuthorName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookInsert) Reset()         { *m = BookInsert{} }
func (m *BookInsert) String() string { return proto.CompactTextString(m) }
func (*BookInsert) ProtoMessage()    {}
func (*BookInsert) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{4}
}

func (m *BookInsert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookInsert.Unmarshal(m, b)
}
func (m *BookInsert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookInsert.Marshal(b, m, deterministic)
}
func (m *BookInsert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookInsert.Merge(m, src)
}
func (m *BookInsert) XXX_Size() int {
	return xxx_messageInfo_BookInsert.Size(m)
}
func (m *BookInsert) XXX_DiscardUnknown() {
	xxx_messageInfo_BookInsert.DiscardUnknown(m)
}

var xxx_messageInfo_BookInsert proto.InternalMessageInfo

func (m *BookInsert) GetBookName() string {
	if m != nil {
		return m.BookName
	}
	return ""
}

func (m *BookInsert) GetAuthorName() string {
	if m != nil {
		return m.AuthorName
	}
	return ""
}

type SomeID struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SomeID) Reset()         { *m = SomeID{} }
func (m *SomeID) String() string { return proto.CompactTextString(m) }
func (*SomeID) ProtoMessage()    {}
func (*SomeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{5}
}

func (m *SomeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SomeID.Unmarshal(m, b)
}
func (m *SomeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SomeID.Marshal(b, m, deterministic)
}
func (m *SomeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SomeID.Merge(m, src)
}
func (m *SomeID) XXX_Size() int {
	return xxx_messageInfo_SomeID.Size(m)
}
func (m *SomeID) XXX_DiscardUnknown() {
	xxx_messageInfo_SomeID.DiscardUnknown(m)
}

var xxx_messageInfo_SomeID proto.InternalMessageInfo

func (m *SomeID) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*Writer)(nil), "protocol.Writer")
	proto.RegisterType((*Book)(nil), "protocol.Book")
	proto.RegisterType((*NothingBooks)(nil), "protocol.NothingBooks")
	proto.RegisterType((*WriterBookName)(nil), "protocol.WriterBookName")
	proto.RegisterType((*BookInsert)(nil), "protocol.BookInsert")
	proto.RegisterType((*SomeID)(nil), "protocol.SomeID")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x8d, 0x43, 0x93, 0xa6, 0x07, 0x8a, 0xd0, 0x51, 0x55, 0x56, 0x06, 0x14, 0x59, 0x0c, 0x19,
	0x50, 0x85, 0xda, 0x81, 0x85, 0x25, 0x55, 0x06, 0xb2, 0x74, 0x30, 0x48, 0xcc, 0x40, 0x22, 0x5a,
	0x41, 0x62, 0xe4, 0xb8, 0x43, 0xff, 0x8c, 0xcf, 0x43, 0xb6, 0xeb, 0x36, 0x14, 0x06, 0x36, 0xdf,
	0xdd, 0x7b, 0xf7, 0xde, 0x3b, 0x43, 0xfc, 0x29, 0x85, 0x12, 0xaf, 0xe2, 0x63, 0x6a, 0x1e, 0x18,
	0xb9, 0x9a, 0x5d, 0x43, 0xf8, 0x24, 0xd7, 0xaa, 0x96, 0x18, 0x83, 0x5f, 0x16, 0x94, 0xa4, 0x24,
	0x0b, 0xb8, 0x5f, 0x16, 0x88, 0x30, 0x58, 0x3e, 0x37, 0x35, 0xf5, 0x53, 0x92, 0x8d, 0xb8, 0x79,
	0xb3, 0x47, 0x18, 0x2c, 0x84, 0x78, 0xff, 0x0f, 0x16, 0x33, 0x08, 0xf3, 0x8d, 0x5a, 0x09, 0x49,
	0x4f, 0x52, 0x92, 0x9d, 0xce, 0xce, 0xa7, 0x7b, 0x13, 0x56, 0x91, 0xef, 0xe6, 0xec, 0x0a, 0xce,
	0x96, 0x42, 0xad, 0xd6, 0xed, 0x9b, 0x5e, 0xde, 0xe1, 0x18, 0x82, 0x6a, 0xd3, 0x34, 0x5b, 0x23,
	0x10, 0x71, 0x5b, 0xb0, 0x3b, 0x88, 0x2d, 0x4f, 0x83, 0x8c, 0xc2, 0xc4, 0x79, 0x37, 0xc0, 0x11,
	0x77, 0x49, 0xfe, 0x72, 0x7e, 0x0f, 0xa0, 0x79, 0x65, 0xdb, 0xd5, 0x52, 0x61, 0x02, 0x91, 0xdb,
	0xb2, 0xe3, 0xee, 0x6b, 0xbc, 0x04, 0xb0, 0xbe, 0x7a, 0x3b, 0x7a, 0x1d, 0x46, 0x21, 0x7c, 0x10,
	0x4d, 0x5d, 0x16, 0xc7, 0x57, 0x98, 0x7d, 0x11, 0x08, 0x6c, 0x82, 0x5b, 0x18, 0x5a, 0x46, 0x87,
	0x93, 0x43, 0xec, 0x7e, 0xc8, 0xe4, 0xd7, 0x39, 0x98, 0x77, 0x43, 0x30, 0x87, 0x0b, 0x3d, 0x5e,
	0x6c, 0x2d, 0x3d, 0x6f, 0x2b, 0xe3, 0x89, 0x1e, 0x83, 0x9d, 0xdb, 0x24, 0x3e, 0x4c, 0x74, 0x8f,
	0x79, 0x38, 0x87, 0x61, 0x5e, 0x55, 0xe6, 0x9b, 0xc6, 0x3f, 0x87, 0x36, 0x7c, 0x5f, 0xd9, 0x06,
	0x61, 0xde, 0x4b, 0x68, 0x5a, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x62, 0x26, 0x1f,
	0x29, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BooksClient is the client API for Books service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BooksClient interface {
	Authors(ctx context.Context, in *NothingBooks, opts ...grpc.CallOption) (Books_AuthorsClient, error)
	BookByAuthorAndName(ctx context.Context, in *WriterBookName, opts ...grpc.CallOption) (*Book, error)
	AddBook(ctx context.Context, in *BookInsert, opts ...grpc.CallOption) (*SomeID, error)
}

type booksClient struct {
	cc *grpc.ClientConn
}

func NewBooksClient(cc *grpc.ClientConn) BooksClient {
	return &booksClient{cc}
}

func (c *booksClient) Authors(ctx context.Context, in *NothingBooks, opts ...grpc.CallOption) (Books_AuthorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Books_serviceDesc.Streams[0], "/protocol.Books/Authors", opts...)
	if err != nil {
		return nil, err
	}
	x := &booksAuthorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Books_AuthorsClient interface {
	Recv() (*Writer, error)
	grpc.ClientStream
}

type booksAuthorsClient struct {
	grpc.ClientStream
}

func (x *booksAuthorsClient) Recv() (*Writer, error) {
	m := new(Writer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *booksClient) BookByAuthorAndName(ctx context.Context, in *WriterBookName, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/protocol.Books/BookByAuthorAndName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksClient) AddBook(ctx context.Context, in *BookInsert, opts ...grpc.CallOption) (*SomeID, error) {
	out := new(SomeID)
	err := c.cc.Invoke(ctx, "/protocol.Books/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServer is the server API for Books service.
type BooksServer interface {
	Authors(*NothingBooks, Books_AuthorsServer) error
	BookByAuthorAndName(context.Context, *WriterBookName) (*Book, error)
	AddBook(context.Context, *BookInsert) (*SomeID, error)
}

func RegisterBooksServer(s *grpc.Server, srv BooksServer) {
	s.RegisterService(&_Books_serviceDesc, srv)
}

func _Books_Authors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NothingBooks)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BooksServer).Authors(m, &booksAuthorsServer{stream})
}

type Books_AuthorsServer interface {
	Send(*Writer) error
	grpc.ServerStream
}

type booksAuthorsServer struct {
	grpc.ServerStream
}

func (x *booksAuthorsServer) Send(m *Writer) error {
	return x.ServerStream.SendMsg(m)
}

func _Books_BookByAuthorAndName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriterBookName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).BookByAuthorAndName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Books/BookByAuthorAndName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).BookByAuthorAndName(ctx, req.(*WriterBookName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Books_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookInsert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Books/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).AddBook(ctx, req.(*BookInsert))
	}
	return interceptor(ctx, in, info, handler)
}

var _Books_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Books",
	HandlerType: (*BooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookByAuthorAndName",
			Handler:    _Books_BookByAuthorAndName_Handler,
		},
		{
			MethodName: "AddBook",
			Handler:    _Books_AddBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Authors",
			Handler:       _Books_Authors_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol.proto",
}
