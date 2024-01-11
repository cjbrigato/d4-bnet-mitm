// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/api/client/v2/voice_service.proto

package client

import (
	protocol "github.com/cjbrigato/logprox/bnet/bgs/protocol"
	v1 "github.com/cjbrigato/logprox/bnet/bgs/protocol/account/v1"
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

type CreateLoginCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentAccountId *v1.AccountId `protobuf:"bytes,1,opt,name=agent_account_id,json=agentAccountId" json:"agent_account_id,omitempty"`
	Version        *uint32       `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (x *CreateLoginCredentialsRequest) Reset() {
	*x = CreateLoginCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginCredentialsRequest) ProtoMessage() {}

func (x *CreateLoginCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginCredentialsRequest.ProtoReflect.Descriptor instead.
func (*CreateLoginCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLoginCredentialsRequest) GetAgentAccountId() *v1.AccountId {
	if x != nil {
		return x.AgentAccountId
	}
	return nil
}

func (x *CreateLoginCredentialsRequest) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

type CreateLoginCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credentials *protocol.VoiceCredentials `protobuf:"bytes,1,opt,name=credentials" json:"credentials,omitempty"`
}

func (x *CreateLoginCredentialsResponse) Reset() {
	*x = CreateLoginCredentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginCredentialsResponse) ProtoMessage() {}

func (x *CreateLoginCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginCredentialsResponse.ProtoReflect.Descriptor instead.
func (*CreateLoginCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLoginCredentialsResponse) GetCredentials() *protocol.VoiceCredentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type CreateChannelSttTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentAccountId *v1.AccountId `protobuf:"bytes,1,opt,name=agent_account_id,json=agentAccountId" json:"agent_account_id,omitempty"`
	ChannelUri     *string       `protobuf:"bytes,2,opt,name=channel_uri,json=channelUri" json:"channel_uri,omitempty"`
	Version        *uint32       `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
}

func (x *CreateChannelSttTokenRequest) Reset() {
	*x = CreateChannelSttTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelSttTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelSttTokenRequest) ProtoMessage() {}

func (x *CreateChannelSttTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelSttTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateChannelSttTokenRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateChannelSttTokenRequest) GetAgentAccountId() *v1.AccountId {
	if x != nil {
		return x.AgentAccountId
	}
	return nil
}

func (x *CreateChannelSttTokenRequest) GetChannelUri() string {
	if x != nil && x.ChannelUri != nil {
		return *x.ChannelUri
	}
	return ""
}

func (x *CreateChannelSttTokenRequest) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

type CreateChannelSttTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (x *CreateChannelSttTokenResponse) Reset() {
	*x = CreateChannelSttTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelSttTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelSttTokenResponse) ProtoMessage() {}

func (x *CreateChannelSttTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelSttTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateChannelSttTokenResponse) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateChannelSttTokenResponse) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

var File_bgs_low_pb_client_api_client_v2_voice_service_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x1a, 0x23, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f,
	0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x52, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x1e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x22, 0xa7, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x53, 0x74, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4c, 0x0a, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x52,
	0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x55, 0x72, 0x69,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x32, 0xb7, 0x02, 0x0a, 0x0c, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x3b, 0x2e,
	0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x90, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x3a, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x74, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67,
	0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x78, 0x2f, 0x62, 0x6e, 0x65, 0x74,
	0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
}

var (
	file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescData = file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDesc
)

func file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescData)
	})
	return file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDescData
}

var file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bgs_low_pb_client_api_client_v2_voice_service_proto_goTypes = []interface{}{
	(*CreateLoginCredentialsRequest)(nil),  // 0: bgs.protocol.voice.v2.client.CreateLoginCredentialsRequest
	(*CreateLoginCredentialsResponse)(nil), // 1: bgs.protocol.voice.v2.client.CreateLoginCredentialsResponse
	(*CreateChannelSttTokenRequest)(nil),   // 2: bgs.protocol.voice.v2.client.CreateChannelSttTokenRequest
	(*CreateChannelSttTokenResponse)(nil),  // 3: bgs.protocol.voice.v2.client.CreateChannelSttTokenResponse
	(*v1.AccountId)(nil),                   // 4: bgs.protocol.account.v1.AccountId
	(*protocol.VoiceCredentials)(nil),      // 5: bgs.protocol.VoiceCredentials
}
var file_bgs_low_pb_client_api_client_v2_voice_service_proto_depIdxs = []int32{
	4, // 0: bgs.protocol.voice.v2.client.CreateLoginCredentialsRequest.agent_account_id:type_name -> bgs.protocol.account.v1.AccountId
	5, // 1: bgs.protocol.voice.v2.client.CreateLoginCredentialsResponse.credentials:type_name -> bgs.protocol.VoiceCredentials
	4, // 2: bgs.protocol.voice.v2.client.CreateChannelSttTokenRequest.agent_account_id:type_name -> bgs.protocol.account.v1.AccountId
	0, // 3: bgs.protocol.voice.v2.client.VoiceService.CreateLoginCredentials:input_type -> bgs.protocol.voice.v2.client.CreateLoginCredentialsRequest
	2, // 4: bgs.protocol.voice.v2.client.VoiceService.CreateChannelSttToken:input_type -> bgs.protocol.voice.v2.client.CreateChannelSttTokenRequest
	1, // 5: bgs.protocol.voice.v2.client.VoiceService.CreateLoginCredentials:output_type -> bgs.protocol.voice.v2.client.CreateLoginCredentialsResponse
	3, // 6: bgs.protocol.voice.v2.client.VoiceService.CreateChannelSttToken:output_type -> bgs.protocol.voice.v2.client.CreateChannelSttTokenResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_api_client_v2_voice_service_proto_init() }
func file_bgs_low_pb_client_api_client_v2_voice_service_proto_init() {
	if File_bgs_low_pb_client_api_client_v2_voice_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginCredentialsRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginCredentialsResponse); i {
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
		file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelSttTokenRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelSttTokenResponse); i {
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
			RawDescriptor: file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bgs_low_pb_client_api_client_v2_voice_service_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_api_client_v2_voice_service_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_api_client_v2_voice_service_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_api_client_v2_voice_service_proto = out.File
	file_bgs_low_pb_client_api_client_v2_voice_service_proto_rawDesc = nil
	file_bgs_low_pb_client_api_client_v2_voice_service_proto_goTypes = nil
	file_bgs_low_pb_client_api_client_v2_voice_service_proto_depIdxs = nil
}
