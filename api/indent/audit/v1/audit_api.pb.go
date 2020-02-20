// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indent/audit/v1/audit_api.proto

package auditv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type WriteRequest struct {
	// Target being written to.
	Target *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Async returns immediately and doesn't wait for events to be written.
	Async bool `protobuf:"varint,5,opt,name=async,proto3" json:"async,omitempty"`
	// Events being written.
	Events               []*Event `protobuf:"bytes,20,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff0992820003f443, []int{0}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *WriteRequest) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

func (m *WriteRequest) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type WriteEventRequest struct {
	// Name of space containing provider.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Name of provider containing Input.
	ProviderName string `protobuf:"bytes,2,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	// Name of Input the event is written to.
	InputName string `protobuf:"bytes,3,opt,name=input_name,json=inputName,proto3" json:"input_name,omitempty"`
	// Event being written.
	Event                *Event   `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteEventRequest) Reset()         { *m = WriteEventRequest{} }
func (m *WriteEventRequest) String() string { return proto.CompactTextString(m) }
func (*WriteEventRequest) ProtoMessage()    {}
func (*WriteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff0992820003f443, []int{1}
}

func (m *WriteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteEventRequest.Unmarshal(m, b)
}
func (m *WriteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteEventRequest.Marshal(b, m, deterministic)
}
func (m *WriteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteEventRequest.Merge(m, src)
}
func (m *WriteEventRequest) XXX_Size() int {
	return xxx_messageInfo_WriteEventRequest.Size(m)
}
func (m *WriteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteEventRequest proto.InternalMessageInfo

func (m *WriteEventRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *WriteEventRequest) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

func (m *WriteEventRequest) GetInputName() string {
	if m != nil {
		return m.InputName
	}
	return ""
}

func (m *WriteEventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type WriteBatchRequest struct {
	// Name of space containing provider.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Name of provider containing Input.
	ProviderName string `protobuf:"bytes,2,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	// Name of Input the event is written to.
	InputName string `protobuf:"bytes,3,opt,name=input_name,json=inputName,proto3" json:"input_name,omitempty"`
	// Async returns immediately and doesn't wait for events to be written.
	Async bool `protobuf:"varint,4,opt,name=async,proto3" json:"async,omitempty"`
	// Events being written.
	Events               []*Event `protobuf:"bytes,20,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteBatchRequest) Reset()         { *m = WriteBatchRequest{} }
func (m *WriteBatchRequest) String() string { return proto.CompactTextString(m) }
func (*WriteBatchRequest) ProtoMessage()    {}
func (*WriteBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff0992820003f443, []int{2}
}

func (m *WriteBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteBatchRequest.Unmarshal(m, b)
}
func (m *WriteBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteBatchRequest.Marshal(b, m, deterministic)
}
func (m *WriteBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteBatchRequest.Merge(m, src)
}
func (m *WriteBatchRequest) XXX_Size() int {
	return xxx_messageInfo_WriteBatchRequest.Size(m)
}
func (m *WriteBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteBatchRequest proto.InternalMessageInfo

func (m *WriteBatchRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *WriteBatchRequest) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

func (m *WriteBatchRequest) GetInputName() string {
	if m != nil {
		return m.InputName
	}
	return ""
}

func (m *WriteBatchRequest) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

func (m *WriteBatchRequest) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteRequest)(nil), "indent.audit.v1.WriteRequest")
	proto.RegisterType((*WriteEventRequest)(nil), "indent.audit.v1.WriteEventRequest")
	proto.RegisterType((*WriteBatchRequest)(nil), "indent.audit.v1.WriteBatchRequest")
}

func init() { proto.RegisterFile("indent/audit/v1/audit_api.proto", fileDescriptor_ff0992820003f443) }

var fileDescriptor_ff0992820003f443 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0x35, 0x49, 0x13, 0xda, 0x69, 0xab, 0xaa, 0x43, 0x15, 0xa2, 0xd0, 0x88, 0xc8, 0x6c,
	0xa2, 0x0a, 0xcd, 0xc8, 0x61, 0xe7, 0x0d, 0x4a, 0xa0, 0x8b, 0xb0, 0x40, 0x91, 0x85, 0xca, 0x1f,
	0x45, 0xaa, 0xa6, 0xce, 0x10, 0x46, 0x22, 0x33, 0xc6, 0x1e, 0x1b, 0x55, 0x51, 0x36, 0x48, 0x9c,
	0xa0, 0x70, 0x01, 0xc4, 0x8a, 0x33, 0x70, 0x02, 0xb6, 0x5c, 0x81, 0x83, 0x20, 0xbf, 0xb1, 0x49,
	0x48, 0xf0, 0x06, 0xa4, 0xee, 0xec, 0xf7, 0x7d, 0x33, 0xdf, 0x6f, 0xde, 0x7b, 0xf8, 0x8e, 0x54,
	0x13, 0xa1, 0x0c, 0xe3, 0xc9, 0x44, 0x1a, 0x96, 0xba, 0xf6, 0xe3, 0x9c, 0x87, 0x92, 0x86, 0x91,
	0x36, 0x9a, 0x1c, 0x58, 0x03, 0x85, 0x3a, 0x4d, 0xdd, 0xd6, 0xf1, 0x54, 0xeb, 0xe9, 0x1b, 0xc1,
	0x78, 0x28, 0x19, 0x57, 0x4a, 0x1b, 0x6e, 0xa4, 0x56, 0xb1, 0xb5, 0xb7, 0x6e, 0xe7, 0x2a, 0xfc,
	0x5d, 0x24, 0xaf, 0x98, 0x98, 0x85, 0xe6, 0x32, 0x17, 0x37, 0xc2, 0x22, 0x11, 0xeb, 0x24, 0x0a,
	0x44, 0x7e, 0xda, 0xf9, 0x80, 0xf0, 0xde, 0xb3, 0x48, 0x1a, 0xe1, 0x8b, 0xb7, 0x89, 0x88, 0x0d,
	0x61, 0xb8, 0x6e, 0x78, 0x34, 0x15, 0xa6, 0x89, 0x3a, 0xa8, 0xbb, 0xdb, 0xbb, 0x45, 0xd7, 0x70,
	0xe8, 0x53, 0x90, 0xfd, 0xdc, 0x46, 0x8e, 0x70, 0x8d, 0xc7, 0x97, 0x2a, 0x68, 0xd6, 0x3a, 0xa8,
	0xbb, 0xed, 0xdb, 0x1f, 0x42, 0x71, 0x5d, 0xa4, 0x42, 0x99, 0xb8, 0x79, 0xd4, 0xa9, 0x76, 0x77,
	0x7b, 0x8d, 0x8d, 0x6b, 0x4e, 0x33, 0xd9, 0xcf, 0x5d, 0xce, 0x17, 0x84, 0x0f, 0x81, 0xc3, 0x96,
	0x73, 0x98, 0x36, 0xc6, 0x71, 0xc8, 0x03, 0x71, 0xae, 0xf8, 0x4c, 0x00, 0xd0, 0x8e, 0xbf, 0x03,
	0x95, 0x27, 0x7c, 0x26, 0xc8, 0x5d, 0xbc, 0x1f, 0x46, 0x3a, 0x95, 0x13, 0x11, 0x59, 0x47, 0x05,
	0x1c, 0x7b, 0x45, 0x11, 0x4c, 0x6d, 0x8c, 0xa5, 0x0a, 0x13, 0x63, 0x1d, 0x55, 0x7b, 0x07, 0x54,
	0x40, 0xbe, 0x87, 0x6b, 0x80, 0xd0, 0xdc, 0x82, 0xe7, 0x96, 0x71, 0x5a, 0x93, 0xf3, 0xad, 0xc0,
	0x1c, 0x70, 0x13, 0xbc, 0xbe, 0x46, 0xcc, 0xdf, 0x5d, 0xde, 0xfa, 0x8f, 0x2e, 0xf7, 0x3e, 0x55,
	0xf1, 0x76, 0x3f, 0x93, 0xfa, 0xa3, 0x21, 0x19, 0xe3, 0x1a, 0x3c, 0x85, 0xb4, 0x37, 0x4e, 0xad,
	0x6e, 0x44, 0xab, 0x41, 0xed, 0x86, 0xd1, 0x62, 0xc3, 0xe8, 0x69, 0xb6, 0x61, 0xce, 0xf1, 0xfb,
	0x1f, 0x3f, 0xaf, 0x2a, 0x0d, 0xe7, 0x30, 0xdb, 0xab, 0xb9, 0xdd, 0x06, 0x9a, 0xd1, 0x2f, 0x3c,
	0x74, 0x42, 0xae, 0x10, 0xc6, 0xcb, 0x81, 0x12, 0xe7, 0xef, 0x19, 0xab, 0xd3, 0x2e, 0x0d, 0x1a,
	0x42, 0xd0, 0x43, 0xc7, 0xcb, 0x82, 0xa0, 0x25, 0x31, 0x9b, 0x2f, 0x1b, 0xbe, 0x60, 0xf3, 0x3f,
	0xda, 0xbb, 0x60, 0xf3, 0x65, 0x27, 0x17, 0xde, 0xbb, 0x2c, 0xc7, 0xb3, 0xf3, 0x23, 0x1f, 0x0b,
	0x2a, 0x98, 0x5f, 0x19, 0xd5, 0xea, 0x70, 0x4b, 0xa9, 0x1e, 0x03, 0xd5, 0x23, 0xe7, 0xc1, 0xbf,
	0x53, 0x41, 0x8e, 0x87, 0x4e, 0x06, 0x2f, 0xf0, 0xcd, 0x40, 0xcf, 0xd6, 0x61, 0x06, 0xfb, 0x76,
	0x56, 0xa1, 0x1c, 0x65, 0xd1, 0x23, 0xf4, 0xf2, 0x06, 0x48, 0xa9, 0xfb, 0xb9, 0x52, 0x1d, 0xf6,
	0x9f, 0x7f, 0xad, 0x1c, 0x0c, 0xed, 0x09, 0xf0, 0xd1, 0x33, 0xf7, 0x7b, 0x51, 0x19, 0x43, 0x65,
	0x7c, 0xe6, 0x5e, 0xd4, 0x01, 0xfb, 0xfe, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xfa, 0xc1,
	0x67, 0x77, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuditAPIClient is the client API for AuditAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditAPIClient interface {
	// Write a new Event to a Target.
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// WriteEvent a new Event into Indent.
	//
	// Deprecated: Use Write.
	WriteEvent(ctx context.Context, in *WriteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// WriteBatch writes multiple Events into Indent.
	//
	// Deprecated: Use Write.
	WriteBatch(ctx context.Context, in *WriteBatchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type auditAPIClient struct {
	cc *grpc.ClientConn
}

func NewAuditAPIClient(cc *grpc.ClientConn) AuditAPIClient {
	return &auditAPIClient{cc}
}

func (c *auditAPIClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/indent.audit.v1.AuditAPI/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditAPIClient) WriteEvent(ctx context.Context, in *WriteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/indent.audit.v1.AuditAPI/WriteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditAPIClient) WriteBatch(ctx context.Context, in *WriteBatchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/indent.audit.v1.AuditAPI/WriteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditAPIServer is the server API for AuditAPI service.
type AuditAPIServer interface {
	// Write a new Event to a Target.
	Write(context.Context, *WriteRequest) (*empty.Empty, error)
	// WriteEvent a new Event into Indent.
	//
	// Deprecated: Use Write.
	WriteEvent(context.Context, *WriteEventRequest) (*empty.Empty, error)
	// WriteBatch writes multiple Events into Indent.
	//
	// Deprecated: Use Write.
	WriteBatch(context.Context, *WriteBatchRequest) (*empty.Empty, error)
}

func RegisterAuditAPIServer(s *grpc.Server, srv AuditAPIServer) {
	s.RegisterService(&_AuditAPI_serviceDesc, srv)
}

func _AuditAPI_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditAPIServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.audit.v1.AuditAPI/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditAPIServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditAPI_WriteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditAPIServer).WriteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.audit.v1.AuditAPI/WriteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditAPIServer).WriteEvent(ctx, req.(*WriteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditAPI_WriteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditAPIServer).WriteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.audit.v1.AuditAPI/WriteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditAPIServer).WriteBatch(ctx, req.(*WriteBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuditAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "indent.audit.v1.AuditAPI",
	HandlerType: (*AuditAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _AuditAPI_Write_Handler,
		},
		{
			MethodName: "WriteEvent",
			Handler:    _AuditAPI_WriteEvent_Handler,
		},
		{
			MethodName: "WriteBatch",
			Handler:    _AuditAPI_WriteBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indent/audit/v1/audit_api.proto",
}
