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
// source: GalleryStartSource.proto

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

type GalleryStartSource int32

const (
	GalleryStartSource_GALLERY_START_SOURCE_BY_NONE  GalleryStartSource = 0
	GalleryStartSource_GALLERY_START_SOURCE_BY_MATCH GalleryStartSource = 1
	GalleryStartSource_GALLERY_START_SOURCE_BY_DRAFT GalleryStartSource = 2
)

// Enum value maps for GalleryStartSource.
var (
	GalleryStartSource_name = map[int32]string{
		0: "GALLERY_START_SOURCE_BY_NONE",
		1: "GALLERY_START_SOURCE_BY_MATCH",
		2: "GALLERY_START_SOURCE_BY_DRAFT",
	}
	GalleryStartSource_value = map[string]int32{
		"GALLERY_START_SOURCE_BY_NONE":  0,
		"GALLERY_START_SOURCE_BY_MATCH": 1,
		"GALLERY_START_SOURCE_BY_DRAFT": 2,
	}
)

func (x GalleryStartSource) Enum() *GalleryStartSource {
	p := new(GalleryStartSource)
	*p = x
	return p
}

func (x GalleryStartSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GalleryStartSource) Descriptor() protoreflect.EnumDescriptor {
	return file_GalleryStartSource_proto_enumTypes[0].Descriptor()
}

func (GalleryStartSource) Type() protoreflect.EnumType {
	return &file_GalleryStartSource_proto_enumTypes[0]
}

func (x GalleryStartSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GalleryStartSource.Descriptor instead.
func (GalleryStartSource) EnumDescriptor() ([]byte, []int) {
	return file_GalleryStartSource_proto_rawDescGZIP(), []int{0}
}

var File_GalleryStartSource_proto protoreflect.FileDescriptor

var file_GalleryStartSource_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x7c, 0x0a, 0x12, 0x47, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x1c, 0x47, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x42, 0x59, 0x5f, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x59,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x42, 0x59,
	0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GalleryStartSource_proto_rawDescOnce sync.Once
	file_GalleryStartSource_proto_rawDescData = file_GalleryStartSource_proto_rawDesc
)

func file_GalleryStartSource_proto_rawDescGZIP() []byte {
	file_GalleryStartSource_proto_rawDescOnce.Do(func() {
		file_GalleryStartSource_proto_rawDescData = protoimpl.X.CompressGZIP(file_GalleryStartSource_proto_rawDescData)
	})
	return file_GalleryStartSource_proto_rawDescData
}

var file_GalleryStartSource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GalleryStartSource_proto_goTypes = []interface{}{
	(GalleryStartSource)(0), // 0: GalleryStartSource
}
var file_GalleryStartSource_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GalleryStartSource_proto_init() }
func file_GalleryStartSource_proto_init() {
	if File_GalleryStartSource_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GalleryStartSource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GalleryStartSource_proto_goTypes,
		DependencyIndexes: file_GalleryStartSource_proto_depIdxs,
		EnumInfos:         file_GalleryStartSource_proto_enumTypes,
	}.Build()
	File_GalleryStartSource_proto = out.File
	file_GalleryStartSource_proto_rawDesc = nil
	file_GalleryStartSource_proto_goTypes = nil
	file_GalleryStartSource_proto_depIdxs = nil
}