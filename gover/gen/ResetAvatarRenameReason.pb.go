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
// source: ResetAvatarRenameReason.proto

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

type ResetAvatarRenameReason int32

const (
	ResetAvatarRenameReason_RESET_AVATAR_RENAME_REASON_NONE          ResetAvatarRenameReason = 0
	ResetAvatarRenameReason_RESET_AVATAR_RENAME_REASON_IP_BLACK_LIST ResetAvatarRenameReason = 1
	ResetAvatarRenameReason_RESET_AVATAR_RENAME_REASON_QUEST_AUDIT   ResetAvatarRenameReason = 2
	ResetAvatarRenameReason_RESET_AVATAR_RENAME_REASON_ACTIVE        ResetAvatarRenameReason = 3
)

// Enum value maps for ResetAvatarRenameReason.
var (
	ResetAvatarRenameReason_name = map[int32]string{
		0: "RESET_AVATAR_RENAME_REASON_NONE",
		1: "RESET_AVATAR_RENAME_REASON_IP_BLACK_LIST",
		2: "RESET_AVATAR_RENAME_REASON_QUEST_AUDIT",
		3: "RESET_AVATAR_RENAME_REASON_ACTIVE",
	}
	ResetAvatarRenameReason_value = map[string]int32{
		"RESET_AVATAR_RENAME_REASON_NONE":          0,
		"RESET_AVATAR_RENAME_REASON_IP_BLACK_LIST": 1,
		"RESET_AVATAR_RENAME_REASON_QUEST_AUDIT":   2,
		"RESET_AVATAR_RENAME_REASON_ACTIVE":        3,
	}
)

func (x ResetAvatarRenameReason) Enum() *ResetAvatarRenameReason {
	p := new(ResetAvatarRenameReason)
	*p = x
	return p
}

func (x ResetAvatarRenameReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResetAvatarRenameReason) Descriptor() protoreflect.EnumDescriptor {
	return file_ResetAvatarRenameReason_proto_enumTypes[0].Descriptor()
}

func (ResetAvatarRenameReason) Type() protoreflect.EnumType {
	return &file_ResetAvatarRenameReason_proto_enumTypes[0]
}

func (x ResetAvatarRenameReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResetAvatarRenameReason.Descriptor instead.
func (ResetAvatarRenameReason) EnumDescriptor() ([]byte, []int) {
	return file_ResetAvatarRenameReason_proto_rawDescGZIP(), []int{0}
}

var File_ResetAvatarRenameReason_proto protoreflect.FileDescriptor

var file_ResetAvatarRenameReason_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xbf, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x1f, 0x52,
	0x45, 0x53, 0x45, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x52, 0x45, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x2c, 0x0a, 0x28, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52,
	0x5f, 0x52, 0x45, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49,
	0x50, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x2a,
	0x0a, 0x26, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x52,
	0x45, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x54, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45,
	0x53, 0x45, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x52, 0x45, 0x4e, 0x41, 0x4d,
	0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ResetAvatarRenameReason_proto_rawDescOnce sync.Once
	file_ResetAvatarRenameReason_proto_rawDescData = file_ResetAvatarRenameReason_proto_rawDesc
)

func file_ResetAvatarRenameReason_proto_rawDescGZIP() []byte {
	file_ResetAvatarRenameReason_proto_rawDescOnce.Do(func() {
		file_ResetAvatarRenameReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_ResetAvatarRenameReason_proto_rawDescData)
	})
	return file_ResetAvatarRenameReason_proto_rawDescData
}

var file_ResetAvatarRenameReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ResetAvatarRenameReason_proto_goTypes = []interface{}{
	(ResetAvatarRenameReason)(0), // 0: ResetAvatarRenameReason
}
var file_ResetAvatarRenameReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ResetAvatarRenameReason_proto_init() }
func file_ResetAvatarRenameReason_proto_init() {
	if File_ResetAvatarRenameReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ResetAvatarRenameReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ResetAvatarRenameReason_proto_goTypes,
		DependencyIndexes: file_ResetAvatarRenameReason_proto_depIdxs,
		EnumInfos:         file_ResetAvatarRenameReason_proto_enumTypes,
	}.Build()
	File_ResetAvatarRenameReason_proto = out.File
	file_ResetAvatarRenameReason_proto_rawDesc = nil
	file_ResetAvatarRenameReason_proto_goTypes = nil
	file_ResetAvatarRenameReason_proto_depIdxs = nil
}