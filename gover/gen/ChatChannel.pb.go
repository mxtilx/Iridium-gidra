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
// source: ChatChannel.proto

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

type ChatChannel int32

const (
	ChatChannel_CHAT_CHANNEL_TEAM             ChatChannel = 0
	ChatChannel_CHAT_CHANNEL_FRIEND           ChatChannel = 1
	ChatChannel_CHAT_CHANNEL_NONE             ChatChannel = 2
	ChatChannel_CHAT_CHANNEL_HIDEANDSEEK      ChatChannel = 3
	ChatChannel_CHAT_CHANNEL_HIDEANDSEEK_TEAM ChatChannel = 5
	ChatChannel_CHAT_CHANNEL_ALL              ChatChannel = 10
	ChatChannel_CHAT_CHANNEL_TEST             ChatChannel = 100
)

// Enum value maps for ChatChannel.
var (
	ChatChannel_name = map[int32]string{
		0:   "CHAT_CHANNEL_TEAM",
		1:   "CHAT_CHANNEL_FRIEND",
		2:   "CHAT_CHANNEL_NONE",
		3:   "CHAT_CHANNEL_HIDEANDSEEK",
		5:   "CHAT_CHANNEL_HIDEANDSEEK_TEAM",
		10:  "CHAT_CHANNEL_ALL",
		100: "CHAT_CHANNEL_TEST",
	}
	ChatChannel_value = map[string]int32{
		"CHAT_CHANNEL_TEAM":             0,
		"CHAT_CHANNEL_FRIEND":           1,
		"CHAT_CHANNEL_NONE":             2,
		"CHAT_CHANNEL_HIDEANDSEEK":      3,
		"CHAT_CHANNEL_HIDEANDSEEK_TEAM": 5,
		"CHAT_CHANNEL_ALL":              10,
		"CHAT_CHANNEL_TEST":             100,
	}
)

func (x ChatChannel) Enum() *ChatChannel {
	p := new(ChatChannel)
	*p = x
	return p
}

func (x ChatChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_ChatChannel_proto_enumTypes[0].Descriptor()
}

func (ChatChannel) Type() protoreflect.EnumType {
	return &file_ChatChannel_proto_enumTypes[0]
}

func (x ChatChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatChannel.Descriptor instead.
func (ChatChannel) EnumDescriptor() ([]byte, []int) {
	return file_ChatChannel_proto_rawDescGZIP(), []int{0}
}

var File_ChatChannel_proto protoreflect.FileDescriptor

var file_ChatChannel_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xc2, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48,
	0x41, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x48,
	0x41, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x48, 0x49, 0x44, 0x45, 0x41,
	0x4e, 0x44, 0x53, 0x45, 0x45, 0x4b, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x48, 0x41, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x48, 0x49, 0x44, 0x45, 0x41, 0x4e, 0x44,
	0x53, 0x45, 0x45, 0x4b, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x43,
	0x48, 0x41, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x41, 0x4c, 0x4c, 0x10,
	0x0a, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChatChannel_proto_rawDescOnce sync.Once
	file_ChatChannel_proto_rawDescData = file_ChatChannel_proto_rawDesc
)

func file_ChatChannel_proto_rawDescGZIP() []byte {
	file_ChatChannel_proto_rawDescOnce.Do(func() {
		file_ChatChannel_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChatChannel_proto_rawDescData)
	})
	return file_ChatChannel_proto_rawDescData
}

var file_ChatChannel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChatChannel_proto_goTypes = []interface{}{
	(ChatChannel)(0), // 0: ChatChannel
}
var file_ChatChannel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChatChannel_proto_init() }
func file_ChatChannel_proto_init() {
	if File_ChatChannel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChatChannel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChatChannel_proto_goTypes,
		DependencyIndexes: file_ChatChannel_proto_depIdxs,
		EnumInfos:         file_ChatChannel_proto_enumTypes,
	}.Build()
	File_ChatChannel_proto = out.File
	file_ChatChannel_proto_rawDesc = nil
	file_ChatChannel_proto_goTypes = nil
	file_ChatChannel_proto_depIdxs = nil
}