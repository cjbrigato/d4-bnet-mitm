// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/api/client/v1/block_list_types.proto

package client

import (
	_ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
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

type BlockedPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BattleTag      *string `protobuf:"bytes,2,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
	CreationTimeUs *uint64 `protobuf:"varint,4,opt,name=creation_time_us,json=creationTimeUs" json:"creation_time_us,omitempty"`
	ModifiedTimeUs *uint64 `protobuf:"varint,5,opt,name=modified_time_us,json=modifiedTimeUs" json:"modified_time_us,omitempty"`
}

func (x *BlockedPlayer) Reset() {
	*x = BlockedPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockedPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockedPlayer) ProtoMessage() {}

func (x *BlockedPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockedPlayer.ProtoReflect.Descriptor instead.
func (*BlockedPlayer) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescGZIP(), []int{0}
}

func (x *BlockedPlayer) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BlockedPlayer) GetBattleTag() string {
	if x != nil && x.BattleTag != nil {
		return *x.BattleTag
	}
	return ""
}

func (x *BlockedPlayer) GetCreationTimeUs() uint64 {
	if x != nil && x.CreationTimeUs != nil {
		return *x.CreationTimeUs
	}
	return 0
}

func (x *BlockedPlayer) GetModifiedTimeUs() uint64 {
	if x != nil && x.ModifiedTimeUs != nil {
		return *x.ModifiedTimeUs
	}
	return 0
}

type BlockPlayerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId *uint64 `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
}

func (x *BlockPlayerOptions) Reset() {
	*x = BlockPlayerOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockPlayerOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockPlayerOptions) ProtoMessage() {}

func (x *BlockPlayerOptions) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockPlayerOptions.ProtoReflect.Descriptor instead.
func (*BlockPlayerOptions) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescGZIP(), []int{1}
}

func (x *BlockPlayerOptions) GetAccountId() uint64 {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return 0
}

type UnblockPlayerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId *uint64 `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
}

func (x *UnblockPlayerOptions) Reset() {
	*x = UnblockPlayerOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnblockPlayerOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockPlayerOptions) ProtoMessage() {}

func (x *UnblockPlayerOptions) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockPlayerOptions.ProtoReflect.Descriptor instead.
func (*UnblockPlayerOptions) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescGZIP(), []int{2}
}

func (x *UnblockPlayerOptions) GetAccountId() uint64 {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return 0
}

type UnblockPlayerAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (x *UnblockPlayerAssignment) Reset() {
	*x = UnblockPlayerAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnblockPlayerAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockPlayerAssignment) ProtoMessage() {}

func (x *UnblockPlayerAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockPlayerAssignment.ProtoReflect.Descriptor instead.
func (*UnblockPlayerAssignment) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescGZIP(), []int{3}
}

func (x *UnblockPlayerAssignment) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

type BlockListState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player []*BlockedPlayer `protobuf:"bytes,1,rep,name=player" json:"player,omitempty"`
}

func (x *BlockListState) Reset() {
	*x = BlockListState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockListState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockListState) ProtoMessage() {}

func (x *BlockListState) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockListState.ProtoReflect.Descriptor instead.
func (*BlockListState) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescGZIP(), []int{4}
}

func (x *BlockListState) GetPlayer() []*BlockedPlayer {
	if x != nil {
		return x.Player
	}
	return nil
}

var File_bgs_low_pb_client_api_client_v1_block_list_types_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDesc = []byte{
	0x0a, 0x36, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x21, 0x62, 0x67, 0x73,
	0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92,
	0x01, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x12,
	0x28, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x55, 0x73, 0x22, 0x33, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x14, 0x55, 0x6e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x29, 0x0a, 0x17, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x64,
	0x34, 0x2d, 0x62, 0x6e, 0x65, 0x74, 0x2d, 0x6d, 0x69, 0x74, 0x6d, 0x2f, 0x62, 0x6e, 0x65, 0x74,
	0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74,
}

var (
	file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescData = file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDesc
)

func file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescData)
	})
	return file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDescData
}

var file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bgs_low_pb_client_api_client_v1_block_list_types_proto_goTypes = []interface{}{
	(*BlockedPlayer)(nil),           // 0: bgs.protocol.block_list.v1.client.BlockedPlayer
	(*BlockPlayerOptions)(nil),      // 1: bgs.protocol.block_list.v1.client.BlockPlayerOptions
	(*UnblockPlayerOptions)(nil),    // 2: bgs.protocol.block_list.v1.client.UnblockPlayerOptions
	(*UnblockPlayerAssignment)(nil), // 3: bgs.protocol.block_list.v1.client.UnblockPlayerAssignment
	(*BlockListState)(nil),          // 4: bgs.protocol.block_list.v1.client.BlockListState
}
var file_bgs_low_pb_client_api_client_v1_block_list_types_proto_depIdxs = []int32{
	0, // 0: bgs.protocol.block_list.v1.client.BlockListState.player:type_name -> bgs.protocol.block_list.v1.client.BlockedPlayer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_api_client_v1_block_list_types_proto_init() }
func file_bgs_low_pb_client_api_client_v1_block_list_types_proto_init() {
	if File_bgs_low_pb_client_api_client_v1_block_list_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockedPlayer); i {
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
		file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockPlayerOptions); i {
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
		file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnblockPlayerOptions); i {
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
		file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnblockPlayerAssignment); i {
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
		file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockListState); i {
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
			RawDescriptor: file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_api_client_v1_block_list_types_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_api_client_v1_block_list_types_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_api_client_v1_block_list_types_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_api_client_v1_block_list_types_proto = out.File
	file_bgs_low_pb_client_api_client_v1_block_list_types_proto_rawDesc = nil
	file_bgs_low_pb_client_api_client_v1_block_list_types_proto_goTypes = nil
	file_bgs_low_pb_client_api_client_v1_block_list_types_proto_depIdxs = nil
}
