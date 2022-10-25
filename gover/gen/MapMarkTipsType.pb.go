// Sorapointa - A server software re-implementation for a certain anime game,
// and avoid sorapointa. Copyright (C) 2022  Sorapointa Team
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
// 	protoc        v3.21.8
// source: MapMarkTipsType.proto

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

type MapMarkTipsType int32

const (
	MapMarkTipsType_MAP_MARK_TIPS_TYPE_DUNGEON_ELEMENT_TRIAL MapMarkTipsType = 0
)

// Enum value maps for MapMarkTipsType.
var (
	MapMarkTipsType_name = map[int32]string{
		0: "MAP_MARK_TIPS_TYPE_DUNGEON_ELEMENT_TRIAL",
	}
	MapMarkTipsType_value = map[string]int32{
		"MAP_MARK_TIPS_TYPE_DUNGEON_ELEMENT_TRIAL": 0,
	}
)

func (x MapMarkTipsType) Enum() *MapMarkTipsType {
	p := new(MapMarkTipsType)
	*p = x
	return p
}

func (x MapMarkTipsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapMarkTipsType) Descriptor() protoreflect.EnumDescriptor {
	return file_MapMarkTipsType_proto_enumTypes[0].Descriptor()
}

func (MapMarkTipsType) Type() protoreflect.EnumType {
	return &file_MapMarkTipsType_proto_enumTypes[0]
}

func (x MapMarkTipsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapMarkTipsType.Descriptor instead.
func (MapMarkTipsType) EnumDescriptor() ([]byte, []int) {
	return file_MapMarkTipsType_proto_rawDescGZIP(), []int{0}
}

var File_MapMarkTipsType_proto protoreflect.FileDescriptor

var file_MapMarkTipsType_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x54, 0x69, 0x70, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x3f, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x4d, 0x61,
	0x72, 0x6b, 0x54, 0x69, 0x70, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x28, 0x4d, 0x41,
	0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x54, 0x49, 0x50, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MapMarkTipsType_proto_rawDescOnce sync.Once
	file_MapMarkTipsType_proto_rawDescData = file_MapMarkTipsType_proto_rawDesc
)

func file_MapMarkTipsType_proto_rawDescGZIP() []byte {
	file_MapMarkTipsType_proto_rawDescOnce.Do(func() {
		file_MapMarkTipsType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MapMarkTipsType_proto_rawDescData)
	})
	return file_MapMarkTipsType_proto_rawDescData
}

var file_MapMarkTipsType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MapMarkTipsType_proto_goTypes = []interface{}{
	(MapMarkTipsType)(0), // 0: MapMarkTipsType
}
var file_MapMarkTipsType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MapMarkTipsType_proto_init() }
func file_MapMarkTipsType_proto_init() {
	if File_MapMarkTipsType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MapMarkTipsType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MapMarkTipsType_proto_goTypes,
		DependencyIndexes: file_MapMarkTipsType_proto_depIdxs,
		EnumInfos:         file_MapMarkTipsType_proto_enumTypes,
	}.Build()
	File_MapMarkTipsType_proto = out.File
	file_MapMarkTipsType_proto_rawDesc = nil
	file_MapMarkTipsType_proto_goTypes = nil
	file_MapMarkTipsType_proto_depIdxs = nil
}