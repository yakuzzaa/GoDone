// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: list.proto

package list_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ListInfo) Reset() {
	*x = ListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInfo) ProtoMessage() {}

func (x *ListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInfo.ProtoReflect.Descriptor instead.
func (*ListInfo) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{0}
}

func (x *ListInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info      *ListInfo              `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{1}
}

func (x *List) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *List) GetInfo() *ListInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *List) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *List) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info      *ItemInfo              `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetInfo() *ItemInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Item) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Item) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ItemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *ItemInfo) Reset() {
	*x = ItemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemInfo) ProtoMessage() {}

func (x *ItemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemInfo.ProtoReflect.Descriptor instead.
func (*ItemInfo) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{3}
}

func (x *ItemInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ItemInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ItemInfo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type ListWithItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  *List   `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListWithItems) Reset() {
	*x = ListWithItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWithItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWithItems) ProtoMessage() {}

func (x *ListWithItems) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWithItems.ProtoReflect.Descriptor instead.
func (*ListWithItems) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{4}
}

func (x *ListWithItems) GetList() *List {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListWithItems) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateList) Reset() {
	*x = UpdateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateList) ProtoMessage() {}

func (x *UpdateList) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateList.ProtoReflect.Descriptor instead.
func (*UpdateList) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateList) GetTitle() *wrapperspb.StringValue {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *UpdateList) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{6}
}

func (x *ListRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lists []*List `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetLists() []*List {
	if x != nil {
		return x.Lists
	}
	return nil
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{8}
}

func (x *DetailRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DetailRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List *ListWithItems `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{9}
}

func (x *DetailResponse) GetList() *ListWithItems {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Info   *ListInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{10}
}

func (x *CreateRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRequest) GetInfo() *ListInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{11}
}

func (x *CreateResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info *UpdateList `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetInfo() *UpdateList {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_list_proto protoreflect.FileDescriptor

var file_list_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x42, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb0, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x56, 0x0a, 0x08, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f,
	0x6e, 0x65, 0x22, 0x51, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x30, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x22, 0x38, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x0e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x1f,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x9e, 0x02, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x31, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x11, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6b, 0x75, 0x7a, 0x7a, 0x61, 0x61, 0x2f, 0x47, 0x6f, 0x44, 0x6f, 0x6e, 0x65, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x3b, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_list_proto_rawDescOnce sync.Once
	file_list_proto_rawDescData = file_list_proto_rawDesc
)

func file_list_proto_rawDescGZIP() []byte {
	file_list_proto_rawDescOnce.Do(func() {
		file_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_proto_rawDescData)
	})
	return file_list_proto_rawDescData
}

var file_list_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_list_proto_goTypes = []any{
	(*ListInfo)(nil),               // 0: list.ListInfo
	(*List)(nil),                   // 1: list.List
	(*Item)(nil),                   // 2: list.Item
	(*ItemInfo)(nil),               // 3: list.ItemInfo
	(*ListWithItems)(nil),          // 4: list.ListWithItems
	(*UpdateList)(nil),             // 5: list.UpdateList
	(*ListRequest)(nil),            // 6: list.ListRequest
	(*ListResponse)(nil),           // 7: list.ListResponse
	(*DetailRequest)(nil),          // 8: list.DetailRequest
	(*DetailResponse)(nil),         // 9: list.DetailResponse
	(*CreateRequest)(nil),          // 10: list.CreateRequest
	(*CreateResponse)(nil),         // 11: list.CreateResponse
	(*UpdateRequest)(nil),          // 12: list.UpdateRequest
	(*DeleteRequest)(nil),          // 13: list.DeleteRequest
	(*timestamppb.Timestamp)(nil),  // 14: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 15: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 16: google.protobuf.Empty
}
var file_list_proto_depIdxs = []int32{
	0,  // 0: list.List.info:type_name -> list.ListInfo
	14, // 1: list.List.created_at:type_name -> google.protobuf.Timestamp
	14, // 2: list.List.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 3: list.Item.info:type_name -> list.ItemInfo
	14, // 4: list.Item.created_at:type_name -> google.protobuf.Timestamp
	14, // 5: list.Item.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 6: list.ListWithItems.list:type_name -> list.List
	2,  // 7: list.ListWithItems.items:type_name -> list.Item
	15, // 8: list.UpdateList.title:type_name -> google.protobuf.StringValue
	15, // 9: list.UpdateList.description:type_name -> google.protobuf.StringValue
	1,  // 10: list.ListResponse.lists:type_name -> list.List
	4,  // 11: list.DetailResponse.list:type_name -> list.ListWithItems
	0,  // 12: list.CreateRequest.info:type_name -> list.ListInfo
	5,  // 13: list.UpdateRequest.info:type_name -> list.UpdateList
	6,  // 14: list.ListV1.List:input_type -> list.ListRequest
	8,  // 15: list.ListV1.GetDetail:input_type -> list.DetailRequest
	10, // 16: list.ListV1.CreateList:input_type -> list.CreateRequest
	12, // 17: list.ListV1.UpdateList:input_type -> list.UpdateRequest
	13, // 18: list.ListV1.DeleteList:input_type -> list.DeleteRequest
	7,  // 19: list.ListV1.List:output_type -> list.ListResponse
	9,  // 20: list.ListV1.GetDetail:output_type -> list.DetailResponse
	11, // 21: list.ListV1.CreateList:output_type -> list.CreateResponse
	16, // 22: list.ListV1.UpdateList:output_type -> google.protobuf.Empty
	16, // 23: list.ListV1.DeleteList:output_type -> google.protobuf.Empty
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_list_proto_init() }
func file_list_proto_init() {
	if File_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_list_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*List); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ItemInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListWithItems); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_list_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_list_proto_goTypes,
		DependencyIndexes: file_list_proto_depIdxs,
		MessageInfos:      file_list_proto_msgTypes,
	}.Build()
	File_list_proto = out.File
	file_list_proto_rawDesc = nil
	file_list_proto_goTypes = nil
	file_list_proto_depIdxs = nil
}
