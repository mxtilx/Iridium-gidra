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
// source: HuntingOfferState.proto

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

type HuntingOfferState int32

const (
	HuntingOfferState_HUNTING_OFFER_STATE_NONE      HuntingOfferState = 0
	HuntingOfferState_HUNTING_OFFER_STATE_STARTED   HuntingOfferState = 1
	HuntingOfferState_HUNTING_OFFER_STATE_UNSTARTED HuntingOfferState = 2
	HuntingOfferState_HUNTING_OFFER_STATE_SUCC      HuntingOfferState = 3
)

// Enum value maps for HuntingOfferState.
var (
	HuntingOfferState_name = map[int32]string{
		0: "HUNTING_OFFER_STATE_NONE",
		1: "HUNTING_OFFER_STATE_STARTED",
		2: "HUNTING_OFFER_STATE_UNSTARTED",
		3: "HUNTING_OFFER_STATE_SUCC",
	}
	HuntingOfferState_value = map[string]int32{
		"HUNTING_OFFER_STATE_NONE":      0,
		"HUNTING_OFFER_STATE_STARTED":   1,
		"HUNTING_OFFER_STATE_UNSTARTED": 2,
		"HUNTING_OFFER_STATE_SUCC":      3,
	}
)

func (x HuntingOfferState) Enum() *HuntingOfferState {
	p := new(HuntingOfferState)
	*p = x
	return p
}

func (x HuntingOfferState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HuntingOfferState) Descriptor() protoreflect.EnumDescriptor {
	return file_HuntingOfferState_proto_enumTypes[0].Descriptor()
}

func (HuntingOfferState) Type() protoreflect.EnumType {
	return &file_HuntingOfferState_proto_enumTypes[0]
}

func (x HuntingOfferState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HuntingOfferState.Descriptor instead.
func (HuntingOfferState) EnumDescriptor() ([]byte, []int) {
	return file_HuntingOfferState_proto_rawDescGZIP(), []int{0}
}

var File_HuntingOfferState_proto protoreflect.FileDescriptor

var file_HuntingOfferState_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x93, 0x01, 0x0a, 0x11, 0x48, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x18, 0x48, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x46, 0x46, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1f, 0x0a,
	0x1b, 0x48, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x21,
	0x0a, 0x1d, 0x48, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x1c, 0x0a, 0x18, 0x48, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x46, 0x46,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x10, 0x03, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HuntingOfferState_proto_rawDescOnce sync.Once
	file_HuntingOfferState_proto_rawDescData = file_HuntingOfferState_proto_rawDesc
)

func file_HuntingOfferState_proto_rawDescGZIP() []byte {
	file_HuntingOfferState_proto_rawDescOnce.Do(func() {
		file_HuntingOfferState_proto_rawDescData = protoimpl.X.CompressGZIP(file_HuntingOfferState_proto_rawDescData)
	})
	return file_HuntingOfferState_proto_rawDescData
}

var file_HuntingOfferState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HuntingOfferState_proto_goTypes = []interface{}{
	(HuntingOfferState)(0), // 0: HuntingOfferState
}
var file_HuntingOfferState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HuntingOfferState_proto_init() }
func file_HuntingOfferState_proto_init() {
	if File_HuntingOfferState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HuntingOfferState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HuntingOfferState_proto_goTypes,
		DependencyIndexes: file_HuntingOfferState_proto_depIdxs,
		EnumInfos:         file_HuntingOfferState_proto_enumTypes,
	}.Build()
	File_HuntingOfferState_proto = out.File
	file_HuntingOfferState_proto_rawDesc = nil
	file_HuntingOfferState_proto_goTypes = nil
	file_HuntingOfferState_proto_depIdxs = nil
}