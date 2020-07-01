// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: ProdService.proto

package Services

import (
	proto "github.com/golang/protobuf/proto"
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

type ProdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag:json:"size",form:"size"
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size" form:"size"`
}

func (x *ProdsRequest) Reset() {
	*x = ProdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdsRequest) ProtoMessage() {}

func (x *ProdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProdService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdsRequest.ProtoReflect.Descriptor instead.
func (*ProdsRequest) Descriptor() ([]byte, []int) {
	return file_ProdService_proto_rawDescGZIP(), []int{0}
}

func (x *ProdsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ProdListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProdModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProdListResponse) Reset() {
	*x = ProdListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdListResponse) ProtoMessage() {}

func (x *ProdListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProdService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdListResponse.ProtoReflect.Descriptor instead.
func (*ProdListResponse) Descriptor() ([]byte, []int) {
	return file_ProdService_proto_rawDescGZIP(), []int{1}
}

func (x *ProdListResponse) GetData() []*ProdModel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ProdService_proto protoreflect.FileDescriptor

var file_ProdService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x3b, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x51, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProdService_proto_rawDescOnce sync.Once
	file_ProdService_proto_rawDescData = file_ProdService_proto_rawDesc
)

func file_ProdService_proto_rawDescGZIP() []byte {
	file_ProdService_proto_rawDescOnce.Do(func() {
		file_ProdService_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProdService_proto_rawDescData)
	})
	return file_ProdService_proto_rawDescData
}

var file_ProdService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ProdService_proto_goTypes = []interface{}{
	(*ProdsRequest)(nil),     // 0: Services.ProdsRequest
	(*ProdListResponse)(nil), // 1: Services.ProdListResponse
	(*ProdModel)(nil),        // 2: Services.ProdModel
}
var file_ProdService_proto_depIdxs = []int32{
	2, // 0: Services.ProdListResponse.data:type_name -> Services.ProdModel
	0, // 1: Services.ProdService.GetProdsList:input_type -> Services.ProdsRequest
	1, // 2: Services.ProdService.GetProdsList:output_type -> Services.ProdListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ProdService_proto_init() }
func file_ProdService_proto_init() {
	if File_ProdService_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ProdService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdsRequest); i {
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
		file_ProdService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdListResponse); i {
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
			RawDescriptor: file_ProdService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ProdService_proto_goTypes,
		DependencyIndexes: file_ProdService_proto_depIdxs,
		MessageInfos:      file_ProdService_proto_msgTypes,
	}.Build()
	File_ProdService_proto = out.File
	file_ProdService_proto_rawDesc = nil
	file_ProdService_proto_goTypes = nil
	file_ProdService_proto_depIdxs = nil
}
