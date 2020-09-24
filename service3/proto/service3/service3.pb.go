// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/service3/service3.proto

package go_micro_service_service3

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

type SerialNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SerialNum) Reset() {
	*x = SerialNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service3_service3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerialNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerialNum) ProtoMessage() {}

func (x *SerialNum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service3_service3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerialNum.ProtoReflect.Descriptor instead.
func (*SerialNum) Descriptor() ([]byte, []int) {
	return file_proto_service3_service3_proto_rawDescGZIP(), []int{0}
}

func (x *SerialNum) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temp1    float32 `protobuf:"fixed32,2,opt,name=temp1,proto3" json:"temp1,omitempty"`
	Temp2    float32 `protobuf:"fixed32,3,opt,name=temp2,proto3" json:"temp2,omitempty"`
	Height   float32 `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
	Latitude float32 `protobuf:"fixed32,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Logitude float32 `protobuf:"fixed32,6,opt,name=logitude,proto3" json:"logitude,omitempty"`
	Status   int32   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service3_service3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service3_service3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_proto_service3_service3_proto_rawDescGZIP(), []int{1}
}

func (x *Values) GetTemp1() float32 {
	if x != nil {
		return x.Temp1
	}
	return 0
}

func (x *Values) GetTemp2() float32 {
	if x != nil {
		return x.Temp2
	}
	return 0
}

func (x *Values) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Values) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Values) GetLogitude() float32 {
	if x != nil {
		return x.Logitude
	}
	return 0
}

func (x *Values) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     int64   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Values *Values `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service3_service3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service3_service3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_service3_service3_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *Response) GetValues() *Values {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_proto_service3_service3_proto protoreflect.FileDescriptor

var file_proto_service3_service3_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x33,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x33, 0x22, 0x21, 0x0a, 0x09, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9c, 0x01,
	0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x6d, 0x70,
	0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x31, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74,
	0x65, 0x6d, 0x70, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x32, 0x62, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x33, 0x12,
	0x56, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x33, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service3_service3_proto_rawDescOnce sync.Once
	file_proto_service3_service3_proto_rawDescData = file_proto_service3_service3_proto_rawDesc
)

func file_proto_service3_service3_proto_rawDescGZIP() []byte {
	file_proto_service3_service3_proto_rawDescOnce.Do(func() {
		file_proto_service3_service3_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service3_service3_proto_rawDescData)
	})
	return file_proto_service3_service3_proto_rawDescData
}

var file_proto_service3_service3_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_service3_service3_proto_goTypes = []interface{}{
	(*SerialNum)(nil), // 0: go.micro.service.service3.SerialNum
	(*Values)(nil),    // 1: go.micro.service.service3.Values
	(*Response)(nil),  // 2: go.micro.service.service3.Response
}
var file_proto_service3_service3_proto_depIdxs = []int32{
	1, // 0: go.micro.service.service3.Response.values:type_name -> go.micro.service.service3.Values
	0, // 1: go.micro.service.service3.Service3.Telemetry:input_type -> go.micro.service.service3.SerialNum
	2, // 2: go.micro.service.service3.Service3.Telemetry:output_type -> go.micro.service.service3.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_service3_service3_proto_init() }
func file_proto_service3_service3_proto_init() {
	if File_proto_service3_service3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service3_service3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerialNum); i {
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
		file_proto_service3_service3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Values); i {
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
		file_proto_service3_service3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_service3_service3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service3_service3_proto_goTypes,
		DependencyIndexes: file_proto_service3_service3_proto_depIdxs,
		MessageInfos:      file_proto_service3_service3_proto_msgTypes,
	}.Build()
	File_proto_service3_service3_proto = out.File
	file_proto_service3_service3_proto_rawDesc = nil
	file_proto_service3_service3_proto_goTypes = nil
	file_proto_service3_service3_proto_depIdxs = nil
}
