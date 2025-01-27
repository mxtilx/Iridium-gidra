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
// source: EchoShellDetailInfo.proto

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

// Obf: DIBPFDJEBMC
type EchoShellDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JCPFGDCAGFC           []uint32                 `protobuf:"varint,5,rep,packed,name=JCPFGDCAGFC,proto3" json:"JCPFGDCAGFC,omitempty"`
	LPEHAMKJFCA           []uint32                 `protobuf:"varint,13,rep,packed,name=LPEHAMKJFCA,proto3" json:"LPEHAMKJFCA,omitempty"`
	FKBCPCENBPP           []uint32                 `protobuf:"varint,2,rep,packed,name=FKBCPCENBPP,proto3" json:"FKBCPCENBPP,omitempty"`
	SummerTimeDungeonList []*SummerTimeDungeonInfo `protobuf:"bytes,9,rep,name=summer_time_dungeon_list,json=summerTimeDungeonList,proto3" json:"summer_time_dungeon_list,omitempty"`
}

func (x *EchoShellDetailInfo) Reset() {
	*x = EchoShellDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EchoShellDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoShellDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoShellDetailInfo) ProtoMessage() {}

func (x *EchoShellDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EchoShellDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoShellDetailInfo.ProtoReflect.Descriptor instead.
func (*EchoShellDetailInfo) Descriptor() ([]byte, []int) {
	return file_EchoShellDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoShellDetailInfo) GetJCPFGDCAGFC() []uint32 {
	if x != nil {
		return x.JCPFGDCAGFC
	}
	return nil
}

func (x *EchoShellDetailInfo) GetLPEHAMKJFCA() []uint32 {
	if x != nil {
		return x.LPEHAMKJFCA
	}
	return nil
}

func (x *EchoShellDetailInfo) GetFKBCPCENBPP() []uint32 {
	if x != nil {
		return x.FKBCPCENBPP
	}
	return nil
}

func (x *EchoShellDetailInfo) GetSummerTimeDungeonList() []*SummerTimeDungeonInfo {
	if x != nil {
		return x.SummerTimeDungeonList
	}
	return nil
}

var File_EchoShellDetailInfo_proto protoreflect.FileDescriptor

var file_EchoShellDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x53, 0x75, 0x6d,
	0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x13, 0x45, 0x63, 0x68,
	0x6f, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x43, 0x50, 0x46, 0x47, 0x44, 0x43, 0x41, 0x47, 0x46, 0x43, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x43, 0x50, 0x46, 0x47, 0x44, 0x43, 0x41, 0x47,
	0x46, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x50, 0x45, 0x48, 0x41, 0x4d, 0x4b, 0x4a, 0x46, 0x43,
	0x41, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x50, 0x45, 0x48, 0x41, 0x4d, 0x4b,
	0x4a, 0x46, 0x43, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4b, 0x42, 0x43, 0x50, 0x43, 0x45, 0x4e,
	0x42, 0x50, 0x50, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4b, 0x42, 0x43, 0x50,
	0x43, 0x45, 0x4e, 0x42, 0x50, 0x50, 0x12, 0x4f, 0x0a, 0x18, 0x73, 0x75, 0x6d, 0x6d, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x15, 0x73, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EchoShellDetailInfo_proto_rawDescOnce sync.Once
	file_EchoShellDetailInfo_proto_rawDescData = file_EchoShellDetailInfo_proto_rawDesc
)

func file_EchoShellDetailInfo_proto_rawDescGZIP() []byte {
	file_EchoShellDetailInfo_proto_rawDescOnce.Do(func() {
		file_EchoShellDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EchoShellDetailInfo_proto_rawDescData)
	})
	return file_EchoShellDetailInfo_proto_rawDescData
}

var file_EchoShellDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EchoShellDetailInfo_proto_goTypes = []interface{}{
	(*EchoShellDetailInfo)(nil),   // 0: EchoShellDetailInfo
	(*SummerTimeDungeonInfo)(nil), // 1: SummerTimeDungeonInfo
}
var file_EchoShellDetailInfo_proto_depIdxs = []int32{
	1, // 0: EchoShellDetailInfo.summer_time_dungeon_list:type_name -> SummerTimeDungeonInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EchoShellDetailInfo_proto_init() }
func file_EchoShellDetailInfo_proto_init() {
	if File_EchoShellDetailInfo_proto != nil {
		return
	}
	file_SummerTimeDungeonInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EchoShellDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoShellDetailInfo); i {
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
			RawDescriptor: file_EchoShellDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EchoShellDetailInfo_proto_goTypes,
		DependencyIndexes: file_EchoShellDetailInfo_proto_depIdxs,
		MessageInfos:      file_EchoShellDetailInfo_proto_msgTypes,
	}.Build()
	File_EchoShellDetailInfo_proto = out.File
	file_EchoShellDetailInfo_proto_rawDesc = nil
	file_EchoShellDetailInfo_proto_goTypes = nil
	file_EchoShellDetailInfo_proto_depIdxs = nil
}
