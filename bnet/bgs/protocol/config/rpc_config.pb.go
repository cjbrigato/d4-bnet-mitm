// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bgs/low/pb/client/rpc_config.proto

package config

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

type RPCMethodConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName        *string  `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName         *string  `protobuf:"bytes,2,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	FixedCallCost      *uint32  `protobuf:"varint,3,opt,name=fixed_call_cost,json=fixedCallCost,def=1" json:"fixed_call_cost,omitempty"`
	FixedPacketSize    *uint32  `protobuf:"varint,4,opt,name=fixed_packet_size,json=fixedPacketSize" json:"fixed_packet_size,omitempty"`
	VariableMultiplier *float32 `protobuf:"fixed32,5,opt,name=variable_multiplier,json=variableMultiplier" json:"variable_multiplier,omitempty"`
	Multiplier         *float32 `protobuf:"fixed32,6,opt,name=multiplier,def=1" json:"multiplier,omitempty"`
	RateLimitCount     *uint32  `protobuf:"varint,7,opt,name=rate_limit_count,json=rateLimitCount" json:"rate_limit_count,omitempty"`
	RateLimitSeconds   *uint32  `protobuf:"varint,8,opt,name=rate_limit_seconds,json=rateLimitSeconds" json:"rate_limit_seconds,omitempty"`
	MaxPacketSize      *uint32  `protobuf:"varint,9,opt,name=max_packet_size,json=maxPacketSize" json:"max_packet_size,omitempty"`
	MaxEncodedSize     *uint32  `protobuf:"varint,10,opt,name=max_encoded_size,json=maxEncodedSize" json:"max_encoded_size,omitempty"`
	Timeout            *float32 `protobuf:"fixed32,11,opt,name=timeout" json:"timeout,omitempty"`
	CapBalance         *uint32  `protobuf:"varint,12,opt,name=cap_balance,json=capBalance" json:"cap_balance,omitempty"`
	IncomePerSecond    *float32 `protobuf:"fixed32,13,opt,name=income_per_second,json=incomePerSecond" json:"income_per_second,omitempty"`
	ServiceHash        *uint32  `protobuf:"varint,14,opt,name=service_hash,json=serviceHash" json:"service_hash,omitempty"`
	MethodId           *uint32  `protobuf:"varint,15,opt,name=method_id,json=methodId" json:"method_id,omitempty"`
}

// Default values for RPCMethodConfig fields.
const (
	Default_RPCMethodConfig_FixedCallCost = uint32(1)
	Default_RPCMethodConfig_Multiplier    = float32(1)
)

func (x *RPCMethodConfig) Reset() {
	*x = RPCMethodConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_rpc_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCMethodConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCMethodConfig) ProtoMessage() {}

func (x *RPCMethodConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_rpc_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCMethodConfig.ProtoReflect.Descriptor instead.
func (*RPCMethodConfig) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_rpc_config_proto_rawDescGZIP(), []int{0}
}

func (x *RPCMethodConfig) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *RPCMethodConfig) GetMethodName() string {
	if x != nil && x.MethodName != nil {
		return *x.MethodName
	}
	return ""
}

func (x *RPCMethodConfig) GetFixedCallCost() uint32 {
	if x != nil && x.FixedCallCost != nil {
		return *x.FixedCallCost
	}
	return Default_RPCMethodConfig_FixedCallCost
}

func (x *RPCMethodConfig) GetFixedPacketSize() uint32 {
	if x != nil && x.FixedPacketSize != nil {
		return *x.FixedPacketSize
	}
	return 0
}

func (x *RPCMethodConfig) GetVariableMultiplier() float32 {
	if x != nil && x.VariableMultiplier != nil {
		return *x.VariableMultiplier
	}
	return 0
}

func (x *RPCMethodConfig) GetMultiplier() float32 {
	if x != nil && x.Multiplier != nil {
		return *x.Multiplier
	}
	return Default_RPCMethodConfig_Multiplier
}

func (x *RPCMethodConfig) GetRateLimitCount() uint32 {
	if x != nil && x.RateLimitCount != nil {
		return *x.RateLimitCount
	}
	return 0
}

func (x *RPCMethodConfig) GetRateLimitSeconds() uint32 {
	if x != nil && x.RateLimitSeconds != nil {
		return *x.RateLimitSeconds
	}
	return 0
}

func (x *RPCMethodConfig) GetMaxPacketSize() uint32 {
	if x != nil && x.MaxPacketSize != nil {
		return *x.MaxPacketSize
	}
	return 0
}

func (x *RPCMethodConfig) GetMaxEncodedSize() uint32 {
	if x != nil && x.MaxEncodedSize != nil {
		return *x.MaxEncodedSize
	}
	return 0
}

func (x *RPCMethodConfig) GetTimeout() float32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *RPCMethodConfig) GetCapBalance() uint32 {
	if x != nil && x.CapBalance != nil {
		return *x.CapBalance
	}
	return 0
}

func (x *RPCMethodConfig) GetIncomePerSecond() float32 {
	if x != nil && x.IncomePerSecond != nil {
		return *x.IncomePerSecond
	}
	return 0
}

func (x *RPCMethodConfig) GetServiceHash() uint32 {
	if x != nil && x.ServiceHash != nil {
		return *x.ServiceHash
	}
	return 0
}

func (x *RPCMethodConfig) GetMethodId() uint32 {
	if x != nil && x.MethodId != nil {
		return *x.MethodId
	}
	return 0
}

type RPCMeterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method          []*RPCMethodConfig `protobuf:"bytes,1,rep,name=method" json:"method,omitempty"`
	IncomePerSecond *uint32            `protobuf:"varint,2,opt,name=income_per_second,json=incomePerSecond,def=1" json:"income_per_second,omitempty"`
	InitialBalance  *uint32            `protobuf:"varint,3,opt,name=initial_balance,json=initialBalance" json:"initial_balance,omitempty"`
	CapBalance      *uint32            `protobuf:"varint,4,opt,name=cap_balance,json=capBalance" json:"cap_balance,omitempty"`
	StartupPeriod   *float32           `protobuf:"fixed32,5,opt,name=startup_period,json=startupPeriod" json:"startup_period,omitempty"`
}

// Default values for RPCMeterConfig fields.
const (
	Default_RPCMeterConfig_IncomePerSecond = uint32(1)
)

func (x *RPCMeterConfig) Reset() {
	*x = RPCMeterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_rpc_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCMeterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCMeterConfig) ProtoMessage() {}

func (x *RPCMeterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_rpc_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCMeterConfig.ProtoReflect.Descriptor instead.
func (*RPCMeterConfig) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_rpc_config_proto_rawDescGZIP(), []int{1}
}

func (x *RPCMeterConfig) GetMethod() []*RPCMethodConfig {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *RPCMeterConfig) GetIncomePerSecond() uint32 {
	if x != nil && x.IncomePerSecond != nil {
		return *x.IncomePerSecond
	}
	return Default_RPCMeterConfig_IncomePerSecond
}

func (x *RPCMeterConfig) GetInitialBalance() uint32 {
	if x != nil && x.InitialBalance != nil {
		return *x.InitialBalance
	}
	return 0
}

func (x *RPCMeterConfig) GetCapBalance() uint32 {
	if x != nil && x.CapBalance != nil {
		return *x.CapBalance
	}
	return 0
}

func (x *RPCMeterConfig) GetStartupPeriod() float32 {
	if x != nil && x.StartupPeriod != nil {
		return *x.StartupPeriod
	}
	return 0
}

var File_bgs_low_pb_client_rpc_config_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_rpc_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xd1, 0x04, 0x0a, 0x0f, 0x52, 0x50,
	0x43, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x0f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x31, 0x52, 0x0d, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0a, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x31,
	0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d,
	0x61, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x6d, 0x61, 0x78, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x70, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x22, 0xee, 0x01,
	0x0a, 0x0e, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3c, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2d,
	0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x31, 0x52, 0x0f, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x70, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x70,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x75, 0x70, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6a, 0x62,
	0x72, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x78, 0x2f, 0x62,
	0x6e, 0x65, 0x74, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
}

var (
	file_bgs_low_pb_client_rpc_config_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_rpc_config_proto_rawDescData = file_bgs_low_pb_client_rpc_config_proto_rawDesc
)

func file_bgs_low_pb_client_rpc_config_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_rpc_config_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_rpc_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_rpc_config_proto_rawDescData)
	})
	return file_bgs_low_pb_client_rpc_config_proto_rawDescData
}

var file_bgs_low_pb_client_rpc_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bgs_low_pb_client_rpc_config_proto_goTypes = []interface{}{
	(*RPCMethodConfig)(nil), // 0: bgs.protocol.config.RPCMethodConfig
	(*RPCMeterConfig)(nil),  // 1: bgs.protocol.config.RPCMeterConfig
}
var file_bgs_low_pb_client_rpc_config_proto_depIdxs = []int32{
	0, // 0: bgs.protocol.config.RPCMeterConfig.method:type_name -> bgs.protocol.config.RPCMethodConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_rpc_config_proto_init() }
func file_bgs_low_pb_client_rpc_config_proto_init() {
	if File_bgs_low_pb_client_rpc_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_rpc_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCMethodConfig); i {
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
		file_bgs_low_pb_client_rpc_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCMeterConfig); i {
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
			RawDescriptor: file_bgs_low_pb_client_rpc_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_rpc_config_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_rpc_config_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_rpc_config_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_rpc_config_proto = out.File
	file_bgs_low_pb_client_rpc_config_proto_rawDesc = nil
	file_bgs_low_pb_client_rpc_config_proto_goTypes = nil
	file_bgs_low_pb_client_rpc_config_proto_depIdxs = nil
}
