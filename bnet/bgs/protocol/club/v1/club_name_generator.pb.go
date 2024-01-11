// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/club_name_generator.proto

package v1

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

type NameGeneratorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameGenerators    []*NameGenerator        `protobuf:"bytes,1,rep,name=name_generators,json=nameGenerators" json:"name_generators,omitempty"`
	ClubTypeScorecard *NameGeneratorScorecard `protobuf:"bytes,2,opt,name=club_type_scorecard,json=clubTypeScorecard" json:"club_type_scorecard,omitempty"`
	LocaleScorecard   *NameGeneratorScorecard `protobuf:"bytes,3,opt,name=locale_scorecard,json=localeScorecard" json:"locale_scorecard,omitempty"`
}

func (x *NameGeneratorConfig) Reset() {
	*x = NameGeneratorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameGeneratorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameGeneratorConfig) ProtoMessage() {}

func (x *NameGeneratorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameGeneratorConfig.ProtoReflect.Descriptor instead.
func (*NameGeneratorConfig) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_name_generator_proto_rawDescGZIP(), []int{0}
}

func (x *NameGeneratorConfig) GetNameGenerators() []*NameGenerator {
	if x != nil {
		return x.NameGenerators
	}
	return nil
}

func (x *NameGeneratorConfig) GetClubTypeScorecard() *NameGeneratorScorecard {
	if x != nil {
		return x.ClubTypeScorecard
	}
	return nil
}

func (x *NameGeneratorConfig) GetLocaleScorecard() *NameGeneratorScorecard {
	if x != nil {
		return x.LocaleScorecard
	}
	return nil
}

type NameGeneratorScorecard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsRequired      *bool   `protobuf:"varint,1,opt,name=is_required,json=isRequired" json:"is_required,omitempty"`
	FullMatch       *uint32 `protobuf:"varint,2,opt,name=full_match,json=fullMatch" json:"full_match,omitempty"`
	PartialMatch    *uint32 `protobuf:"varint,3,opt,name=partial_match,json=partialMatch" json:"partial_match,omitempty"`
	PartialFallback *uint32 `protobuf:"varint,4,opt,name=partial_fallback,json=partialFallback" json:"partial_fallback,omitempty"`
	FullFallback    *uint32 `protobuf:"varint,5,opt,name=full_fallback,json=fullFallback" json:"full_fallback,omitempty"`
}

func (x *NameGeneratorScorecard) Reset() {
	*x = NameGeneratorScorecard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameGeneratorScorecard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameGeneratorScorecard) ProtoMessage() {}

func (x *NameGeneratorScorecard) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameGeneratorScorecard.ProtoReflect.Descriptor instead.
func (*NameGeneratorScorecard) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_name_generator_proto_rawDescGZIP(), []int{1}
}

func (x *NameGeneratorScorecard) GetIsRequired() bool {
	if x != nil && x.IsRequired != nil {
		return *x.IsRequired
	}
	return false
}

func (x *NameGeneratorScorecard) GetFullMatch() uint32 {
	if x != nil && x.FullMatch != nil {
		return *x.FullMatch
	}
	return 0
}

func (x *NameGeneratorScorecard) GetPartialMatch() uint32 {
	if x != nil && x.PartialMatch != nil {
		return *x.PartialMatch
	}
	return 0
}

func (x *NameGeneratorScorecard) GetPartialFallback() uint32 {
	if x != nil && x.PartialFallback != nil {
		return *x.PartialFallback
	}
	return 0
}

func (x *NameGeneratorScorecard) GetFullFallback() uint32 {
	if x != nil && x.FullFallback != nil {
		return *x.FullFallback
	}
	return 0
}

type NameGenerator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names        []string                    `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	Replacements []*NameGeneratorReplacement `protobuf:"bytes,2,rep,name=replacements" json:"replacements,omitempty"`
	ClubTypes    []*UniqueClubType           `protobuf:"bytes,3,rep,name=club_types,json=clubTypes" json:"club_types,omitempty"`
	Locales      []string                    `protobuf:"bytes,4,rep,name=locales" json:"locales,omitempty"`
}

func (x *NameGenerator) Reset() {
	*x = NameGenerator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameGenerator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameGenerator) ProtoMessage() {}

func (x *NameGenerator) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameGenerator.ProtoReflect.Descriptor instead.
func (*NameGenerator) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_name_generator_proto_rawDescGZIP(), []int{2}
}

func (x *NameGenerator) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *NameGenerator) GetReplacements() []*NameGeneratorReplacement {
	if x != nil {
		return x.Replacements
	}
	return nil
}

func (x *NameGenerator) GetClubTypes() []*UniqueClubType {
	if x != nil {
		return x.ClubTypes
	}
	return nil
}

func (x *NameGenerator) GetLocales() []string {
	if x != nil {
		return x.Locales
	}
	return nil
}

type NameGeneratorReplacement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Options []string `protobuf:"bytes,2,rep,name=options" json:"options,omitempty"`
}

func (x *NameGeneratorReplacement) Reset() {
	*x = NameGeneratorReplacement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameGeneratorReplacement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameGeneratorReplacement) ProtoMessage() {}

func (x *NameGeneratorReplacement) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_club_name_generator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameGeneratorReplacement.ProtoReflect.Descriptor instead.
func (*NameGeneratorReplacement) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_club_name_generator_proto_rawDescGZIP(), []int{3}
}

func (x *NameGeneratorReplacement) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *NameGeneratorReplacement) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_bgs_low_pb_client_club_name_generator_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_club_name_generator_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x1a, 0x21, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x13, 0x4e, 0x61, 0x6d, 0x65, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4c,
	0x0a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x6e, 0x61,
	0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x5c, 0x0a, 0x13,
	0x63, 0x6c, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x11, 0x63, 0x6c, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x12, 0x57, 0x0a, 0x10, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x52, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x22, 0xcd, 0x01, 0x0a, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x46, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x22, 0xd8, 0x01, 0x0a, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x0c, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x43, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x43, 0x6c, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x44,
	0x0a, 0x18, 0x4e, 0x61, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67,
	0x70, 0x72, 0x6f, 0x78, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x76, 0x31,
}

var (
	file_bgs_low_pb_client_club_name_generator_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_club_name_generator_proto_rawDescData = file_bgs_low_pb_client_club_name_generator_proto_rawDesc
)

func file_bgs_low_pb_client_club_name_generator_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_club_name_generator_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_club_name_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_club_name_generator_proto_rawDescData)
	})
	return file_bgs_low_pb_client_club_name_generator_proto_rawDescData
}

var file_bgs_low_pb_client_club_name_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bgs_low_pb_client_club_name_generator_proto_goTypes = []interface{}{
	(*NameGeneratorConfig)(nil),      // 0: bgs.protocol.club.v1.NameGeneratorConfig
	(*NameGeneratorScorecard)(nil),   // 1: bgs.protocol.club.v1.NameGeneratorScorecard
	(*NameGenerator)(nil),            // 2: bgs.protocol.club.v1.NameGenerator
	(*NameGeneratorReplacement)(nil), // 3: bgs.protocol.club.v1.NameGeneratorReplacement
	(*UniqueClubType)(nil),           // 4: bgs.protocol.club.v1.UniqueClubType
}
var file_bgs_low_pb_client_club_name_generator_proto_depIdxs = []int32{
	2, // 0: bgs.protocol.club.v1.NameGeneratorConfig.name_generators:type_name -> bgs.protocol.club.v1.NameGenerator
	1, // 1: bgs.protocol.club.v1.NameGeneratorConfig.club_type_scorecard:type_name -> bgs.protocol.club.v1.NameGeneratorScorecard
	1, // 2: bgs.protocol.club.v1.NameGeneratorConfig.locale_scorecard:type_name -> bgs.protocol.club.v1.NameGeneratorScorecard
	3, // 3: bgs.protocol.club.v1.NameGenerator.replacements:type_name -> bgs.protocol.club.v1.NameGeneratorReplacement
	4, // 4: bgs.protocol.club.v1.NameGenerator.club_types:type_name -> bgs.protocol.club.v1.UniqueClubType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_club_name_generator_proto_init() }
func file_bgs_low_pb_client_club_name_generator_proto_init() {
	if File_bgs_low_pb_client_club_name_generator_proto != nil {
		return
	}
	file_bgs_low_pb_client_club_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_club_name_generator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameGeneratorConfig); i {
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
		file_bgs_low_pb_client_club_name_generator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameGeneratorScorecard); i {
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
		file_bgs_low_pb_client_club_name_generator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameGenerator); i {
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
		file_bgs_low_pb_client_club_name_generator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameGeneratorReplacement); i {
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
			RawDescriptor: file_bgs_low_pb_client_club_name_generator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_club_name_generator_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_club_name_generator_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_club_name_generator_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_club_name_generator_proto = out.File
	file_bgs_low_pb_client_club_name_generator_proto_rawDesc = nil
	file_bgs_low_pb_client_club_name_generator_proto_goTypes = nil
	file_bgs_low_pb_client_club_name_generator_proto_depIdxs = nil
}
