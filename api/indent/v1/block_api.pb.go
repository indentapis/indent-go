// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indent/v1/block_api.proto

package indentv1

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

type ListBlocksRequest struct {
	// Name of Space containing Blocks.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// LabelSelector specifies which Blocks should be returned.
	LabelSelector string `protobuf:"bytes,2,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	// Latest returns the most recent variation matched by LabelSelector.
	Latest bool `protobuf:"varint,3,opt,name=latest,proto3" json:"latest,omitempty"`
	// Among constrains listed blocks by name.
	Among []string `protobuf:"bytes,4,rep,name=among,proto3" json:"among,omitempty"`
	// Max number of Blocks to be returned.
	PageSize int32 `protobuf:"varint,10,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for current page position for this list of Blocks.
	PageToken            string   `protobuf:"bytes,11,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlocksRequest) Reset()         { *m = ListBlocksRequest{} }
func (m *ListBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*ListBlocksRequest) ProtoMessage()    {}
func (*ListBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{0}
}

func (m *ListBlocksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlocksRequest.Unmarshal(m, b)
}
func (m *ListBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlocksRequest.Marshal(b, m, deterministic)
}
func (m *ListBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlocksRequest.Merge(m, src)
}
func (m *ListBlocksRequest) XXX_Size() int {
	return xxx_messageInfo_ListBlocksRequest.Size(m)
}
func (m *ListBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlocksRequest proto.InternalMessageInfo

func (m *ListBlocksRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *ListBlocksRequest) GetLabelSelector() string {
	if m != nil {
		return m.LabelSelector
	}
	return ""
}

func (m *ListBlocksRequest) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
}

func (m *ListBlocksRequest) GetAmong() []string {
	if m != nil {
		return m.Among
	}
	return nil
}

func (m *ListBlocksRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListBlocksRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListBlocksResponse struct {
	// Paginated list of Blocks in a Space.
	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	// Page cursor for list of Blocks.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlocksResponse) Reset()         { *m = ListBlocksResponse{} }
func (m *ListBlocksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBlocksResponse) ProtoMessage()    {}
func (*ListBlocksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{1}
}

func (m *ListBlocksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlocksResponse.Unmarshal(m, b)
}
func (m *ListBlocksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlocksResponse.Marshal(b, m, deterministic)
}
func (m *ListBlocksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlocksResponse.Merge(m, src)
}
func (m *ListBlocksResponse) XXX_Size() int {
	return xxx_messageInfo_ListBlocksResponse.Size(m)
}
func (m *ListBlocksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlocksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlocksResponse proto.InternalMessageInfo

func (m *ListBlocksResponse) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *ListBlocksResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetBlockRequest struct {
	// Name of space containing Block.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Resource name of Block to retrieve.
	BlockName            string   `protobuf:"bytes,2,opt,name=block_name,json=blockName,proto3" json:"block_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockRequest) Reset()         { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()    {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{2}
}

func (m *GetBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockRequest.Unmarshal(m, b)
}
func (m *GetBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockRequest.Merge(m, src)
}
func (m *GetBlockRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockRequest.Size(m)
}
func (m *GetBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockRequest proto.InternalMessageInfo

func (m *GetBlockRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *GetBlockRequest) GetBlockName() string {
	if m != nil {
		return m.BlockName
	}
	return ""
}

type CreateBlockRequest struct {
	// Space where Block should be created.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Ephemeral prevents the Block from being persisted.
	Ephemeral bool `protobuf:"varint,5,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	// Announce allows watches to be evaluated for this change.
	Announce bool `protobuf:"varint,15,opt,name=announce,proto3" json:"announce,omitempty"`
	// Block being created, name is ignored and will be autogenerated.
	Block                *Block   `protobuf:"bytes,40,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlockRequest) Reset()         { *m = CreateBlockRequest{} }
func (m *CreateBlockRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlockRequest) ProtoMessage()    {}
func (*CreateBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{3}
}

func (m *CreateBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlockRequest.Unmarshal(m, b)
}
func (m *CreateBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlockRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlockRequest.Merge(m, src)
}
func (m *CreateBlockRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlockRequest.Size(m)
}
func (m *CreateBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlockRequest proto.InternalMessageInfo

func (m *CreateBlockRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *CreateBlockRequest) GetEphemeral() bool {
	if m != nil {
		return m.Ephemeral
	}
	return false
}

func (m *CreateBlockRequest) GetAnnounce() bool {
	if m != nil {
		return m.Announce
	}
	return false
}

func (m *CreateBlockRequest) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type UpdateBlockRequest struct {
	// SpaceName where Block should be updated.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// BlockName is the name of the block being updated.
	BlockName string `protobuf:"bytes,3,opt,name=block_name,json=blockName,proto3" json:"block_name,omitempty"`
	// Ephemeral prevents changes to the Block from being persisted.
	Ephemeral bool `protobuf:"varint,5,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	// Announce allows watches to be evaluated for this change.
	Announce bool `protobuf:"varint,15,opt,name=announce,proto3" json:"announce,omitempty"`
	// Block being updated, name is ignored.
	Block                *Block   `protobuf:"bytes,40,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlockRequest) Reset()         { *m = UpdateBlockRequest{} }
func (m *UpdateBlockRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBlockRequest) ProtoMessage()    {}
func (*UpdateBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{4}
}

func (m *UpdateBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlockRequest.Unmarshal(m, b)
}
func (m *UpdateBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlockRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlockRequest.Merge(m, src)
}
func (m *UpdateBlockRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBlockRequest.Size(m)
}
func (m *UpdateBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlockRequest proto.InternalMessageInfo

func (m *UpdateBlockRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *UpdateBlockRequest) GetBlockName() string {
	if m != nil {
		return m.BlockName
	}
	return ""
}

func (m *UpdateBlockRequest) GetEphemeral() bool {
	if m != nil {
		return m.Ephemeral
	}
	return false
}

func (m *UpdateBlockRequest) GetAnnounce() bool {
	if m != nil {
		return m.Announce
	}
	return false
}

func (m *UpdateBlockRequest) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type DeleteBlockRequest struct {
	// Name of Space containing Block.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Name of the Block to be deleted.
	BlockName            string   `protobuf:"bytes,2,opt,name=block_name,json=blockName,proto3" json:"block_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlockRequest) Reset()         { *m = DeleteBlockRequest{} }
func (m *DeleteBlockRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlockRequest) ProtoMessage()    {}
func (*DeleteBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{5}
}

func (m *DeleteBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlockRequest.Unmarshal(m, b)
}
func (m *DeleteBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlockRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlockRequest.Merge(m, src)
}
func (m *DeleteBlockRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlockRequest.Size(m)
}
func (m *DeleteBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlockRequest proto.InternalMessageInfo

func (m *DeleteBlockRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *DeleteBlockRequest) GetBlockName() string {
	if m != nil {
		return m.BlockName
	}
	return ""
}

type BulkDeleteBlocksRequest struct {
	// SpaceName where Blocks should be deleted.
	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	// Blocks to be deleted.
	BlockNames           []string `protobuf:"bytes,40,rep,name=block_names,json=blockNames,proto3" json:"block_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkDeleteBlocksRequest) Reset()         { *m = BulkDeleteBlocksRequest{} }
func (m *BulkDeleteBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*BulkDeleteBlocksRequest) ProtoMessage()    {}
func (*BulkDeleteBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{6}
}

func (m *BulkDeleteBlocksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDeleteBlocksRequest.Unmarshal(m, b)
}
func (m *BulkDeleteBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDeleteBlocksRequest.Marshal(b, m, deterministic)
}
func (m *BulkDeleteBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDeleteBlocksRequest.Merge(m, src)
}
func (m *BulkDeleteBlocksRequest) XXX_Size() int {
	return xxx_messageInfo_BulkDeleteBlocksRequest.Size(m)
}
func (m *BulkDeleteBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDeleteBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDeleteBlocksRequest proto.InternalMessageInfo

func (m *BulkDeleteBlocksRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *BulkDeleteBlocksRequest) GetBlockNames() []string {
	if m != nil {
		return m.BlockNames
	}
	return nil
}

type BulkDeleteBlocksResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkDeleteBlocksResponse) Reset()         { *m = BulkDeleteBlocksResponse{} }
func (m *BulkDeleteBlocksResponse) String() string { return proto.CompactTextString(m) }
func (*BulkDeleteBlocksResponse) ProtoMessage()    {}
func (*BulkDeleteBlocksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5445ae8e00ca6a, []int{7}
}

func (m *BulkDeleteBlocksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDeleteBlocksResponse.Unmarshal(m, b)
}
func (m *BulkDeleteBlocksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDeleteBlocksResponse.Marshal(b, m, deterministic)
}
func (m *BulkDeleteBlocksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDeleteBlocksResponse.Merge(m, src)
}
func (m *BulkDeleteBlocksResponse) XXX_Size() int {
	return xxx_messageInfo_BulkDeleteBlocksResponse.Size(m)
}
func (m *BulkDeleteBlocksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDeleteBlocksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDeleteBlocksResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListBlocksRequest)(nil), "indent.v1.ListBlocksRequest")
	proto.RegisterType((*ListBlocksResponse)(nil), "indent.v1.ListBlocksResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "indent.v1.GetBlockRequest")
	proto.RegisterType((*CreateBlockRequest)(nil), "indent.v1.CreateBlockRequest")
	proto.RegisterType((*UpdateBlockRequest)(nil), "indent.v1.UpdateBlockRequest")
	proto.RegisterType((*DeleteBlockRequest)(nil), "indent.v1.DeleteBlockRequest")
	proto.RegisterType((*BulkDeleteBlocksRequest)(nil), "indent.v1.BulkDeleteBlocksRequest")
	proto.RegisterType((*BulkDeleteBlocksResponse)(nil), "indent.v1.BulkDeleteBlocksResponse")
}

func init() {
	proto.RegisterFile("indent/v1/block_api.proto", fileDescriptor_4e5445ae8e00ca6a)
}

var fileDescriptor_4e5445ae8e00ca6a = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0xd5, 0x24, 0x2f, 0x55, 0x7c, 0xa3, 0xbe, 0xc2, 0x08, 0x81, 0x31, 0x2d, 0x58, 0xae, 0xa8,
	0xdc, 0x22, 0x62, 0xd2, 0x22, 0x16, 0x65, 0x45, 0x41, 0x42, 0x95, 0x10, 0x44, 0x2e, 0x54, 0x05,
	0x21, 0x45, 0x93, 0xf4, 0x12, 0xac, 0x3a, 0x1e, 0x37, 0x33, 0x89, 0x50, 0xab, 0x6e, 0xd8, 0xb2,
	0x44, 0xfc, 0x00, 0x4b, 0xd6, 0xfc, 0x01, 0x3b, 0x96, 0xf0, 0x0b, 0x7c, 0x08, 0xf2, 0x8c, 0x1b,
	0xbb, 0x75, 0x1a, 0x1a, 0x21, 0xb1, 0xf3, 0x9c, 0x3b, 0x39, 0xe7, 0x9e, 0x3b, 0x67, 0x32, 0x70,
	0x35, 0x88, 0x76, 0x31, 0x92, 0xde, 0xb0, 0xe1, 0xb5, 0x43, 0xde, 0xd9, 0x6b, 0xb1, 0x38, 0xa8,
	0xc7, 0x7d, 0x2e, 0x39, 0x35, 0x74, 0xa9, 0x3e, 0x6c, 0x58, 0xf3, 0x5d, 0xce, 0xbb, 0x21, 0x7a,
	0x2c, 0x0e, 0x3c, 0x16, 0x45, 0x5c, 0x32, 0x19, 0xf0, 0x48, 0xe8, 0x8d, 0x56, 0x8e, 0xa3, 0x8f,
	0x82, 0x0f, 0xfa, 0x1d, 0x4c, 0x4b, 0xce, 0x37, 0x02, 0x17, 0x9f, 0x04, 0x42, 0x6e, 0x24, 0xdc,
	0xc2, 0xc7, 0xfd, 0x01, 0x0a, 0x49, 0x17, 0x00, 0x44, 0xcc, 0x3a, 0xd8, 0x8a, 0x58, 0x0f, 0x4d,
	0x62, 0x13, 0xd7, 0xf0, 0x0d, 0x85, 0x3c, 0x65, 0x3d, 0xa4, 0x37, 0xe1, 0xff, 0x90, 0xb5, 0x31,
	0x6c, 0x09, 0x0c, 0xb1, 0x23, 0x79, 0xdf, 0x2c, 0xa9, 0x2d, 0xb3, 0x0a, 0xdd, 0x4a, 0x41, 0x7a,
	0x19, 0x66, 0x42, 0x26, 0x51, 0x48, 0xb3, 0x6c, 0x13, 0xb7, 0xea, 0xa7, 0x2b, 0x7a, 0x09, 0x2a,
	0xac, 0xc7, 0xa3, 0xae, 0xf9, 0x9f, 0x5d, 0x76, 0x0d, 0x5f, 0x2f, 0xe8, 0x35, 0x30, 0x62, 0xd6,
	0xc5, 0x96, 0x08, 0x0e, 0xd0, 0x04, 0x9b, 0xb8, 0x15, 0xbf, 0x9a, 0x00, 0x5b, 0xc1, 0x01, 0x26,
	0x0d, 0xa9, 0xa2, 0xe4, 0x7b, 0x18, 0x99, 0x35, 0xdd, 0x50, 0x82, 0x3c, 0x4f, 0x00, 0xe7, 0x0d,
	0xd0, 0xbc, 0x09, 0x11, 0xf3, 0x48, 0x20, 0x75, 0x61, 0x46, 0x8d, 0x4c, 0x98, 0xc4, 0x2e, 0xbb,
	0xb5, 0xd5, 0x0b, 0xf5, 0xd1, 0xc0, 0xea, 0x6a, 0xab, 0x9f, 0xd6, 0xe9, 0x12, 0xcc, 0x45, 0xf8,
	0x4e, 0xb6, 0x72, 0x1a, 0xa9, 0xa3, 0x04, 0x6e, 0x8e, 0x74, 0x9e, 0xc1, 0xdc, 0x63, 0xd4, 0x32,
	0xe7, 0x1c, 0xd5, 0x02, 0x80, 0x3e, 0x36, 0x55, 0xd6, 0xa4, 0x86, 0x42, 0x92, 0xb2, 0xf3, 0x89,
	0x00, 0x7d, 0xd8, 0x47, 0x26, 0x71, 0x1a, 0xd2, 0x79, 0x30, 0x30, 0x7e, 0x8b, 0x3d, 0xec, 0xb3,
	0xd0, 0xac, 0xa8, 0xd9, 0x66, 0x00, 0xb5, 0xa0, 0x9a, 0x44, 0x60, 0x10, 0x75, 0xd0, 0x9c, 0x53,
	0xc5, 0xd1, 0x9a, 0x2e, 0x41, 0x45, 0x89, 0x9b, 0xae, 0x4d, 0xc6, 0x4e, 0x44, 0x97, 0x9d, 0xaf,
	0x04, 0xe8, 0x8b, 0x78, 0x77, 0xca, 0xbe, 0x4e, 0x9a, 0x2d, 0x9f, 0x32, 0xfb, 0x0f, 0xda, 0xf6,
	0x81, 0x3e, 0xc2, 0x10, 0xff, 0xa6, 0xeb, 0xc2, 0x11, 0xbd, 0x84, 0x2b, 0x1b, 0x83, 0x70, 0x2f,
	0xc7, 0x7b, 0xde, 0x6b, 0x72, 0x03, 0x6a, 0x19, 0xb1, 0x30, 0x5d, 0x95, 0x76, 0x18, 0x31, 0x0b,
	0xc7, 0x02, 0xb3, 0x48, 0xad, 0xc3, 0xbb, 0xfa, 0xa3, 0x02, 0x55, 0x05, 0x3d, 0x68, 0x6e, 0xd2,
	0x7d, 0x80, 0x2c, 0xdf, 0x74, 0x3e, 0x67, 0xbf, 0x70, 0x77, 0xad, 0x85, 0x33, 0xaa, 0x9a, 0xd7,
	0x59, 0x7a, 0xff, 0xf3, 0xd7, 0xc7, 0x92, 0x4d, 0xaf, 0x27, 0xff, 0x06, 0xaa, 0x57, 0xe1, 0x1d,
	0x66, 0x2e, 0x8e, 0xbc, 0xf4, 0x4a, 0xf4, 0xa0, 0x7a, 0x1c, 0x75, 0x6a, 0xe5, 0x28, 0x4f, 0xe5,
	0xdf, 0x2a, 0x9c, 0x85, 0xb3, 0xa6, 0x14, 0x6e, 0xd3, 0x5b, 0x93, 0x15, 0xbc, 0xc3, 0x6c, 0x38,
	0x47, 0x34, 0x80, 0x5a, 0xee, 0x1e, 0xd0, 0xbc, 0x89, 0xe2, 0xfd, 0x18, 0x23, 0xba, 0xac, 0x44,
	0x17, 0xd7, 0xc9, 0x8a, 0xf3, 0x27, 0x67, 0x43, 0xa8, 0xe5, 0xa2, 0x7d, 0x42, 0xaa, 0x18, 0xf9,
	0x31, 0x52, 0xf7, 0x94, 0xd4, 0x9d, 0x75, 0xb2, 0x62, 0x4d, 0x65, 0x51, 0x40, 0x2d, 0x77, 0xd2,
	0x27, 0x74, 0x8b, 0xa1, 0x3d, 0x7b, 0xae, 0x2b, 0x53, 0x89, 0x7e, 0x20, 0x00, 0x59, 0xc6, 0xa8,
	0x93, 0x67, 0x1d, 0x9f, 0x6a, 0x6b, 0x71, 0xe2, 0x9e, 0x34, 0x46, 0x77, 0x55, 0x33, 0xf5, 0x64,
	0xde, 0xcb, 0x93, 0xfb, 0x59, 0x6f, 0x67, 0x1c, 0x3e, 0xcc, 0x76, 0x78, 0x2f, 0xe3, 0xdf, 0x98,
	0xd5, 0x11, 0x8f, 0x83, 0x66, 0xf2, 0x1a, 0x35, 0xc9, 0x2b, 0x3a, 0x7a, 0xaa, 0xee, 0xeb, 0xaf,
	0x61, 0xe3, 0x73, 0xa9, 0xbc, 0xb9, 0xb3, 0xf3, 0xa5, 0x64, 0x6c, 0xea, 0x9f, 0x6d, 0x37, 0xbe,
	0x1f, 0x7f, 0xbf, 0xde, 0x6e, 0xb4, 0x67, 0xd4, 0x43, 0xb6, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0x96, 0x5b, 0x92, 0x29, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlockAPIClient is the client API for BlockAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockAPIClient interface {
	// List the Blocks in the given space.
	ListBlocks(ctx context.Context, in *ListBlocksRequest, opts ...grpc.CallOption) (*ListBlocksResponse, error)
	// Retrieve specified Block by name and Space.
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*Block, error)
	// Create a new Block within a space.
	CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*Block, error)
	// UpdateBlock allows modification of the contents of a Block.
	UpdateBlock(ctx context.Context, in *UpdateBlockRequest, opts ...grpc.CallOption) (*Block, error)
	// Permanently destroy a Block.
	DeleteBlock(ctx context.Context, in *DeleteBlockRequest, opts ...grpc.CallOption) (*Block, error)
	// Permanently destroy Blocks.
	BulkDelete(ctx context.Context, in *BulkDeleteBlocksRequest, opts ...grpc.CallOption) (*BulkDeleteBlocksResponse, error)
}

type blockAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockAPIClient(cc grpc.ClientConnInterface) BlockAPIClient {
	return &blockAPIClient{cc}
}

func (c *blockAPIClient) ListBlocks(ctx context.Context, in *ListBlocksRequest, opts ...grpc.CallOption) (*ListBlocksResponse, error) {
	out := new(ListBlocksResponse)
	err := c.cc.Invoke(ctx, "/indent.v1.BlockAPI/ListBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockAPIClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/indent.v1.BlockAPI/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockAPIClient) CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/indent.v1.BlockAPI/CreateBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockAPIClient) UpdateBlock(ctx context.Context, in *UpdateBlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/indent.v1.BlockAPI/UpdateBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockAPIClient) DeleteBlock(ctx context.Context, in *DeleteBlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/indent.v1.BlockAPI/DeleteBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockAPIClient) BulkDelete(ctx context.Context, in *BulkDeleteBlocksRequest, opts ...grpc.CallOption) (*BulkDeleteBlocksResponse, error) {
	out := new(BulkDeleteBlocksResponse)
	err := c.cc.Invoke(ctx, "/indent.v1.BlockAPI/BulkDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockAPIServer is the server API for BlockAPI service.
type BlockAPIServer interface {
	// List the Blocks in the given space.
	ListBlocks(context.Context, *ListBlocksRequest) (*ListBlocksResponse, error)
	// Retrieve specified Block by name and Space.
	GetBlock(context.Context, *GetBlockRequest) (*Block, error)
	// Create a new Block within a space.
	CreateBlock(context.Context, *CreateBlockRequest) (*Block, error)
	// UpdateBlock allows modification of the contents of a Block.
	UpdateBlock(context.Context, *UpdateBlockRequest) (*Block, error)
	// Permanently destroy a Block.
	DeleteBlock(context.Context, *DeleteBlockRequest) (*Block, error)
	// Permanently destroy Blocks.
	BulkDelete(context.Context, *BulkDeleteBlocksRequest) (*BulkDeleteBlocksResponse, error)
}

// UnimplementedBlockAPIServer can be embedded to have forward compatible implementations.
type UnimplementedBlockAPIServer struct {
}

func (*UnimplementedBlockAPIServer) ListBlocks(ctx context.Context, req *ListBlocksRequest) (*ListBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlocks not implemented")
}
func (*UnimplementedBlockAPIServer) GetBlock(ctx context.Context, req *GetBlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedBlockAPIServer) CreateBlock(ctx context.Context, req *CreateBlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlock not implemented")
}
func (*UnimplementedBlockAPIServer) UpdateBlock(ctx context.Context, req *UpdateBlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlock not implemented")
}
func (*UnimplementedBlockAPIServer) DeleteBlock(ctx context.Context, req *DeleteBlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlock not implemented")
}
func (*UnimplementedBlockAPIServer) BulkDelete(ctx context.Context, req *BulkDeleteBlocksRequest) (*BulkDeleteBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkDelete not implemented")
}

func RegisterBlockAPIServer(s *grpc.Server, srv BlockAPIServer) {
	s.RegisterService(&_BlockAPI_serviceDesc, srv)
}

func _BlockAPI_ListBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockAPIServer).ListBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.BlockAPI/ListBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockAPIServer).ListBlocks(ctx, req.(*ListBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockAPI_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockAPIServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.BlockAPI/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockAPIServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockAPI_CreateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockAPIServer).CreateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.BlockAPI/CreateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockAPIServer).CreateBlock(ctx, req.(*CreateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockAPI_UpdateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockAPIServer).UpdateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.BlockAPI/UpdateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockAPIServer).UpdateBlock(ctx, req.(*UpdateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockAPI_DeleteBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockAPIServer).DeleteBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.BlockAPI/DeleteBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockAPIServer).DeleteBlock(ctx, req.(*DeleteBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockAPI_BulkDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockAPIServer).BulkDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.BlockAPI/BulkDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockAPIServer).BulkDelete(ctx, req.(*BulkDeleteBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "indent.v1.BlockAPI",
	HandlerType: (*BlockAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBlocks",
			Handler:    _BlockAPI_ListBlocks_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BlockAPI_GetBlock_Handler,
		},
		{
			MethodName: "CreateBlock",
			Handler:    _BlockAPI_CreateBlock_Handler,
		},
		{
			MethodName: "UpdateBlock",
			Handler:    _BlockAPI_UpdateBlock_Handler,
		},
		{
			MethodName: "DeleteBlock",
			Handler:    _BlockAPI_DeleteBlock_Handler,
		},
		{
			MethodName: "BulkDelete",
			Handler:    _BlockAPI_BulkDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indent/v1/block_api.proto",
}
