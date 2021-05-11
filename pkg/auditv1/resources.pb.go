// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indent/audit/v1/resources.proto

package auditv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Kind of Target the Event is being written to.
type Target_Kind int32

const (
	Target_INPUT  Target_Kind = 0
	Target_OUTPUT Target_Kind = 1
)

var Target_Kind_name = map[int32]string{
	0: "INPUT",
	1: "OUTPUT",
}

var Target_Kind_value = map[string]int32{
	"INPUT":  0,
	"OUTPUT": 1,
}

func (x Target_Kind) String() string {
	return proto.EnumName(Target_Kind_name, int32(x))
}

func (Target_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3bf215d7d77adf9c, []int{3, 0}
}

// Event contains information captured at a specific moment of time, typically relating to an authorization decision or instance of access.
type Event struct {
	Event                string                 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Meta                 *Meta                  `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	Reason               string                 `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id                   string                 `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
	ExternalId           string                 `protobuf:"bytes,13,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	SessionId            string                 `protobuf:"bytes,20,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Actor                *Resource              `protobuf:"bytes,27,opt,name=actor,proto3" json:"actor,omitempty"`
	Resources            []*Resource            `protobuf:"bytes,30,rep,name=resources,proto3" json:"resources,omitempty"`
	XOriginal            []byte                 `protobuf:"bytes,35,opt,name=_original,json=Original,proto3" json:"_original,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bf215d7d77adf9c, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Event) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Event) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Event) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Event) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Event) GetActor() *Resource {
	if m != nil {
		return m.Actor
	}
	return nil
}

func (m *Event) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Event) GetXOriginal() []byte {
	if m != nil {
		return m.XOriginal
	}
	return nil
}

// Meta contains metadata about an Event.
type Meta struct {
	Labels               map[string]string `protobuf:"bytes,20,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bf215d7d77adf9c, []int{1}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Resource related to the Event.
type Resource struct {
	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	AltIds      []string `protobuf:"bytes,3,rep,name=alt_ids,json=altIds,proto3" json:"alt_ids,omitempty"`
	Kind        string   `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Email       string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// Arbitrary string-addressable labels
	Labels               map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bf215d7d77adf9c, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Resource) GetAltIds() []string {
	if m != nil {
		return m.AltIds
	}
	return nil
}

func (m *Resource) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Resource) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Resource) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Target is the intended recipient of the contained Events.
type Target struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 Target_Kind `protobuf:"varint,2,opt,name=kind,proto3,enum=indent.audit.v1.Target_Kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bf215d7d77adf9c, []int{3}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Target) GetKind() Target_Kind {
	if m != nil {
		return m.Kind
	}
	return Target_INPUT
}

func init() {
	proto.RegisterEnum("indent.audit.v1.Target_Kind", Target_Kind_name, Target_Kind_value)
	proto.RegisterType((*Event)(nil), "indent.audit.v1.Event")
	proto.RegisterType((*Meta)(nil), "indent.audit.v1.Meta")
	proto.RegisterMapType((map[string]string)(nil), "indent.audit.v1.Meta.LabelsEntry")
	proto.RegisterType((*Resource)(nil), "indent.audit.v1.Resource")
	proto.RegisterMapType((map[string]string)(nil), "indent.audit.v1.Resource.LabelsEntry")
	proto.RegisterType((*Target)(nil), "indent.audit.v1.Target")
}

func init() {
	proto.RegisterFile("indent/audit/v1/resources.proto", fileDescriptor_3bf215d7d77adf9c)
}

var fileDescriptor_3bf215d7d77adf9c = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x76, 0x92, 0xd6, 0x93, 0x92, 0x56, 0x4b, 0x80, 0x25, 0xa5, 0x24, 0x0d, 0x42, 0x0a,
	0x17, 0x9b, 0x94, 0x03, 0x2d, 0x12, 0x87, 0x56, 0xea, 0xc1, 0x02, 0xda, 0xc8, 0x4a, 0x2b, 0x54,
	0x55, 0x8a, 0x36, 0xd9, 0x25, 0x5a, 0xd5, 0x1f, 0x91, 0x77, 0x13, 0x11, 0x89, 0x5f, 0xc3, 0x91,
	0x9f, 0xc2, 0xbf, 0xe1, 0xca, 0x09, 0xed, 0xd8, 0x4e, 0x51, 0x0b, 0x9c, 0x7a, 0x9b, 0x8f, 0xf7,
	0xc6, 0xf3, 0xde, 0xac, 0xa1, 0x2d, 0x13, 0x2e, 0x12, 0xed, 0xb3, 0x39, 0x97, 0xda, 0x5f, 0xf4,
	0xfd, 0x4c, 0xa8, 0x74, 0x9e, 0x4d, 0x84, 0xf2, 0x66, 0x59, 0xaa, 0x53, 0xb2, 0x99, 0x03, 0x3c,
	0x04, 0x78, 0x8b, 0x7e, 0xab, 0x3d, 0x4d, 0xd3, 0x69, 0x24, 0x7c, 0x6c, 0x8f, 0xe7, 0x9f, 0x7d,
	0x2d, 0x63, 0xa1, 0x34, 0x8b, 0x67, 0x39, 0xa3, 0xfb, 0xd3, 0x86, 0xea, 0xf1, 0x42, 0x24, 0x9a,
	0x34, 0xa1, 0x2a, 0x4c, 0x40, 0xad, 0x8e, 0xd5, 0x73, 0xc3, 0x3c, 0x21, 0x2f, 0xa1, 0x12, 0x0b,
	0xcd, 0xa8, 0xdd, 0xb1, 0x7a, 0xf5, 0xbd, 0x87, 0xde, 0x8d, 0x0f, 0x78, 0x1f, 0x85, 0x66, 0x21,
	0x42, 0xc8, 0x23, 0xa8, 0x65, 0x82, 0xa9, 0x34, 0xa1, 0x0e, 0x4e, 0x28, 0x32, 0xb2, 0x0f, 0xee,
	0xea, 0xab, 0xb4, 0x8a, 0x73, 0x5a, 0x5e, 0xbe, 0x97, 0x57, 0xee, 0xe5, 0x0d, 0x4b, 0x44, 0x78,
	0x0d, 0x26, 0x0d, 0xb0, 0x25, 0xa7, 0x2e, 0x4e, 0xb3, 0x25, 0x27, 0x6d, 0xa8, 0x8b, 0x2f, 0x5a,
	0x64, 0x09, 0x8b, 0x46, 0x92, 0xd3, 0xfb, 0xd8, 0x80, 0xb2, 0x14, 0x70, 0xb2, 0x03, 0xa0, 0x84,
	0x52, 0x32, 0x4d, 0x4c, 0xbf, 0x89, 0x7d, 0xb7, 0xa8, 0x04, 0x9c, 0xf8, 0x50, 0x65, 0x13, 0x9d,
	0x66, 0x74, 0x1b, 0xb7, 0x78, 0x72, 0x4b, 0x4d, 0x58, 0xf8, 0x19, 0xe6, 0x38, 0xf2, 0x06, 0xdc,
	0x95, 0xc5, 0xf4, 0x59, 0xc7, 0xf9, 0x3f, 0xe9, 0x1a, 0x4b, 0xb6, 0xc1, 0x1d, 0xa5, 0x99, 0x9c,
	0xca, 0x84, 0x45, 0xf4, 0x79, 0xc7, 0xea, 0x6d, 0x84, 0xeb, 0xa7, 0x45, 0xde, 0xfd, 0x0a, 0x15,
	0x63, 0x1b, 0x39, 0x80, 0x5a, 0xc4, 0xc6, 0x22, 0x52, 0xb4, 0x89, 0xa3, 0x77, 0xff, 0xea, 0xae,
	0xf7, 0x01, 0x31, 0xc7, 0x89, 0xce, 0x96, 0x61, 0x41, 0x68, 0x1d, 0x40, 0xfd, 0x8f, 0x32, 0xd9,
	0x02, 0xe7, 0x4a, 0x2c, 0x8b, 0xcb, 0x99, 0xd0, 0x5c, 0x73, 0xc1, 0xa2, 0xb9, 0xc0, 0xc3, 0xb9,
	0x61, 0x9e, 0xbc, 0xb5, 0xf7, 0xad, 0xee, 0x2f, 0x0b, 0xd6, 0xcb, 0x95, 0x0b, 0x87, 0xad, 0x95,
	0xc3, 0xbb, 0xb0, 0xc1, 0xa5, 0x9a, 0x45, 0x6c, 0x39, 0x4a, 0x58, 0x5c, 0xb2, 0xeb, 0x45, 0xed,
	0x84, 0xc5, 0x82, 0x3c, 0x86, 0x35, 0x16, 0xe9, 0x91, 0xe4, 0x8a, 0x3a, 0x1d, 0xc7, 0xdc, 0x99,
	0x45, 0x3a, 0xe0, 0x8a, 0x10, 0xa8, 0x5c, 0xc9, 0x84, 0xd3, 0x0a, 0x72, 0x30, 0xc6, 0x47, 0x15,
	0x33, 0x19, 0xe1, 0xdd, 0xcd, 0xa3, 0x32, 0x09, 0x79, 0xb7, 0x12, 0x0e, 0x28, 0xfc, 0xc5, 0x3f,
	0x3d, 0xbd, 0x6b, 0xf1, 0x31, 0xd4, 0x86, 0x2c, 0x9b, 0x0a, 0x6d, 0xb6, 0x45, 0x85, 0x39, 0x0d,
	0x63, 0xf2, 0xaa, 0x50, 0x60, 0x68, 0x8d, 0xbd, 0xa7, 0xb7, 0xb6, 0xca, 0xa9, 0xde, 0x7b, 0x99,
	0xf0, 0x5c, 0x5f, 0x77, 0x07, 0x2a, 0x26, 0x23, 0x2e, 0x54, 0x83, 0x93, 0xc1, 0xd9, 0x70, 0xeb,
	0x1e, 0x01, 0xa8, 0x9d, 0x9e, 0x0d, 0x4d, 0x6c, 0x1d, 0x5d, 0xc0, 0x83, 0x49, 0x1a, 0xdf, 0x9c,
	0x73, 0xd4, 0x28, 0xe5, 0xa9, 0x81, 0x79, 0xff, 0x03, 0xeb, 0x62, 0x0d, 0x7b, 0x8b, 0xfe, 0x37,
	0xdb, 0x09, 0x0e, 0x3f, 0x7d, 0xb7, 0x37, 0x83, 0x9c, 0x72, 0x88, 0x94, 0xf3, 0xfe, 0x8f, 0xb2,
	0x72, 0x89, 0x95, 0xcb, 0xf3, 0xfe, 0xb8, 0x86, 0xff, 0xce, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0xb0, 0xa7, 0xa1, 0x15, 0x04, 0x00, 0x00,
}
