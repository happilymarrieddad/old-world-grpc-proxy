// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: proto/itemtypes/item-types.proto

package itemtypes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetItemTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetItemTypeRequest) Reset() {
	*x = GetItemTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemTypeRequest) ProtoMessage() {}

func (x *GetItemTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemTypeRequest.ProtoReflect.Descriptor instead.
func (*GetItemTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemTypeRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *GetItemTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetItemTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT    string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	GameId string `protobuf:"bytes,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
	Limit  int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetItemTypesRequest) Reset() {
	*x = GetItemTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemTypesRequest) ProtoMessage() {}

func (x *GetItemTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemTypesRequest.ProtoReflect.Descriptor instead.
func (*GetItemTypesRequest) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemTypesRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *GetItemTypesRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *GetItemTypesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetItemTypesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetItemTypesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemTypes []*ItemType `protobuf:"bytes,1,rep,name=itemTypes,proto3" json:"itemTypes,omitempty"`
	Count     int64       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetItemTypesReply) Reset() {
	*x = GetItemTypesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemTypesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemTypesReply) ProtoMessage() {}

func (x *GetItemTypesReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemTypesReply.ProtoReflect.Descriptor instead.
func (*GetItemTypesReply) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{2}
}

func (x *GetItemTypesReply) GetItemTypes() []*ItemType {
	if x != nil {
		return x.ItemTypes
	}
	return nil
}

func (x *GetItemTypesReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateItemTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT    string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GameId string `protobuf:"bytes,3,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (x *CreateItemTypeRequest) Reset() {
	*x = CreateItemTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItemTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemTypeRequest) ProtoMessage() {}

func (x *CreateItemTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateItemTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{3}
}

func (x *CreateItemTypeRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *CreateItemTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateItemTypeRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type UpdateItemTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT  string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateItemTypeRequest) Reset() {
	*x = UpdateItemTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemTypeRequest) ProtoMessage() {}

func (x *UpdateItemTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateItemTypeRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *UpdateItemTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateItemTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteItemTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteItemTypeRequest) Reset() {
	*x = DeleteItemTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemTypeRequest) ProtoMessage() {}

func (x *DeleteItemTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemTypeRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteItemTypeRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *DeleteItemTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ItemType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GameId    string                 `protobuf:"bytes,3,opt,name=gameId,proto3" json:"gameId,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ItemType) Reset() {
	*x = ItemType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemType) ProtoMessage() {}

func (x *ItemType) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemType.ProtoReflect.Descriptor instead.
func (*ItemType) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{6}
}

func (x *ItemType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemType) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *ItemType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type EmptyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReply) Reset() {
	*x = EmptyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_itemtypes_item_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReply) ProtoMessage() {}

func (x *EmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_itemtypes_item_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReply.ProtoReflect.Descriptor instead.
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return file_proto_itemtypes_item_types_proto_rawDescGZIP(), []int{7}
}

var File_proto_itemtypes_item_types_proto protoreflect.FileDescriptor

var file_proto_itemtypes_item_types_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4a, 0x57, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x87, 0x03, 0x0a, 0x0b, 0x56, 0x31, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x2e, 0x69,
	0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61,
	0x70, 0x70, 0x69, 0x6c, 0x79, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x64, 0x61, 0x64, 0x2f,
	0x6f, 0x6c, 0x64, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x33, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_itemtypes_item_types_proto_rawDescOnce sync.Once
	file_proto_itemtypes_item_types_proto_rawDescData = file_proto_itemtypes_item_types_proto_rawDesc
)

func file_proto_itemtypes_item_types_proto_rawDescGZIP() []byte {
	file_proto_itemtypes_item_types_proto_rawDescOnce.Do(func() {
		file_proto_itemtypes_item_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_itemtypes_item_types_proto_rawDescData)
	})
	return file_proto_itemtypes_item_types_proto_rawDescData
}

var file_proto_itemtypes_item_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_itemtypes_item_types_proto_goTypes = []interface{}{
	(*GetItemTypeRequest)(nil),    // 0: itemtypes.GetItemTypeRequest
	(*GetItemTypesRequest)(nil),   // 1: itemtypes.GetItemTypesRequest
	(*GetItemTypesReply)(nil),     // 2: itemtypes.GetItemTypesReply
	(*CreateItemTypeRequest)(nil), // 3: itemtypes.CreateItemTypeRequest
	(*UpdateItemTypeRequest)(nil), // 4: itemtypes.UpdateItemTypeRequest
	(*DeleteItemTypeRequest)(nil), // 5: itemtypes.DeleteItemTypeRequest
	(*ItemType)(nil),              // 6: itemtypes.ItemType
	(*EmptyReply)(nil),            // 7: itemtypes.EmptyReply
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_proto_itemtypes_item_types_proto_depIdxs = []int32{
	6, // 0: itemtypes.GetItemTypesReply.itemTypes:type_name -> itemtypes.ItemType
	8, // 1: itemtypes.ItemType.createdAt:type_name -> google.protobuf.Timestamp
	1, // 2: itemtypes.V1ItemTypes.GetItemTypes:input_type -> itemtypes.GetItemTypesRequest
	0, // 3: itemtypes.V1ItemTypes.GetItemType:input_type -> itemtypes.GetItemTypeRequest
	3, // 4: itemtypes.V1ItemTypes.CreateItemType:input_type -> itemtypes.CreateItemTypeRequest
	4, // 5: itemtypes.V1ItemTypes.UpdateItemType:input_type -> itemtypes.UpdateItemTypeRequest
	5, // 6: itemtypes.V1ItemTypes.DeleteItemType:input_type -> itemtypes.DeleteItemTypeRequest
	2, // 7: itemtypes.V1ItemTypes.GetItemTypes:output_type -> itemtypes.GetItemTypesReply
	6, // 8: itemtypes.V1ItemTypes.GetItemType:output_type -> itemtypes.ItemType
	6, // 9: itemtypes.V1ItemTypes.CreateItemType:output_type -> itemtypes.ItemType
	7, // 10: itemtypes.V1ItemTypes.UpdateItemType:output_type -> itemtypes.EmptyReply
	7, // 11: itemtypes.V1ItemTypes.DeleteItemType:output_type -> itemtypes.EmptyReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_itemtypes_item_types_proto_init() }
func file_proto_itemtypes_item_types_proto_init() {
	if File_proto_itemtypes_item_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_itemtypes_item_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemTypeRequest); i {
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
		file_proto_itemtypes_item_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemTypesRequest); i {
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
		file_proto_itemtypes_item_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemTypesReply); i {
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
		file_proto_itemtypes_item_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItemTypeRequest); i {
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
		file_proto_itemtypes_item_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemTypeRequest); i {
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
		file_proto_itemtypes_item_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemTypeRequest); i {
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
		file_proto_itemtypes_item_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemType); i {
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
		file_proto_itemtypes_item_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReply); i {
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
			RawDescriptor: file_proto_itemtypes_item_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_itemtypes_item_types_proto_goTypes,
		DependencyIndexes: file_proto_itemtypes_item_types_proto_depIdxs,
		MessageInfos:      file_proto_itemtypes_item_types_proto_msgTypes,
	}.Build()
	File_proto_itemtypes_item_types_proto = out.File
	file_proto_itemtypes_item_types_proto_rawDesc = nil
	file_proto_itemtypes_item_types_proto_goTypes = nil
	file_proto_itemtypes_item_types_proto_depIdxs = nil
}
