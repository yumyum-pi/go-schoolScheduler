// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: sequence.proto

package models

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type GRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerID string   `protobuf:"bytes,1,opt,name=serverID,proto3" json:"serverID,omitempty"`
	ClientID string   `protobuf:"bytes,2,opt,name=clientID,proto3" json:"clientID,omitempty"`
	GSize    int32    `protobuf:"varint,3,opt,name=gSize,proto3" json:"gSize,omitempty"`
	NNType   int32    `protobuf:"varint,4,opt,name=nNType,proto3" json:"nNType,omitempty"`
	Pkgs     [][]byte `protobuf:"bytes,5,rep,name=pkgs,proto3" json:"pkgs,omitempty"`
}

func (x *GRequest) Reset() {
	*x = GRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRequest) ProtoMessage() {}

func (x *GRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sequence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRequest.ProtoReflect.Descriptor instead.
func (*GRequest) Descriptor() ([]byte, []int) {
	return file_sequence_proto_rawDescGZIP(), []int{0}
}

func (x *GRequest) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *GRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *GRequest) GetGSize() int32 {
	if x != nil {
		return x.GSize
	}
	return 0
}

func (x *GRequest) GetNNType() int32 {
	if x != nil {
		return x.NNType
	}
	return 0
}

func (x *GRequest) GetPkgs() [][]byte {
	if x != nil {
		return x.Pkgs
	}
	return nil
}

type GResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NError int32    `protobuf:"varint,1,opt,name=nError,proto3" json:"nError,omitempty"`
	Pkgs   [][]byte `protobuf:"bytes,2,rep,name=pkgs,proto3" json:"pkgs,omitempty"`
}

func (x *GResponse) Reset() {
	*x = GResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GResponse) ProtoMessage() {}

func (x *GResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sequence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GResponse.ProtoReflect.Descriptor instead.
func (*GResponse) Descriptor() ([]byte, []int) {
	return file_sequence_proto_rawDescGZIP(), []int{1}
}

func (x *GResponse) GetNError() int32 {
	if x != nil {
		return x.NError
	}
	return 0
}

func (x *GResponse) GetPkgs() [][]byte {
	if x != nil {
		return x.Pkgs
	}
	return nil
}

var File_sequence_proto protoreflect.FileDescriptor

var file_sequence_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x47, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x4e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x4e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6b, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x6b, 0x67, 0x73, 0x22,
	0x37, 0x0a, 0x09, 0x47, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6b, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x04, 0x70, 0x6b, 0x67, 0x73, 0x32, 0x3f, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x54, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x6d, 0x79, 0x75, 0x6d, 0x2d, 0x70,
	0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sequence_proto_rawDescOnce sync.Once
	file_sequence_proto_rawDescData = file_sequence_proto_rawDesc
)

func file_sequence_proto_rawDescGZIP() []byte {
	file_sequence_proto_rawDescOnce.Do(func() {
		file_sequence_proto_rawDescData = protoimpl.X.CompressGZIP(file_sequence_proto_rawDescData)
	})
	return file_sequence_proto_rawDescData
}

var file_sequence_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sequence_proto_goTypes = []interface{}{
	(*GRequest)(nil),  // 0: models.GRequest
	(*GResponse)(nil), // 1: models.GResponse
}
var file_sequence_proto_depIdxs = []int32{
	0, // 0: models.Sequence.GenerateTT:input_type -> models.GRequest
	1, // 1: models.Sequence.GenerateTT:output_type -> models.GResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sequence_proto_init() }
func file_sequence_proto_init() {
	if File_sequence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sequence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRequest); i {
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
		file_sequence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GResponse); i {
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
			RawDescriptor: file_sequence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sequence_proto_goTypes,
		DependencyIndexes: file_sequence_proto_depIdxs,
		MessageInfos:      file_sequence_proto_msgTypes,
	}.Build()
	File_sequence_proto = out.File
	file_sequence_proto_rawDesc = nil
	file_sequence_proto_goTypes = nil
	file_sequence_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SequenceClient is the client API for Sequence service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SequenceClient interface {
	GenerateTT(ctx context.Context, in *GRequest, opts ...grpc.CallOption) (*GResponse, error)
}

type sequenceClient struct {
	cc grpc.ClientConnInterface
}

func NewSequenceClient(cc grpc.ClientConnInterface) SequenceClient {
	return &sequenceClient{cc}
}

func (c *sequenceClient) GenerateTT(ctx context.Context, in *GRequest, opts ...grpc.CallOption) (*GResponse, error) {
	out := new(GResponse)
	err := c.cc.Invoke(ctx, "/models.Sequence/GenerateTT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequenceServer is the server API for Sequence service.
type SequenceServer interface {
	GenerateTT(context.Context, *GRequest) (*GResponse, error)
}

// UnimplementedSequenceServer can be embedded to have forward compatible implementations.
type UnimplementedSequenceServer struct {
}

func (*UnimplementedSequenceServer) GenerateTT(context.Context, *GRequest) (*GResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTT not implemented")
}

func RegisterSequenceServer(s *grpc.Server, srv SequenceServer) {
	s.RegisterService(&_Sequence_serviceDesc, srv)
}

func _Sequence_GenerateTT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GenerateTT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Sequence/GenerateTT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GenerateTT(ctx, req.(*GRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sequence_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Sequence",
	HandlerType: (*SequenceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTT",
			Handler:    _Sequence_GenerateTT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sequence.proto",
}
