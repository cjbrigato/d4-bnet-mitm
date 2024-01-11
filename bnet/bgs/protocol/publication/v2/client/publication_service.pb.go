// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/api/client/v2/publication_service.proto

package client

import (
	protocol "github.com/cjbrigato/logprox/bnet/bgs/protocol"
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

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target  []*Target         `protobuf:"bytes,1,rep,name=target" json:"target,omitempty"`
	Options *SubscribeOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetTarget() []*Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *SubscribeRequest) GetOptions() *SubscribeOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failures []*SubscribeResult `protobuf:"bytes,1,rep,name=failures" json:"failures,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeResponse) GetFailures() []*SubscribeResult {
	if x != nil {
		return x.Failures
	}
	return nil
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target []*Target `protobuf:"bytes,1,rep,name=target" json:"target,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescGZIP(), []int{2}
}

func (x *UnsubscribeRequest) GetTarget() []*Target {
	if x != nil {
		return x.Target
	}
	return nil
}

var File_bgs_low_pb_client_api_client_v2_publication_service_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x21, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x37, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x4e, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x64, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x62, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x12, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x32, 0xeb, 0x01, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x34, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x36, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x70, 0x72,
	0x6f, 0x78, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
}

var (
	file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescData = file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDesc
)

func file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescData)
	})
	return file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDescData
}

var file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bgs_low_pb_client_api_client_v2_publication_service_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),   // 0: bgs.protocol.publication.v2.client.SubscribeRequest
	(*SubscribeResponse)(nil),  // 1: bgs.protocol.publication.v2.client.SubscribeResponse
	(*UnsubscribeRequest)(nil), // 2: bgs.protocol.publication.v2.client.UnsubscribeRequest
	(*Target)(nil),             // 3: bgs.protocol.publication.v2.client.Target
	(*SubscribeOptions)(nil),   // 4: bgs.protocol.publication.v2.client.SubscribeOptions
	(*SubscribeResult)(nil),    // 5: bgs.protocol.publication.v2.client.SubscribeResult
	(*protocol.NoData)(nil),    // 6: bgs.protocol.NoData
}
var file_bgs_low_pb_client_api_client_v2_publication_service_proto_depIdxs = []int32{
	3, // 0: bgs.protocol.publication.v2.client.SubscribeRequest.target:type_name -> bgs.protocol.publication.v2.client.Target
	4, // 1: bgs.protocol.publication.v2.client.SubscribeRequest.options:type_name -> bgs.protocol.publication.v2.client.SubscribeOptions
	5, // 2: bgs.protocol.publication.v2.client.SubscribeResponse.failures:type_name -> bgs.protocol.publication.v2.client.SubscribeResult
	3, // 3: bgs.protocol.publication.v2.client.UnsubscribeRequest.target:type_name -> bgs.protocol.publication.v2.client.Target
	0, // 4: bgs.protocol.publication.v2.client.PublicationService.Subscribe:input_type -> bgs.protocol.publication.v2.client.SubscribeRequest
	2, // 5: bgs.protocol.publication.v2.client.PublicationService.Unsubscribe:input_type -> bgs.protocol.publication.v2.client.UnsubscribeRequest
	1, // 6: bgs.protocol.publication.v2.client.PublicationService.Subscribe:output_type -> bgs.protocol.publication.v2.client.SubscribeResponse
	6, // 7: bgs.protocol.publication.v2.client.PublicationService.Unsubscribe:output_type -> bgs.protocol.NoData
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_api_client_v2_publication_service_proto_init() }
func file_bgs_low_pb_client_api_client_v2_publication_service_proto_init() {
	if File_bgs_low_pb_client_api_client_v2_publication_service_proto != nil {
		return
	}
	file_bgs_low_pb_client_api_client_v2_publication_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
			RawDescriptor: file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bgs_low_pb_client_api_client_v2_publication_service_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_api_client_v2_publication_service_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_api_client_v2_publication_service_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_api_client_v2_publication_service_proto = out.File
	file_bgs_low_pb_client_api_client_v2_publication_service_proto_rawDesc = nil
	file_bgs_low_pb_client_api_client_v2_publication_service_proto_goTypes = nil
	file_bgs_low_pb_client_api_client_v2_publication_service_proto_depIdxs = nil
}
