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
// source: Unk3000_ILLNKBDNGKP.proto

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

type Unk3000_ILLNKBDNGKP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_PHKHIPLDOOA []*Unk3000_HGBNOCJBDEK `protobuf:"bytes,5,rep,name=Unk2700_PHKHIPLDOOA,json=Unk2700PHKHIPLDOOA,proto3" json:"Unk2700_PHKHIPLDOOA,omitempty"`
	Unk3000_AIENCMLMCBE []*Unk3000_DCHMAMFIFOF `protobuf:"bytes,7,rep,name=Unk3000_AIENCMLMCBE,json=Unk3000AIENCMLMCBE,proto3" json:"Unk3000_AIENCMLMCBE,omitempty"`
}

func (x *Unk3000_ILLNKBDNGKP) Reset() {
	*x = Unk3000_ILLNKBDNGKP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_ILLNKBDNGKP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_ILLNKBDNGKP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_ILLNKBDNGKP) ProtoMessage() {}

func (x *Unk3000_ILLNKBDNGKP) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_ILLNKBDNGKP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_ILLNKBDNGKP.ProtoReflect.Descriptor instead.
func (*Unk3000_ILLNKBDNGKP) Descriptor() ([]byte, []int) {
	return file_Unk3000_ILLNKBDNGKP_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_ILLNKBDNGKP) GetUnk2700_PHKHIPLDOOA() []*Unk3000_HGBNOCJBDEK {
	if x != nil {
		return x.Unk2700_PHKHIPLDOOA
	}
	return nil
}

func (x *Unk3000_ILLNKBDNGKP) GetUnk3000_AIENCMLMCBE() []*Unk3000_DCHMAMFIFOF {
	if x != nil {
		return x.Unk3000_AIENCMLMCBE
	}
	return nil
}

var File_Unk3000_ILLNKBDNGKP_proto protoreflect.FileDescriptor

var file_Unk3000_ILLNKBDNGKP_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x49, 0x4c, 0x4c, 0x4e, 0x4b, 0x42,
	0x44, 0x4e, 0x47, 0x4b, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x44, 0x43, 0x48, 0x4d, 0x41, 0x4d, 0x46, 0x49, 0x46, 0x4f, 0x46,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f,
	0x48, 0x47, 0x42, 0x4e, 0x4f, 0x43, 0x4a, 0x42, 0x44, 0x45, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x49, 0x4c,
	0x4c, 0x4e, 0x4b, 0x42, 0x44, 0x4e, 0x47, 0x4b, 0x50, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x48, 0x4b, 0x48, 0x49, 0x50, 0x4c, 0x44, 0x4f, 0x4f, 0x41,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x5f, 0x48, 0x47, 0x42, 0x4e, 0x4f, 0x43, 0x4a, 0x42, 0x44, 0x45, 0x4b, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x50, 0x48, 0x4b, 0x48, 0x49, 0x50, 0x4c, 0x44, 0x4f, 0x4f, 0x41,
	0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x49, 0x45, 0x4e,
	0x43, 0x4d, 0x4c, 0x4d, 0x43, 0x42, 0x45, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x44, 0x43, 0x48, 0x4d, 0x41, 0x4d, 0x46, 0x49,
	0x46, 0x4f, 0x46, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x41, 0x49, 0x45, 0x4e,
	0x43, 0x4d, 0x4c, 0x4d, 0x43, 0x42, 0x45, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_ILLNKBDNGKP_proto_rawDescOnce sync.Once
	file_Unk3000_ILLNKBDNGKP_proto_rawDescData = file_Unk3000_ILLNKBDNGKP_proto_rawDesc
)

func file_Unk3000_ILLNKBDNGKP_proto_rawDescGZIP() []byte {
	file_Unk3000_ILLNKBDNGKP_proto_rawDescOnce.Do(func() {
		file_Unk3000_ILLNKBDNGKP_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_ILLNKBDNGKP_proto_rawDescData)
	})
	return file_Unk3000_ILLNKBDNGKP_proto_rawDescData
}

var file_Unk3000_ILLNKBDNGKP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3000_ILLNKBDNGKP_proto_goTypes = []interface{}{
	(*Unk3000_ILLNKBDNGKP)(nil), // 0: Unk3000_ILLNKBDNGKP
	(*Unk3000_HGBNOCJBDEK)(nil), // 1: Unk3000_HGBNOCJBDEK
	(*Unk3000_DCHMAMFIFOF)(nil), // 2: Unk3000_DCHMAMFIFOF
}
var file_Unk3000_ILLNKBDNGKP_proto_depIdxs = []int32{
	1, // 0: Unk3000_ILLNKBDNGKP.Unk2700_PHKHIPLDOOA:type_name -> Unk3000_HGBNOCJBDEK
	2, // 1: Unk3000_ILLNKBDNGKP.Unk3000_AIENCMLMCBE:type_name -> Unk3000_DCHMAMFIFOF
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk3000_ILLNKBDNGKP_proto_init() }
func file_Unk3000_ILLNKBDNGKP_proto_init() {
	if File_Unk3000_ILLNKBDNGKP_proto != nil {
		return
	}
	file_Unk3000_DCHMAMFIFOF_proto_init()
	file_Unk3000_HGBNOCJBDEK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_ILLNKBDNGKP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_ILLNKBDNGKP); i {
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
			RawDescriptor: file_Unk3000_ILLNKBDNGKP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_ILLNKBDNGKP_proto_goTypes,
		DependencyIndexes: file_Unk3000_ILLNKBDNGKP_proto_depIdxs,
		MessageInfos:      file_Unk3000_ILLNKBDNGKP_proto_msgTypes,
	}.Build()
	File_Unk3000_ILLNKBDNGKP_proto = out.File
	file_Unk3000_ILLNKBDNGKP_proto_rawDesc = nil
	file_Unk3000_ILLNKBDNGKP_proto_goTypes = nil
	file_Unk3000_ILLNKBDNGKP_proto_depIdxs = nil
}