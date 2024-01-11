// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/api/client/v2/session_service.proto

package client

import (
	protocol "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
	v2 "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/account/v2"
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

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameAccount *v2.GameAccountHandle `protobuf:"bytes,1,opt,name=game_account,json=gameAccount" json:"game_account,omitempty"`
	Options     *CreateOptions        `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionRequest) GetGameAccount() *v2.GameAccountHandle {
	if x != nil {
		return x.GameAccount
	}
	return nil
}

func (x *CreateSessionRequest) GetOptions() *CreateOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type CreateSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId *SessionId        `protobuf:"bytes,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Variables *SessionVariables `protobuf:"bytes,2,opt,name=variables" json:"variables,omitempty"`
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSessionResponse) GetSessionId() *SessionId {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *CreateSessionResponse) GetVariables() *SessionVariables {
	if x != nil {
		return x.Variables
	}
	return nil
}

type RestoreSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity     *v2.Identity    `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	SessionKey   []byte          `protobuf:"bytes,2,opt,name=session_key,json=sessionKey" json:"session_key,omitempty"`
	Options      *RestoreOptions `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	StateOptions *StateOptions   `protobuf:"bytes,4,opt,name=state_options,json=stateOptions" json:"state_options,omitempty"`
}

func (x *RestoreSessionRequest) Reset() {
	*x = RestoreSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreSessionRequest) ProtoMessage() {}

func (x *RestoreSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreSessionRequest.ProtoReflect.Descriptor instead.
func (*RestoreSessionRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{2}
}

func (x *RestoreSessionRequest) GetIdentity() *v2.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *RestoreSessionRequest) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

func (x *RestoreSessionRequest) GetOptions() *RestoreOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *RestoreSessionRequest) GetStateOptions() *StateOptions {
	if x != nil {
		return x.StateOptions
	}
	return nil
}

type RestoreSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId *SessionId        `protobuf:"bytes,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Variables *SessionVariables `protobuf:"bytes,2,opt,name=variables" json:"variables,omitempty"`
}

func (x *RestoreSessionResponse) Reset() {
	*x = RestoreSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreSessionResponse) ProtoMessage() {}

func (x *RestoreSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreSessionResponse.ProtoReflect.Descriptor instead.
func (*RestoreSessionResponse) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{3}
}

func (x *RestoreSessionResponse) GetSessionId() *SessionId {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *RestoreSessionResponse) GetVariables() *SessionVariables {
	if x != nil {
		return x.Variables
	}
	return nil
}

type DestroySessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DestroySessionRequest) Reset() {
	*x = DestroySessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroySessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroySessionRequest) ProtoMessage() {}

func (x *DestroySessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroySessionRequest.ProtoReflect.Descriptor instead.
func (*DestroySessionRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{4}
}

type GetSignedSessionStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSignedSessionStateRequest) Reset() {
	*x = GetSignedSessionStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignedSessionStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignedSessionStateRequest) ProtoMessage() {}

func (x *GetSignedSessionStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignedSessionStateRequest.ProtoReflect.Descriptor instead.
func (*GetSignedSessionStateRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{5}
}

type GetSignedSessionStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedToken *string `protobuf:"bytes,1,opt,name=signed_token,json=signedToken" json:"signed_token,omitempty"`
}

func (x *GetSignedSessionStateResponse) Reset() {
	*x = GetSignedSessionStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignedSessionStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignedSessionStateResponse) ProtoMessage() {}

func (x *GetSignedSessionStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignedSessionStateResponse.ProtoReflect.Descriptor instead.
func (*GetSignedSessionStateResponse) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetSignedSessionStateResponse) GetSignedToken() string {
	if x != nil && x.SignedToken != nil {
		return *x.SignedToken
	}
	return ""
}

type GetSessionStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSessionStateRequest) Reset() {
	*x = GetSessionStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSessionStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionStateRequest) ProtoMessage() {}

func (x *GetSessionStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionStateRequest.ProtoReflect.Descriptor instead.
func (*GetSessionStateRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{7}
}

type GetSessionStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *SessionState `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (x *GetSessionStateResponse) Reset() {
	*x = GetSessionStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSessionStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionStateResponse) ProtoMessage() {}

func (x *GetSessionStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionStateResponse.ProtoReflect.Descriptor instead.
func (*GetSessionStateResponse) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetSessionStateResponse) GetSession() *SessionState {
	if x != nil {
		return x.Session
	}
	return nil
}

type MarkSessionAliveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkSessionAliveRequest) Reset() {
	*x = MarkSessionAliveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkSessionAliveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkSessionAliveRequest) ProtoMessage() {}

func (x *MarkSessionAliveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkSessionAliveRequest.ProtoReflect.Descriptor instead.
func (*MarkSessionAliveRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP(), []int{9}
}

var File_bgs_low_pb_client_api_client_v2_session_service_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x33, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77,
	0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x62, 0x67,
	0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77,
	0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x62, 0x67,
	0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xae, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x0b, 0x67, 0x61, 0x6d, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xb1, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x22, 0x94, 0x02, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d,
	0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x48,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x16,
	0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x4e, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x18, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x17, 0x4d, 0x61,
	0x72, 0x6b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xed, 0x05, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x4e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x94, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x3c, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x36, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x37, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e,
	0x6f, 0x44, 0x61, 0x74, 0x61, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x64, 0x34,
	0x2d, 0x62, 0x6e, 0x65, 0x74, 0x2d, 0x6d, 0x69, 0x74, 0x6d, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f,
	0x62, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
}

var (
	file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescData = file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDesc
)

func file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescData)
	})
	return file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDescData
}

var file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bgs_low_pb_client_api_client_v2_session_service_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil),          // 0: bgs.protocol.session.v2.client.CreateSessionRequest
	(*CreateSessionResponse)(nil),         // 1: bgs.protocol.session.v2.client.CreateSessionResponse
	(*RestoreSessionRequest)(nil),         // 2: bgs.protocol.session.v2.client.RestoreSessionRequest
	(*RestoreSessionResponse)(nil),        // 3: bgs.protocol.session.v2.client.RestoreSessionResponse
	(*DestroySessionRequest)(nil),         // 4: bgs.protocol.session.v2.client.DestroySessionRequest
	(*GetSignedSessionStateRequest)(nil),  // 5: bgs.protocol.session.v2.client.GetSignedSessionStateRequest
	(*GetSignedSessionStateResponse)(nil), // 6: bgs.protocol.session.v2.client.GetSignedSessionStateResponse
	(*GetSessionStateRequest)(nil),        // 7: bgs.protocol.session.v2.client.GetSessionStateRequest
	(*GetSessionStateResponse)(nil),       // 8: bgs.protocol.session.v2.client.GetSessionStateResponse
	(*MarkSessionAliveRequest)(nil),       // 9: bgs.protocol.session.v2.client.MarkSessionAliveRequest
	(*v2.GameAccountHandle)(nil),          // 10: bgs.protocol.account.v2.GameAccountHandle
	(*CreateOptions)(nil),                 // 11: bgs.protocol.session.v2.client.CreateOptions
	(*SessionId)(nil),                     // 12: bgs.protocol.session.v2.client.SessionId
	(*SessionVariables)(nil),              // 13: bgs.protocol.session.v2.client.SessionVariables
	(*v2.Identity)(nil),                   // 14: bgs.protocol.account.v2.Identity
	(*RestoreOptions)(nil),                // 15: bgs.protocol.session.v2.client.RestoreOptions
	(*StateOptions)(nil),                  // 16: bgs.protocol.session.v2.client.StateOptions
	(*SessionState)(nil),                  // 17: bgs.protocol.session.v2.client.SessionState
	(*protocol.NoData)(nil),               // 18: bgs.protocol.NoData
}
var file_bgs_low_pb_client_api_client_v2_session_service_proto_depIdxs = []int32{
	10, // 0: bgs.protocol.session.v2.client.CreateSessionRequest.game_account:type_name -> bgs.protocol.account.v2.GameAccountHandle
	11, // 1: bgs.protocol.session.v2.client.CreateSessionRequest.options:type_name -> bgs.protocol.session.v2.client.CreateOptions
	12, // 2: bgs.protocol.session.v2.client.CreateSessionResponse.session_id:type_name -> bgs.protocol.session.v2.client.SessionId
	13, // 3: bgs.protocol.session.v2.client.CreateSessionResponse.variables:type_name -> bgs.protocol.session.v2.client.SessionVariables
	14, // 4: bgs.protocol.session.v2.client.RestoreSessionRequest.identity:type_name -> bgs.protocol.account.v2.Identity
	15, // 5: bgs.protocol.session.v2.client.RestoreSessionRequest.options:type_name -> bgs.protocol.session.v2.client.RestoreOptions
	16, // 6: bgs.protocol.session.v2.client.RestoreSessionRequest.state_options:type_name -> bgs.protocol.session.v2.client.StateOptions
	12, // 7: bgs.protocol.session.v2.client.RestoreSessionResponse.session_id:type_name -> bgs.protocol.session.v2.client.SessionId
	13, // 8: bgs.protocol.session.v2.client.RestoreSessionResponse.variables:type_name -> bgs.protocol.session.v2.client.SessionVariables
	17, // 9: bgs.protocol.session.v2.client.GetSessionStateResponse.session:type_name -> bgs.protocol.session.v2.client.SessionState
	0,  // 10: bgs.protocol.session.v2.client.SessionService.CreateSession:input_type -> bgs.protocol.session.v2.client.CreateSessionRequest
	2,  // 11: bgs.protocol.session.v2.client.SessionService.RestoreSession:input_type -> bgs.protocol.session.v2.client.RestoreSessionRequest
	4,  // 12: bgs.protocol.session.v2.client.SessionService.DestroySession:input_type -> bgs.protocol.session.v2.client.DestroySessionRequest
	5,  // 13: bgs.protocol.session.v2.client.SessionService.GetSignedSessionState:input_type -> bgs.protocol.session.v2.client.GetSignedSessionStateRequest
	7,  // 14: bgs.protocol.session.v2.client.SessionService.GetSessionState:input_type -> bgs.protocol.session.v2.client.GetSessionStateRequest
	9,  // 15: bgs.protocol.session.v2.client.SessionService.MarkSessionAlive:input_type -> bgs.protocol.session.v2.client.MarkSessionAliveRequest
	1,  // 16: bgs.protocol.session.v2.client.SessionService.CreateSession:output_type -> bgs.protocol.session.v2.client.CreateSessionResponse
	3,  // 17: bgs.protocol.session.v2.client.SessionService.RestoreSession:output_type -> bgs.protocol.session.v2.client.RestoreSessionResponse
	18, // 18: bgs.protocol.session.v2.client.SessionService.DestroySession:output_type -> bgs.protocol.NoData
	6,  // 19: bgs.protocol.session.v2.client.SessionService.GetSignedSessionState:output_type -> bgs.protocol.session.v2.client.GetSignedSessionStateResponse
	8,  // 20: bgs.protocol.session.v2.client.SessionService.GetSessionState:output_type -> bgs.protocol.session.v2.client.GetSessionStateResponse
	18, // 21: bgs.protocol.session.v2.client.SessionService.MarkSessionAlive:output_type -> bgs.protocol.NoData
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_api_client_v2_session_service_proto_init() }
func file_bgs_low_pb_client_api_client_v2_session_service_proto_init() {
	if File_bgs_low_pb_client_api_client_v2_session_service_proto != nil {
		return
	}
	file_bgs_low_pb_client_api_client_v2_session_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponse); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreSessionRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreSessionResponse); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroySessionRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignedSessionStateRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignedSessionStateResponse); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSessionStateRequest); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSessionStateResponse); i {
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
		file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkSessionAliveRequest); i {
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
			RawDescriptor: file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bgs_low_pb_client_api_client_v2_session_service_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_api_client_v2_session_service_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_api_client_v2_session_service_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_api_client_v2_session_service_proto = out.File
	file_bgs_low_pb_client_api_client_v2_session_service_proto_rawDesc = nil
	file_bgs_low_pb_client_api_client_v2_session_service_proto_goTypes = nil
	file_bgs_low_pb_client_api_client_v2_session_service_proto_depIdxs = nil
}
