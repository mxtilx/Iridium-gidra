// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.3
// source: GCGMsgGameOver.proto

package gen

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

type GCGMsgGameOver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndReason       GCGEndReason `protobuf:"varint,6,opt,name=end_reason,json=endReason,proto3,enum=GCGEndReason" json:"end_reason,omitempty"`
	WinControllerId uint32       `protobuf:"varint,3,opt,name=win_controller_id,json=winControllerId,proto3" json:"win_controller_id,omitempty"`
}

func (x *GCGMsgGameOver) Reset() {
	*x = GCGMsgGameOver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgGameOver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgGameOver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgGameOver) ProtoMessage() {}

func (x *GCGMsgGameOver) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgGameOver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgGameOver.ProtoReflect.Descriptor instead.
func (*GCGMsgGameOver) Descriptor() ([]byte, []int) {
	return file_GCGMsgGameOver_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgGameOver) GetEndReason() GCGEndReason {
	if x != nil {
		return x.EndReason
	}
	return GCGEndReason_GCG_END_REASON_DEFAULT
}

func (x *GCGMsgGameOver) GetWinControllerId() uint32 {
	if x != nil {
		return x.WinControllerId
	}
	return 0
}

var File_GCGMsgGameOver_proto protoreflect.FileDescriptor

var file_GCGMsgGameOver_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x47, 0x43, 0x47, 0x45, 0x6e, 0x64, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0e, 0x47, 0x43,
	0x47, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x0a,
	0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x47, 0x43, 0x47, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52,
	0x09, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x69,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x77, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGMsgGameOver_proto_rawDescOnce sync.Once
	file_GCGMsgGameOver_proto_rawDescData = file_GCGMsgGameOver_proto_rawDesc
)

func file_GCGMsgGameOver_proto_rawDescGZIP() []byte {
	file_GCGMsgGameOver_proto_rawDescOnce.Do(func() {
		file_GCGMsgGameOver_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgGameOver_proto_rawDescData)
	})
	return file_GCGMsgGameOver_proto_rawDescData
}

var file_GCGMsgGameOver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgGameOver_proto_goTypes = []interface{}{
	(*GCGMsgGameOver)(nil), // 0: GCGMsgGameOver
	(GCGEndReason)(0),      // 1: GCGEndReason
}
var file_GCGMsgGameOver_proto_depIdxs = []int32{
	1, // 0: GCGMsgGameOver.end_reason:type_name -> GCGEndReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgGameOver_proto_init() }
func file_GCGMsgGameOver_proto_init() {
	if File_GCGMsgGameOver_proto != nil {
		return
	}
	file_GCGEndReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgGameOver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgGameOver); i {
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
			RawDescriptor: file_GCGMsgGameOver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgGameOver_proto_goTypes,
		DependencyIndexes: file_GCGMsgGameOver_proto_depIdxs,
		MessageInfos:      file_GCGMsgGameOver_proto_msgTypes,
	}.Build()
	File_GCGMsgGameOver_proto = out.File
	file_GCGMsgGameOver_proto_rawDesc = nil
	file_GCGMsgGameOver_proto_goTypes = nil
	file_GCGMsgGameOver_proto_depIdxs = nil
}