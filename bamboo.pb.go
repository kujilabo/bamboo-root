// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: bamboo.proto

package bamboo

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

type ResponseType int32

const (
	ResponseType_HEARTBEAT ResponseType = 0
	ResponseType_DATA      ResponseType = 1
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "HEARTBEAT",
		1: "DATA",
	}
	ResponseType_value = map[string]int32{
		"HEARTBEAT": 0,
		"DATA":      1,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_bamboo_proto_enumTypes[0].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_bamboo_proto_enumTypes[0]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_bamboo_proto_rawDescGZIP(), []int{0}
}

type WorkerParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carrier              map[string]string `protobuf:"bytes,1,rep,name=carrier,proto3" json:"carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResultChannel        string            `protobuf:"bytes,3,opt,name=resultChannel,proto3" json:"resultChannel,omitempty"`
	HeartbeatIntervalSec int32             `protobuf:"varint,4,opt,name=heartbeatIntervalSec,proto3" json:"heartbeatIntervalSec,omitempty"`
	JobTimeoutSec        int32             `protobuf:"varint,5,opt,name=jobTimeoutSec,proto3" json:"jobTimeoutSec,omitempty"`
	Data                 []byte            `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WorkerParameter) Reset() {
	*x = WorkerParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bamboo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerParameter) ProtoMessage() {}

func (x *WorkerParameter) ProtoReflect() protoreflect.Message {
	mi := &file_bamboo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerParameter.ProtoReflect.Descriptor instead.
func (*WorkerParameter) Descriptor() ([]byte, []int) {
	return file_bamboo_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerParameter) GetCarrier() map[string]string {
	if x != nil {
		return x.Carrier
	}
	return nil
}

func (x *WorkerParameter) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *WorkerParameter) GetResultChannel() string {
	if x != nil {
		return x.ResultChannel
	}
	return ""
}

func (x *WorkerParameter) GetHeartbeatIntervalSec() int32 {
	if x != nil {
		return x.HeartbeatIntervalSec
	}
	return 0
}

func (x *WorkerParameter) GetJobTimeoutSec() int32 {
	if x != nil {
		return x.JobTimeoutSec
	}
	return 0
}

func (x *WorkerParameter) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type WorkerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ResponseType `protobuf:"varint,1,opt,name=type,proto3,enum=bamboo.ResponseType" json:"type,omitempty"`
	Data []byte       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WorkerResponse) Reset() {
	*x = WorkerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bamboo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerResponse) ProtoMessage() {}

func (x *WorkerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bamboo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerResponse.ProtoReflect.Descriptor instead.
func (*WorkerResponse) Descriptor() ([]byte, []int) {
	return file_bamboo_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerResponse) GetType() ResponseType {
	if x != nil {
		return x.Type
	}
	return ResponseType_HEARTBEAT
}

func (x *WorkerResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_bamboo_proto protoreflect.FileDescriptor

var file_bamboo_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x6d, 0x62, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x61, 0x6d, 0x62, 0x6f, 0x6f, 0x22, 0x9d, 0x03, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x07, 0x63, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x61,
	0x6d, 0x62, 0x6f, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x61,
	0x6d, 0x62, 0x6f, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x32, 0x0a, 0x14, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x53, 0x65, 0x63, 0x12, 0x24, 0x0a, 0x0d, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x53, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6a, 0x6f, 0x62,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3a,
	0x0a, 0x0c, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4e, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x6d, 0x62, 0x6f, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x27, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42,
	0x45, 0x41, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x42,
	0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75,
	0x6a, 0x69, 0x6c, 0x61, 0x62, 0x6f, 0x2f, 0x62, 0x61, 0x6d, 0x62, 0x6f, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bamboo_proto_rawDescOnce sync.Once
	file_bamboo_proto_rawDescData = file_bamboo_proto_rawDesc
)

func file_bamboo_proto_rawDescGZIP() []byte {
	file_bamboo_proto_rawDescOnce.Do(func() {
		file_bamboo_proto_rawDescData = protoimpl.X.CompressGZIP(file_bamboo_proto_rawDescData)
	})
	return file_bamboo_proto_rawDescData
}

var file_bamboo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bamboo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bamboo_proto_goTypes = []interface{}{
	(ResponseType)(0),       // 0: bamboo.ResponseType
	(*WorkerParameter)(nil), // 1: bamboo.WorkerParameter
	(*WorkerResponse)(nil),  // 2: bamboo.WorkerResponse
	nil,                     // 3: bamboo.WorkerParameter.CarrierEntry
	nil,                     // 4: bamboo.WorkerParameter.HeadersEntry
}
var file_bamboo_proto_depIdxs = []int32{
	3, // 0: bamboo.WorkerParameter.carrier:type_name -> bamboo.WorkerParameter.CarrierEntry
	4, // 1: bamboo.WorkerParameter.headers:type_name -> bamboo.WorkerParameter.HeadersEntry
	0, // 2: bamboo.WorkerResponse.type:type_name -> bamboo.ResponseType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bamboo_proto_init() }
func file_bamboo_proto_init() {
	if File_bamboo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bamboo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerParameter); i {
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
		file_bamboo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerResponse); i {
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
			RawDescriptor: file_bamboo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bamboo_proto_goTypes,
		DependencyIndexes: file_bamboo_proto_depIdxs,
		EnumInfos:         file_bamboo_proto_enumTypes,
		MessageInfos:      file_bamboo_proto_msgTypes,
	}.Build()
	File_bamboo_proto = out.File
	file_bamboo_proto_rawDesc = nil
	file_bamboo_proto_goTypes = nil
	file_bamboo_proto_depIdxs = nil
}
