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
// source: Unk3000_LBNFMLHLBIH.proto

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

type Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN int32

const (
	Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN_OBSTACLE_SHAPE_CAPSULE Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN = 0
	Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN_OBSTACLE_SHAPE_BOX     Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN = 1
)

// Enum value maps for Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN.
var (
	Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN_name = map[int32]string{
		0: "Unk3000_GPHBIBGMHJN_OBSTACLE_SHAPE_CAPSULE",
		1: "Unk3000_GPHBIBGMHJN_OBSTACLE_SHAPE_BOX",
	}
	Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN_value = map[string]int32{
		"Unk3000_GPHBIBGMHJN_OBSTACLE_SHAPE_CAPSULE": 0,
		"Unk3000_GPHBIBGMHJN_OBSTACLE_SHAPE_BOX":     1,
	}
)

func (x Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN) Enum() *Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN {
	p := new(Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN)
	*p = x
	return p
}

func (x Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN) Descriptor() protoreflect.EnumDescriptor {
	return file_Unk3000_LBNFMLHLBIH_proto_enumTypes[0].Descriptor()
}

func (Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN) Type() protoreflect.EnumType {
	return &file_Unk3000_LBNFMLHLBIH_proto_enumTypes[0]
}

func (x Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN.Descriptor instead.
func (Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN) EnumDescriptor() ([]byte, []int) {
	return file_Unk3000_LBNFMLHLBIH_proto_rawDescGZIP(), []int{0, 0}
}

type Unk3000_LBNFMLHLBIH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN `protobuf:"varint,2,opt,name=type,proto3,enum=Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN" json:"type,omitempty"`
	Unk3000_MFHLAJACMFA int32                                   `protobuf:"varint,11,opt,name=Unk3000_MFHLAJACMFA,json=Unk3000MFHLAJACMFA,proto3" json:"Unk3000_MFHLAJACMFA,omitempty"`
	Rotation            *MathQuaternion                         `protobuf:"bytes,7,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Center              *Vector                                 `protobuf:"bytes,13,opt,name=center,proto3" json:"center,omitempty"`
	Unk3000_LNHPLNEBBIP *Vector                                 `protobuf:"bytes,14,opt,name=Unk3000_LNHPLNEBBIP,json=Unk3000LNHPLNEBBIP,proto3" json:"Unk3000_LNHPLNEBBIP,omitempty"`
}

func (x *Unk3000_LBNFMLHLBIH) Reset() {
	*x = Unk3000_LBNFMLHLBIH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_LBNFMLHLBIH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_LBNFMLHLBIH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_LBNFMLHLBIH) ProtoMessage() {}

func (x *Unk3000_LBNFMLHLBIH) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_LBNFMLHLBIH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_LBNFMLHLBIH.ProtoReflect.Descriptor instead.
func (*Unk3000_LBNFMLHLBIH) Descriptor() ([]byte, []int) {
	return file_Unk3000_LBNFMLHLBIH_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_LBNFMLHLBIH) GetType() Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN {
	if x != nil {
		return x.Type
	}
	return Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN_OBSTACLE_SHAPE_CAPSULE
}

func (x *Unk3000_LBNFMLHLBIH) GetUnk3000_MFHLAJACMFA() int32 {
	if x != nil {
		return x.Unk3000_MFHLAJACMFA
	}
	return 0
}

func (x *Unk3000_LBNFMLHLBIH) GetRotation() *MathQuaternion {
	if x != nil {
		return x.Rotation
	}
	return nil
}

func (x *Unk3000_LBNFMLHLBIH) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *Unk3000_LBNFMLHLBIH) GetUnk3000_LNHPLNEBBIP() *Vector {
	if x != nil {
		return x.Unk3000_LNHPLNEBBIP
	}
	return nil
}

var File_Unk3000_LBNFMLHLBIH_proto protoreflect.FileDescriptor

var file_Unk3000_LBNFMLHLBIH_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4c, 0x42, 0x4e, 0x46, 0x4d, 0x4c,
	0x48, 0x4c, 0x42, 0x49, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4d, 0x61, 0x74,
	0x68, 0x51, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xff, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4c, 0x42, 0x4e, 0x46,
	0x4d, 0x4c, 0x48, 0x4c, 0x42, 0x49, 0x48, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f,
	0x4c, 0x42, 0x4e, 0x46, 0x4d, 0x4c, 0x48, 0x4c, 0x42, 0x49, 0x48, 0x2e, 0x55, 0x6e, 0x6b, 0x33,
	0x30, 0x30, 0x30, 0x5f, 0x47, 0x50, 0x48, 0x42, 0x49, 0x42, 0x47, 0x4d, 0x48, 0x4a, 0x4e, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x5f, 0x4d, 0x46, 0x48, 0x4c, 0x41, 0x4a, 0x41, 0x43, 0x4d, 0x46, 0x41, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x4d, 0x46, 0x48, 0x4c, 0x41,
	0x4a, 0x41, 0x43, 0x4d, 0x46, 0x41, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x51,
	0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f,
	0x4c, 0x4e, 0x48, 0x50, 0x4c, 0x4e, 0x45, 0x42, 0x42, 0x49, 0x50, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x30, 0x30, 0x30, 0x4c, 0x4e, 0x48, 0x50, 0x4c, 0x4e, 0x45, 0x42, 0x42, 0x49, 0x50, 0x22, 0x71,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x47, 0x50, 0x48, 0x42, 0x49, 0x42,
	0x47, 0x4d, 0x48, 0x4a, 0x4e, 0x12, 0x2e, 0x0a, 0x2a, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x5f, 0x47, 0x50, 0x48, 0x42, 0x49, 0x42, 0x47, 0x4d, 0x48, 0x4a, 0x4e, 0x5f, 0x4f, 0x42, 0x53,
	0x54, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x50, 0x53,
	0x55, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x2a, 0x0a, 0x26, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x5f, 0x47, 0x50, 0x48, 0x42, 0x49, 0x42, 0x47, 0x4d, 0x48, 0x4a, 0x4e, 0x5f, 0x4f, 0x42, 0x53,
	0x54, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x58, 0x10,
	0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Unk3000_LBNFMLHLBIH_proto_rawDescOnce sync.Once
	file_Unk3000_LBNFMLHLBIH_proto_rawDescData = file_Unk3000_LBNFMLHLBIH_proto_rawDesc
)

func file_Unk3000_LBNFMLHLBIH_proto_rawDescGZIP() []byte {
	file_Unk3000_LBNFMLHLBIH_proto_rawDescOnce.Do(func() {
		file_Unk3000_LBNFMLHLBIH_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_LBNFMLHLBIH_proto_rawDescData)
	})
	return file_Unk3000_LBNFMLHLBIH_proto_rawDescData
}

var file_Unk3000_LBNFMLHLBIH_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Unk3000_LBNFMLHLBIH_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3000_LBNFMLHLBIH_proto_goTypes = []interface{}{
	(Unk3000_LBNFMLHLBIH_Unk3000_GPHBIBGMHJN)(0), // 0: Unk3000_LBNFMLHLBIH.Unk3000_GPHBIBGMHJN
	(*Unk3000_LBNFMLHLBIH)(nil),                  // 1: Unk3000_LBNFMLHLBIH
	(*MathQuaternion)(nil),                       // 2: MathQuaternion
	(*Vector)(nil),                               // 3: Vector
}
var file_Unk3000_LBNFMLHLBIH_proto_depIdxs = []int32{
	0, // 0: Unk3000_LBNFMLHLBIH.type:type_name -> Unk3000_LBNFMLHLBIH.Unk3000_GPHBIBGMHJN
	2, // 1: Unk3000_LBNFMLHLBIH.rotation:type_name -> MathQuaternion
	3, // 2: Unk3000_LBNFMLHLBIH.center:type_name -> Vector
	3, // 3: Unk3000_LBNFMLHLBIH.Unk3000_LNHPLNEBBIP:type_name -> Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Unk3000_LBNFMLHLBIH_proto_init() }
func file_Unk3000_LBNFMLHLBIH_proto_init() {
	if File_Unk3000_LBNFMLHLBIH_proto != nil {
		return
	}
	file_MathQuaternion_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_LBNFMLHLBIH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_LBNFMLHLBIH); i {
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
			RawDescriptor: file_Unk3000_LBNFMLHLBIH_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_LBNFMLHLBIH_proto_goTypes,
		DependencyIndexes: file_Unk3000_LBNFMLHLBIH_proto_depIdxs,
		EnumInfos:         file_Unk3000_LBNFMLHLBIH_proto_enumTypes,
		MessageInfos:      file_Unk3000_LBNFMLHLBIH_proto_msgTypes,
	}.Build()
	File_Unk3000_LBNFMLHLBIH_proto = out.File
	file_Unk3000_LBNFMLHLBIH_proto_rawDesc = nil
	file_Unk3000_LBNFMLHLBIH_proto_goTypes = nil
	file_Unk3000_LBNFMLHLBIH_proto_depIdxs = nil
}