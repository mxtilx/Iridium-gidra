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
// source: Unk2700_DCBEFDDECOJ.proto

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

// CmdId: 8858
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_DCBEFDDECOJ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_INIBKFPMCFO []*Unk2700_PKAPCOBGIJL `protobuf:"bytes,8,rep,name=Unk2700_INIBKFPMCFO,json=Unk2700INIBKFPMCFO,proto3" json:"Unk2700_INIBKFPMCFO,omitempty"`
	LevelId             uint32                 `protobuf:"varint,1,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	Unk2700_CBPNPEBMPOH bool                   `protobuf:"varint,15,opt,name=Unk2700_CBPNPEBMPOH,json=Unk2700CBPNPEBMPOH,proto3" json:"Unk2700_CBPNPEBMPOH,omitempty"`
	DifficultyId        uint32                 `protobuf:"varint,11,opt,name=difficulty_id,json=difficultyId,proto3" json:"difficulty_id,omitempty"`
	Unk2700_EONPKLLJHPH []*Unk2700_ADIGEBEIJBA `protobuf:"bytes,3,rep,name=Unk2700_EONPKLLJHPH,json=Unk2700EONPKLLJHPH,proto3" json:"Unk2700_EONPKLLJHPH,omitempty"`
	Unk2700_FIGHJIFINKI uint32                 `protobuf:"varint,7,opt,name=Unk2700_FIGHJIFINKI,json=Unk2700FIGHJIFINKI,proto3" json:"Unk2700_FIGHJIFINKI,omitempty"`
}

func (x *Unk2700_DCBEFDDECOJ) Reset() {
	*x = Unk2700_DCBEFDDECOJ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_DCBEFDDECOJ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_DCBEFDDECOJ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_DCBEFDDECOJ) ProtoMessage() {}

func (x *Unk2700_DCBEFDDECOJ) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_DCBEFDDECOJ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_DCBEFDDECOJ.ProtoReflect.Descriptor instead.
func (*Unk2700_DCBEFDDECOJ) Descriptor() ([]byte, []int) {
	return file_Unk2700_DCBEFDDECOJ_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_DCBEFDDECOJ) GetUnk2700_INIBKFPMCFO() []*Unk2700_PKAPCOBGIJL {
	if x != nil {
		return x.Unk2700_INIBKFPMCFO
	}
	return nil
}

func (x *Unk2700_DCBEFDDECOJ) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *Unk2700_DCBEFDDECOJ) GetUnk2700_CBPNPEBMPOH() bool {
	if x != nil {
		return x.Unk2700_CBPNPEBMPOH
	}
	return false
}

func (x *Unk2700_DCBEFDDECOJ) GetDifficultyId() uint32 {
	if x != nil {
		return x.DifficultyId
	}
	return 0
}

func (x *Unk2700_DCBEFDDECOJ) GetUnk2700_EONPKLLJHPH() []*Unk2700_ADIGEBEIJBA {
	if x != nil {
		return x.Unk2700_EONPKLLJHPH
	}
	return nil
}

func (x *Unk2700_DCBEFDDECOJ) GetUnk2700_FIGHJIFINKI() uint32 {
	if x != nil {
		return x.Unk2700_FIGHJIFINKI
	}
	return 0
}

var File_Unk2700_DCBEFDDECOJ_proto protoreflect.FileDescriptor

var file_Unk2700_DCBEFDDECOJ_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x43, 0x42, 0x45, 0x46, 0x44,
	0x44, 0x45, 0x43, 0x4f, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x44, 0x49, 0x47, 0x45, 0x42, 0x45, 0x49, 0x4a, 0x42, 0x41,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x50, 0x4b, 0x41, 0x50, 0x43, 0x4f, 0x42, 0x47, 0x49, 0x4a, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x43,
	0x42, 0x45, 0x46, 0x44, 0x44, 0x45, 0x43, 0x4f, 0x4a, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x49, 0x4e, 0x49, 0x42, 0x4b, 0x46, 0x50, 0x4d, 0x43, 0x46, 0x4f,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x50, 0x4b, 0x41, 0x50, 0x43, 0x4f, 0x42, 0x47, 0x49, 0x4a, 0x4c, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x4e, 0x49, 0x42, 0x4b, 0x46, 0x50, 0x4d, 0x43, 0x46, 0x4f,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x42, 0x50, 0x4e, 0x50, 0x45, 0x42, 0x4d, 0x50,
	0x4f, 0x48, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x43, 0x42, 0x50, 0x4e, 0x50, 0x45, 0x42, 0x4d, 0x50, 0x4f, 0x48, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x4f, 0x4e,
	0x50, 0x4b, 0x4c, 0x4c, 0x4a, 0x48, 0x50, 0x48, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x44, 0x49, 0x47, 0x45, 0x42, 0x45,
	0x49, 0x4a, 0x42, 0x41, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x45, 0x4f, 0x4e,
	0x50, 0x4b, 0x4c, 0x4c, 0x4a, 0x48, 0x50, 0x48, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x46, 0x49, 0x47, 0x48, 0x4a, 0x49, 0x46, 0x49, 0x4e, 0x4b, 0x49, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x46, 0x49,
	0x47, 0x48, 0x4a, 0x49, 0x46, 0x49, 0x4e, 0x4b, 0x49, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_DCBEFDDECOJ_proto_rawDescOnce sync.Once
	file_Unk2700_DCBEFDDECOJ_proto_rawDescData = file_Unk2700_DCBEFDDECOJ_proto_rawDesc
)

func file_Unk2700_DCBEFDDECOJ_proto_rawDescGZIP() []byte {
	file_Unk2700_DCBEFDDECOJ_proto_rawDescOnce.Do(func() {
		file_Unk2700_DCBEFDDECOJ_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_DCBEFDDECOJ_proto_rawDescData)
	})
	return file_Unk2700_DCBEFDDECOJ_proto_rawDescData
}

var file_Unk2700_DCBEFDDECOJ_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_DCBEFDDECOJ_proto_goTypes = []interface{}{
	(*Unk2700_DCBEFDDECOJ)(nil), // 0: Unk2700_DCBEFDDECOJ
	(*Unk2700_PKAPCOBGIJL)(nil), // 1: Unk2700_PKAPCOBGIJL
	(*Unk2700_ADIGEBEIJBA)(nil), // 2: Unk2700_ADIGEBEIJBA
}
var file_Unk2700_DCBEFDDECOJ_proto_depIdxs = []int32{
	1, // 0: Unk2700_DCBEFDDECOJ.Unk2700_INIBKFPMCFO:type_name -> Unk2700_PKAPCOBGIJL
	2, // 1: Unk2700_DCBEFDDECOJ.Unk2700_EONPKLLJHPH:type_name -> Unk2700_ADIGEBEIJBA
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk2700_DCBEFDDECOJ_proto_init() }
func file_Unk2700_DCBEFDDECOJ_proto_init() {
	if File_Unk2700_DCBEFDDECOJ_proto != nil {
		return
	}
	file_Unk2700_ADIGEBEIJBA_proto_init()
	file_Unk2700_PKAPCOBGIJL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_DCBEFDDECOJ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_DCBEFDDECOJ); i {
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
			RawDescriptor: file_Unk2700_DCBEFDDECOJ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_DCBEFDDECOJ_proto_goTypes,
		DependencyIndexes: file_Unk2700_DCBEFDDECOJ_proto_depIdxs,
		MessageInfos:      file_Unk2700_DCBEFDDECOJ_proto_msgTypes,
	}.Build()
	File_Unk2700_DCBEFDDECOJ_proto = out.File
	file_Unk2700_DCBEFDDECOJ_proto_rawDesc = nil
	file_Unk2700_DCBEFDDECOJ_proto_goTypes = nil
	file_Unk2700_DCBEFDDECOJ_proto_depIdxs = nil
}