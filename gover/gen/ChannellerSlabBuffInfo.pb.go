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
// source: ChannellerSlabBuffInfo.proto

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

// Obf: CAGONPCNIDL
type ChannellerSlabBuffInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AGGNKCMFEJK    *ChannellerSlabBuffSchemeInfo `protobuf:"bytes,6,opt,name=AGGNKCMFEJK,proto3" json:"AGGNKCMFEJK,omitempty"`
	AssistInfoList []*ChannellerSlabAssistInfo   `protobuf:"bytes,5,rep,name=assist_info_list,json=assistInfoList,proto3" json:"assist_info_list,omitempty"`
	AHBMAAMGCND    *ChannellerSlabBuffSchemeInfo `protobuf:"bytes,2,opt,name=AHBMAAMGCND,proto3" json:"AHBMAAMGCND,omitempty"`
	BuffIdList     []uint32                      `protobuf:"varint,13,rep,packed,name=buff_id_list,json=buffIdList,proto3" json:"buff_id_list,omitempty"`
}

func (x *ChannellerSlabBuffInfo) Reset() {
	*x = ChannellerSlabBuffInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChannellerSlabBuffInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannellerSlabBuffInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannellerSlabBuffInfo) ProtoMessage() {}

func (x *ChannellerSlabBuffInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChannellerSlabBuffInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannellerSlabBuffInfo.ProtoReflect.Descriptor instead.
func (*ChannellerSlabBuffInfo) Descriptor() ([]byte, []int) {
	return file_ChannellerSlabBuffInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChannellerSlabBuffInfo) GetAGGNKCMFEJK() *ChannellerSlabBuffSchemeInfo {
	if x != nil {
		return x.AGGNKCMFEJK
	}
	return nil
}

func (x *ChannellerSlabBuffInfo) GetAssistInfoList() []*ChannellerSlabAssistInfo {
	if x != nil {
		return x.AssistInfoList
	}
	return nil
}

func (x *ChannellerSlabBuffInfo) GetAHBMAAMGCND() *ChannellerSlabBuffSchemeInfo {
	if x != nil {
		return x.AHBMAAMGCND
	}
	return nil
}

func (x *ChannellerSlabBuffInfo) GetBuffIdList() []uint32 {
	if x != nil {
		return x.BuffIdList
	}
	return nil
}

var File_ChannellerSlabBuffInfo_proto protoreflect.FileDescriptor

var file_ChannellerSlabBuffInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62,
	0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x42, 0x75,
	0x66, 0x66, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c,
	0x61, 0x62, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x53, 0x6c, 0x61, 0x62, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a,
	0x0b, 0x41, 0x47, 0x47, 0x4e, 0x4b, 0x43, 0x4d, 0x46, 0x45, 0x4a, 0x4b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53,
	0x6c, 0x61, 0x62, 0x42, 0x75, 0x66, 0x66, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x41, 0x47, 0x47, 0x4e, 0x4b, 0x43, 0x4d, 0x46, 0x45, 0x4a, 0x4b, 0x12, 0x43,
	0x0a, 0x10, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x48, 0x42, 0x4d, 0x41, 0x41, 0x4d, 0x47, 0x43,
	0x4e, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x42, 0x75, 0x66, 0x66, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x41, 0x48, 0x42, 0x4d, 0x41, 0x41, 0x4d,
	0x47, 0x43, 0x4e, 0x44, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x75, 0x66, 0x66,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChannellerSlabBuffInfo_proto_rawDescOnce sync.Once
	file_ChannellerSlabBuffInfo_proto_rawDescData = file_ChannellerSlabBuffInfo_proto_rawDesc
)

func file_ChannellerSlabBuffInfo_proto_rawDescGZIP() []byte {
	file_ChannellerSlabBuffInfo_proto_rawDescOnce.Do(func() {
		file_ChannellerSlabBuffInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChannellerSlabBuffInfo_proto_rawDescData)
	})
	return file_ChannellerSlabBuffInfo_proto_rawDescData
}

var file_ChannellerSlabBuffInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChannellerSlabBuffInfo_proto_goTypes = []interface{}{
	(*ChannellerSlabBuffInfo)(nil),       // 0: ChannellerSlabBuffInfo
	(*ChannellerSlabBuffSchemeInfo)(nil), // 1: ChannellerSlabBuffSchemeInfo
	(*ChannellerSlabAssistInfo)(nil),     // 2: ChannellerSlabAssistInfo
}
var file_ChannellerSlabBuffInfo_proto_depIdxs = []int32{
	1, // 0: ChannellerSlabBuffInfo.AGGNKCMFEJK:type_name -> ChannellerSlabBuffSchemeInfo
	2, // 1: ChannellerSlabBuffInfo.assist_info_list:type_name -> ChannellerSlabAssistInfo
	1, // 2: ChannellerSlabBuffInfo.AHBMAAMGCND:type_name -> ChannellerSlabBuffSchemeInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ChannellerSlabBuffInfo_proto_init() }
func file_ChannellerSlabBuffInfo_proto_init() {
	if File_ChannellerSlabBuffInfo_proto != nil {
		return
	}
	file_ChannellerSlabBuffSchemeInfo_proto_init()
	file_ChannellerSlabAssistInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChannellerSlabBuffInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannellerSlabBuffInfo); i {
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
			RawDescriptor: file_ChannellerSlabBuffInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChannellerSlabBuffInfo_proto_goTypes,
		DependencyIndexes: file_ChannellerSlabBuffInfo_proto_depIdxs,
		MessageInfos:      file_ChannellerSlabBuffInfo_proto_msgTypes,
	}.Build()
	File_ChannellerSlabBuffInfo_proto = out.File
	file_ChannellerSlabBuffInfo_proto_rawDesc = nil
	file_ChannellerSlabBuffInfo_proto_goTypes = nil
	file_ChannellerSlabBuffInfo_proto_depIdxs = nil
}