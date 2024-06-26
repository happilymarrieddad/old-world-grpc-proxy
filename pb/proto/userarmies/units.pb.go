// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: proto/userarmies/units.proto

package userarmies

import (
	unittypes "github.com/happilymarrieddad/grpc-proxy/pb/proto/unittypes"
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

type ArmyUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserArmyId   string                 `protobuf:"bytes,2,opt,name=userArmyId,proto3" json:"userArmyId,omitempty"`
	UserArmyName string                 `protobuf:"bytes,3,opt,name=userArmyName,proto3" json:"userArmyName,omitempty"`
	UnitTypeId   string                 `protobuf:"bytes,4,opt,name=unitTypeId,proto3" json:"unitTypeId,omitempty"`
	UnitType     *unittypes.UnitType    `protobuf:"bytes,5,opt,name=unitType,proto3" json:"unitType,omitempty"`
	Points       int64                  `protobuf:"varint,6,opt,name=points,proto3" json:"points,omitempty"`
	Quantity     int64                  `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ArmyUnit) Reset() {
	*x = ArmyUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userarmies_units_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArmyUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArmyUnit) ProtoMessage() {}

func (x *ArmyUnit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userarmies_units_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArmyUnit.ProtoReflect.Descriptor instead.
func (*ArmyUnit) Descriptor() ([]byte, []int) {
	return file_proto_userarmies_units_proto_rawDescGZIP(), []int{0}
}

func (x *ArmyUnit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArmyUnit) GetUserArmyId() string {
	if x != nil {
		return x.UserArmyId
	}
	return ""
}

func (x *ArmyUnit) GetUserArmyName() string {
	if x != nil {
		return x.UserArmyName
	}
	return ""
}

func (x *ArmyUnit) GetUnitTypeId() string {
	if x != nil {
		return x.UnitTypeId
	}
	return ""
}

func (x *ArmyUnit) GetUnitType() *unittypes.UnitType {
	if x != nil {
		return x.UnitType
	}
	return nil
}

func (x *ArmyUnit) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *ArmyUnit) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ArmyUnit) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_proto_userarmies_units_proto protoreflect.FileDescriptor

var file_proto_userarmies_units_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x72, 0x6d, 0x69,
	0x65, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x61, 0x72, 0x6d, 0x69, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x08, 0x41, 0x72, 0x6d,
	0x79, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x72, 0x6d,
	0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x72, 0x6d, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x41, 0x72, 0x6d,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x72, 0x6d, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x69,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x6e, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x75, 0x6e, 0x69,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x6e,
	0x69, 0x74, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x70, 0x70, 0x69, 0x6c, 0x79, 0x6d, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x64, 0x64, 0x61, 0x64, 0x2f, 0x6f, 0x6c, 0x64, 0x2d, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x33, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x72, 0x6d, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_userarmies_units_proto_rawDescOnce sync.Once
	file_proto_userarmies_units_proto_rawDescData = file_proto_userarmies_units_proto_rawDesc
)

func file_proto_userarmies_units_proto_rawDescGZIP() []byte {
	file_proto_userarmies_units_proto_rawDescOnce.Do(func() {
		file_proto_userarmies_units_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_userarmies_units_proto_rawDescData)
	})
	return file_proto_userarmies_units_proto_rawDescData
}

var file_proto_userarmies_units_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_userarmies_units_proto_goTypes = []interface{}{
	(*ArmyUnit)(nil),              // 0: userarmies.ArmyUnit
	(*unittypes.UnitType)(nil),    // 1: unittypes.UnitType
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_userarmies_units_proto_depIdxs = []int32{
	1, // 0: userarmies.ArmyUnit.unitType:type_name -> unittypes.UnitType
	2, // 1: userarmies.ArmyUnit.createdAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_userarmies_units_proto_init() }
func file_proto_userarmies_units_proto_init() {
	if File_proto_userarmies_units_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_userarmies_units_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArmyUnit); i {
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
			RawDescriptor: file_proto_userarmies_units_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_userarmies_units_proto_goTypes,
		DependencyIndexes: file_proto_userarmies_units_proto_depIdxs,
		MessageInfos:      file_proto_userarmies_units_proto_msgTypes,
	}.Build()
	File_proto_userarmies_units_proto = out.File
	file_proto_userarmies_units_proto_rawDesc = nil
	file_proto_userarmies_units_proto_goTypes = nil
	file_proto_userarmies_units_proto_depIdxs = nil
}
