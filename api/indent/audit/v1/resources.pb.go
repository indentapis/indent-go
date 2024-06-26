// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indent/audit/v1/resources.proto

package auditv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "go.indent.com/indent-go/api/indent/log"
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
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Space                string                 `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	DisplayName          string                 `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Labels               map[string]string      `protobuf:"bytes,20,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreateTime           *timestamppb.Timestamp `protobuf:"bytes,30,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamppb.Timestamp `protobuf:"bytes,31,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime           *timestamppb.Timestamp `protobuf:"bytes,32,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	ExpireTime           *timestamppb.Timestamp `protobuf:"bytes,35,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	StartTime            *timestamppb.Timestamp `protobuf:"bytes,40,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamppb.Timestamp `protobuf:"bytes,45,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
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

func (m *Meta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Meta) GetSpace() string {
	if m != nil {
		return m.Space
	}
	return ""
}

func (m *Meta) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Meta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Meta) GetCreateTime() *timestamppb.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Meta) GetUpdateTime() *timestamppb.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Meta) GetDeleteTime() *timestamppb.Timestamp {
	if m != nil {
		return m.DeleteTime
	}
	return nil
}

func (m *Meta) GetExpireTime() *timestamppb.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func (m *Meta) GetStartTime() *timestamppb.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Meta) GetEndTime() *timestamppb.Timestamp {
	if m != nil {
		return m.EndTime
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
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0xc7, 0x7f, 0xbb, 0xd9, 0x4d, 0xbb, 0x93, 0xfe, 0xda, 0xca, 0x04, 0xd5, 0xa4, 0x7f, 0x92,
	0xa6, 0x42, 0x84, 0x03, 0x1b, 0x52, 0x84, 0xa0, 0x54, 0x42, 0x6a, 0x51, 0x0f, 0x11, 0xd0, 0x56,
	0xab, 0xb4, 0x42, 0xa8, 0x52, 0xe4, 0xc6, 0x26, 0x32, 0xdd, 0x3f, 0xd1, 0xda, 0x89, 0xda, 0x33,
	0x2f, 0xc1, 0x99, 0x23, 0x17, 0xde, 0x81, 0x23, 0x4f, 0xc0, 0xe3, 0x20, 0xdb, 0xbb, 0x69, 0x48,
	0x0b, 0x39, 0x70, 0x58, 0xc9, 0x9e, 0xf9, 0x7e, 0x66, 0xc7, 0x33, 0x63, 0x43, 0x95, 0xc7, 0x94,
	0xc5, 0xb2, 0x49, 0x86, 0x94, 0xcb, 0xe6, 0xa8, 0xd5, 0x4c, 0x99, 0x48, 0x86, 0x69, 0x8f, 0x09,
	0x7f, 0x90, 0x26, 0x32, 0x41, 0x4b, 0x46, 0xe0, 0x6b, 0x81, 0x3f, 0x6a, 0x55, 0xaa, 0xfd, 0x24,
	0xe9, 0x87, 0xac, 0xa9, 0xdd, 0xe7, 0xc3, 0x0f, 0x4d, 0xc9, 0x23, 0x26, 0x24, 0x89, 0x06, 0x86,
	0xa8, 0x94, 0xb3, 0x90, 0x61, 0xd2, 0x57, 0x9f, 0xb1, 0xd6, 0x3f, 0x15, 0xc0, 0x3d, 0x18, 0xb1,
	0x58, 0xa2, 0x0a, 0xb8, 0x4c, 0x2d, 0xb0, 0x55, 0xb3, 0x1a, 0xde, 0xbe, 0xf3, 0xf3, 0x7b, 0xd5,
	0x0a, 0x8c, 0x09, 0x3d, 0x04, 0x27, 0x62, 0x92, 0x60, 0xbb, 0x66, 0x35, 0x4a, 0xdb, 0x77, 0xfd,
	0xa9, 0x9f, 0xfb, 0x6f, 0x99, 0x24, 0x81, 0x96, 0xa0, 0x35, 0x28, 0xa6, 0x8c, 0x88, 0x24, 0xc6,
	0x85, 0x89, 0x38, 0x99, 0x0d, 0xbd, 0x04, 0x6f, 0x9c, 0x17, 0x76, 0x75, 0xb4, 0x8a, 0x6f, 0x32,
	0xf7, 0xf3, 0xcc, 0xfd, 0x4e, 0xae, 0xc8, 0xe0, 0x6b, 0x04, 0x2d, 0x82, 0xcd, 0x29, 0xf6, 0x54,
	0xe4, 0xc0, 0xe6, 0x14, 0x55, 0xa1, 0xc4, 0x2e, 0x25, 0x4b, 0x63, 0x12, 0x76, 0x39, 0xc5, 0xff,
	0x6b, 0x07, 0xe4, 0xa6, 0x36, 0x45, 0xeb, 0x00, 0x82, 0x09, 0xc1, 0x93, 0x58, 0xf9, 0xcb, 0xda,
	0xef, 0x65, 0x96, 0x36, 0x45, 0x4d, 0x70, 0x49, 0x4f, 0x26, 0x29, 0x5e, 0xd5, 0xb9, 0xdc, 0xbb,
	0x71, 0xb2, 0x20, 0xab, 0x7b, 0x60, 0x74, 0xe8, 0x19, 0x78, 0xe3, 0x56, 0xe0, 0x8d, 0x5a, 0xe1,
	0xef, 0xd0, 0xb5, 0x16, 0xad, 0x82, 0xd7, 0x4d, 0x52, 0xde, 0xe7, 0x31, 0x09, 0xf1, 0x56, 0xcd,
	0x6a, 0x2c, 0x04, 0xf3, 0x47, 0xd9, 0xbe, 0xfe, 0xcd, 0x01, 0x47, 0xd5, 0x10, 0x21, 0x70, 0x62,
	0x12, 0x31, 0xd3, 0x83, 0x40, 0xaf, 0x51, 0x19, 0x5c, 0x31, 0x20, 0x3d, 0xa6, 0xab, 0xef, 0x05,
	0x66, 0x83, 0x36, 0x61, 0x81, 0x72, 0x31, 0x08, 0xc9, 0x55, 0x57, 0x13, 0xae, 0x76, 0x96, 0x32,
	0xdb, 0xa1, 0x02, 0x77, 0xa0, 0x18, 0x92, 0x73, 0x16, 0x0a, 0x5c, 0xd6, 0x89, 0x6e, 0xde, 0xda,
	0x37, 0xff, 0x8d, 0xd6, 0x1c, 0xc4, 0x32, 0xbd, 0x0a, 0x32, 0x00, 0xed, 0x42, 0xa9, 0x97, 0x32,
	0x22, 0x59, 0x57, 0xd5, 0x1e, 0x6f, 0xcc, 0xea, 0x54, 0x00, 0x46, 0xae, 0x0c, 0x0a, 0x1e, 0x0e,
	0xe8, 0x18, 0xae, 0xce, 0x86, 0x8d, 0x3c, 0x87, 0x29, 0x0b, 0x59, 0x0e, 0xd7, 0x66, 0xc3, 0x46,
	0x9e, 0xc3, 0xec, 0x72, 0xc0, 0xd3, 0x0c, 0xde, 0x9a, 0x0d, 0x1b, 0xb9, 0x86, 0x77, 0x00, 0x84,
	0x24, 0xa9, 0x34, 0x6c, 0x63, 0x26, 0xeb, 0x69, 0xb5, 0x46, 0x9f, 0xc2, 0x3c, 0x8b, 0xa9, 0x01,
	0x1f, 0xcd, 0x04, 0xe7, 0x58, 0x4c, 0xd5, 0xae, 0xb2, 0x03, 0xa5, 0x89, 0xe2, 0xa3, 0x65, 0x28,
	0x5c, 0xb0, 0xab, 0xac, 0xf7, 0x6a, 0xa9, 0x5a, 0x3f, 0x22, 0xe1, 0x70, 0xdc, 0x7a, 0xbd, 0x79,
	0x61, 0x3f, 0xb7, 0xea, 0x9f, 0x6d, 0x98, 0xcf, 0xc7, 0x0c, 0x95, 0xf5, 0xad, 0x98, 0xbc, 0xb7,
	0xea, 0x6e, 0x3c, 0x98, 0x9a, 0x10, 0x7b, 0xc2, 0xff, 0xdb, 0x9c, 0xac, 0xc0, 0x1c, 0x09, 0x65,
	0x97, 0x53, 0x81, 0x0b, 0xb5, 0x42, 0xc3, 0x0b, 0x8a, 0x24, 0x94, 0x6d, 0x2a, 0x10, 0x06, 0xe7,
	0x82, 0xc7, 0x14, 0x3b, 0x13, 0xa4, 0xb6, 0xe8, 0xc7, 0x22, 0x22, 0x3c, 0x34, 0x63, 0x37, 0x7e,
	0x2c, 0x94, 0x09, 0xbd, 0x1a, 0x8f, 0x1d, 0xe8, 0xb1, 0xbb, 0xff, 0xc7, 0xfb, 0x31, 0x39, 0x7a,
	0xf9, 0x43, 0x61, 0xd0, 0x7f, 0x29, 0x4d, 0x04, 0xc5, 0x0e, 0x49, 0xfb, 0x4c, 0xde, 0x7a, 0x9b,
	0x1e, 0x67, 0x67, 0x52, 0xd8, 0xe2, 0xf6, 0xda, 0x8d, 0xdc, 0x0c, 0xea, 0xbf, 0xe6, 0x31, 0x35,
	0x67, 0xad, 0xaf, 0x83, 0xa3, 0x76, 0xc8, 0x03, 0xb7, 0x7d, 0x78, 0x7c, 0xd2, 0x59, 0xfe, 0x0f,
	0x01, 0x14, 0x8f, 0x4e, 0x3a, 0x6a, 0x6d, 0xed, 0x7f, 0x84, 0x3b, 0xbd, 0x24, 0x9a, 0x8e, 0xb3,
	0xbf, 0x98, 0x1f, 0x52, 0x1c, 0xab, 0x09, 0x38, 0xb6, 0xde, 0xaf, 0x4c, 0xbd, 0xe9, 0xbb, 0x7a,
	0x31, 0x6a, 0x7d, 0xb1, 0x0b, 0xed, 0xbd, 0x77, 0x5f, 0xed, 0xa5, 0xb6, 0x09, 0xb1, 0xa7, 0x43,
	0x9c, 0xb6, 0x7e, 0xe4, 0x96, 0x33, 0x6d, 0x39, 0x3b, 0x6d, 0x9d, 0x17, 0xf5, 0x34, 0x3d, 0xf9,
	0x15, 0x00, 0x00, 0xff, 0xff, 0xff, 0x03, 0x47, 0x46, 0x1f, 0x06, 0x00, 0x00,
}
