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
// source: OFOHJIEDJKD.proto

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

type OFOHJIEDJKD int32

const (
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseInvalid OFOHJIEDJKD = 0
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseStart   OFOHJIEDJKD = 1
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseDraw    OFOHJIEDJKD = 2
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseOnStage OFOHJIEDJKD = 3
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseDice    OFOHJIEDJKD = 4
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseMain    OFOHJIEDJKD = 5
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseEnd     OFOHJIEDJKD = 6
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseDie     OFOHJIEDJKD = 7
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseFin     OFOHJIEDJKD = 8
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhasePreMain OFOHJIEDJKD = 9
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseReroll  OFOHJIEDJKD = 10
	OFOHJIEDJKD_OFOHJIEDJKD_GcgPhaseRedraw  OFOHJIEDJKD = 11
)

// Enum value maps for OFOHJIEDJKD.
var (
	OFOHJIEDJKD_name = map[int32]string{
		0:  "OFOHJIEDJKD_GcgPhaseInvalid",
		1:  "OFOHJIEDJKD_GcgPhaseStart",
		2:  "OFOHJIEDJKD_GcgPhaseDraw",
		3:  "OFOHJIEDJKD_GcgPhaseOnStage",
		4:  "OFOHJIEDJKD_GcgPhaseDice",
		5:  "OFOHJIEDJKD_GcgPhaseMain",
		6:  "OFOHJIEDJKD_GcgPhaseEnd",
		7:  "OFOHJIEDJKD_GcgPhaseDie",
		8:  "OFOHJIEDJKD_GcgPhaseFin",
		9:  "OFOHJIEDJKD_GcgPhasePreMain",
		10: "OFOHJIEDJKD_GcgPhaseReroll",
		11: "OFOHJIEDJKD_GcgPhaseRedraw",
	}
	OFOHJIEDJKD_value = map[string]int32{
		"OFOHJIEDJKD_GcgPhaseInvalid": 0,
		"OFOHJIEDJKD_GcgPhaseStart":   1,
		"OFOHJIEDJKD_GcgPhaseDraw":    2,
		"OFOHJIEDJKD_GcgPhaseOnStage": 3,
		"OFOHJIEDJKD_GcgPhaseDice":    4,
		"OFOHJIEDJKD_GcgPhaseMain":    5,
		"OFOHJIEDJKD_GcgPhaseEnd":     6,
		"OFOHJIEDJKD_GcgPhaseDie":     7,
		"OFOHJIEDJKD_GcgPhaseFin":     8,
		"OFOHJIEDJKD_GcgPhasePreMain": 9,
		"OFOHJIEDJKD_GcgPhaseReroll":  10,
		"OFOHJIEDJKD_GcgPhaseRedraw":  11,
	}
)

func (x OFOHJIEDJKD) Enum() *OFOHJIEDJKD {
	p := new(OFOHJIEDJKD)
	*p = x
	return p
}

func (x OFOHJIEDJKD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OFOHJIEDJKD) Descriptor() protoreflect.EnumDescriptor {
	return file_OFOHJIEDJKD_proto_enumTypes[0].Descriptor()
}

func (OFOHJIEDJKD) Type() protoreflect.EnumType {
	return &file_OFOHJIEDJKD_proto_enumTypes[0]
}

func (x OFOHJIEDJKD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OFOHJIEDJKD.Descriptor instead.
func (OFOHJIEDJKD) EnumDescriptor() ([]byte, []int) {
	return file_OFOHJIEDJKD_proto_rawDescGZIP(), []int{0}
}

var File_OFOHJIEDJKD_proto protoreflect.FileDescriptor

var file_OFOHJIEDJKD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a, 0x4b, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x80, 0x03, 0x0a, 0x0b, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44,
	0x4a, 0x4b, 0x44, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a,
	0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44,
	0x4a, 0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a,
	0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x44, 0x72, 0x61, 0x77, 0x10,
	0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a, 0x4b, 0x44,
	0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a, 0x4b,
	0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x44, 0x69, 0x63, 0x65, 0x10, 0x04,
	0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a, 0x4b, 0x44, 0x5f,
	0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x10, 0x05, 0x12, 0x1b,
	0x0a, 0x17, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a, 0x4b, 0x44, 0x5f, 0x47, 0x63,
	0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x4f,
	0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45, 0x44, 0x4a, 0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68,
	0x61, 0x73, 0x65, 0x44, 0x69, 0x65, 0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x46, 0x4f, 0x48,
	0x4a, 0x49, 0x45, 0x44, 0x4a, 0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x46, 0x69, 0x6e, 0x10, 0x08, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49, 0x45,
	0x44, 0x4a, 0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x50, 0x72, 0x65,
	0x4d, 0x61, 0x69, 0x6e, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49,
	0x45, 0x44, 0x4a, 0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x72, 0x6f, 0x6c, 0x6c, 0x10, 0x0a, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x46, 0x4f, 0x48, 0x4a, 0x49,
	0x45, 0x44, 0x4a, 0x4b, 0x44, 0x5f, 0x47, 0x63, 0x67, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x64, 0x72, 0x61, 0x77, 0x10, 0x0b, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OFOHJIEDJKD_proto_rawDescOnce sync.Once
	file_OFOHJIEDJKD_proto_rawDescData = file_OFOHJIEDJKD_proto_rawDesc
)

func file_OFOHJIEDJKD_proto_rawDescGZIP() []byte {
	file_OFOHJIEDJKD_proto_rawDescOnce.Do(func() {
		file_OFOHJIEDJKD_proto_rawDescData = protoimpl.X.CompressGZIP(file_OFOHJIEDJKD_proto_rawDescData)
	})
	return file_OFOHJIEDJKD_proto_rawDescData
}

var file_OFOHJIEDJKD_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_OFOHJIEDJKD_proto_goTypes = []interface{}{
	(OFOHJIEDJKD)(0), // 0: OFOHJIEDJKD
}
var file_OFOHJIEDJKD_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OFOHJIEDJKD_proto_init() }
func file_OFOHJIEDJKD_proto_init() {
	if File_OFOHJIEDJKD_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_OFOHJIEDJKD_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OFOHJIEDJKD_proto_goTypes,
		DependencyIndexes: file_OFOHJIEDJKD_proto_depIdxs,
		EnumInfos:         file_OFOHJIEDJKD_proto_enumTypes,
	}.Build()
	File_OFOHJIEDJKD_proto = out.File
	file_OFOHJIEDJKD_proto_rawDesc = nil
	file_OFOHJIEDJKD_proto_goTypes = nil
	file_OFOHJIEDJKD_proto_depIdxs = nil
}