// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/api/client/v2/report_service.proto

package v2

import (
	protocol "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
	v1 "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/account/v1"
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

type SubmitReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId         *v1.AccountId  `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	UserDescription *string        `protobuf:"bytes,2,opt,name=user_description,json=userDescription" json:"user_description,omitempty"`
	UserOptions     *UserOptions   `protobuf:"bytes,10,opt,name=user_options,json=userOptions" json:"user_options,omitempty"`
	ClubOptions     *ClubOptions   `protobuf:"bytes,11,opt,name=club_options,json=clubOptions" json:"club_options,omitempty"`
	EntityOptions   *EntityOptions `protobuf:"bytes,20,opt,name=entity_options,json=entityOptions" json:"entity_options,omitempty"`
}

func (x *SubmitReportRequest) Reset() {
	*x = SubmitReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_api_client_v2_report_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitReportRequest) ProtoMessage() {}

func (x *SubmitReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_api_client_v2_report_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitReportRequest.ProtoReflect.Descriptor instead.
func (*SubmitReportRequest) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitReportRequest) GetAgentId() *v1.AccountId {
	if x != nil {
		return x.AgentId
	}
	return nil
}

func (x *SubmitReportRequest) GetUserDescription() string {
	if x != nil && x.UserDescription != nil {
		return *x.UserDescription
	}
	return ""
}

func (x *SubmitReportRequest) GetUserOptions() *UserOptions {
	if x != nil {
		return x.UserOptions
	}
	return nil
}

func (x *SubmitReportRequest) GetClubOptions() *ClubOptions {
	if x != nil {
		return x.ClubOptions
	}
	return nil
}

func (x *SubmitReportRequest) GetEntityOptions() *EntityOptions {
	if x != nil {
		return x.EntityOptions
	}
	return nil
}

var File_bgs_low_pb_client_api_client_v2_report_service_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x25,
	0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x62, 0x67, 0x73, 0x2f, 0x6c,
	0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a,
	0x13, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x62, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c,
	0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x62, 0x0a, 0x0d,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2b, 0x2e,
	0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6a, 0x62, 0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x64, 0x34, 0x2d, 0x62, 0x6e, 0x65, 0x74,
	0x2d, 0x6d, 0x69, 0x74, 0x6d, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76,
	0x32,
}

var (
	file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescData = file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDesc
)

func file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescData)
	})
	return file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDescData
}

var file_bgs_low_pb_client_api_client_v2_report_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bgs_low_pb_client_api_client_v2_report_service_proto_goTypes = []interface{}{
	(*SubmitReportRequest)(nil), // 0: bgs.protocol.report.v2.SubmitReportRequest
	(*v1.AccountId)(nil),        // 1: bgs.protocol.account.v1.AccountId
	(*UserOptions)(nil),         // 2: bgs.protocol.report.v2.UserOptions
	(*ClubOptions)(nil),         // 3: bgs.protocol.report.v2.ClubOptions
	(*EntityOptions)(nil),       // 4: bgs.protocol.report.v2.EntityOptions
	(*protocol.NoData)(nil),     // 5: bgs.protocol.NoData
}
var file_bgs_low_pb_client_api_client_v2_report_service_proto_depIdxs = []int32{
	1, // 0: bgs.protocol.report.v2.SubmitReportRequest.agent_id:type_name -> bgs.protocol.account.v1.AccountId
	2, // 1: bgs.protocol.report.v2.SubmitReportRequest.user_options:type_name -> bgs.protocol.report.v2.UserOptions
	3, // 2: bgs.protocol.report.v2.SubmitReportRequest.club_options:type_name -> bgs.protocol.report.v2.ClubOptions
	4, // 3: bgs.protocol.report.v2.SubmitReportRequest.entity_options:type_name -> bgs.protocol.report.v2.EntityOptions
	0, // 4: bgs.protocol.report.v2.ReportService.SubmitReport:input_type -> bgs.protocol.report.v2.SubmitReportRequest
	5, // 5: bgs.protocol.report.v2.ReportService.SubmitReport:output_type -> bgs.protocol.NoData
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_api_client_v2_report_service_proto_init() }
func file_bgs_low_pb_client_api_client_v2_report_service_proto_init() {
	if File_bgs_low_pb_client_api_client_v2_report_service_proto != nil {
		return
	}
	file_bgs_low_pb_client_api_client_v2_report_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_api_client_v2_report_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitReportRequest); i {
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
			RawDescriptor: file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bgs_low_pb_client_api_client_v2_report_service_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_api_client_v2_report_service_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_api_client_v2_report_service_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_api_client_v2_report_service_proto = out.File
	file_bgs_low_pb_client_api_client_v2_report_service_proto_rawDesc = nil
	file_bgs_low_pb_client_api_client_v2_report_service_proto_goTypes = nil
	file_bgs_low_pb_client_api_client_v2_report_service_proto_depIdxs = nil
}
