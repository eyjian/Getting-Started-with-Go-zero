// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.11.4
// source: add.proto

package add

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

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddReq) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResp) Reset() {
	*x = AddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResp) ProtoMessage() {}

func (x *AddResp) ProtoReflect() protoreflect.Message {
	mi := &file_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResp.ProtoReflect.Descriptor instead.
func (*AddResp) Descriptor() ([]byte, []int) {
	return file_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddResp) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_add_proto protoreflect.FileDescriptor

var file_add_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x64, 0x64,
	0x22, 0x24, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x1b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x73, 0x75, 0x6d, 0x32, 0x29, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x03,
	0x61, 0x64, 0x64, 0x12, 0x0b, 0x2e, 0x61, 0x64, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x61, 0x64, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x64, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_add_proto_rawDescOnce sync.Once
	file_add_proto_rawDescData = file_add_proto_rawDesc
)

func file_add_proto_rawDescGZIP() []byte {
	file_add_proto_rawDescOnce.Do(func() {
		file_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_add_proto_rawDescData)
	})
	return file_add_proto_rawDescData
}

var file_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),  // 0: add.AddReq
	(*AddResp)(nil), // 1: add.AddResp
}
var file_add_proto_depIdxs = []int32{
	0, // 0: add.Adder.add:input_type -> add.AddReq
	1, // 1: add.Adder.add:output_type -> add.AddResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_add_proto_init() }
func file_add_proto_init() {
	if File_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResp); i {
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
			RawDescriptor: file_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_add_proto_goTypes,
		DependencyIndexes: file_add_proto_depIdxs,
		MessageInfos:      file_add_proto_msgTypes,
	}.Build()
	File_add_proto = out.File
	file_add_proto_rawDesc = nil
	file_add_proto_goTypes = nil
	file_add_proto_depIdxs = nil
}
