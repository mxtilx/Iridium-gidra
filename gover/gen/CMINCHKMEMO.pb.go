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
// source: CMINCHKMEMO.proto

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

type CMINCHKMEMO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DLFJNEEBLOG bool    `protobuf:"varint,8,opt,name=DLFJNEEBLOG,proto3" json:"DLFJNEEBLOG,omitempty"`
	JNEKCJANDGG float32 `protobuf:"fixed32,3,opt,name=JNEKCJANDGG,proto3" json:"JNEKCJANDGG,omitempty"`
	AvatarId    uint32  `protobuf:"varint,12,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	DENNPCBBNFD string  `protobuf:"bytes,13,opt,name=DENNPCBBNFD,proto3" json:"DENNPCBBNFD,omitempty"`
	ABJCKKIJDJH float32 `protobuf:"fixed32,4,opt,name=ABJCKKIJDJH,proto3" json:"ABJCKKIJDJH,omitempty"`
	HNGHMNBGHLG float32 `protobuf:"fixed32,6,opt,name=HNGHMNBGHLG,proto3" json:"HNGHMNBGHLG,omitempty"`
}

func (x *CMINCHKMEMO) Reset() {
	*x = CMINCHKMEMO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CMINCHKMEMO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMINCHKMEMO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMINCHKMEMO) ProtoMessage() {}

func (x *CMINCHKMEMO) ProtoReflect() protoreflect.Message {
	mi := &file_CMINCHKMEMO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMINCHKMEMO.ProtoReflect.Descriptor instead.
func (*CMINCHKMEMO) Descriptor() ([]byte, []int) {
	return file_CMINCHKMEMO_proto_rawDescGZIP(), []int{0}
}

func (x *CMINCHKMEMO) GetDLFJNEEBLOG() bool {
	if x != nil {
		return x.DLFJNEEBLOG
	}
	return false
}

func (x *CMINCHKMEMO) GetJNEKCJANDGG() float32 {
	if x != nil {
		return x.JNEKCJANDGG
	}
	return 0
}

func (x *CMINCHKMEMO) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *CMINCHKMEMO) GetDENNPCBBNFD() string {
	if x != nil {
		return x.DENNPCBBNFD
	}
	return ""
}

func (x *CMINCHKMEMO) GetABJCKKIJDJH() float32 {
	if x != nil {
		return x.ABJCKKIJDJH
	}
	return 0
}

func (x *CMINCHKMEMO) GetHNGHMNBGHLG() float32 {
	if x != nil {
		return x.HNGHMNBGHLG
	}
	return 0
}

var File_CMINCHKMEMO_proto protoreflect.FileDescriptor

var file_CMINCHKMEMO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x4d, 0x49, 0x4e, 0x43, 0x48, 0x4b, 0x4d, 0x45, 0x4d, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0b, 0x43, 0x4d, 0x49, 0x4e, 0x43, 0x48, 0x4b, 0x4d,
	0x45, 0x4d, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4c, 0x46, 0x4a, 0x4e, 0x45, 0x45, 0x42, 0x4c,
	0x4f, 0x47, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x4c, 0x46, 0x4a, 0x4e, 0x45,
	0x45, 0x42, 0x4c, 0x4f, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4e, 0x45, 0x4b, 0x43, 0x4a, 0x41,
	0x4e, 0x44, 0x47, 0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4a, 0x4e, 0x45, 0x4b,
	0x43, 0x4a, 0x41, 0x4e, 0x44, 0x47, 0x47, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x45, 0x4e, 0x4e, 0x50, 0x43, 0x42, 0x42,
	0x4e, 0x46, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x45, 0x4e, 0x4e, 0x50,
	0x43, 0x42, 0x42, 0x4e, 0x46, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x42, 0x4a, 0x43, 0x4b, 0x4b,
	0x49, 0x4a, 0x44, 0x4a, 0x48, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x41, 0x42, 0x4a,
	0x43, 0x4b, 0x4b, 0x49, 0x4a, 0x44, 0x4a, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4e, 0x47, 0x48,
	0x4d, 0x4e, 0x42, 0x47, 0x48, 0x4c, 0x47, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x48,
	0x4e, 0x47, 0x48, 0x4d, 0x4e, 0x42, 0x47, 0x48, 0x4c, 0x47, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CMINCHKMEMO_proto_rawDescOnce sync.Once
	file_CMINCHKMEMO_proto_rawDescData = file_CMINCHKMEMO_proto_rawDesc
)

func file_CMINCHKMEMO_proto_rawDescGZIP() []byte {
	file_CMINCHKMEMO_proto_rawDescOnce.Do(func() {
		file_CMINCHKMEMO_proto_rawDescData = protoimpl.X.CompressGZIP(file_CMINCHKMEMO_proto_rawDescData)
	})
	return file_CMINCHKMEMO_proto_rawDescData
}

var file_CMINCHKMEMO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CMINCHKMEMO_proto_goTypes = []interface{}{
	(*CMINCHKMEMO)(nil), // 0: CMINCHKMEMO
}
var file_CMINCHKMEMO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CMINCHKMEMO_proto_init() }
func file_CMINCHKMEMO_proto_init() {
	if File_CMINCHKMEMO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CMINCHKMEMO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMINCHKMEMO); i {
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
			RawDescriptor: file_CMINCHKMEMO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CMINCHKMEMO_proto_goTypes,
		DependencyIndexes: file_CMINCHKMEMO_proto_depIdxs,
		MessageInfos:      file_CMINCHKMEMO_proto_msgTypes,
	}.Build()
	File_CMINCHKMEMO_proto = out.File
	file_CMINCHKMEMO_proto_rawDesc = nil
	file_CMINCHKMEMO_proto_goTypes = nil
	file_CMINCHKMEMO_proto_depIdxs = nil
}