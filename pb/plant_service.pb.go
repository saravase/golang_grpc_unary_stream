// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.8.0
// source: plant_service.proto

package pb

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

type CreatePlantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plant *Plant `protobuf:"bytes,1,opt,name=plant,proto3" json:"plant,omitempty"`
}

func (x *CreatePlantRequest) Reset() {
	*x = CreatePlantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plant_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlantRequest) ProtoMessage() {}

func (x *CreatePlantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plant_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlantRequest.ProtoReflect.Descriptor instead.
func (*CreatePlantRequest) Descriptor() ([]byte, []int) {
	return file_plant_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlantRequest) GetPlant() *Plant {
	if x != nil {
		return x.Plant
	}
	return nil
}

type CreatePlantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePlantResponse) Reset() {
	*x = CreatePlantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plant_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlantResponse) ProtoMessage() {}

func (x *CreatePlantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plant_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlantResponse.ProtoReflect.Descriptor instead.
func (*CreatePlantResponse) Descriptor() ([]byte, []int) {
	return file_plant_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlantResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_plant_service_proto protoreflect.FileDescriptor

var file_plant_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x1a, 0x13, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x6e,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x22, 0x25, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x88, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x04, 0x5a, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plant_service_proto_rawDescOnce sync.Once
	file_plant_service_proto_rawDescData = file_plant_service_proto_rawDesc
)

func file_plant_service_proto_rawDescGZIP() []byte {
	file_plant_service_proto_rawDescOnce.Do(func() {
		file_plant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_plant_service_proto_rawDescData)
	})
	return file_plant_service_proto_rawDescData
}

var file_plant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_plant_service_proto_goTypes = []interface{}{
	(*CreatePlantRequest)(nil),  // 0: golang_grpc_unary_stream.plant.CreatePlantRequest
	(*CreatePlantResponse)(nil), // 1: golang_grpc_unary_stream.plant.CreatePlantResponse
	(*Plant)(nil),               // 2: golang_grpc_unary_stream.plant.Plant
}
var file_plant_service_proto_depIdxs = []int32{
	2, // 0: golang_grpc_unary_stream.plant.CreatePlantRequest.plant:type_name -> golang_grpc_unary_stream.plant.Plant
	0, // 1: golang_grpc_unary_stream.plant.PlantService.CreatePlant:input_type -> golang_grpc_unary_stream.plant.CreatePlantRequest
	1, // 2: golang_grpc_unary_stream.plant.PlantService.CreatePlant:output_type -> golang_grpc_unary_stream.plant.CreatePlantResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_plant_service_proto_init() }
func file_plant_service_proto_init() {
	if File_plant_service_proto != nil {
		return
	}
	file_plant_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plant_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlantRequest); i {
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
		file_plant_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlantResponse); i {
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
			RawDescriptor: file_plant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plant_service_proto_goTypes,
		DependencyIndexes: file_plant_service_proto_depIdxs,
		MessageInfos:      file_plant_service_proto_msgTypes,
	}.Build()
	File_plant_service_proto = out.File
	file_plant_service_proto_rawDesc = nil
	file_plant_service_proto_goTypes = nil
	file_plant_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlantServiceClient is the client API for PlantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlantServiceClient interface {
	CreatePlant(ctx context.Context, in *CreatePlantRequest, opts ...grpc.CallOption) (*CreatePlantResponse, error)
}

type plantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlantServiceClient(cc grpc.ClientConnInterface) PlantServiceClient {
	return &plantServiceClient{cc}
}

func (c *plantServiceClient) CreatePlant(ctx context.Context, in *CreatePlantRequest, opts ...grpc.CallOption) (*CreatePlantResponse, error) {
	out := new(CreatePlantResponse)
	err := c.cc.Invoke(ctx, "/golang_grpc_unary_stream.plant.PlantService/CreatePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlantServiceServer is the server API for PlantService service.
type PlantServiceServer interface {
	CreatePlant(context.Context, *CreatePlantRequest) (*CreatePlantResponse, error)
}

// UnimplementedPlantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlantServiceServer struct {
}

func (*UnimplementedPlantServiceServer) CreatePlant(context.Context, *CreatePlantRequest) (*CreatePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlant not implemented")
}

func RegisterPlantServiceServer(s *grpc.Server, srv PlantServiceServer) {
	s.RegisterService(&_PlantService_serviceDesc, srv)
}

func _PlantService_CreatePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantServiceServer).CreatePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golang_grpc_unary_stream.plant.PlantService/CreatePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantServiceServer).CreatePlant(ctx, req.(*CreatePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "golang_grpc_unary_stream.plant.PlantService",
	HandlerType: (*PlantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlant",
			Handler:    _PlantService_CreatePlant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plant_service.proto",
}
