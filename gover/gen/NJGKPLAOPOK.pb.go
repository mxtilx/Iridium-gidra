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
// source: NJGKPLAOPOK.proto

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

// CmdId: 9210
type NJGKPLAOPOK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ALIKJHIECLN []*ANAEEOOIDJJ `protobuf:"bytes,11,rep,name=ALIKJHIECLN,proto3" json:"ALIKJHIECLN,omitempty"`
	JNIHNFJHDGD uint32         `protobuf:"varint,15,opt,name=JNIHNFJHDGD,proto3" json:"JNIHNFJHDGD,omitempty"`
	LNCDJCIKOGC *LLOOADCMMPP   `protobuf:"bytes,4,opt,name=LNCDJCIKOGC,proto3" json:"LNCDJCIKOGC,omitempty"`
	Retcode     int32          `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GroupId     uint32         `protobuf:"varint,12,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *NJGKPLAOPOK) Reset() {
	*x = NJGKPLAOPOK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NJGKPLAOPOK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NJGKPLAOPOK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NJGKPLAOPOK) ProtoMessage() {}

func (x *NJGKPLAOPOK) ProtoReflect() protoreflect.Message {
	mi := &file_NJGKPLAOPOK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NJGKPLAOPOK.ProtoReflect.Descriptor instead.
func (*NJGKPLAOPOK) Descriptor() ([]byte, []int) {
	return file_NJGKPLAOPOK_proto_rawDescGZIP(), []int{0}
}

func (x *NJGKPLAOPOK) GetALIKJHIECLN() []*ANAEEOOIDJJ {
	if x != nil {
		return x.ALIKJHIECLN
	}
	return nil
}

func (x *NJGKPLAOPOK) GetJNIHNFJHDGD() uint32 {
	if x != nil {
		return x.JNIHNFJHDGD
	}
	return 0
}

func (x *NJGKPLAOPOK) GetLNCDJCIKOGC() *LLOOADCMMPP {
	if x != nil {
		return x.LNCDJCIKOGC
	}
	return nil
}

func (x *NJGKPLAOPOK) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *NJGKPLAOPOK) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_NJGKPLAOPOK_proto protoreflect.FileDescriptor

var file_NJGKPLAOPOK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x4a, 0x47, 0x4b, 0x50, 0x4c, 0x41, 0x4f, 0x50, 0x4f, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4e, 0x41, 0x45, 0x45, 0x4f, 0x4f, 0x49, 0x44, 0x4a, 0x4a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x4c, 0x4f, 0x4f, 0x41, 0x44, 0x43, 0x4d,
	0x4d, 0x50, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x4e, 0x4a,
	0x47, 0x4b, 0x50, 0x4c, 0x41, 0x4f, 0x50, 0x4f, 0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x41, 0x4c, 0x49,
	0x4b, 0x4a, 0x48, 0x49, 0x45, 0x43, 0x4c, 0x4e, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x41, 0x4e, 0x41, 0x45, 0x45, 0x4f, 0x4f, 0x49, 0x44, 0x4a, 0x4a, 0x52, 0x0b, 0x41, 0x4c,
	0x49, 0x4b, 0x4a, 0x48, 0x49, 0x45, 0x43, 0x4c, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4e, 0x49,
	0x48, 0x4e, 0x46, 0x4a, 0x48, 0x44, 0x47, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4a, 0x4e, 0x49, 0x48, 0x4e, 0x46, 0x4a, 0x48, 0x44, 0x47, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x4c,
	0x4e, 0x43, 0x44, 0x4a, 0x43, 0x49, 0x4b, 0x4f, 0x47, 0x43, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4c, 0x4c, 0x4f, 0x4f, 0x41, 0x44, 0x43, 0x4d, 0x4d, 0x50, 0x50, 0x52, 0x0b,
	0x4c, 0x4e, 0x43, 0x44, 0x4a, 0x43, 0x49, 0x4b, 0x4f, 0x47, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NJGKPLAOPOK_proto_rawDescOnce sync.Once
	file_NJGKPLAOPOK_proto_rawDescData = file_NJGKPLAOPOK_proto_rawDesc
)

func file_NJGKPLAOPOK_proto_rawDescGZIP() []byte {
	file_NJGKPLAOPOK_proto_rawDescOnce.Do(func() {
		file_NJGKPLAOPOK_proto_rawDescData = protoimpl.X.CompressGZIP(file_NJGKPLAOPOK_proto_rawDescData)
	})
	return file_NJGKPLAOPOK_proto_rawDescData
}

var file_NJGKPLAOPOK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NJGKPLAOPOK_proto_goTypes = []interface{}{
	(*NJGKPLAOPOK)(nil), // 0: NJGKPLAOPOK
	(*ANAEEOOIDJJ)(nil), // 1: ANAEEOOIDJJ
	(*LLOOADCMMPP)(nil), // 2: LLOOADCMMPP
}
var file_NJGKPLAOPOK_proto_depIdxs = []int32{
	1, // 0: NJGKPLAOPOK.ALIKJHIECLN:type_name -> ANAEEOOIDJJ
	2, // 1: NJGKPLAOPOK.LNCDJCIKOGC:type_name -> LLOOADCMMPP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_NJGKPLAOPOK_proto_init() }
func file_NJGKPLAOPOK_proto_init() {
	if File_NJGKPLAOPOK_proto != nil {
		return
	}
	file_ANAEEOOIDJJ_proto_init()
	file_LLOOADCMMPP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_NJGKPLAOPOK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NJGKPLAOPOK); i {
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
			RawDescriptor: file_NJGKPLAOPOK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NJGKPLAOPOK_proto_goTypes,
		DependencyIndexes: file_NJGKPLAOPOK_proto_depIdxs,
		MessageInfos:      file_NJGKPLAOPOK_proto_msgTypes,
	}.Build()
	File_NJGKPLAOPOK_proto = out.File
	file_NJGKPLAOPOK_proto_rawDesc = nil
	file_NJGKPLAOPOK_proto_goTypes = nil
	file_NJGKPLAOPOK_proto_depIdxs = nil
}
