// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/attribute_types.proto

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

type AttributeFilter_Operation int32

const (
	AttributeFilter_MATCH_NONE              AttributeFilter_Operation = 0
	AttributeFilter_MATCH_ANY               AttributeFilter_Operation = 1
	AttributeFilter_MATCH_ALL               AttributeFilter_Operation = 2
	AttributeFilter_MATCH_ALL_MOST_SPECIFIC AttributeFilter_Operation = 3
)

// Enum value maps for AttributeFilter_Operation.
var (
	AttributeFilter_Operation_name = map[int32]string{
		0: "MATCH_NONE",
		1: "MATCH_ANY",
		2: "MATCH_ALL",
		3: "MATCH_ALL_MOST_SPECIFIC",
	}
	AttributeFilter_Operation_value = map[string]int32{
		"MATCH_NONE":              0,
		"MATCH_ANY":               1,
		"MATCH_ALL":               2,
		"MATCH_ALL_MOST_SPECIFIC": 3,
	}
)

func (x AttributeFilter_Operation) Enum() *AttributeFilter_Operation {
	p := new(AttributeFilter_Operation)
	*p = x
	return p
}

func (x AttributeFilter_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeFilter_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_bgs_low_pb_client_attribute_types_proto_enumTypes[0].Descriptor()
}

func (AttributeFilter_Operation) Type() protoreflect.EnumType {
	return &file_bgs_low_pb_client_attribute_types_proto_enumTypes[0]
}

func (x AttributeFilter_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AttributeFilter_Operation) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AttributeFilter_Operation(num)
	return nil
}

// Deprecated: Use AttributeFilter_Operation.Descriptor instead.
func (AttributeFilter_Operation) EnumDescriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_attribute_types_proto_rawDescGZIP(), []int{2, 0}
}

type Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolValue     *bool     `protobuf:"varint,2,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	IntValue      *int64    `protobuf:"varint,3,opt,name=int_value,json=intValue" json:"int_value,omitempty"`
	FloatValue    *float64  `protobuf:"fixed64,4,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	StringValue   *string   `protobuf:"bytes,5,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	BlobValue     []byte    `protobuf:"bytes,6,opt,name=blob_value,json=blobValue" json:"blob_value,omitempty"`
	MessageValue  []byte    `protobuf:"bytes,7,opt,name=message_value,json=messageValue" json:"message_value,omitempty"`
	FourccValue   *string   `protobuf:"bytes,8,opt,name=fourcc_value,json=fourccValue" json:"fourcc_value,omitempty"`
	UintValue     *uint64   `protobuf:"varint,9,opt,name=uint_value,json=uintValue" json:"uint_value,omitempty"`
	EntityIdValue *EntityId `protobuf:"bytes,10,opt,name=entity_id_value,json=entityIdValue" json:"entity_id_value,omitempty"`
}

func (x *Variant) Reset() {
	*x = Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_attribute_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_attribute_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_attribute_types_proto_rawDescGZIP(), []int{0}
}

func (x *Variant) GetBoolValue() bool {
	if x != nil && x.BoolValue != nil {
		return *x.BoolValue
	}
	return false
}

func (x *Variant) GetIntValue() int64 {
	if x != nil && x.IntValue != nil {
		return *x.IntValue
	}
	return 0
}

func (x *Variant) GetFloatValue() float64 {
	if x != nil && x.FloatValue != nil {
		return *x.FloatValue
	}
	return 0
}

func (x *Variant) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Variant) GetBlobValue() []byte {
	if x != nil {
		return x.BlobValue
	}
	return nil
}

func (x *Variant) GetMessageValue() []byte {
	if x != nil {
		return x.MessageValue
	}
	return nil
}

func (x *Variant) GetFourccValue() string {
	if x != nil && x.FourccValue != nil {
		return *x.FourccValue
	}
	return ""
}

func (x *Variant) GetUintValue() uint64 {
	if x != nil && x.UintValue != nil {
		return *x.UintValue
	}
	return 0
}

func (x *Variant) GetEntityIdValue() *EntityId {
	if x != nil {
		return x.EntityIdValue
	}
	return nil
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value *Variant `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_attribute_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_attribute_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_attribute_types_proto_rawDescGZIP(), []int{1}
}

func (x *Attribute) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Attribute) GetValue() *Variant {
	if x != nil {
		return x.Value
	}
	return nil
}

type AttributeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op        *AttributeFilter_Operation `protobuf:"varint,1,req,name=op,enum=bgs.protocol.AttributeFilter_Operation" json:"op,omitempty"`
	Attribute []*Attribute               `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
}

func (x *AttributeFilter) Reset() {
	*x = AttributeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_attribute_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeFilter) ProtoMessage() {}

func (x *AttributeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_attribute_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeFilter.ProtoReflect.Descriptor instead.
func (*AttributeFilter) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_attribute_types_proto_rawDescGZIP(), []int{2}
}

func (x *AttributeFilter) GetOp() AttributeFilter_Operation {
	if x != nil && x.Op != nil {
		return *x.Op
	}
	return AttributeFilter_MATCH_NONE
}

func (x *AttributeFilter) GetAttribute() []*Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

var File_bgs_low_pb_client_attribute_types_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_attribute_types_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x24, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77,
	0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02,
	0x0a, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x62, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x62, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x6f, 0x75, 0x72, 0x63, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x75, 0x72, 0x63, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x3e, 0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x52, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x4c, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd9, 0x01,
	0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x37, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x22, 0x56, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x0a, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x4d, 0x4f, 0x53, 0x54, 0x5f, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x10, 0x03, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74,
	0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x78, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x62,
	0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
}

var (
	file_bgs_low_pb_client_attribute_types_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_attribute_types_proto_rawDescData = file_bgs_low_pb_client_attribute_types_proto_rawDesc
)

func file_bgs_low_pb_client_attribute_types_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_attribute_types_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_attribute_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_attribute_types_proto_rawDescData)
	})
	return file_bgs_low_pb_client_attribute_types_proto_rawDescData
}

var file_bgs_low_pb_client_attribute_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bgs_low_pb_client_attribute_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bgs_low_pb_client_attribute_types_proto_goTypes = []interface{}{
	(AttributeFilter_Operation)(0), // 0: bgs.protocol.AttributeFilter.Operation
	(*Variant)(nil),                // 1: bgs.protocol.Variant
	(*Attribute)(nil),              // 2: bgs.protocol.Attribute
	(*AttributeFilter)(nil),        // 3: bgs.protocol.AttributeFilter
	(*EntityId)(nil),               // 4: bgs.protocol.EntityId
}
var file_bgs_low_pb_client_attribute_types_proto_depIdxs = []int32{
	4, // 0: bgs.protocol.Variant.entity_id_value:type_name -> bgs.protocol.EntityId
	1, // 1: bgs.protocol.Attribute.value:type_name -> bgs.protocol.Variant
	0, // 2: bgs.protocol.AttributeFilter.op:type_name -> bgs.protocol.AttributeFilter.Operation
	2, // 3: bgs.protocol.AttributeFilter.attribute:type_name -> bgs.protocol.Attribute
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_attribute_types_proto_init() }
func file_bgs_low_pb_client_attribute_types_proto_init() {
	if File_bgs_low_pb_client_attribute_types_proto != nil {
		return
	}
	file_bgs_low_pb_client_entity_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_attribute_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant); i {
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
		file_bgs_low_pb_client_attribute_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_bgs_low_pb_client_attribute_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeFilter); i {
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
			RawDescriptor: file_bgs_low_pb_client_attribute_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_attribute_types_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_attribute_types_proto_depIdxs,
		EnumInfos:         file_bgs_low_pb_client_attribute_types_proto_enumTypes,
		MessageInfos:      file_bgs_low_pb_client_attribute_types_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_attribute_types_proto = out.File
	file_bgs_low_pb_client_attribute_types_proto_rawDesc = nil
	file_bgs_low_pb_client_attribute_types_proto_goTypes = nil
	file_bgs_low_pb_client_attribute_types_proto_depIdxs = nil
}
