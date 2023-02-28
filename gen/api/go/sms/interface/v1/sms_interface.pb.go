// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: sms/interface/v1/sms_interface.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SendSmsCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *SendSmsCodeRequest) Reset() {
	*x = SendSmsCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsCodeRequest) ProtoMessage() {}

func (x *SendSmsCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsCodeRequest.ProtoReflect.Descriptor instead.
func (*SendSmsCodeRequest) Descriptor() ([]byte, []int) {
	return file_sms_interface_v1_sms_interface_proto_rawDescGZIP(), []int{0}
}

func (x *SendSmsCodeRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SendSmsCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendSmsCodeReply) Reset() {
	*x = SendSmsCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsCodeReply) ProtoMessage() {}

func (x *SendSmsCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsCodeReply.ProtoReflect.Descriptor instead.
func (*SendSmsCodeReply) Descriptor() ([]byte, []int) {
	return file_sms_interface_v1_sms_interface_proto_rawDescGZIP(), []int{1}
}

type VerifySmsCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifySmsCodeRequest) Reset() {
	*x = VerifySmsCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifySmsCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifySmsCodeRequest) ProtoMessage() {}

func (x *VerifySmsCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifySmsCodeRequest.ProtoReflect.Descriptor instead.
func (*VerifySmsCodeRequest) Descriptor() ([]byte, []int) {
	return file_sms_interface_v1_sms_interface_proto_rawDescGZIP(), []int{2}
}

func (x *VerifySmsCodeRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *VerifySmsCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifySmsCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifySmsCodeReply) Reset() {
	*x = VerifySmsCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifySmsCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifySmsCodeReply) ProtoMessage() {}

func (x *VerifySmsCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_sms_interface_v1_sms_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifySmsCodeReply.ProtoReflect.Descriptor instead.
func (*VerifySmsCodeReply) Descriptor() ([]byte, []int) {
	return file_sms_interface_v1_sms_interface_proto_rawDescGZIP(), []int{3}
}

var File_sms_interface_v1_sms_interface_proto protoreflect.FileDescriptor

var file_sms_interface_v1_sms_interface_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4c, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53,
	0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x06, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x6d, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x8a, 0x02, 0x0a, 0x03, 0x53, 0x6d,
	0x73, 0x12, 0x7c, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x84, 0x01, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53,
	0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x6d, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6c, 0x69, 0x66, 0x65, 0x69, 0x65, 0x72, 0x2f, 0x76,
	0x76, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_interface_v1_sms_interface_proto_rawDescOnce sync.Once
	file_sms_interface_v1_sms_interface_proto_rawDescData = file_sms_interface_v1_sms_interface_proto_rawDesc
)

func file_sms_interface_v1_sms_interface_proto_rawDescGZIP() []byte {
	file_sms_interface_v1_sms_interface_proto_rawDescOnce.Do(func() {
		file_sms_interface_v1_sms_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_interface_v1_sms_interface_proto_rawDescData)
	})
	return file_sms_interface_v1_sms_interface_proto_rawDescData
}

var file_sms_interface_v1_sms_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sms_interface_v1_sms_interface_proto_goTypes = []interface{}{
	(*SendSmsCodeRequest)(nil),   // 0: api.sms.interface.v1.SendSmsCodeRequest
	(*SendSmsCodeReply)(nil),     // 1: api.sms.interface.v1.SendSmsCodeReply
	(*VerifySmsCodeRequest)(nil), // 2: api.sms.interface.v1.VerifySmsCodeRequest
	(*VerifySmsCodeReply)(nil),   // 3: api.sms.interface.v1.VerifySmsCodeReply
}
var file_sms_interface_v1_sms_interface_proto_depIdxs = []int32{
	0, // 0: api.sms.interface.v1.Sms.SendSmsCode:input_type -> api.sms.interface.v1.SendSmsCodeRequest
	2, // 1: api.sms.interface.v1.Sms.VerifySmsCode:input_type -> api.sms.interface.v1.VerifySmsCodeRequest
	1, // 2: api.sms.interface.v1.Sms.SendSmsCode:output_type -> api.sms.interface.v1.SendSmsCodeReply
	3, // 3: api.sms.interface.v1.Sms.VerifySmsCode:output_type -> api.sms.interface.v1.VerifySmsCodeReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_interface_v1_sms_interface_proto_init() }
func file_sms_interface_v1_sms_interface_proto_init() {
	if File_sms_interface_v1_sms_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_interface_v1_sms_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsCodeRequest); i {
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
		file_sms_interface_v1_sms_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsCodeReply); i {
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
		file_sms_interface_v1_sms_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifySmsCodeRequest); i {
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
		file_sms_interface_v1_sms_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifySmsCodeReply); i {
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
			RawDescriptor: file_sms_interface_v1_sms_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_interface_v1_sms_interface_proto_goTypes,
		DependencyIndexes: file_sms_interface_v1_sms_interface_proto_depIdxs,
		MessageInfos:      file_sms_interface_v1_sms_interface_proto_msgTypes,
	}.Build()
	File_sms_interface_v1_sms_interface_proto = out.File
	file_sms_interface_v1_sms_interface_proto_rawDesc = nil
	file_sms_interface_v1_sms_interface_proto_goTypes = nil
	file_sms_interface_v1_sms_interface_proto_depIdxs = nil
}
