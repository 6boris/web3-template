// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: service/data/defi.proto

package data

import (
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

type GetDefiPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceToken string `protobuf:"bytes,1,opt,name=source_token,json=sourceToken,proto3" json:"source_token,omitempty"`
	TargetToken string `protobuf:"bytes,2,opt,name=target_token,json=targetToken,proto3" json:"target_token,omitempty"`
}

func (x *GetDefiPriceRequest) Reset() {
	*x = GetDefiPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_data_defi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefiPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefiPriceRequest) ProtoMessage() {}

func (x *GetDefiPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_data_defi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefiPriceRequest.ProtoReflect.Descriptor instead.
func (*GetDefiPriceRequest) Descriptor() ([]byte, []int) {
	return file_service_data_defi_proto_rawDescGZIP(), []int{0}
}

func (x *GetDefiPriceRequest) GetSourceToken() string {
	if x != nil {
		return x.SourceToken
	}
	return ""
}

func (x *GetDefiPriceRequest) GetTargetToken() string {
	if x != nil {
		return x.TargetToken
	}
	return ""
}

type GetDefiPriceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *DefiPriceItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetDefiPriceReply) Reset() {
	*x = GetDefiPriceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_data_defi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefiPriceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefiPriceReply) ProtoMessage() {}

func (x *GetDefiPriceReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_data_defi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefiPriceReply.ProtoReflect.Descriptor instead.
func (*GetDefiPriceReply) Descriptor() ([]byte, []int) {
	return file_service_data_defi_proto_rawDescGZIP(), []int{1}
}

func (x *GetDefiPriceReply) GetItem() *DefiPriceItem {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_service_data_defi_proto protoreflect.FileDescriptor

var file_service_data_defi_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64,
	0x65, 0x66, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x48, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x32, 0x62, 0x0a, 0x04, 0x44,
	0x65, 0x66, 0x69, 0x12, 0x5a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x66, 0x69, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x30, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x50, 0x01, 0x5a, 0x1a, 0x77, 0x65, 0x62, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x3b, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_data_defi_proto_rawDescOnce sync.Once
	file_service_data_defi_proto_rawDescData = file_service_data_defi_proto_rawDesc
)

func file_service_data_defi_proto_rawDescGZIP() []byte {
	file_service_data_defi_proto_rawDescOnce.Do(func() {
		file_service_data_defi_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_data_defi_proto_rawDescData)
	})
	return file_service_data_defi_proto_rawDescData
}

var file_service_data_defi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_data_defi_proto_goTypes = []interface{}{
	(*GetDefiPriceRequest)(nil), // 0: api.service.data.GetDefiPriceRequest
	(*GetDefiPriceReply)(nil),   // 1: api.service.data.GetDefiPriceReply
	(*DefiPriceItem)(nil),       // 2: api.service.data.DefiPriceItem
}
var file_service_data_defi_proto_depIdxs = []int32{
	2, // 0: api.service.data.GetDefiPriceReply.item:type_name -> api.service.data.DefiPriceItem
	0, // 1: api.service.data.Defi.GetDefiPrice:input_type -> api.service.data.GetDefiPriceRequest
	1, // 2: api.service.data.Defi.GetDefiPrice:output_type -> api.service.data.GetDefiPriceReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_data_defi_proto_init() }
func file_service_data_defi_proto_init() {
	if File_service_data_defi_proto != nil {
		return
	}
	file_service_data_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_data_defi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefiPriceRequest); i {
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
		file_service_data_defi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefiPriceReply); i {
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
			RawDescriptor: file_service_data_defi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_data_defi_proto_goTypes,
		DependencyIndexes: file_service_data_defi_proto_depIdxs,
		MessageInfos:      file_service_data_defi_proto_msgTypes,
	}.Build()
	File_service_data_defi_proto = out.File
	file_service_data_defi_proto_rawDesc = nil
	file_service_data_defi_proto_goTypes = nil
	file_service_data_defi_proto_depIdxs = nil
}
