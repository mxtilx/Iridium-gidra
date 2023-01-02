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
// source: EvtBeingHealedNotify.proto

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

type EvtBeingHealedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RealHealAmount float32 `protobuf:"fixed32,2,opt,name=real_heal_amount,json=realHealAmount,proto3" json:"real_heal_amount,omitempty"`
	HealAmount     float32 `protobuf:"fixed32,3,opt,name=heal_amount,json=healAmount,proto3" json:"heal_amount,omitempty"`
	TargetId       uint32  `protobuf:"varint,4,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	SourceId       uint32  `protobuf:"varint,6,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
}

func (x *EvtBeingHealedNotify) Reset() {
	*x = EvtBeingHealedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtBeingHealedNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtBeingHealedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtBeingHealedNotify) ProtoMessage() {}

func (x *EvtBeingHealedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtBeingHealedNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtBeingHealedNotify.ProtoReflect.Descriptor instead.
func (*EvtBeingHealedNotify) Descriptor() ([]byte, []int) {
	return file_EvtBeingHealedNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtBeingHealedNotify) GetRealHealAmount() float32 {
	if x != nil {
		return x.RealHealAmount
	}
	return 0
}

func (x *EvtBeingHealedNotify) GetHealAmount() float32 {
	if x != nil {
		return x.HealAmount
	}
	return 0
}

func (x *EvtBeingHealedNotify) GetTargetId() uint32 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *EvtBeingHealedNotify) GetSourceId() uint32 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

var File_EvtBeingHealedNotify_proto protoreflect.FileDescriptor

var file_EvtBeingHealedNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a,
	0x14, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0e, 0x72, 0x65, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtBeingHealedNotify_proto_rawDescOnce sync.Once
	file_EvtBeingHealedNotify_proto_rawDescData = file_EvtBeingHealedNotify_proto_rawDesc
)

func file_EvtBeingHealedNotify_proto_rawDescGZIP() []byte {
	file_EvtBeingHealedNotify_proto_rawDescOnce.Do(func() {
		file_EvtBeingHealedNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtBeingHealedNotify_proto_rawDescData)
	})
	return file_EvtBeingHealedNotify_proto_rawDescData
}

var file_EvtBeingHealedNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtBeingHealedNotify_proto_goTypes = []interface{}{
	(*EvtBeingHealedNotify)(nil), // 0: EvtBeingHealedNotify
}
var file_EvtBeingHealedNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EvtBeingHealedNotify_proto_init() }
func file_EvtBeingHealedNotify_proto_init() {
	if File_EvtBeingHealedNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EvtBeingHealedNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtBeingHealedNotify); i {
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
			RawDescriptor: file_EvtBeingHealedNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtBeingHealedNotify_proto_goTypes,
		DependencyIndexes: file_EvtBeingHealedNotify_proto_depIdxs,
		MessageInfos:      file_EvtBeingHealedNotify_proto_msgTypes,
	}.Build()
	File_EvtBeingHealedNotify_proto = out.File
	file_EvtBeingHealedNotify_proto_rawDesc = nil
	file_EvtBeingHealedNotify_proto_goTypes = nil
	file_EvtBeingHealedNotify_proto_depIdxs = nil
}