// https://github.com/SlushinPS/beach-simulator
// Copyright (C) 2023 Slushy Team
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
// 	protoc        v3.21.5
// source: ChannellerSlabTakeoffBuffReq.proto

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

// CmdId: 8112
// Obf: KDBPPMGGOGJ
type ChannellerSlabTakeoffBuffReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuffId uint32 `protobuf:"varint,7,opt,name=buff_id,json=buffId,proto3" json:"buff_id,omitempty"`
	SlotId uint32 `protobuf:"varint,15,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	IsMp   bool   `protobuf:"varint,5,opt,name=is_mp,json=isMp,proto3" json:"is_mp,omitempty"`
}

func (x *ChannellerSlabTakeoffBuffReq) Reset() {
	*x = ChannellerSlabTakeoffBuffReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChannellerSlabTakeoffBuffReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannellerSlabTakeoffBuffReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannellerSlabTakeoffBuffReq) ProtoMessage() {}

func (x *ChannellerSlabTakeoffBuffReq) ProtoReflect() protoreflect.Message {
	mi := &file_ChannellerSlabTakeoffBuffReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannellerSlabTakeoffBuffReq.ProtoReflect.Descriptor instead.
func (*ChannellerSlabTakeoffBuffReq) Descriptor() ([]byte, []int) {
	return file_ChannellerSlabTakeoffBuffReq_proto_rawDescGZIP(), []int{0}
}

func (x *ChannellerSlabTakeoffBuffReq) GetBuffId() uint32 {
	if x != nil {
		return x.BuffId
	}
	return 0
}

func (x *ChannellerSlabTakeoffBuffReq) GetSlotId() uint32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *ChannellerSlabTakeoffBuffReq) GetIsMp() bool {
	if x != nil {
		return x.IsMp
	}
	return false
}

var File_ChannellerSlabTakeoffBuffReq_proto protoreflect.FileDescriptor

var file_ChannellerSlabTakeoffBuffReq_proto_rawDesc = []byte{
	0x0a, 0x22, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62,
	0x54, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x42, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x1c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x42, 0x75, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4d, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChannellerSlabTakeoffBuffReq_proto_rawDescOnce sync.Once
	file_ChannellerSlabTakeoffBuffReq_proto_rawDescData = file_ChannellerSlabTakeoffBuffReq_proto_rawDesc
)

func file_ChannellerSlabTakeoffBuffReq_proto_rawDescGZIP() []byte {
	file_ChannellerSlabTakeoffBuffReq_proto_rawDescOnce.Do(func() {
		file_ChannellerSlabTakeoffBuffReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChannellerSlabTakeoffBuffReq_proto_rawDescData)
	})
	return file_ChannellerSlabTakeoffBuffReq_proto_rawDescData
}

var file_ChannellerSlabTakeoffBuffReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChannellerSlabTakeoffBuffReq_proto_goTypes = []interface{}{
	(*ChannellerSlabTakeoffBuffReq)(nil), // 0: ChannellerSlabTakeoffBuffReq
}
var file_ChannellerSlabTakeoffBuffReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChannellerSlabTakeoffBuffReq_proto_init() }
func file_ChannellerSlabTakeoffBuffReq_proto_init() {
	if File_ChannellerSlabTakeoffBuffReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChannellerSlabTakeoffBuffReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannellerSlabTakeoffBuffReq); i {
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
			RawDescriptor: file_ChannellerSlabTakeoffBuffReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChannellerSlabTakeoffBuffReq_proto_goTypes,
		DependencyIndexes: file_ChannellerSlabTakeoffBuffReq_proto_depIdxs,
		MessageInfos:      file_ChannellerSlabTakeoffBuffReq_proto_msgTypes,
	}.Build()
	File_ChannellerSlabTakeoffBuffReq_proto = out.File
	file_ChannellerSlabTakeoffBuffReq_proto_rawDesc = nil
	file_ChannellerSlabTakeoffBuffReq_proto_goTypes = nil
	file_ChannellerSlabTakeoffBuffReq_proto_depIdxs = nil
}