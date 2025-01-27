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
// source: IFNOHBONEGM.proto

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

type IFNOHBONEGM_JNHPFDEDNEI int32

const (
	IFNOHBONEGM_HMNKHDMCEKN_StatusFail    IFNOHBONEGM_JNHPFDEDNEI = 0
	IFNOHBONEGM_HMNKHDMCEKN_StatusSucc    IFNOHBONEGM_JNHPFDEDNEI = 1
	IFNOHBONEGM_HMNKHDMCEKN_StatusPartial IFNOHBONEGM_JNHPFDEDNEI = 2
)

// Enum value maps for IFNOHBONEGM_JNHPFDEDNEI.
var (
	IFNOHBONEGM_JNHPFDEDNEI_name = map[int32]string{
		0: "HMNKHDMCEKN_StatusFail",
		1: "HMNKHDMCEKN_StatusSucc",
		2: "HMNKHDMCEKN_StatusPartial",
	}
	IFNOHBONEGM_JNHPFDEDNEI_value = map[string]int32{
		"HMNKHDMCEKN_StatusFail":    0,
		"HMNKHDMCEKN_StatusSucc":    1,
		"HMNKHDMCEKN_StatusPartial": 2,
	}
)

func (x IFNOHBONEGM_JNHPFDEDNEI) Enum() *IFNOHBONEGM_JNHPFDEDNEI {
	p := new(IFNOHBONEGM_JNHPFDEDNEI)
	*p = x
	return p
}

func (x IFNOHBONEGM_JNHPFDEDNEI) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IFNOHBONEGM_JNHPFDEDNEI) Descriptor() protoreflect.EnumDescriptor {
	return file_IFNOHBONEGM_proto_enumTypes[0].Descriptor()
}

func (IFNOHBONEGM_JNHPFDEDNEI) Type() protoreflect.EnumType {
	return &file_IFNOHBONEGM_proto_enumTypes[0]
}

func (x IFNOHBONEGM_JNHPFDEDNEI) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IFNOHBONEGM_JNHPFDEDNEI.Descriptor instead.
func (IFNOHBONEGM_JNHPFDEDNEI) EnumDescriptor() ([]byte, []int) {
	return file_IFNOHBONEGM_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 2397
type IFNOHBONEGM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Corners     []*Vector               `protobuf:"bytes,4,rep,name=corners,proto3" json:"corners,omitempty"`
	QueryStatus IFNOHBONEGM_JNHPFDEDNEI `protobuf:"varint,1,opt,name=query_status,json=queryStatus,proto3,enum=IFNOHBONEGM_JNHPFDEDNEI" json:"query_status,omitempty"`
	Retcode     int32                   `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MFJENPLIKPF []*FFELGFFLAAC          `protobuf:"bytes,6,rep,name=MFJENPLIKPF,proto3" json:"MFJENPLIKPF,omitempty"`
	QueryId     int32                   `protobuf:"varint,12,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	DKPKIMPJIOD []*FFELGFFLAAC          `protobuf:"bytes,5,rep,name=DKPKIMPJIOD,proto3" json:"DKPKIMPJIOD,omitempty"`
	FOPNKJMGKBI []*Vector               `protobuf:"bytes,9,rep,name=FOPNKJMGKBI,proto3" json:"FOPNKJMGKBI,omitempty"`
}

func (x *IFNOHBONEGM) Reset() {
	*x = IFNOHBONEGM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IFNOHBONEGM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFNOHBONEGM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFNOHBONEGM) ProtoMessage() {}

func (x *IFNOHBONEGM) ProtoReflect() protoreflect.Message {
	mi := &file_IFNOHBONEGM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFNOHBONEGM.ProtoReflect.Descriptor instead.
func (*IFNOHBONEGM) Descriptor() ([]byte, []int) {
	return file_IFNOHBONEGM_proto_rawDescGZIP(), []int{0}
}

func (x *IFNOHBONEGM) GetCorners() []*Vector {
	if x != nil {
		return x.Corners
	}
	return nil
}

func (x *IFNOHBONEGM) GetQueryStatus() IFNOHBONEGM_JNHPFDEDNEI {
	if x != nil {
		return x.QueryStatus
	}
	return IFNOHBONEGM_HMNKHDMCEKN_StatusFail
}

func (x *IFNOHBONEGM) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *IFNOHBONEGM) GetMFJENPLIKPF() []*FFELGFFLAAC {
	if x != nil {
		return x.MFJENPLIKPF
	}
	return nil
}

func (x *IFNOHBONEGM) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *IFNOHBONEGM) GetDKPKIMPJIOD() []*FFELGFFLAAC {
	if x != nil {
		return x.DKPKIMPJIOD
	}
	return nil
}

func (x *IFNOHBONEGM) GetFOPNKJMGKBI() []*Vector {
	if x != nil {
		return x.FOPNKJMGKBI
	}
	return nil
}

var File_IFNOHBONEGM_proto protoreflect.FileDescriptor

var file_IFNOHBONEGM_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x46, 0x4e, 0x4f, 0x48, 0x42, 0x4f, 0x4e, 0x45, 0x47, 0x4d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x46, 0x46, 0x45, 0x4c, 0x47, 0x46, 0x46, 0x4c, 0x41, 0x41, 0x43, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x0b, 0x49, 0x46, 0x4e, 0x4f, 0x48, 0x42, 0x4f,
	0x4e, 0x45, 0x47, 0x4d, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07,
	0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x49, 0x46, 0x4e, 0x4f, 0x48, 0x42, 0x4f, 0x4e, 0x45, 0x47, 0x4d, 0x2e, 0x4a, 0x4e, 0x48, 0x50,
	0x46, 0x44, 0x45, 0x44, 0x4e, 0x45, 0x49, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e,
	0x0a, 0x0b, 0x4d, 0x46, 0x4a, 0x45, 0x4e, 0x50, 0x4c, 0x49, 0x4b, 0x50, 0x46, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x46, 0x45, 0x4c, 0x47, 0x46, 0x46, 0x4c, 0x41, 0x41,
	0x43, 0x52, 0x0b, 0x4d, 0x46, 0x4a, 0x45, 0x4e, 0x50, 0x4c, 0x49, 0x4b, 0x50, 0x46, 0x12, 0x19,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x4b, 0x50,
	0x4b, 0x49, 0x4d, 0x50, 0x4a, 0x49, 0x4f, 0x44, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x46, 0x46, 0x45, 0x4c, 0x47, 0x46, 0x46, 0x4c, 0x41, 0x41, 0x43, 0x52, 0x0b, 0x44, 0x4b,
	0x50, 0x4b, 0x49, 0x4d, 0x50, 0x4a, 0x49, 0x4f, 0x44, 0x12, 0x29, 0x0a, 0x0b, 0x46, 0x4f, 0x50,
	0x4e, 0x4b, 0x4a, 0x4d, 0x47, 0x4b, 0x42, 0x49, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x46, 0x4f, 0x50, 0x4e, 0x4b, 0x4a, 0x4d,
	0x47, 0x4b, 0x42, 0x49, 0x22, 0x64, 0x0a, 0x0b, 0x4a, 0x4e, 0x48, 0x50, 0x46, 0x44, 0x45, 0x44,
	0x4e, 0x45, 0x49, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x4d, 0x4e, 0x4b, 0x48, 0x44, 0x4d, 0x43, 0x45,
	0x4b, 0x4e, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x48, 0x4d, 0x4e, 0x4b, 0x48, 0x44, 0x4d, 0x43, 0x45, 0x4b, 0x4e, 0x5f, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x63, 0x63, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x48,
	0x4d, 0x4e, 0x4b, 0x48, 0x44, 0x4d, 0x43, 0x45, 0x4b, 0x4e, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IFNOHBONEGM_proto_rawDescOnce sync.Once
	file_IFNOHBONEGM_proto_rawDescData = file_IFNOHBONEGM_proto_rawDesc
)

func file_IFNOHBONEGM_proto_rawDescGZIP() []byte {
	file_IFNOHBONEGM_proto_rawDescOnce.Do(func() {
		file_IFNOHBONEGM_proto_rawDescData = protoimpl.X.CompressGZIP(file_IFNOHBONEGM_proto_rawDescData)
	})
	return file_IFNOHBONEGM_proto_rawDescData
}

var file_IFNOHBONEGM_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_IFNOHBONEGM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IFNOHBONEGM_proto_goTypes = []interface{}{
	(IFNOHBONEGM_JNHPFDEDNEI)(0), // 0: IFNOHBONEGM.JNHPFDEDNEI
	(*IFNOHBONEGM)(nil),          // 1: IFNOHBONEGM
	(*Vector)(nil),               // 2: Vector
	(*FFELGFFLAAC)(nil),          // 3: FFELGFFLAAC
}
var file_IFNOHBONEGM_proto_depIdxs = []int32{
	2, // 0: IFNOHBONEGM.corners:type_name -> Vector
	0, // 1: IFNOHBONEGM.query_status:type_name -> IFNOHBONEGM.JNHPFDEDNEI
	3, // 2: IFNOHBONEGM.MFJENPLIKPF:type_name -> FFELGFFLAAC
	3, // 3: IFNOHBONEGM.DKPKIMPJIOD:type_name -> FFELGFFLAAC
	2, // 4: IFNOHBONEGM.FOPNKJMGKBI:type_name -> Vector
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_IFNOHBONEGM_proto_init() }
func file_IFNOHBONEGM_proto_init() {
	if File_IFNOHBONEGM_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_FFELGFFLAAC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IFNOHBONEGM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IFNOHBONEGM); i {
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
			RawDescriptor: file_IFNOHBONEGM_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IFNOHBONEGM_proto_goTypes,
		DependencyIndexes: file_IFNOHBONEGM_proto_depIdxs,
		EnumInfos:         file_IFNOHBONEGM_proto_enumTypes,
		MessageInfos:      file_IFNOHBONEGM_proto_msgTypes,
	}.Build()
	File_IFNOHBONEGM_proto = out.File
	file_IFNOHBONEGM_proto_rawDesc = nil
	file_IFNOHBONEGM_proto_goTypes = nil
	file_IFNOHBONEGM_proto_depIdxs = nil
}
