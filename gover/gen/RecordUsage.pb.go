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
// source: RecordUsage.proto

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

type RecordUsage int32

const (
	RecordUsage_RECORD_USAGE_UGC_RECORD_USAGE_NONE    RecordUsage = 0
	RecordUsage_RECORD_USAGE_UGC_RECORD_USAGE_IMPORT  RecordUsage = 1
	RecordUsage_RECORD_USAGE_UGC_RECORD_USAGE_PLAY    RecordUsage = 2
	RecordUsage_RECORD_USAGE_UGC_RECORD_USAGE_TRIAL   RecordUsage = 3
	RecordUsage_RECORD_USAGE_UGC_RECORD_USAGE_COMPARE RecordUsage = 4
)

// Enum value maps for RecordUsage.
var (
	RecordUsage_name = map[int32]string{
		0: "RECORD_USAGE_UGC_RECORD_USAGE_NONE",
		1: "RECORD_USAGE_UGC_RECORD_USAGE_IMPORT",
		2: "RECORD_USAGE_UGC_RECORD_USAGE_PLAY",
		3: "RECORD_USAGE_UGC_RECORD_USAGE_TRIAL",
		4: "RECORD_USAGE_UGC_RECORD_USAGE_COMPARE",
	}
	RecordUsage_value = map[string]int32{
		"RECORD_USAGE_UGC_RECORD_USAGE_NONE":    0,
		"RECORD_USAGE_UGC_RECORD_USAGE_IMPORT":  1,
		"RECORD_USAGE_UGC_RECORD_USAGE_PLAY":    2,
		"RECORD_USAGE_UGC_RECORD_USAGE_TRIAL":   3,
		"RECORD_USAGE_UGC_RECORD_USAGE_COMPARE": 4,
	}
)

func (x RecordUsage) Enum() *RecordUsage {
	p := new(RecordUsage)
	*p = x
	return p
}

func (x RecordUsage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordUsage) Descriptor() protoreflect.EnumDescriptor {
	return file_RecordUsage_proto_enumTypes[0].Descriptor()
}

func (RecordUsage) Type() protoreflect.EnumType {
	return &file_RecordUsage_proto_enumTypes[0]
}

func (x RecordUsage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordUsage.Descriptor instead.
func (RecordUsage) EnumDescriptor() ([]byte, []int) {
	return file_RecordUsage_proto_rawDescGZIP(), []int{0}
}

var File_RecordUsage_proto protoreflect.FileDescriptor

var file_RecordUsage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xdb, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x55, 0x47, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x52,
	0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x47, 0x43, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x4d, 0x50,
	0x4f, 0x52, 0x54, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x47, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x02, 0x12, 0x27, 0x0a,
	0x23, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x47,
	0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x52, 0x49, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x47, 0x43, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52,
	0x44, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x10,
	0x04, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RecordUsage_proto_rawDescOnce sync.Once
	file_RecordUsage_proto_rawDescData = file_RecordUsage_proto_rawDesc
)

func file_RecordUsage_proto_rawDescGZIP() []byte {
	file_RecordUsage_proto_rawDescOnce.Do(func() {
		file_RecordUsage_proto_rawDescData = protoimpl.X.CompressGZIP(file_RecordUsage_proto_rawDescData)
	})
	return file_RecordUsage_proto_rawDescData
}

var file_RecordUsage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RecordUsage_proto_goTypes = []interface{}{
	(RecordUsage)(0), // 0: RecordUsage
}
var file_RecordUsage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RecordUsage_proto_init() }
func file_RecordUsage_proto_init() {
	if File_RecordUsage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RecordUsage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RecordUsage_proto_goTypes,
		DependencyIndexes: file_RecordUsage_proto_depIdxs,
		EnumInfos:         file_RecordUsage_proto_enumTypes,
	}.Build()
	File_RecordUsage_proto = out.File
	file_RecordUsage_proto_rawDesc = nil
	file_RecordUsage_proto_goTypes = nil
	file_RecordUsage_proto_depIdxs = nil
}