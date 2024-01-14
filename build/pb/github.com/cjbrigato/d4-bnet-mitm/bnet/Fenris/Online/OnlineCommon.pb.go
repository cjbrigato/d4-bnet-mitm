// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: OnlineCommon.proto

package Online

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

type GameHandle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameInstanceId *uint32 `protobuf:"varint,1,opt,name=game_instance_id,json=gameInstanceId" json:"game_instance_id,omitempty"`
	MatchmakerId   *uint64 `protobuf:"varint,2,opt,name=matchmaker_id,json=matchmakerId" json:"matchmaker_id,omitempty"`
	GameServerGuid *uint64 `protobuf:"varint,3,opt,name=game_server_guid,json=gameServerGuid" json:"game_server_guid,omitempty"`
}

func (x *GameHandle) Reset() {
	*x = GameHandle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OnlineCommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameHandle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameHandle) ProtoMessage() {}

func (x *GameHandle) ProtoReflect() protoreflect.Message {
	mi := &file_OnlineCommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameHandle.ProtoReflect.Descriptor instead.
func (*GameHandle) Descriptor() ([]byte, []int) {
	return file_OnlineCommon_proto_rawDescGZIP(), []int{0}
}

func (x *GameHandle) GetGameInstanceId() uint32 {
	if x != nil && x.GameInstanceId != nil {
		return *x.GameInstanceId
	}
	return 0
}

func (x *GameHandle) GetMatchmakerId() uint64 {
	if x != nil && x.MatchmakerId != nil {
		return *x.MatchmakerId
	}
	return 0
}

func (x *GameHandle) GetGameServerGuid() uint64 {
	if x != nil && x.GameServerGuid != nil {
		return *x.GameServerGuid
	}
	return 0
}

type PlatformAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Platform *uint32 `protobuf:"varint,2,opt,name=platform" json:"platform,omitempty"`
}

func (x *PlatformAccount) Reset() {
	*x = PlatformAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OnlineCommon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformAccount) ProtoMessage() {}

func (x *PlatformAccount) ProtoReflect() protoreflect.Message {
	mi := &file_OnlineCommon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformAccount.ProtoReflect.Descriptor instead.
func (*PlatformAccount) Descriptor() ([]byte, []int) {
	return file_OnlineCommon_proto_rawDescGZIP(), []int{1}
}

func (x *PlatformAccount) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *PlatformAccount) GetPlatform() uint32 {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return 0
}

type PlatformProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Price         *uint64 `protobuf:"varint,2,opt,name=price" json:"price,omitempty"`
	OriginalPrice *uint64 `protobuf:"varint,3,opt,name=original_price,json=originalPrice" json:"original_price,omitempty"`
	CurrencyCode  *string `protobuf:"bytes,6,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
}

func (x *PlatformProduct) Reset() {
	*x = PlatformProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OnlineCommon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformProduct) ProtoMessage() {}

func (x *PlatformProduct) ProtoReflect() protoreflect.Message {
	mi := &file_OnlineCommon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformProduct.ProtoReflect.Descriptor instead.
func (*PlatformProduct) Descriptor() ([]byte, []int) {
	return file_OnlineCommon_proto_rawDescGZIP(), []int{2}
}

func (x *PlatformProduct) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *PlatformProduct) GetPrice() uint64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *PlatformProduct) GetOriginalPrice() uint64 {
	if x != nil && x.OriginalPrice != nil {
		return *x.OriginalPrice
	}
	return 0
}

func (x *PlatformProduct) GetCurrencyCode() string {
	if x != nil && x.CurrencyCode != nil {
		return *x.CurrencyCode
	}
	return ""
}

var File_OnlineCommon_proto protoreflect.FileDescriptor

var file_OnlineCommon_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x46, 0x65, 0x6e, 0x72, 0x69, 0x73, 0x2e, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x67, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0f, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x64, 0x34, 0x2d, 0x62, 0x6e, 0x65, 0x74,
	0x2d, 0x6d, 0x69, 0x74, 0x6d, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x46, 0x65, 0x6e, 0x72, 0x69,
	0x73, 0x2f, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
}

var (
	file_OnlineCommon_proto_rawDescOnce sync.Once
	file_OnlineCommon_proto_rawDescData = file_OnlineCommon_proto_rawDesc
)

func file_OnlineCommon_proto_rawDescGZIP() []byte {
	file_OnlineCommon_proto_rawDescOnce.Do(func() {
		file_OnlineCommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_OnlineCommon_proto_rawDescData)
	})
	return file_OnlineCommon_proto_rawDescData
}

var file_OnlineCommon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_OnlineCommon_proto_goTypes = []interface{}{
	(*GameHandle)(nil),      // 0: Fenris.Online.GameHandle
	(*PlatformAccount)(nil), // 1: Fenris.Online.PlatformAccount
	(*PlatformProduct)(nil), // 2: Fenris.Online.PlatformProduct
}
var file_OnlineCommon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OnlineCommon_proto_init() }
func file_OnlineCommon_proto_init() {
	if File_OnlineCommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OnlineCommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameHandle); i {
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
		file_OnlineCommon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformAccount); i {
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
		file_OnlineCommon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformProduct); i {
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
			RawDescriptor: file_OnlineCommon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OnlineCommon_proto_goTypes,
		DependencyIndexes: file_OnlineCommon_proto_depIdxs,
		MessageInfos:      file_OnlineCommon_proto_msgTypes,
	}.Build()
	File_OnlineCommon_proto = out.File
	file_OnlineCommon_proto_rawDesc = nil
	file_OnlineCommon_proto_goTypes = nil
	file_OnlineCommon_proto_depIdxs = nil
}
