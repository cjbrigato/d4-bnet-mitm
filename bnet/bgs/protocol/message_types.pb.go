// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/message_types.proto

package protocol

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

type TypingIndicator int32

const (
	TypingIndicator_TYPING_START TypingIndicator = 0
	TypingIndicator_TYPING_STOP  TypingIndicator = 1
)

// Enum value maps for TypingIndicator.
var (
	TypingIndicator_name = map[int32]string{
		0: "TYPING_START",
		1: "TYPING_STOP",
	}
	TypingIndicator_value = map[string]int32{
		"TYPING_START": 0,
		"TYPING_STOP":  1,
	}
)

func (x TypingIndicator) Enum() *TypingIndicator {
	p := new(TypingIndicator)
	*p = x
	return p
}

func (x TypingIndicator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypingIndicator) Descriptor() protoreflect.EnumDescriptor {
	return file_bgs_low_pb_client_message_types_proto_enumTypes[0].Descriptor()
}

func (TypingIndicator) Type() protoreflect.EnumType {
	return &file_bgs_low_pb_client_message_types_proto_enumTypes[0]
}

func (x TypingIndicator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TypingIndicator) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TypingIndicator(num)
	return nil
}

// Deprecated: Use TypingIndicator.Descriptor instead.
func (TypingIndicator) EnumDescriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_message_types_proto_rawDescGZIP(), []int{0}
}

type MessageId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch    *uint64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	Position *uint64 `protobuf:"varint,2,opt,name=position" json:"position,omitempty"`
}

func (x *MessageId) Reset() {
	*x = MessageId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_message_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageId) ProtoMessage() {}

func (x *MessageId) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_message_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageId.ProtoReflect.Descriptor instead.
func (*MessageId) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_message_types_proto_rawDescGZIP(), []int{0}
}

func (x *MessageId) GetEpoch() uint64 {
	if x != nil && x.Epoch != nil {
		return *x.Epoch
	}
	return 0
}

func (x *MessageId) GetPosition() uint64 {
	if x != nil && x.Position != nil {
		return *x.Position
	}
	return 0
}

var File_bgs_low_pb_client_message_types_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_message_types_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x3d, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x34, 0x0a, 0x0f, 0x54, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61,
	0x74, 0x6f, 0x2f, 0x64, 0x34, 0x2d, 0x62, 0x6e, 0x65, 0x74, 0x2d, 0x6d, 0x69, 0x74, 0x6d, 0x2f,
	0x62, 0x6e, 0x65, 0x74, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c,
}

var (
	file_bgs_low_pb_client_message_types_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_message_types_proto_rawDescData = file_bgs_low_pb_client_message_types_proto_rawDesc
)

func file_bgs_low_pb_client_message_types_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_message_types_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_message_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_message_types_proto_rawDescData)
	})
	return file_bgs_low_pb_client_message_types_proto_rawDescData
}

var file_bgs_low_pb_client_message_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bgs_low_pb_client_message_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bgs_low_pb_client_message_types_proto_goTypes = []interface{}{
	(TypingIndicator)(0), // 0: bgs.protocol.TypingIndicator
	(*MessageId)(nil),    // 1: bgs.protocol.MessageId
}
var file_bgs_low_pb_client_message_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_message_types_proto_init() }
func file_bgs_low_pb_client_message_types_proto_init() {
	if File_bgs_low_pb_client_message_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_message_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageId); i {
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
			RawDescriptor: file_bgs_low_pb_client_message_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_message_types_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_message_types_proto_depIdxs,
		EnumInfos:         file_bgs_low_pb_client_message_types_proto_enumTypes,
		MessageInfos:      file_bgs_low_pb_client_message_types_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_message_types_proto = out.File
	file_bgs_low_pb_client_message_types_proto_rawDesc = nil
	file_bgs_low_pb_client_message_types_proto_goTypes = nil
	file_bgs_low_pb_client_message_types_proto_depIdxs = nil
}
