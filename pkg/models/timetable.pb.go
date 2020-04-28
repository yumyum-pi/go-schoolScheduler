// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: pkg/models/timetable.proto

package models

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TimeTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NDays    int32    `protobuf:"varint,1,opt,name=nDays,proto3" json:"nDays,omitempty"`       // no of day in a week
	NPeriods int32    `protobuf:"varint,2,opt,name=nPeriods,proto3" json:"nPeriods,omitempty"` // no of periods in a week
	Period   [][]byte `protobuf:"bytes,3,rep,name=period,proto3" json:"period,omitempty"`      // period package
}

func (x *TimeTable) Reset() {
	*x = TimeTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_models_timetable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeTable) ProtoMessage() {}

func (x *TimeTable) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_models_timetable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeTable.ProtoReflect.Descriptor instead.
func (*TimeTable) Descriptor() ([]byte, []int) {
	return file_pkg_models_timetable_proto_rawDescGZIP(), []int{0}
}

func (x *TimeTable) GetNDays() int32 {
	if x != nil {
		return x.NDays
	}
	return 0
}

func (x *TimeTable) GetNPeriods() int32 {
	if x != nil {
		return x.NPeriods
	}
	return 0
}

func (x *TimeTable) GetPeriod() [][]byte {
	if x != nil {
		return x.Period
	}
	return nil
}

var File_pkg_models_timetable_proto protoreflect.FileDescriptor

var file_pkg_models_timetable_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x22, 0x55, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_models_timetable_proto_rawDescOnce sync.Once
	file_pkg_models_timetable_proto_rawDescData = file_pkg_models_timetable_proto_rawDesc
)

func file_pkg_models_timetable_proto_rawDescGZIP() []byte {
	file_pkg_models_timetable_proto_rawDescOnce.Do(func() {
		file_pkg_models_timetable_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_models_timetable_proto_rawDescData)
	})
	return file_pkg_models_timetable_proto_rawDescData
}

var file_pkg_models_timetable_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_models_timetable_proto_goTypes = []interface{}{
	(*TimeTable)(nil), // 0: models.TimeTable
}
var file_pkg_models_timetable_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_models_timetable_proto_init() }
func file_pkg_models_timetable_proto_init() {
	if File_pkg_models_timetable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_models_timetable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeTable); i {
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
			RawDescriptor: file_pkg_models_timetable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_models_timetable_proto_goTypes,
		DependencyIndexes: file_pkg_models_timetable_proto_depIdxs,
		MessageInfos:      file_pkg_models_timetable_proto_msgTypes,
	}.Build()
	File_pkg_models_timetable_proto = out.File
	file_pkg_models_timetable_proto_rawDesc = nil
	file_pkg_models_timetable_proto_goTypes = nil
	file_pkg_models_timetable_proto_depIdxs = nil
}
