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
// source: GCGLevelType.proto

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

type GCGLevelType int32

const (
	GCGLevelType_GCG_LEVEL_TYPE_NONE        GCGLevelType = 0
	GCGLevelType_GCG_LEVEL_TYPE_CONST       GCGLevelType = 1
	GCGLevelType_GCG_LEVEL_TYPE_WEEK        GCGLevelType = 2
	GCGLevelType_GCG_LEVEL_TYPE_WORLD       GCGLevelType = 3
	GCGLevelType_GCG_LEVEL_TYPE_BOSS        GCGLevelType = 4
	GCGLevelType_GCG_LEVEL_TYPE_CHARACTER   GCGLevelType = 5
	GCGLevelType_GCG_LEVEL_TYPE_BREAK       GCGLevelType = 6
	GCGLevelType_GCG_LEVEL_TYPE_QUEST       GCGLevelType = 7
	GCGLevelType_GCG_LEVEL_TYPE_GUIDE_GROUP GCGLevelType = 8
)

// Enum value maps for GCGLevelType.
var (
	GCGLevelType_name = map[int32]string{
		0: "GCG_LEVEL_TYPE_NONE",
		1: "GCG_LEVEL_TYPE_CONST",
		2: "GCG_LEVEL_TYPE_WEEK",
		3: "GCG_LEVEL_TYPE_WORLD",
		4: "GCG_LEVEL_TYPE_BOSS",
		5: "GCG_LEVEL_TYPE_CHARACTER",
		6: "GCG_LEVEL_TYPE_BREAK",
		7: "GCG_LEVEL_TYPE_QUEST",
		8: "GCG_LEVEL_TYPE_GUIDE_GROUP",
	}
	GCGLevelType_value = map[string]int32{
		"GCG_LEVEL_TYPE_NONE":        0,
		"GCG_LEVEL_TYPE_CONST":       1,
		"GCG_LEVEL_TYPE_WEEK":        2,
		"GCG_LEVEL_TYPE_WORLD":       3,
		"GCG_LEVEL_TYPE_BOSS":        4,
		"GCG_LEVEL_TYPE_CHARACTER":   5,
		"GCG_LEVEL_TYPE_BREAK":       6,
		"GCG_LEVEL_TYPE_QUEST":       7,
		"GCG_LEVEL_TYPE_GUIDE_GROUP": 8,
	}
)

func (x GCGLevelType) Enum() *GCGLevelType {
	p := new(GCGLevelType)
	*p = x
	return p
}

func (x GCGLevelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GCGLevelType) Descriptor() protoreflect.EnumDescriptor {
	return file_GCGLevelType_proto_enumTypes[0].Descriptor()
}

func (GCGLevelType) Type() protoreflect.EnumType {
	return &file_GCGLevelType_proto_enumTypes[0]
}

func (x GCGLevelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GCGLevelType.Descriptor instead.
func (GCGLevelType) EnumDescriptor() ([]byte, []int) {
	return file_GCGLevelType_proto_rawDescGZIP(), []int{0}
}

var File_GCGLevelType_proto protoreflect.FileDescriptor

var file_GCGLevelType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x47, 0x43, 0x47, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xff, 0x01, 0x0a, 0x0c, 0x47, 0x43, 0x47, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x43, 0x47, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18,
	0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x43, 0x47, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x47,
	0x43, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f,
	0x53, 0x53, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x43, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x41, 0x43, 0x54, 0x45, 0x52,
	0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14,
	0x47, 0x43, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x43, 0x47, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x55, 0x49, 0x44, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x10, 0x08, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGLevelType_proto_rawDescOnce sync.Once
	file_GCGLevelType_proto_rawDescData = file_GCGLevelType_proto_rawDesc
)

func file_GCGLevelType_proto_rawDescGZIP() []byte {
	file_GCGLevelType_proto_rawDescOnce.Do(func() {
		file_GCGLevelType_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGLevelType_proto_rawDescData)
	})
	return file_GCGLevelType_proto_rawDescData
}

var file_GCGLevelType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GCGLevelType_proto_goTypes = []interface{}{
	(GCGLevelType)(0), // 0: GCGLevelType
}
var file_GCGLevelType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGLevelType_proto_init() }
func file_GCGLevelType_proto_init() {
	if File_GCGLevelType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGLevelType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGLevelType_proto_goTypes,
		DependencyIndexes: file_GCGLevelType_proto_depIdxs,
		EnumInfos:         file_GCGLevelType_proto_enumTypes,
	}.Build()
	File_GCGLevelType_proto = out.File
	file_GCGLevelType_proto_rawDesc = nil
	file_GCGLevelType_proto_goTypes = nil
	file_GCGLevelType_proto_depIdxs = nil
}
