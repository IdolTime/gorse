// Copyright 2020 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: protocol.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NodeType int32

const (
	NodeType_ServerNode NodeType = 0
	NodeType_WorkerNode NodeType = 1
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "ServerNode",
		1: "WorkerNode",
	}
	NodeType_value = map[string]int32{
		"ServerNode": 0,
		"WorkerNode": 1,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_proto_enumTypes[0].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_protocol_proto_enumTypes[0]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config           string   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	UserIndexVersion int64    `protobuf:"varint,2,opt,name=user_index_version,json=userIndexVersion,proto3" json:"user_index_version,omitempty"`
	CfVersion        int64    `protobuf:"varint,3,opt,name=cf_version,json=cfVersion,proto3" json:"cf_version,omitempty"`
	FmVersion        int64    `protobuf:"varint,4,opt,name=fm_version,json=fmVersion,proto3" json:"fm_version,omitempty"`
	Me               string   `protobuf:"bytes,5,opt,name=me,proto3" json:"me,omitempty"`
	Servers          []string `protobuf:"bytes,6,rep,name=servers,proto3" json:"servers,omitempty"`
	Workers          []string `protobuf:"bytes,7,rep,name=workers,proto3" json:"workers,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Meta) GetUserIndexVersion() int64 {
	if x != nil {
		return x.UserIndexVersion
	}
	return 0
}

func (x *Meta) GetCfVersion() int64 {
	if x != nil {
		return x.CfVersion
	}
	return 0
}

func (x *Meta) GetFmVersion() int64 {
	if x != nil {
		return x.FmVersion
	}
	return 0
}

func (x *Meta) GetMe() string {
	if x != nil {
		return x.Me
	}
	return ""
}

func (x *Meta) GetServers() []string {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *Meta) GetWorkers() []string {
	if x != nil {
		return x.Workers
	}
	return nil
}

type UserIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   int64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`                     // user index version
	UserIndex []byte `protobuf:"bytes,2,opt,name=user_index,json=userIndex,proto3" json:"user_index,omitempty"` // user index data
}

func (x *UserIndex) Reset() {
	*x = UserIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIndex) ProtoMessage() {}

func (x *UserIndex) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIndex.ProtoReflect.Descriptor instead.
func (*UserIndex) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *UserIndex) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UserIndex) GetUserIndex() []byte {
	if x != nil {
		return x.UserIndex
	}
	return nil
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"` // model version
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`        // model name
	Model   []byte `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`      // model data
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *Model) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetModel() []byte {
	if x != nil {
		return x.Model
	}
	return nil
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeType NodeType `protobuf:"varint,1,opt,name=node_type,json=nodeType,proto3,enum=protocol.NodeType" json:"node_type,omitempty"`
	NodeName string   `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *NodeInfo) GetNodeType() NodeType {
	if x != nil {
		return x.NodeType
	}
	return NodeType_ServerNode
}

func (x *NodeInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

var File_protocol_proto protoreflect.FileDescriptor

var file_protocol_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xce, 0x01, 0x0a, 0x04, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x12, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x66, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x66, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6d, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x6d,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x44, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x4b, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x58,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x2a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x10, 0x01, 0x32, 0xdf, 0x01, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x54, 0x52, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x52, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x65, 0x6e, 0x67, 0x68, 0x61, 0x6f, 0x7a, 0x2f, 0x67,
	0x6f, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_rawDescOnce sync.Once
	file_protocol_proto_rawDescData = file_protocol_proto_rawDesc
)

func file_protocol_proto_rawDescGZIP() []byte {
	file_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_rawDescData)
	})
	return file_protocol_proto_rawDescData
}

var file_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_proto_goTypes = []interface{}{
	(NodeType)(0),     // 0: protocol.NodeType
	(*Meta)(nil),      // 1: protocol.Meta
	(*UserIndex)(nil), // 2: protocol.UserIndex
	(*Model)(nil),     // 3: protocol.Model
	(*NodeInfo)(nil),  // 4: protocol.NodeInfo
}
var file_protocol_proto_depIdxs = []int32{
	0, // 0: protocol.NodeInfo.node_type:type_name -> protocol.NodeType
	4, // 1: protocol.Master.GetMeta:input_type -> protocol.NodeInfo
	4, // 2: protocol.Master.GetUserIndex:input_type -> protocol.NodeInfo
	4, // 3: protocol.Master.GetCTRModel:input_type -> protocol.NodeInfo
	4, // 4: protocol.Master.GetPRModel:input_type -> protocol.NodeInfo
	1, // 5: protocol.Master.GetMeta:output_type -> protocol.Meta
	2, // 6: protocol.Master.GetUserIndex:output_type -> protocol.UserIndex
	3, // 7: protocol.Master.GetCTRModel:output_type -> protocol.Model
	3, // 8: protocol.Master.GetPRModel:output_type -> protocol.Model
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_proto_init() }
func file_protocol_proto_init() {
	if File_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIndex); i {
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
		file_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
			RawDescriptor: file_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_proto_depIdxs,
		EnumInfos:         file_protocol_proto_enumTypes,
		MessageInfos:      file_protocol_proto_msgTypes,
	}.Build()
	File_protocol_proto = out.File
	file_protocol_proto_rawDesc = nil
	file_protocol_proto_goTypes = nil
	file_protocol_proto_depIdxs = nil
}
