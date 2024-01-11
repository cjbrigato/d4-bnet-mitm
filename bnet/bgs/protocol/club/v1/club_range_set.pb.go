// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/club_range_set.proto

package v1

import (
	protocol "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
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

type ClubTypeRangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameRange        *protocol.UnsignedIntRange `protobuf:"bytes,2,opt,name=name_range,json=nameRange" json:"name_range,omitempty"`
	DescriptionRange *protocol.UnsignedIntRange `protobuf:"bytes,3,opt,name=description_range,json=descriptionRange" json:"description_range,omitempty"`
	BroadcastRange   *protocol.UnsignedIntRange `protobuf:"bytes,4,opt,name=broadcast_range,json=broadcastRange" json:"broadcast_range,omitempty"`
	ShortNameRange   *protocol.UnsignedIntRange `protobuf:"bytes,7,opt,name=short_name_range,json=shortNameRange" json:"short_name_range,omitempty"`
	Member           *ClubMemberRangeSet        `protobuf:"bytes,25,opt,name=member" json:"member,omitempty"`
	Stream           *ClubStreamRangeSet        `protobuf:"bytes,26,opt,name=stream" json:"stream,omitempty"`
	Invitation       *ClubInvitationRangeSet    `protobuf:"bytes,27,opt,name=invitation" json:"invitation,omitempty"`
	Suggestion       *ClubSuggestionRangeSet    `protobuf:"bytes,28,opt,name=suggestion" json:"suggestion,omitempty"`
	Ticket           *ClubTicketRangeSet        `protobuf:"bytes,29,opt,name=ticket" json:"ticket,omitempty"`
	Ban              *ClubBanRangeSet           `protobuf:"bytes,30,opt,name=ban" json:"ban,omitempty"`
}

func (x *ClubTypeRangeSet) Reset() {
	*x = ClubTypeRangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubTypeRangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubTypeRangeSet) ProtoMessage() {}

func (x *ClubTypeRangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubTypeRangeSet.ProtoReflect.Descriptor instead.
func (*ClubTypeRangeSet) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP(), []int{0}
}

func (x *ClubTypeRangeSet) GetNameRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.NameRange
	}
	return nil
}

func (x *ClubTypeRangeSet) GetDescriptionRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.DescriptionRange
	}
	return nil
}

func (x *ClubTypeRangeSet) GetBroadcastRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.BroadcastRange
	}
	return nil
}

func (x *ClubTypeRangeSet) GetShortNameRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.ShortNameRange
	}
	return nil
}

func (x *ClubTypeRangeSet) GetMember() *ClubMemberRangeSet {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *ClubTypeRangeSet) GetStream() *ClubStreamRangeSet {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (x *ClubTypeRangeSet) GetInvitation() *ClubInvitationRangeSet {
	if x != nil {
		return x.Invitation
	}
	return nil
}

func (x *ClubTypeRangeSet) GetSuggestion() *ClubSuggestionRangeSet {
	if x != nil {
		return x.Suggestion
	}
	return nil
}

func (x *ClubTypeRangeSet) GetTicket() *ClubTicketRangeSet {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *ClubTypeRangeSet) GetBan() *ClubBanRangeSet {
	if x != nil {
		return x.Ban
	}
	return nil
}

type ClubMemberRangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count               *protocol.UnsignedIntRange `protobuf:"bytes,1,opt,name=count" json:"count,omitempty"`
	Voice               *protocol.UnsignedIntRange `protobuf:"bytes,3,opt,name=voice" json:"voice,omitempty"`
	StreamSubscriptions *protocol.UnsignedIntRange `protobuf:"bytes,5,opt,name=stream_subscriptions,json=streamSubscriptions" json:"stream_subscriptions,omitempty"`
	NoteRange           *protocol.UnsignedIntRange `protobuf:"bytes,7,opt,name=note_range,json=noteRange" json:"note_range,omitempty"`
}

func (x *ClubMemberRangeSet) Reset() {
	*x = ClubMemberRangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubMemberRangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubMemberRangeSet) ProtoMessage() {}

func (x *ClubMemberRangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubMemberRangeSet.ProtoReflect.Descriptor instead.
func (*ClubMemberRangeSet) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP(), []int{1}
}

func (x *ClubMemberRangeSet) GetCount() *protocol.UnsignedIntRange {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *ClubMemberRangeSet) GetVoice() *protocol.UnsignedIntRange {
	if x != nil {
		return x.Voice
	}
	return nil
}

func (x *ClubMemberRangeSet) GetStreamSubscriptions() *protocol.UnsignedIntRange {
	if x != nil {
		return x.StreamSubscriptions
	}
	return nil
}

func (x *ClubMemberRangeSet) GetNoteRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.NoteRange
	}
	return nil
}

type ClubStreamRangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count        *protocol.UnsignedIntRange `protobuf:"bytes,1,opt,name=count" json:"count,omitempty"`
	NameRange    *protocol.UnsignedIntRange `protobuf:"bytes,3,opt,name=name_range,json=nameRange" json:"name_range,omitempty"`
	SubjectRange *protocol.UnsignedIntRange `protobuf:"bytes,4,opt,name=subject_range,json=subjectRange" json:"subject_range,omitempty"`
	MessageRange *protocol.UnsignedIntRange `protobuf:"bytes,5,opt,name=message_range,json=messageRange" json:"message_range,omitempty"`
}

func (x *ClubStreamRangeSet) Reset() {
	*x = ClubStreamRangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubStreamRangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubStreamRangeSet) ProtoMessage() {}

func (x *ClubStreamRangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubStreamRangeSet.ProtoReflect.Descriptor instead.
func (*ClubStreamRangeSet) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP(), []int{2}
}

func (x *ClubStreamRangeSet) GetCount() *protocol.UnsignedIntRange {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *ClubStreamRangeSet) GetNameRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.NameRange
	}
	return nil
}

func (x *ClubStreamRangeSet) GetSubjectRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.SubjectRange
	}
	return nil
}

func (x *ClubStreamRangeSet) GetMessageRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.MessageRange
	}
	return nil
}

type ClubInvitationRangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count *protocol.UnsignedIntRange `protobuf:"bytes,1,opt,name=count" json:"count,omitempty"`
}

func (x *ClubInvitationRangeSet) Reset() {
	*x = ClubInvitationRangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubInvitationRangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubInvitationRangeSet) ProtoMessage() {}

func (x *ClubInvitationRangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubInvitationRangeSet.ProtoReflect.Descriptor instead.
func (*ClubInvitationRangeSet) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP(), []int{3}
}

func (x *ClubInvitationRangeSet) GetCount() *protocol.UnsignedIntRange {
	if x != nil {
		return x.Count
	}
	return nil
}

type ClubSuggestionRangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count *protocol.UnsignedIntRange `protobuf:"bytes,1,opt,name=count" json:"count,omitempty"`
}

func (x *ClubSuggestionRangeSet) Reset() {
	*x = ClubSuggestionRangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubSuggestionRangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubSuggestionRangeSet) ProtoMessage() {}

func (x *ClubSuggestionRangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubSuggestionRangeSet.ProtoReflect.Descriptor instead.
func (*ClubSuggestionRangeSet) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP(), []int{4}
}

func (x *ClubSuggestionRangeSet) GetCount() *protocol.UnsignedIntRange {
	if x != nil {
		return x.Count
	}
	return nil
}

type ClubTicketRangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count *protocol.UnsignedIntRange `protobuf:"bytes,1,opt,name=count" json:"count,omitempty"`
}

func (x *ClubTicketRangeSet) Reset() {
	*x = ClubTicketRangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubTicketRangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubTicketRangeSet) ProtoMessage() {}

func (x *ClubTicketRangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubTicketRangeSet.ProtoReflect.Descriptor instead.
func (*ClubTicketRangeSet) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP(), []int{5}
}

func (x *ClubTicketRangeSet) GetCount() *protocol.UnsignedIntRange {
	if x != nil {
		return x.Count
	}
	return nil
}

type ClubBanRangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       *protocol.UnsignedIntRange `protobuf:"bytes,1,opt,name=count" json:"count,omitempty"`
	ReasonRange *protocol.UnsignedIntRange `protobuf:"bytes,3,opt,name=reason_range,json=reasonRange" json:"reason_range,omitempty"`
}

func (x *ClubBanRangeSet) Reset() {
	*x = ClubBanRangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubBanRangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubBanRangeSet) ProtoMessage() {}

func (x *ClubBanRangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_range_set_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubBanRangeSet.ProtoReflect.Descriptor instead.
func (*ClubBanRangeSet) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP(), []int{6}
}

func (x *ClubBanRangeSet) GetCount() *protocol.UnsignedIntRange {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *ClubBanRangeSet) GetReasonRange() *protocol.UnsignedIntRange {
	if x != nil {
		return x.ReasonRange
	}
	return nil
}

var File_bgs_low_pb_client_club_range_set_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_club_range_set_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x2f,
	0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcc, 0x05, 0x0a, 0x10, 0x43, 0x6c, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x10,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x47, 0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x0e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x62, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x4c, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x1d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x03, 0x62, 0x61, 0x6e, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x42, 0x61,
	0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x03, 0x62, 0x61, 0x6e, 0x22, 0x92,
	0x02, 0x0a, 0x12, 0x43, 0x6c, 0x75, 0x62, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x51, 0x0a, 0x14, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55,
	0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x13, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0x93, 0x02, 0x0a, 0x12, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3d, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x43, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e,
	0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x16, 0x43, 0x6c, 0x75,
	0x62, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x16, 0x43, 0x6c, 0x75,
	0x62, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x12, 0x43, 0x6c, 0x75,
	0x62, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x34, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x62, 0x42, 0x61,
	0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x41, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x64, 0x34, 0x2d, 0x62, 0x6e,
	0x65, 0x74, 0x2d, 0x6d, 0x69, 0x74, 0x6d, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x62, 0x67, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x76,
	0x31,
}

var (
	file_bgs_low_pb_client_club_range_set_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_club_range_set_proto_rawDescData = file_bgs_low_pb_client_club_range_set_proto_rawDesc
)

func file_bgs_low_pb_client_club_range_set_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_club_range_set_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_club_range_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_club_range_set_proto_rawDescData)
	})
	return file_bgs_low_pb_client_club_range_set_proto_rawDescData
}

var file_bgs_low_pb_client_club_range_set_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bgs_low_pb_client_club_range_set_proto_goTypes = []interface{}{
	(*ClubTypeRangeSet)(nil),          // 0: bgs.protocol.club.v1.ClubTypeRangeSet
	(*ClubMemberRangeSet)(nil),        // 1: bgs.protocol.club.v1.ClubMemberRangeSet
	(*ClubStreamRangeSet)(nil),        // 2: bgs.protocol.club.v1.ClubStreamRangeSet
	(*ClubInvitationRangeSet)(nil),    // 3: bgs.protocol.club.v1.ClubInvitationRangeSet
	(*ClubSuggestionRangeSet)(nil),    // 4: bgs.protocol.club.v1.ClubSuggestionRangeSet
	(*ClubTicketRangeSet)(nil),        // 5: bgs.protocol.club.v1.ClubTicketRangeSet
	(*ClubBanRangeSet)(nil),           // 6: bgs.protocol.club.v1.ClubBanRangeSet
	(*protocol.UnsignedIntRange)(nil), // 7: bgs.protocol.UnsignedIntRange
}
var file_bgs_low_pb_client_club_range_set_proto_depIdxs = []int32{
	7,  // 0: bgs.protocol.club.v1.ClubTypeRangeSet.name_range:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 1: bgs.protocol.club.v1.ClubTypeRangeSet.description_range:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 2: bgs.protocol.club.v1.ClubTypeRangeSet.broadcast_range:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 3: bgs.protocol.club.v1.ClubTypeRangeSet.short_name_range:type_name -> bgs.protocol.UnsignedIntRange
	1,  // 4: bgs.protocol.club.v1.ClubTypeRangeSet.member:type_name -> bgs.protocol.club.v1.ClubMemberRangeSet
	2,  // 5: bgs.protocol.club.v1.ClubTypeRangeSet.stream:type_name -> bgs.protocol.club.v1.ClubStreamRangeSet
	3,  // 6: bgs.protocol.club.v1.ClubTypeRangeSet.invitation:type_name -> bgs.protocol.club.v1.ClubInvitationRangeSet
	4,  // 7: bgs.protocol.club.v1.ClubTypeRangeSet.suggestion:type_name -> bgs.protocol.club.v1.ClubSuggestionRangeSet
	5,  // 8: bgs.protocol.club.v1.ClubTypeRangeSet.ticket:type_name -> bgs.protocol.club.v1.ClubTicketRangeSet
	6,  // 9: bgs.protocol.club.v1.ClubTypeRangeSet.ban:type_name -> bgs.protocol.club.v1.ClubBanRangeSet
	7,  // 10: bgs.protocol.club.v1.ClubMemberRangeSet.count:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 11: bgs.protocol.club.v1.ClubMemberRangeSet.voice:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 12: bgs.protocol.club.v1.ClubMemberRangeSet.stream_subscriptions:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 13: bgs.protocol.club.v1.ClubMemberRangeSet.note_range:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 14: bgs.protocol.club.v1.ClubStreamRangeSet.count:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 15: bgs.protocol.club.v1.ClubStreamRangeSet.name_range:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 16: bgs.protocol.club.v1.ClubStreamRangeSet.subject_range:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 17: bgs.protocol.club.v1.ClubStreamRangeSet.message_range:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 18: bgs.protocol.club.v1.ClubInvitationRangeSet.count:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 19: bgs.protocol.club.v1.ClubSuggestionRangeSet.count:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 20: bgs.protocol.club.v1.ClubTicketRangeSet.count:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 21: bgs.protocol.club.v1.ClubBanRangeSet.count:type_name -> bgs.protocol.UnsignedIntRange
	7,  // 22: bgs.protocol.club.v1.ClubBanRangeSet.reason_range:type_name -> bgs.protocol.UnsignedIntRange
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_club_range_set_proto_init() }
func file_bgs_low_pb_client_club_range_set_proto_init() {
	if File_bgs_low_pb_client_club_range_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_club_range_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubTypeRangeSet); i {
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
		file_bgs_low_pb_client_club_range_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubMemberRangeSet); i {
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
		file_bgs_low_pb_client_club_range_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubStreamRangeSet); i {
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
		file_bgs_low_pb_client_club_range_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubInvitationRangeSet); i {
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
		file_bgs_low_pb_client_club_range_set_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubSuggestionRangeSet); i {
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
		file_bgs_low_pb_client_club_range_set_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubTicketRangeSet); i {
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
		file_bgs_low_pb_client_club_range_set_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubBanRangeSet); i {
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
			RawDescriptor: file_bgs_low_pb_client_club_range_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_club_range_set_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_club_range_set_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_club_range_set_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_club_range_set_proto = out.File
	file_bgs_low_pb_client_club_range_set_proto_rawDesc = nil
	file_bgs_low_pb_client_club_range_set_proto_goTypes = nil
	file_bgs_low_pb_client_club_range_set_proto_depIdxs = nil
}
