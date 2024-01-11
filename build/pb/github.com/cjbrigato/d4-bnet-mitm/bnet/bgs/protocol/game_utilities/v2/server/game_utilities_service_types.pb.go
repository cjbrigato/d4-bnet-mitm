// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/api/server/v2/game_utilities_service_types.proto

package server

import (
	_ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
	v2 "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/account/v2"
	v21 "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/v2"
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

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientAddress     *string      `protobuf:"bytes,1,opt,name=client_address,json=clientAddress" json:"client_address,omitempty"`
	PrivilegedNetwork *bool        `protobuf:"varint,2,opt,name=privileged_network,json=privilegedNetwork" json:"privileged_network,omitempty"`
	User              *v2.Identity `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescGZIP(), []int{0}
}

func (x *ClientInfo) GetClientAddress() string {
	if x != nil && x.ClientAddress != nil {
		return *x.ClientAddress
	}
	return ""
}

func (x *ClientInfo) GetPrivilegedNetwork() bool {
	if x != nil && x.PrivilegedNetwork != nil {
		return *x.PrivilegedNetwork
	}
	return false
}

func (x *ClientInfo) GetUser() *v2.Identity {
	if x != nil {
		return x.User
	}
	return nil
}

type Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrationId *string          `protobuf:"bytes,1,opt,name=registration_id,json=registrationId" json:"registration_id,omitempty"`
	Attribute      []*v21.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	CreationTimeS  *uint64          `protobuf:"varint,3,opt,name=creation_time_s,json=creationTimeS" json:"creation_time_s,omitempty"`
	ModifiedTimeS  *uint64          `protobuf:"varint,4,opt,name=modified_time_s,json=modifiedTimeS" json:"modified_time_s,omitempty"`
}

func (x *Registration) Reset() {
	*x = Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registration) ProtoMessage() {}

func (x *Registration) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registration.ProtoReflect.Descriptor instead.
func (*Registration) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescGZIP(), []int{1}
}

func (x *Registration) GetRegistrationId() string {
	if x != nil && x.RegistrationId != nil {
		return *x.RegistrationId
	}
	return ""
}

func (x *Registration) GetAttribute() []*v21.Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

func (x *Registration) GetCreationTimeS() uint64 {
	if x != nil && x.CreationTimeS != nil {
		return *x.CreationTimeS
	}
	return 0
}

func (x *Registration) GetModifiedTimeS() uint64 {
	if x != nil && x.ModifiedTimeS != nil {
		return *x.ModifiedTimeS
	}
	return 0
}

type GetRegistrationsContinuation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    *uint64 `protobuf:"varint,1,opt,name=token" json:"token,omitempty"`
	Position *string `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
}

func (x *GetRegistrationsContinuation) Reset() {
	*x = GetRegistrationsContinuation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistrationsContinuation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistrationsContinuation) ProtoMessage() {}

func (x *GetRegistrationsContinuation) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistrationsContinuation.ProtoReflect.Descriptor instead.
func (*GetRegistrationsContinuation) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescGZIP(), []int{2}
}

func (x *GetRegistrationsContinuation) GetToken() uint64 {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return 0
}

func (x *GetRegistrationsContinuation) GetPosition() string {
	if x != nil && x.Position != nil {
		return *x.Position
	}
	return ""
}

type GetRegistrationsOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter     *v21.AttributeFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	MaxServers *uint32              `protobuf:"varint,2,opt,name=max_servers,json=maxServers" json:"max_servers,omitempty"`
}

func (x *GetRegistrationsOptions) Reset() {
	*x = GetRegistrationsOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistrationsOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistrationsOptions) ProtoMessage() {}

func (x *GetRegistrationsOptions) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistrationsOptions.ProtoReflect.Descriptor instead.
func (*GetRegistrationsOptions) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescGZIP(), []int{3}
}

func (x *GetRegistrationsOptions) GetFilter() *v21.AttributeFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetRegistrationsOptions) GetMaxServers() uint32 {
	if x != nil && x.MaxServers != nil {
		return *x.MaxServers
	}
	return 0
}

var File_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDesc = []byte{
	0x0a, 0x42, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x21, 0x62, 0x67, 0x73,
	0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33,
	0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x0a, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x5f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x35, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xc1, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x22, 0x50, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x74, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x64, 0x34, 0x2d, 0x62, 0x6e,
	0x65, 0x74, 0x2d, 0x6d, 0x69, 0x74, 0x6d, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x62, 0x67, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x75,
	0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72,
}

var (
	file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescData = file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDesc
)

func file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescData)
	})
	return file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDescData
}

var file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_goTypes = []interface{}{
	(*ClientInfo)(nil),                   // 0: bgs.protocol.game_utilities.v2.server.ClientInfo
	(*Registration)(nil),                 // 1: bgs.protocol.game_utilities.v2.server.Registration
	(*GetRegistrationsContinuation)(nil), // 2: bgs.protocol.game_utilities.v2.server.GetRegistrationsContinuation
	(*GetRegistrationsOptions)(nil),      // 3: bgs.protocol.game_utilities.v2.server.GetRegistrationsOptions
	(*v2.Identity)(nil),                  // 4: bgs.protocol.account.v2.Identity
	(*v21.Attribute)(nil),                // 5: bgs.protocol.v2.Attribute
	(*v21.AttributeFilter)(nil),          // 6: bgs.protocol.v2.AttributeFilter
}
var file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_depIdxs = []int32{
	4, // 0: bgs.protocol.game_utilities.v2.server.ClientInfo.user:type_name -> bgs.protocol.account.v2.Identity
	5, // 1: bgs.protocol.game_utilities.v2.server.Registration.attribute:type_name -> bgs.protocol.v2.Attribute
	6, // 2: bgs.protocol.game_utilities.v2.server.GetRegistrationsOptions.filter:type_name -> bgs.protocol.v2.AttributeFilter
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_init() }
func file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_init() {
	if File_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
		file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registration); i {
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
		file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistrationsContinuation); i {
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
		file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistrationsOptions); i {
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
			RawDescriptor: file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto = out.File
	file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_rawDesc = nil
	file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_goTypes = nil
	file_bgs_low_pb_client_api_server_v2_game_utilities_service_types_proto_depIdxs = nil
}
