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
// source: OKJMFFNHFCA.proto

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

type OKJMFFNHFCA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AHJEEFHMHGN bool   `protobuf:"varint,2,opt,name=AHJEEFHMHGN,proto3" json:"AHJEEFHMHGN,omitempty"`
	AACHDBDODFG bool   `protobuf:"varint,15,opt,name=AACHDBDODFG,proto3" json:"AACHDBDODFG,omitempty"`
	EIPOBNOIHHD bool   `protobuf:"varint,1,opt,name=EIPOBNOIHHD,proto3" json:"EIPOBNOIHHD,omitempty"`
	LMKDFJMIHPJ uint32 `protobuf:"varint,13,opt,name=LMKDFJMIHPJ,proto3" json:"LMKDFJMIHPJ,omitempty"`
	GGJDMIHGAKA uint32 `protobuf:"varint,10,opt,name=GGJDMIHGAKA,proto3" json:"GGJDMIHGAKA,omitempty"`
}

func (x *OKJMFFNHFCA) Reset() {
	*x = OKJMFFNHFCA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OKJMFFNHFCA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OKJMFFNHFCA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OKJMFFNHFCA) ProtoMessage() {}

func (x *OKJMFFNHFCA) ProtoReflect() protoreflect.Message {
	mi := &file_OKJMFFNHFCA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OKJMFFNHFCA.ProtoReflect.Descriptor instead.
func (*OKJMFFNHFCA) Descriptor() ([]byte, []int) {
	return file_OKJMFFNHFCA_proto_rawDescGZIP(), []int{0}
}

func (x *OKJMFFNHFCA) GetAHJEEFHMHGN() bool {
	if x != nil {
		return x.AHJEEFHMHGN
	}
	return false
}

func (x *OKJMFFNHFCA) GetAACHDBDODFG() bool {
	if x != nil {
		return x.AACHDBDODFG
	}
	return false
}

func (x *OKJMFFNHFCA) GetEIPOBNOIHHD() bool {
	if x != nil {
		return x.EIPOBNOIHHD
	}
	return false
}

func (x *OKJMFFNHFCA) GetLMKDFJMIHPJ() uint32 {
	if x != nil {
		return x.LMKDFJMIHPJ
	}
	return 0
}

func (x *OKJMFFNHFCA) GetGGJDMIHGAKA() uint32 {
	if x != nil {
		return x.GGJDMIHGAKA
	}
	return 0
}

var File_OKJMFFNHFCA_proto protoreflect.FileDescriptor

var file_OKJMFFNHFCA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x4b, 0x4a, 0x4d, 0x46, 0x46, 0x4e, 0x48, 0x46, 0x43, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x4f, 0x4b, 0x4a, 0x4d, 0x46, 0x46, 0x4e, 0x48,
	0x46, 0x43, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x48, 0x4a, 0x45, 0x45, 0x46, 0x48, 0x4d, 0x48,
	0x47, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x48, 0x4a, 0x45, 0x45, 0x46,
	0x48, 0x4d, 0x48, 0x47, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x41, 0x43, 0x48, 0x44, 0x42, 0x44,
	0x4f, 0x44, 0x46, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x41, 0x43, 0x48,
	0x44, 0x42, 0x44, 0x4f, 0x44, 0x46, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x49, 0x50, 0x4f, 0x42,
	0x4e, 0x4f, 0x49, 0x48, 0x48, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x45, 0x49,
	0x50, 0x4f, 0x42, 0x4e, 0x4f, 0x49, 0x48, 0x48, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4d, 0x4b,
	0x44, 0x46, 0x4a, 0x4d, 0x49, 0x48, 0x50, 0x4a, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4c, 0x4d, 0x4b, 0x44, 0x46, 0x4a, 0x4d, 0x49, 0x48, 0x50, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x47,
	0x47, 0x4a, 0x44, 0x4d, 0x49, 0x48, 0x47, 0x41, 0x4b, 0x41, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x47, 0x47, 0x4a, 0x44, 0x4d, 0x49, 0x48, 0x47, 0x41, 0x4b, 0x41, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OKJMFFNHFCA_proto_rawDescOnce sync.Once
	file_OKJMFFNHFCA_proto_rawDescData = file_OKJMFFNHFCA_proto_rawDesc
)

func file_OKJMFFNHFCA_proto_rawDescGZIP() []byte {
	file_OKJMFFNHFCA_proto_rawDescOnce.Do(func() {
		file_OKJMFFNHFCA_proto_rawDescData = protoimpl.X.CompressGZIP(file_OKJMFFNHFCA_proto_rawDescData)
	})
	return file_OKJMFFNHFCA_proto_rawDescData
}

var file_OKJMFFNHFCA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OKJMFFNHFCA_proto_goTypes = []interface{}{
	(*OKJMFFNHFCA)(nil), // 0: OKJMFFNHFCA
}
var file_OKJMFFNHFCA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OKJMFFNHFCA_proto_init() }
func file_OKJMFFNHFCA_proto_init() {
	if File_OKJMFFNHFCA_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OKJMFFNHFCA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OKJMFFNHFCA); i {
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
			RawDescriptor: file_OKJMFFNHFCA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OKJMFFNHFCA_proto_goTypes,
		DependencyIndexes: file_OKJMFFNHFCA_proto_depIdxs,
		MessageInfos:      file_OKJMFFNHFCA_proto_msgTypes,
	}.Build()
	File_OKJMFFNHFCA_proto = out.File
	file_OKJMFFNHFCA_proto_rawDesc = nil
	file_OKJMFFNHFCA_proto_goTypes = nil
	file_OKJMFFNHFCA_proto_depIdxs = nil
}