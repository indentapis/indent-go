// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indent/v1/task_api.proto

package indentv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "go.indent.com/indent-go/api/indent/audit/v1"
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

// Phase of the Task.
type Task_Phase int32

const (
	Task_UNKNOWN Task_Phase = 0
	Task_OPEN    Task_Phase = 1
	Task_CLOSED  Task_Phase = 2
	Task_DONE    Task_Phase = 3
)

var Task_Phase_name = map[int32]string{
	0: "UNKNOWN",
	1: "OPEN",
	2: "CLOSED",
	3: "DONE",
}

var Task_Phase_value = map[string]int32{
	"UNKNOWN": 0,
	"OPEN":    1,
	"CLOSED":  2,
	"DONE":    3,
}

func (x Task_Phase) String() string {
	return proto.EnumName(Task_Phase_name, int32(x))
}

func (Task_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{7, 0}
}

type ListTasksRequest struct {
	// Space of Tasks to retrieve.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Max number of Tasks to be returned.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Among constrains listed tasks by name.
	Among                []string `protobuf:"bytes,4,rep,name=among,proto3" json:"among,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTasksRequest) Reset()         { *m = ListTasksRequest{} }
func (m *ListTasksRequest) String() string { return proto.CompactTextString(m) }
func (*ListTasksRequest) ProtoMessage()    {}
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{0}
}

func (m *ListTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTasksRequest.Unmarshal(m, b)
}
func (m *ListTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTasksRequest.Marshal(b, m, deterministic)
}
func (m *ListTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTasksRequest.Merge(m, src)
}
func (m *ListTasksRequest) XXX_Size() int {
	return xxx_messageInfo_ListTasksRequest.Size(m)
}
func (m *ListTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTasksRequest proto.InternalMessageInfo

func (m *ListTasksRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *ListTasksRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTasksRequest) GetAmong() []string {
	if m != nil {
		return m.Among
	}
	return nil
}

type ListTasksResponse struct {
	// Paginated list of Tasks that are children of the parent resource.
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTasksResponse) Reset()         { *m = ListTasksResponse{} }
func (m *ListTasksResponse) String() string { return proto.CompactTextString(m) }
func (*ListTasksResponse) ProtoMessage()    {}
func (*ListTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{1}
}

func (m *ListTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTasksResponse.Unmarshal(m, b)
}
func (m *ListTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTasksResponse.Marshal(b, m, deterministic)
}
func (m *ListTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTasksResponse.Merge(m, src)
}
func (m *ListTasksResponse) XXX_Size() int {
	return xxx_messageInfo_ListTasksResponse.Size(m)
}
func (m *ListTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTasksResponse proto.InternalMessageInfo

func (m *ListTasksResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type GetTaskRequest struct {
	// Space of Task to retrieve.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Name of Task.
	TaskName             string   `protobuf:"bytes,5,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{2}
}

func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *GetTaskRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

type CreateTaskRequest struct {
	// Space to create Task in;
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Task being created.
	Task                 *Task    `protobuf:"bytes,5,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{3}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *CreateTaskRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateTaskRequest struct {
	// Resource name of Task to update.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Name of Task.
	TaskName string `protobuf:"bytes,5,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	// Task with changes to be implemented.
	Task                 *Task    `protobuf:"bytes,10,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{4}
}

func (m *UpdateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskRequest.Unmarshal(m, b)
}
func (m *UpdateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskRequest.Merge(m, src)
}
func (m *UpdateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskRequest.Size(m)
}
func (m *UpdateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskRequest proto.InternalMessageInfo

func (m *UpdateTaskRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *UpdateTaskRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *UpdateTaskRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type DeleteTaskRequest struct {
	// Name of the Task to be deleted.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Name of Task.
	TaskName             string   `protobuf:"bytes,5,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRequest) Reset()         { *m = DeleteTaskRequest{} }
func (m *DeleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRequest) ProtoMessage()    {}
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{5}
}

func (m *DeleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRequest.Unmarshal(m, b)
}
func (m *DeleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRequest.Merge(m, src)
}
func (m *DeleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRequest.Size(m)
}
func (m *DeleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRequest proto.InternalMessageInfo

func (m *DeleteTaskRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *DeleteTaskRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

// Details about a Petition on a Task
type TaskPetitionDetails struct {
	Scope                string   `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	PetitionID           string   `protobuf:"bytes,2,opt,name=petitionID,proto3" json:"petitionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskPetitionDetails) Reset()         { *m = TaskPetitionDetails{} }
func (m *TaskPetitionDetails) String() string { return proto.CompactTextString(m) }
func (*TaskPetitionDetails) ProtoMessage()    {}
func (*TaskPetitionDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{6}
}

func (m *TaskPetitionDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskPetitionDetails.Unmarshal(m, b)
}
func (m *TaskPetitionDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskPetitionDetails.Marshal(b, m, deterministic)
}
func (m *TaskPetitionDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskPetitionDetails.Merge(m, src)
}
func (m *TaskPetitionDetails) XXX_Size() int {
	return xxx_messageInfo_TaskPetitionDetails.Size(m)
}
func (m *TaskPetitionDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskPetitionDetails.DiscardUnknown(m)
}

var xxx_messageInfo_TaskPetitionDetails proto.InternalMessageInfo

func (m *TaskPetitionDetails) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *TaskPetitionDetails) GetPetitionID() string {
	if m != nil {
		return m.PetitionID
	}
	return ""
}

// Task in a Organization's hierarchy.
type Task struct {
	// Resource name of a Task.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of Space of Task.
	SpaceName string `protobuf:"bytes,5,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Meta contains generic information about the Task.
	Meta *v1.Meta `protobuf:"bytes,15,opt,name=meta,proto3" json:"meta,omitempty"`
	// Text that should be shown to users to represent a Task.
	DisplayName string `protobuf:"bytes,20,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Description provides a short summary of the Task and what's necessary to complete it.
	Description string `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`
	// Assignee is responsible for completing the Task.
	Assignee *v1.Resource `protobuf:"bytes,30,opt,name=assignee,proto3" json:"assignee,omitempty"`
	Phase    Task_Phase   `protobuf:"varint,35,opt,name=phase,proto3,enum=indent.v1.Task_Phase" json:"phase,omitempty"`
	// Form data associated with a Task.
	Form *Form `protobuf:"bytes,40,opt,name=form,proto3" json:"form,omitempty"`
	// Details about the scope of a Task.
	//
	// Types that are valid to be assigned to Details:
	//	*Task_PetitionDetails
	Details              isTask_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_947671edfa2a898b, []int{7}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *Task) GetMeta() *v1.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Task) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetAssignee() *v1.Resource {
	if m != nil {
		return m.Assignee
	}
	return nil
}

func (m *Task) GetPhase() Task_Phase {
	if m != nil {
		return m.Phase
	}
	return Task_UNKNOWN
}

func (m *Task) GetForm() *Form {
	if m != nil {
		return m.Form
	}
	return nil
}

type isTask_Details interface {
	isTask_Details()
}

type Task_PetitionDetails struct {
	PetitionDetails *TaskPetitionDetails `protobuf:"bytes,45,opt,name=petition_details,json=petitionDetails,proto3,oneof"`
}

func (*Task_PetitionDetails) isTask_Details() {}

func (m *Task) GetDetails() isTask_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Task) GetPetitionDetails() *TaskPetitionDetails {
	if x, ok := m.GetDetails().(*Task_PetitionDetails); ok {
		return x.PetitionDetails
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Task) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Task_PetitionDetails)(nil),
	}
}

func init() {
	proto.RegisterEnum("indent.v1.Task_Phase", Task_Phase_name, Task_Phase_value)
	proto.RegisterType((*ListTasksRequest)(nil), "indent.v1.ListTasksRequest")
	proto.RegisterType((*ListTasksResponse)(nil), "indent.v1.ListTasksResponse")
	proto.RegisterType((*GetTaskRequest)(nil), "indent.v1.GetTaskRequest")
	proto.RegisterType((*CreateTaskRequest)(nil), "indent.v1.CreateTaskRequest")
	proto.RegisterType((*UpdateTaskRequest)(nil), "indent.v1.UpdateTaskRequest")
	proto.RegisterType((*DeleteTaskRequest)(nil), "indent.v1.DeleteTaskRequest")
	proto.RegisterType((*TaskPetitionDetails)(nil), "indent.v1.TaskPetitionDetails")
	proto.RegisterType((*Task)(nil), "indent.v1.Task")
}

func init() {
	proto.RegisterFile("indent/v1/task_api.proto", fileDescriptor_947671edfa2a898b)
}

var fileDescriptor_947671edfa2a898b = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0x66, 0xfb, 0x43, 0xbb, 0xa7, 0x08, 0xed, 0x88, 0xc9, 0xb6, 0xfc, 0xd5, 0x25, 0x24, 0x05,
	0x62, 0x9b, 0x22, 0xde, 0xe0, 0x95, 0x50, 0xd4, 0x06, 0x6c, 0xeb, 0x22, 0x3f, 0x31, 0x26, 0xcd,
	0xd8, 0x8e, 0x65, 0x42, 0x77, 0x77, 0xe8, 0x2c, 0x24, 0x42, 0xb8, 0xf1, 0x15, 0x7c, 0x03, 0x2f,
	0x8d, 0x0f, 0xe1, 0xb5, 0xb7, 0xbe, 0x82, 0x0f, 0x62, 0x66, 0x76, 0xdb, 0x6d, 0x77, 0xb1, 0x81,
	0xc4, 0xbb, 0xdd, 0x73, 0xce, 0x7c, 0xdf, 0x77, 0xe6, 0xfc, 0x0c, 0x68, 0xd4, 0x6a, 0x13, 0xcb,
	0x29, 0x5d, 0x96, 0x4b, 0x0e, 0xe6, 0x67, 0x4d, 0xcc, 0x68, 0x91, 0xf5, 0x6c, 0xc7, 0x46, 0xaa,
	0xeb, 0x29, 0x5e, 0x96, 0x73, 0xf3, 0x1d, 0xdb, 0xee, 0x74, 0x49, 0x09, 0x33, 0x5a, 0xc2, 0x96,
	0x65, 0x3b, 0xd8, 0xa1, 0xb6, 0xc5, 0xdd, 0xc0, 0xdc, 0x92, 0x07, 0x81, 0x2f, 0xda, 0x54, 0x02,
	0xf5, 0x08, 0xb7, 0x2f, 0x7a, 0x2d, 0xd2, 0x0f, 0xc8, 0xfa, 0x1c, 0x01, 0x97, 0xde, 0x86, 0xf4,
	0x3e, 0xe5, 0xce, 0x3b, 0xcc, 0xcf, 0xb8, 0x41, 0xce, 0x2f, 0x08, 0x77, 0xd0, 0x02, 0x00, 0x67,
	0xb8, 0x45, 0x9a, 0x16, 0x36, 0x89, 0xa6, 0xe4, 0x95, 0x82, 0x6a, 0xa8, 0xd2, 0x52, 0xc3, 0x26,
	0x41, 0x73, 0xa0, 0x32, 0xdc, 0x21, 0x4d, 0x4e, 0xaf, 0x88, 0x16, 0xc9, 0x2b, 0x85, 0xb8, 0x91,
	0x14, 0x86, 0x03, 0x7a, 0x45, 0xd0, 0x2c, 0xc4, 0xb1, 0x69, 0x5b, 0x1d, 0x2d, 0x96, 0x8f, 0x16,
	0x54, 0xc3, 0xfd, 0xd1, 0xb7, 0x20, 0x33, 0xc4, 0xc2, 0x99, 0x6d, 0x71, 0x82, 0x56, 0x20, 0x2e,
	0x32, 0xe6, 0x9a, 0x92, 0x8f, 0x16, 0x52, 0x1b, 0x33, 0xc5, 0x41, 0xbe, 0x45, 0x11, 0x68, 0xb8,
	0x5e, 0x7d, 0x1f, 0xa6, 0x5f, 0x11, 0x79, 0xf4, 0xee, 0xfa, 0xe4, 0x4d, 0x4a, 0x6f, 0x5c, 0x7a,
	0x93, 0xc2, 0x20, 0x9c, 0xfa, 0x31, 0x64, 0x76, 0x7a, 0x04, 0x3b, 0xe4, 0x1e, 0x80, 0xcb, 0x10,
	0x13, 0xe7, 0x25, 0xd6, 0x2d, 0x3a, 0xa5, 0x53, 0x77, 0x20, 0x73, 0xc8, 0xda, 0xf7, 0x03, 0x1e,
	0xa7, 0x74, 0xc0, 0x0a, 0xe3, 0x58, 0xeb, 0x90, 0xa9, 0x90, 0x2e, 0xf9, 0x6f, 0xac, 0xfa, 0x1e,
	0x3c, 0x14, 0x50, 0x0d, 0xe2, 0x50, 0xd1, 0x62, 0x15, 0xe2, 0x60, 0xda, 0xe5, 0xa2, 0xac, 0xbc,
	0x65, 0xb3, 0x3e, 0x9a, 0xfb, 0x83, 0x16, 0x01, 0x98, 0x17, 0x58, 0xad, 0xc8, 0x56, 0x50, 0x8d,
	0x21, 0x8b, 0xfe, 0x33, 0x0a, 0x31, 0x81, 0x86, 0x10, 0xc4, 0x86, 0xb4, 0xc8, 0xef, 0x80, 0xca,
	0x78, 0x50, 0xe5, 0x2a, 0xc4, 0x4c, 0xe2, 0x60, 0x6d, 0x46, 0xa6, 0xff, 0xa8, 0x9f, 0xbe, 0xec,
	0x71, 0x71, 0x09, 0x6f, 0x88, 0x83, 0x0d, 0x19, 0x82, 0x1e, 0xc3, 0x54, 0x9b, 0x72, 0xd6, 0xc5,
	0x9f, 0x5d, 0xac, 0x59, 0x89, 0x95, 0xf2, 0x6c, 0x12, 0x2d, 0x0f, 0xa9, 0x36, 0xe1, 0xad, 0x1e,
	0x65, 0x42, 0x9a, 0x96, 0xf5, 0x22, 0x7c, 0x13, 0x7a, 0x06, 0x49, 0xcc, 0x39, 0xed, 0x58, 0x84,
	0x68, 0x8b, 0x92, 0x33, 0x1b, 0xe2, 0x34, 0xbc, 0xe1, 0x31, 0x06, 0xa1, 0x68, 0x1d, 0xe2, 0xec,
	0x14, 0x73, 0xa2, 0x2d, 0xe7, 0x95, 0xc2, 0xb4, 0xaf, 0xd3, 0x2b, 0x53, 0xb1, 0x21, 0x9c, 0x86,
	0x1b, 0x23, 0x4a, 0xfa, 0xc9, 0xee, 0x99, 0x5a, 0x21, 0x54, 0xd2, 0x97, 0x76, 0xcf, 0x34, 0xa4,
	0x13, 0xed, 0x41, 0xba, 0x7f, 0x85, 0xcd, 0xb6, 0x7b, 0xfd, 0xda, 0x13, 0x79, 0x60, 0x31, 0x00,
	0x1e, 0x28, 0xd2, 0xeb, 0x09, 0x63, 0x86, 0x8d, 0x9a, 0xf4, 0x4d, 0x88, 0x4b, 0x05, 0x28, 0x05,
	0x89, 0xc3, 0xda, 0x5e, 0xad, 0x7e, 0x5c, 0x4b, 0x4f, 0xa0, 0x24, 0xc4, 0xea, 0x8d, 0xdd, 0x5a,
	0x5a, 0x41, 0x00, 0x93, 0x3b, 0xfb, 0xf5, 0x83, 0xdd, 0x4a, 0x3a, 0x22, 0xac, 0x95, 0x7a, 0x6d,
	0x37, 0x1d, 0xdd, 0x56, 0x21, 0xe1, 0x31, 0x6f, 0xfc, 0x88, 0x41, 0x42, 0x70, 0xbd, 0x68, 0x54,
	0x91, 0x09, 0xea, 0x60, 0x8a, 0xd1, 0xdc, 0x90, 0x98, 0xe0, 0x06, 0xc9, 0xcd, 0xdf, 0xee, 0x74,
	0x07, 0x5f, 0x5f, 0xf9, 0xf2, 0xfb, 0xcf, 0xd7, 0xc8, 0x12, 0x5a, 0x10, 0x0b, 0x49, 0x56, 0x9c,
	0x97, 0xae, 0xfd, 0x5e, 0xb8, 0x91, 0x8b, 0x90, 0xa3, 0x53, 0x48, 0x78, 0x83, 0x8f, 0xb2, 0x43,
	0x78, 0xa3, 0xcb, 0x20, 0x17, 0x1c, 0x0c, 0xbd, 0x2c, 0xd1, 0xd7, 0xd1, 0xea, 0x58, 0xf4, 0xd2,
	0xf5, 0x60, 0x06, 0x6e, 0x10, 0x01, 0xf0, 0x97, 0x02, 0x1a, 0x16, 0x1f, 0xda, 0x15, 0x61, 0xbe,
	0x82, 0xe4, 0xd3, 0xf5, 0xf1, 0xd9, 0x6c, 0x29, 0x6b, 0xe8, 0x1c, 0xc0, 0x5f, 0x11, 0x23, 0x34,
	0xa1, 0xcd, 0x11, 0xa6, 0xd9, 0x94, 0x34, 0xc5, 0x8d, 0xbb, 0xa7, 0x25, 0x28, 0x2d, 0x00, 0x7f,
	0x3f, 0x8c, 0x50, 0x86, 0xd6, 0xc6, 0x3f, 0x6f, 0x72, 0xed, 0xee, 0x94, 0xdb, 0x6f, 0xe1, 0x41,
	0xcb, 0x36, 0x7d, 0xa0, 0xed, 0x29, 0xd9, 0x3c, 0x8c, 0x36, 0xc4, 0x6b, 0xd3, 0x50, 0xde, 0xa3,
	0xc1, 0x53, 0xf4, 0xdc, 0xfd, 0xba, 0x2c, 0x7f, 0x8b, 0x44, 0xab, 0x27, 0x27, 0xdf, 0x23, 0x6a,
	0xd5, 0x3d, 0x75, 0x54, 0xfe, 0xd5, 0xff, 0xfe, 0x70, 0x54, 0xfe, 0x38, 0x29, 0x1f, 0xaa, 0xa7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x31, 0xa9, 0x07, 0x29, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaskAPIClient is the client API for TaskAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskAPIClient interface {
	// List the Tasks within a space.
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
	// Retrieve specified Task by name.
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error)
	// Create a new Task.
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	// Update configuration of a Task.
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	// Permanently destroy a Task.
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*Task, error)
}

type taskAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskAPIClient(cc grpc.ClientConnInterface) TaskAPIClient {
	return &taskAPIClient{cc}
}

func (c *taskAPIClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, "/indent.v1.TaskAPI/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskAPIClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/indent.v1.TaskAPI/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskAPIClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/indent.v1.TaskAPI/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskAPIClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/indent.v1.TaskAPI/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskAPIClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/indent.v1.TaskAPI/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskAPIServer is the server API for TaskAPI service.
type TaskAPIServer interface {
	// List the Tasks within a space.
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	// Retrieve specified Task by name.
	GetTask(context.Context, *GetTaskRequest) (*Task, error)
	// Create a new Task.
	CreateTask(context.Context, *CreateTaskRequest) (*Task, error)
	// Update configuration of a Task.
	UpdateTask(context.Context, *UpdateTaskRequest) (*Task, error)
	// Permanently destroy a Task.
	DeleteTask(context.Context, *DeleteTaskRequest) (*Task, error)
}

// UnimplementedTaskAPIServer can be embedded to have forward compatible implementations.
type UnimplementedTaskAPIServer struct {
}

func (*UnimplementedTaskAPIServer) ListTasks(ctx context.Context, req *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (*UnimplementedTaskAPIServer) GetTask(ctx context.Context, req *GetTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedTaskAPIServer) CreateTask(ctx context.Context, req *CreateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedTaskAPIServer) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (*UnimplementedTaskAPIServer) DeleteTask(ctx context.Context, req *DeleteTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}

func RegisterTaskAPIServer(s *grpc.Server, srv TaskAPIServer) {
	s.RegisterService(&_TaskAPI_serviceDesc, srv)
}

func _TaskAPI_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.TaskAPI/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).ListTasks(ctx, req.(*ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskAPI_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.TaskAPI/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskAPI_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.TaskAPI/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskAPI_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.TaskAPI/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskAPI_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.TaskAPI/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "indent.v1.TaskAPI",
	HandlerType: (*TaskAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTasks",
			Handler:    _TaskAPI_ListTasks_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskAPI_GetTask_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _TaskAPI_CreateTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskAPI_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TaskAPI_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indent/v1/task_api.proto",
}