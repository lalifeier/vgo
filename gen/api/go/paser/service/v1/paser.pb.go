// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: paser/service/v1/paser.proto

package v1

import (
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

type PaserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PaserReq) Reset() {
	*x = PaserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paser_service_v1_paser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaserReq) ProtoMessage() {}

func (x *PaserReq) ProtoReflect() protoreflect.Message {
	mi := &file_paser_service_v1_paser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaserReq.ProtoReflect.Descriptor instead.
func (*PaserReq) Descriptor() ([]byte, []int) {
	return file_paser_service_v1_paser_proto_rawDescGZIP(), []int{0}
}

func (x *PaserReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Part struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Ext  string `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *Part) Reset() {
	*x = Part{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paser_service_v1_paser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_paser_service_v1_paser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_paser_service_v1_paser_proto_rawDescGZIP(), []int{1}
}

func (x *Part) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Part) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Part) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quality string  `protobuf:"bytes,1,opt,name=quality,proto3" json:"quality,omitempty"`
	Parts   []*Part `protobuf:"bytes,2,rep,name=parts,proto3" json:"parts,omitempty"`
	Size    int64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Ext     string  `protobuf:"bytes,4,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paser_service_v1_paser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_paser_service_v1_paser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_paser_service_v1_paser_proto_rawDescGZIP(), []int{2}
}

func (x *Stream) GetQuality() string {
	if x != nil {
		return x.Quality
	}
	return ""
}

func (x *Stream) GetParts() []*Part {
	if x != nil {
		return x.Parts
	}
	return nil
}

func (x *Stream) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Stream) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type PaserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site    string             `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	Title   string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type    string             `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Streams map[string]*Stream `protobuf:"bytes,4,rep,name=streams,proto3" json:"streams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Url     string             `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PaserReply) Reset() {
	*x = PaserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paser_service_v1_paser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaserReply) ProtoMessage() {}

func (x *PaserReply) ProtoReflect() protoreflect.Message {
	mi := &file_paser_service_v1_paser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaserReply.ProtoReflect.Descriptor instead.
func (*PaserReply) Descriptor() ([]byte, []int) {
	return file_paser_service_v1_paser_proto_rawDescGZIP(), []int{3}
}

func (x *PaserReply) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *PaserReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PaserReply) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PaserReply) GetStreams() map[string]*Stream {
	if x != nil {
		return x.Streams
	}
	return nil
}

func (x *PaserReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_paser_service_v1_paser_proto protoreflect.FileDescriptor

var file_paser_service_v1_paser_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3e, 0x0a, 0x04,
	0x50, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x22, 0x76, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x2c, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x78, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x43, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x54, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x61, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x5d,
	0x0a, 0x05, 0x50, 0x61, 0x73, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x05, 0x50, 0x61, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x70, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70,
	0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x73, 0x65, 0x72, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6c, 0x69,
	0x66, 0x65, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x76, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x73, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paser_service_v1_paser_proto_rawDescOnce sync.Once
	file_paser_service_v1_paser_proto_rawDescData = file_paser_service_v1_paser_proto_rawDesc
)

func file_paser_service_v1_paser_proto_rawDescGZIP() []byte {
	file_paser_service_v1_paser_proto_rawDescOnce.Do(func() {
		file_paser_service_v1_paser_proto_rawDescData = protoimpl.X.CompressGZIP(file_paser_service_v1_paser_proto_rawDescData)
	})
	return file_paser_service_v1_paser_proto_rawDescData
}

var file_paser_service_v1_paser_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_paser_service_v1_paser_proto_goTypes = []interface{}{
	(*PaserReq)(nil),   // 0: paser.service.v1.PaserReq
	(*Part)(nil),       // 1: paser.service.v1.Part
	(*Stream)(nil),     // 2: paser.service.v1.Stream
	(*PaserReply)(nil), // 3: paser.service.v1.PaserReply
	nil,                // 4: paser.service.v1.PaserReply.StreamsEntry
}
var file_paser_service_v1_paser_proto_depIdxs = []int32{
	1, // 0: paser.service.v1.Stream.parts:type_name -> paser.service.v1.Part
	4, // 1: paser.service.v1.PaserReply.streams:type_name -> paser.service.v1.PaserReply.StreamsEntry
	2, // 2: paser.service.v1.PaserReply.StreamsEntry.value:type_name -> paser.service.v1.Stream
	0, // 3: paser.service.v1.Paser.Paser:input_type -> paser.service.v1.PaserReq
	3, // 4: paser.service.v1.Paser.Paser:output_type -> paser.service.v1.PaserReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_paser_service_v1_paser_proto_init() }
func file_paser_service_v1_paser_proto_init() {
	if File_paser_service_v1_paser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paser_service_v1_paser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaserReq); i {
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
		file_paser_service_v1_paser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Part); i {
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
		file_paser_service_v1_paser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_paser_service_v1_paser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaserReply); i {
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
			RawDescriptor: file_paser_service_v1_paser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paser_service_v1_paser_proto_goTypes,
		DependencyIndexes: file_paser_service_v1_paser_proto_depIdxs,
		MessageInfos:      file_paser_service_v1_paser_proto_msgTypes,
	}.Build()
	File_paser_service_v1_paser_proto = out.File
	file_paser_service_v1_paser_proto_rawDesc = nil
	file_paser_service_v1_paser_proto_goTypes = nil
	file_paser_service_v1_paser_proto_depIdxs = nil
}
