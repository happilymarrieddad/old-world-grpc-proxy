// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: proto/statistics/statistics.proto

package statistics

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

type GetStatisticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStatisticRequest) Reset() {
	*x = GetStatisticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatisticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatisticRequest) ProtoMessage() {}

func (x *GetStatisticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatisticRequest.ProtoReflect.Descriptor instead.
func (*GetStatisticRequest) Descriptor() ([]byte, []int) {
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatisticRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *GetStatisticRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetStatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT    string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	GameId string `protobuf:"bytes,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
	Limit  int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetStatisticsRequest) Reset() {
	*x = GetStatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatisticsRequest) ProtoMessage() {}

func (x *GetStatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatisticsRequest.ProtoReflect.Descriptor instead.
func (*GetStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatisticsRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *GetStatisticsRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *GetStatisticsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetStatisticsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetStatisticsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statistics []*Statistic `protobuf:"bytes,1,rep,name=statistics,proto3" json:"statistics,omitempty"`
	Count      int64        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetStatisticsReply) Reset() {
	*x = GetStatisticsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatisticsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatisticsReply) ProtoMessage() {}

func (x *GetStatisticsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatisticsReply.ProtoReflect.Descriptor instead.
func (*GetStatisticsReply) Descriptor() ([]byte, []int) {
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{2}
}

func (x *GetStatisticsReply) GetStatistics() []*Statistic {
	if x != nil {
		return x.Statistics
	}
	return nil
}

func (x *GetStatisticsReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateStatisticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT     string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Display string `protobuf:"bytes,3,opt,name=display,proto3" json:"display,omitempty"`
	GameId  string `protobuf:"bytes,4,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (x *CreateStatisticRequest) Reset() {
	*x = CreateStatisticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStatisticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStatisticRequest) ProtoMessage() {}

func (x *CreateStatisticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStatisticRequest.ProtoReflect.Descriptor instead.
func (*CreateStatisticRequest) Descriptor() ([]byte, []int) {
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStatisticRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *CreateStatisticRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStatisticRequest) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

func (x *CreateStatisticRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type UpdateStatisticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT     string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Display string `protobuf:"bytes,4,opt,name=display,proto3" json:"display,omitempty"`
}

func (x *UpdateStatisticRequest) Reset() {
	*x = UpdateStatisticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatisticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatisticRequest) ProtoMessage() {}

func (x *UpdateStatisticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatisticRequest.ProtoReflect.Descriptor instead.
func (*UpdateStatisticRequest) Descriptor() ([]byte, []int) {
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStatisticRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *UpdateStatisticRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStatisticRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStatisticRequest) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

type DeleteStatisticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JWT string `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStatisticRequest) Reset() {
	*x = DeleteStatisticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStatisticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStatisticRequest) ProtoMessage() {}

func (x *DeleteStatisticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStatisticRequest.ProtoReflect.Descriptor instead.
func (*DeleteStatisticRequest) Descriptor() ([]byte, []int) {
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteStatisticRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *DeleteStatisticRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Statistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Display   string                 `protobuf:"bytes,3,opt,name=display,proto3" json:"display,omitempty"`
	GameId    string                 `protobuf:"bytes,4,opt,name=gameId,proto3" json:"gameId,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Position  int64                  `protobuf:"varint,6,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Statistic) Reset() {
	*x = Statistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistic) ProtoMessage() {}

func (x *Statistic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistic.ProtoReflect.Descriptor instead.
func (*Statistic) Descriptor() ([]byte, []int) {
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{6}
}

func (x *Statistic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Statistic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Statistic) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

func (x *Statistic) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *Statistic) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Statistic) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

type EmptyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReply) Reset() {
	*x = EmptyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statistics_statistics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReply) ProtoMessage() {}

func (x *EmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statistics_statistics_proto_msgTypes[7]
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
	return file_proto_statistics_statistics_proto_rawDescGZIP(), []int{7}
}

var File_proto_statistics_statistics_proto protoreflect.FileDescriptor

var file_proto_statistics_statistics_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4a, 0x57, 0x54, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x35, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x68,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x3a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4a, 0x57, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0c,
	0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x9e, 0x03, 0x0a,
	0x0c, 0x56, 0x31, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x53, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x20,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x12, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12,
	0x22, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12,
	0x22, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12,
	0x22, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x70, 0x70,
	0x69, 0x6c, 0x79, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x64, 0x61, 0x64, 0x2f, 0x6f, 0x6c,
	0x64, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x33, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_statistics_statistics_proto_rawDescOnce sync.Once
	file_proto_statistics_statistics_proto_rawDescData = file_proto_statistics_statistics_proto_rawDesc
)

func file_proto_statistics_statistics_proto_rawDescGZIP() []byte {
	file_proto_statistics_statistics_proto_rawDescOnce.Do(func() {
		file_proto_statistics_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_statistics_statistics_proto_rawDescData)
	})
	return file_proto_statistics_statistics_proto_rawDescData
}

var file_proto_statistics_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_statistics_statistics_proto_goTypes = []interface{}{
	(*GetStatisticRequest)(nil),    // 0: statistics.GetStatisticRequest
	(*GetStatisticsRequest)(nil),   // 1: statistics.GetStatisticsRequest
	(*GetStatisticsReply)(nil),     // 2: statistics.GetStatisticsReply
	(*CreateStatisticRequest)(nil), // 3: statistics.CreateStatisticRequest
	(*UpdateStatisticRequest)(nil), // 4: statistics.UpdateStatisticRequest
	(*DeleteStatisticRequest)(nil), // 5: statistics.DeleteStatisticRequest
	(*Statistic)(nil),              // 6: statistics.Statistic
	(*EmptyReply)(nil),             // 7: statistics.EmptyReply
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_proto_statistics_statistics_proto_depIdxs = []int32{
	6, // 0: statistics.GetStatisticsReply.statistics:type_name -> statistics.Statistic
	8, // 1: statistics.Statistic.createdAt:type_name -> google.protobuf.Timestamp
	1, // 2: statistics.V1Statistics.GetStatistics:input_type -> statistics.GetStatisticsRequest
	0, // 3: statistics.V1Statistics.GetStatistic:input_type -> statistics.GetStatisticRequest
	3, // 4: statistics.V1Statistics.CreateStatistic:input_type -> statistics.CreateStatisticRequest
	4, // 5: statistics.V1Statistics.UpdateStatistic:input_type -> statistics.UpdateStatisticRequest
	5, // 6: statistics.V1Statistics.DeleteStatistic:input_type -> statistics.DeleteStatisticRequest
	2, // 7: statistics.V1Statistics.GetStatistics:output_type -> statistics.GetStatisticsReply
	6, // 8: statistics.V1Statistics.GetStatistic:output_type -> statistics.Statistic
	6, // 9: statistics.V1Statistics.CreateStatistic:output_type -> statistics.Statistic
	6, // 10: statistics.V1Statistics.UpdateStatistic:output_type -> statistics.Statistic
	7, // 11: statistics.V1Statistics.DeleteStatistic:output_type -> statistics.EmptyReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_statistics_statistics_proto_init() }
func file_proto_statistics_statistics_proto_init() {
	if File_proto_statistics_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_statistics_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatisticRequest); i {
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
		file_proto_statistics_statistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatisticsRequest); i {
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
		file_proto_statistics_statistics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatisticsReply); i {
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
		file_proto_statistics_statistics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStatisticRequest); i {
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
		file_proto_statistics_statistics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatisticRequest); i {
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
		file_proto_statistics_statistics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStatisticRequest); i {
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
		file_proto_statistics_statistics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statistic); i {
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
		file_proto_statistics_statistics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_statistics_statistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_statistics_statistics_proto_goTypes,
		DependencyIndexes: file_proto_statistics_statistics_proto_depIdxs,
		MessageInfos:      file_proto_statistics_statistics_proto_msgTypes,
	}.Build()
	File_proto_statistics_statistics_proto = out.File
	file_proto_statistics_statistics_proto_rawDesc = nil
	file_proto_statistics_statistics_proto_goTypes = nil
	file_proto_statistics_statistics_proto_depIdxs = nil
}
