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
// source: MPDJNGPJIKJ.proto

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

// CmdId: 22599
type MPDJNGPJIKJ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AffixList   []uint32 `protobuf:"varint,5,rep,packed,name=affix_list,json=affixList,proto3" json:"affix_list,omitempty"`
	OPGBGHPHELB uint32   `protobuf:"varint,6,opt,name=OPGBGHPHELB,proto3" json:"OPGBGHPHELB,omitempty"`
	LevelId     uint32   `protobuf:"varint,12,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
}

func (x *MPDJNGPJIKJ) Reset() {
	*x = MPDJNGPJIKJ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MPDJNGPJIKJ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MPDJNGPJIKJ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MPDJNGPJIKJ) ProtoMessage() {}

func (x *MPDJNGPJIKJ) ProtoReflect() protoreflect.Message {
	mi := &file_MPDJNGPJIKJ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MPDJNGPJIKJ.ProtoReflect.Descriptor instead.
func (*MPDJNGPJIKJ) Descriptor() ([]byte, []int) {
	return file_MPDJNGPJIKJ_proto_rawDescGZIP(), []int{0}
}

func (x *MPDJNGPJIKJ) GetAffixList() []uint32 {
	if x != nil {
		return x.AffixList
	}
	return nil
}

func (x *MPDJNGPJIKJ) GetOPGBGHPHELB() uint32 {
	if x != nil {
		return x.OPGBGHPHELB
	}
	return 0
}

func (x *MPDJNGPJIKJ) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

var File_MPDJNGPJIKJ_proto protoreflect.FileDescriptor

var file_MPDJNGPJIKJ_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x50, 0x44, 0x4a, 0x4e, 0x47, 0x50, 0x4a, 0x49, 0x4b, 0x4a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0b, 0x4d, 0x50, 0x44, 0x4a, 0x4e, 0x47, 0x50, 0x4a, 0x49,
	0x4b, 0x4a, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x66, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x50, 0x47, 0x42, 0x47, 0x48, 0x50, 0x48, 0x45, 0x4c, 0x42,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x50, 0x47, 0x42, 0x47, 0x48, 0x50, 0x48,
	0x45, 0x4c, 0x42, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MPDJNGPJIKJ_proto_rawDescOnce sync.Once
	file_MPDJNGPJIKJ_proto_rawDescData = file_MPDJNGPJIKJ_proto_rawDesc
)

func file_MPDJNGPJIKJ_proto_rawDescGZIP() []byte {
	file_MPDJNGPJIKJ_proto_rawDescOnce.Do(func() {
		file_MPDJNGPJIKJ_proto_rawDescData = protoimpl.X.CompressGZIP(file_MPDJNGPJIKJ_proto_rawDescData)
	})
	return file_MPDJNGPJIKJ_proto_rawDescData
}

var file_MPDJNGPJIKJ_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MPDJNGPJIKJ_proto_goTypes = []interface{}{
	(*MPDJNGPJIKJ)(nil), // 0: MPDJNGPJIKJ
}
var file_MPDJNGPJIKJ_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MPDJNGPJIKJ_proto_init() }
func file_MPDJNGPJIKJ_proto_init() {
	if File_MPDJNGPJIKJ_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MPDJNGPJIKJ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MPDJNGPJIKJ); i {
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
			RawDescriptor: file_MPDJNGPJIKJ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MPDJNGPJIKJ_proto_goTypes,
		DependencyIndexes: file_MPDJNGPJIKJ_proto_depIdxs,
		MessageInfos:      file_MPDJNGPJIKJ_proto_msgTypes,
	}.Build()
	File_MPDJNGPJIKJ_proto = out.File
	file_MPDJNGPJIKJ_proto_rawDesc = nil
	file_MPDJNGPJIKJ_proto_goTypes = nil
	file_MPDJNGPJIKJ_proto_depIdxs = nil
}