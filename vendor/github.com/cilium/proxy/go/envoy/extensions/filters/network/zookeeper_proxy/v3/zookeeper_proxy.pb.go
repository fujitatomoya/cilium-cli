// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.24.4
// source: envoy/extensions/filters/network/zookeeper_proxy/v3/zookeeper_proxy.proto

package zookeeper_proxyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LatencyThresholdOverride_Opcode int32

const (
	LatencyThresholdOverride_Connect              LatencyThresholdOverride_Opcode = 0
	LatencyThresholdOverride_Create               LatencyThresholdOverride_Opcode = 1
	LatencyThresholdOverride_Delete               LatencyThresholdOverride_Opcode = 2
	LatencyThresholdOverride_Exists               LatencyThresholdOverride_Opcode = 3
	LatencyThresholdOverride_GetData              LatencyThresholdOverride_Opcode = 4
	LatencyThresholdOverride_SetData              LatencyThresholdOverride_Opcode = 5
	LatencyThresholdOverride_GetAcl               LatencyThresholdOverride_Opcode = 6
	LatencyThresholdOverride_SetAcl               LatencyThresholdOverride_Opcode = 7
	LatencyThresholdOverride_GetChildren          LatencyThresholdOverride_Opcode = 8
	LatencyThresholdOverride_Sync                 LatencyThresholdOverride_Opcode = 9
	LatencyThresholdOverride_Ping                 LatencyThresholdOverride_Opcode = 10
	LatencyThresholdOverride_GetChildren2         LatencyThresholdOverride_Opcode = 11
	LatencyThresholdOverride_Check                LatencyThresholdOverride_Opcode = 12
	LatencyThresholdOverride_Multi                LatencyThresholdOverride_Opcode = 13
	LatencyThresholdOverride_Create2              LatencyThresholdOverride_Opcode = 14
	LatencyThresholdOverride_Reconfig             LatencyThresholdOverride_Opcode = 15
	LatencyThresholdOverride_CheckWatches         LatencyThresholdOverride_Opcode = 16
	LatencyThresholdOverride_RemoveWatches        LatencyThresholdOverride_Opcode = 17
	LatencyThresholdOverride_CreateContainer      LatencyThresholdOverride_Opcode = 18
	LatencyThresholdOverride_CreateTtl            LatencyThresholdOverride_Opcode = 19
	LatencyThresholdOverride_Close                LatencyThresholdOverride_Opcode = 20
	LatencyThresholdOverride_SetAuth              LatencyThresholdOverride_Opcode = 21
	LatencyThresholdOverride_SetWatches           LatencyThresholdOverride_Opcode = 22
	LatencyThresholdOverride_GetEphemerals        LatencyThresholdOverride_Opcode = 23
	LatencyThresholdOverride_GetAllChildrenNumber LatencyThresholdOverride_Opcode = 24
	LatencyThresholdOverride_SetWatches2          LatencyThresholdOverride_Opcode = 25
	LatencyThresholdOverride_AddWatch             LatencyThresholdOverride_Opcode = 26
)

// Enum value maps for LatencyThresholdOverride_Opcode.
var (
	LatencyThresholdOverride_Opcode_name = map[int32]string{
		0:  "Connect",
		1:  "Create",
		2:  "Delete",
		3:  "Exists",
		4:  "GetData",
		5:  "SetData",
		6:  "GetAcl",
		7:  "SetAcl",
		8:  "GetChildren",
		9:  "Sync",
		10: "Ping",
		11: "GetChildren2",
		12: "Check",
		13: "Multi",
		14: "Create2",
		15: "Reconfig",
		16: "CheckWatches",
		17: "RemoveWatches",
		18: "CreateContainer",
		19: "CreateTtl",
		20: "Close",
		21: "SetAuth",
		22: "SetWatches",
		23: "GetEphemerals",
		24: "GetAllChildrenNumber",
		25: "SetWatches2",
		26: "AddWatch",
	}
	LatencyThresholdOverride_Opcode_value = map[string]int32{
		"Connect":              0,
		"Create":               1,
		"Delete":               2,
		"Exists":               3,
		"GetData":              4,
		"SetData":              5,
		"GetAcl":               6,
		"SetAcl":               7,
		"GetChildren":          8,
		"Sync":                 9,
		"Ping":                 10,
		"GetChildren2":         11,
		"Check":                12,
		"Multi":                13,
		"Create2":              14,
		"Reconfig":             15,
		"CheckWatches":         16,
		"RemoveWatches":        17,
		"CreateContainer":      18,
		"CreateTtl":            19,
		"Close":                20,
		"SetAuth":              21,
		"SetWatches":           22,
		"GetEphemerals":        23,
		"GetAllChildrenNumber": 24,
		"SetWatches2":          25,
		"AddWatch":             26,
	}
)

func (x LatencyThresholdOverride_Opcode) Enum() *LatencyThresholdOverride_Opcode {
	p := new(LatencyThresholdOverride_Opcode)
	*p = x
	return p
}

func (x LatencyThresholdOverride_Opcode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LatencyThresholdOverride_Opcode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_enumTypes[0].Descriptor()
}

func (LatencyThresholdOverride_Opcode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_enumTypes[0]
}

func (x LatencyThresholdOverride_Opcode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LatencyThresholdOverride_Opcode.Descriptor instead.
func (LatencyThresholdOverride_Opcode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescGZIP(), []int{1, 0}
}

// [#next-free-field: 10]
type ZooKeeperProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_zookeeper_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// [#not-implemented-hide:] The optional path to use for writing ZooKeeper access logs.
	// If the access log field is empty, access logs will not be written.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Messages — requests, responses and events — that are bigger than this value will
	// be ignored. If it is not set, the default value is 1Mb.
	//
	// The value here should match the jute.maxbuffer property in your cluster configuration:
	//
	// https://zookeeper.apache.org/doc/r3.4.10/zookeeperAdmin.html#Unsafe+Options
	//
	// if that is set. If it isn't, ZooKeeper's default is also 1Mb.
	MaxPacketBytes *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_packet_bytes,json=maxPacketBytes,proto3" json:"max_packet_bytes,omitempty"`
	// Whether to emit latency threshold metrics. If not set, it defaults to false.
	// If false, setting “default_latency_threshold“ and “latency_threshold_overrides“ will not have effect.
	EnableLatencyThresholdMetrics bool `protobuf:"varint,4,opt,name=enable_latency_threshold_metrics,json=enableLatencyThresholdMetrics,proto3" json:"enable_latency_threshold_metrics,omitempty"`
	// The default latency threshold to decide the fast/slow responses and emit metrics (used for error budget calculation).
	//
	// https://sre.google/workbook/implementing-slos/
	//
	// If it is not set, the default value is 100 milliseconds.
	DefaultLatencyThreshold *durationpb.Duration `protobuf:"bytes,5,opt,name=default_latency_threshold,json=defaultLatencyThreshold,proto3" json:"default_latency_threshold,omitempty"`
	// List of latency threshold overrides for opcodes.
	// If the threshold override of one opcode is not set, it will fallback to the default latency
	// threshold.
	// Specifying latency threshold overrides multiple times for one opcode is not allowed.
	LatencyThresholdOverrides []*LatencyThresholdOverride `protobuf:"bytes,6,rep,name=latency_threshold_overrides,json=latencyThresholdOverrides,proto3" json:"latency_threshold_overrides,omitempty"`
	// Whether to emit per opcode request bytes metrics. If not set, it defaults to false.
	EnablePerOpcodeRequestBytesMetrics bool `protobuf:"varint,7,opt,name=enable_per_opcode_request_bytes_metrics,json=enablePerOpcodeRequestBytesMetrics,proto3" json:"enable_per_opcode_request_bytes_metrics,omitempty"`
	// Whether to emit per opcode response bytes metrics. If not set, it defaults to false.
	EnablePerOpcodeResponseBytesMetrics bool `protobuf:"varint,8,opt,name=enable_per_opcode_response_bytes_metrics,json=enablePerOpcodeResponseBytesMetrics,proto3" json:"enable_per_opcode_response_bytes_metrics,omitempty"`
	// Whether to emit per opcode decoder error metrics. If not set, it defaults to false.
	EnablePerOpcodeDecoderErrorMetrics bool `protobuf:"varint,9,opt,name=enable_per_opcode_decoder_error_metrics,json=enablePerOpcodeDecoderErrorMetrics,proto3" json:"enable_per_opcode_decoder_error_metrics,omitempty"`
}

func (x *ZooKeeperProxy) Reset() {
	*x = ZooKeeperProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZooKeeperProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZooKeeperProxy) ProtoMessage() {}

func (x *ZooKeeperProxy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZooKeeperProxy.ProtoReflect.Descriptor instead.
func (*ZooKeeperProxy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *ZooKeeperProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *ZooKeeperProxy) GetAccessLog() string {
	if x != nil {
		return x.AccessLog
	}
	return ""
}

func (x *ZooKeeperProxy) GetMaxPacketBytes() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxPacketBytes
	}
	return nil
}

func (x *ZooKeeperProxy) GetEnableLatencyThresholdMetrics() bool {
	if x != nil {
		return x.EnableLatencyThresholdMetrics
	}
	return false
}

func (x *ZooKeeperProxy) GetDefaultLatencyThreshold() *durationpb.Duration {
	if x != nil {
		return x.DefaultLatencyThreshold
	}
	return nil
}

func (x *ZooKeeperProxy) GetLatencyThresholdOverrides() []*LatencyThresholdOverride {
	if x != nil {
		return x.LatencyThresholdOverrides
	}
	return nil
}

func (x *ZooKeeperProxy) GetEnablePerOpcodeRequestBytesMetrics() bool {
	if x != nil {
		return x.EnablePerOpcodeRequestBytesMetrics
	}
	return false
}

func (x *ZooKeeperProxy) GetEnablePerOpcodeResponseBytesMetrics() bool {
	if x != nil {
		return x.EnablePerOpcodeResponseBytesMetrics
	}
	return false
}

func (x *ZooKeeperProxy) GetEnablePerOpcodeDecoderErrorMetrics() bool {
	if x != nil {
		return x.EnablePerOpcodeDecoderErrorMetrics
	}
	return false
}

type LatencyThresholdOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ZooKeeper opcodes. Can be found as part of the ZooKeeper source code:
	//
	// https://github.com/apache/zookeeper/blob/master/zookeeper-server/src/main/java/org/apache/zookeeper/ZooDefs.java
	Opcode LatencyThresholdOverride_Opcode `protobuf:"varint,1,opt,name=opcode,proto3,enum=envoy.extensions.filters.network.zookeeper_proxy.v3.LatencyThresholdOverride_Opcode" json:"opcode,omitempty"`
	// The latency threshold override of certain opcode.
	Threshold *durationpb.Duration `protobuf:"bytes,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (x *LatencyThresholdOverride) Reset() {
	*x = LatencyThresholdOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatencyThresholdOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatencyThresholdOverride) ProtoMessage() {}

func (x *LatencyThresholdOverride) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatencyThresholdOverride.ProtoReflect.Descriptor instead.
func (*LatencyThresholdOverride) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *LatencyThresholdOverride) GetOpcode() LatencyThresholdOverride_Opcode {
	if x != nil {
		return x.Opcode
	}
	return LatencyThresholdOverride_Connect
}

func (x *LatencyThresholdOverride) GetThreshold() *durationpb.Duration {
	if x != nil {
		return x.Threshold
	}
	return nil
}

var File_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDesc = []byte{
	0x0a, 0x49, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x6f,
	0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x06, 0x0a, 0x0e,
	0x5a, 0x6f, 0x6f, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x28,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x46, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0e, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x47, 0x0a, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x63, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0xaa, 0x01, 0x06, 0x32, 0x04,
	0x10, 0xc0, 0x84, 0x3d, 0x52, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x8d, 0x01,
	0x0a, 0x1b, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x52, 0x19, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x53, 0x0a,
	0x27, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x22,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x55, 0x0a, 0x28, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x23, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x4f,
	0x70, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x53, 0x0a, 0x27, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x22, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x50, 0x65, 0x72, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x3a, 0x4a,
	0x9a, 0xc5, 0x88, 0x1e, 0x45, 0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6f, 0x4b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0xee, 0x04, 0x0a, 0x18, 0x4c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x76, 0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x54, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x47, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0xfa,
	0x42, 0x0b, 0xaa, 0x01, 0x08, 0x08, 0x01, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52, 0x09, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0x90, 0x03, 0x0a, 0x06, 0x4f, 0x70, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x10,
	0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x10, 0x05, 0x12, 0x0a,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x63, 0x6c, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65,
	0x74, 0x41, 0x63, 0x6c, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x10,
	0x09, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x32, 0x10, 0x0b, 0x12, 0x09, 0x0a,
	0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x10, 0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x32, 0x10, 0x0e,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x10, 0x0f, 0x12, 0x10,
	0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x10, 0x10,
	0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x10, 0x11, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x10, 0x12, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x74, 0x6c, 0x10, 0x13, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x10, 0x15, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x10, 0x16, 0x12,
	0x11, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x73,
	0x10, 0x17, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x32, 0x10, 0x19, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x64, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x10, 0x1a, 0x42, 0xd0, 0x01, 0x0a, 0x41,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a,
	0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x33, 0x42, 0x13, 0x5a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x33, 0x3b, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescData = file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_goTypes = []interface{}{
	(LatencyThresholdOverride_Opcode)(0), // 0: envoy.extensions.filters.network.zookeeper_proxy.v3.LatencyThresholdOverride.Opcode
	(*ZooKeeperProxy)(nil),               // 1: envoy.extensions.filters.network.zookeeper_proxy.v3.ZooKeeperProxy
	(*LatencyThresholdOverride)(nil),     // 2: envoy.extensions.filters.network.zookeeper_proxy.v3.LatencyThresholdOverride
	(*wrapperspb.UInt32Value)(nil),       // 3: google.protobuf.UInt32Value
	(*durationpb.Duration)(nil),          // 4: google.protobuf.Duration
}
var file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.network.zookeeper_proxy.v3.ZooKeeperProxy.max_packet_bytes:type_name -> google.protobuf.UInt32Value
	4, // 1: envoy.extensions.filters.network.zookeeper_proxy.v3.ZooKeeperProxy.default_latency_threshold:type_name -> google.protobuf.Duration
	2, // 2: envoy.extensions.filters.network.zookeeper_proxy.v3.ZooKeeperProxy.latency_threshold_overrides:type_name -> envoy.extensions.filters.network.zookeeper_proxy.v3.LatencyThresholdOverride
	0, // 3: envoy.extensions.filters.network.zookeeper_proxy.v3.LatencyThresholdOverride.opcode:type_name -> envoy.extensions.filters.network.zookeeper_proxy.v3.LatencyThresholdOverride.Opcode
	4, // 4: envoy.extensions.filters.network.zookeeper_proxy.v3.LatencyThresholdOverride.threshold:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_init() }
func file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_init() {
	if File_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZooKeeperProxy); i {
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
		file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatencyThresholdOverride); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto = out.File
	file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_network_zookeeper_proxy_v3_zookeeper_proxy_proto_depIdxs = nil
}
