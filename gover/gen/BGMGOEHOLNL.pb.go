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
// source: BGMGOEHOLNL.proto

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

// CmdId: 7626
type BGMGOEHOLNL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DDJCLIAMDHB uint32       `protobuf:"varint,2,opt,name=DDJCLIAMDHB,proto3" json:"DDJCLIAMDHB,omitempty"`
	Retcode     int32        `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MEGGIEFEENH *IIPPNNPPDEG `protobuf:"bytes,13,opt,name=MEGGIEFEENH,proto3" json:"MEGGIEFEENH,omitempty"`
	ScheduleId  uint32       `protobuf:"varint,1,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *BGMGOEHOLNL) Reset() {
	*x = BGMGOEHOLNL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BGMGOEHOLNL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BGMGOEHOLNL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BGMGOEHOLNL) ProtoMessage() {}

func (x *BGMGOEHOLNL) ProtoReflect() protoreflect.Message {
	mi := &file_BGMGOEHOLNL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BGMGOEHOLNL.ProtoReflect.Descriptor instead.
func (*BGMGOEHOLNL) Descriptor() ([]byte, []int) {
	return file_BGMGOEHOLNL_proto_rawDescGZIP(), []int{0}
}

func (x *BGMGOEHOLNL) GetDDJCLIAMDHB() uint32 {
	if x != nil {
		return x.DDJCLIAMDHB
	}
	return 0
}

func (x *BGMGOEHOLNL) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *BGMGOEHOLNL) GetMEGGIEFEENH() *IIPPNNPPDEG {
	if x != nil {
		return x.MEGGIEFEENH
	}
	return nil
}

func (x *BGMGOEHOLNL) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

var File_BGMGOEHOLNL_proto protoreflect.FileDescriptor

var file_BGMGOEHOLNL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x47, 0x4d, 0x47, 0x4f, 0x45, 0x48, 0x4f, 0x4c, 0x4e, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x49, 0x50, 0x50, 0x4e, 0x4e, 0x50, 0x50, 0x44, 0x45, 0x47,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x42, 0x47, 0x4d, 0x47, 0x4f,
	0x45, 0x48, 0x4f, 0x4c, 0x4e, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x44, 0x4a, 0x43, 0x4c, 0x49,
	0x41, 0x4d, 0x44, 0x48, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x44, 0x4a,
	0x43, 0x4c, 0x49, 0x41, 0x4d, 0x44, 0x48, 0x42, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x45, 0x47, 0x47, 0x49, 0x45, 0x46, 0x45, 0x45, 0x4e,
	0x48, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x49, 0x50, 0x50, 0x4e, 0x4e,
	0x50, 0x50, 0x44, 0x45, 0x47, 0x52, 0x0b, 0x4d, 0x45, 0x47, 0x47, 0x49, 0x45, 0x46, 0x45, 0x45,
	0x4e, 0x48, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_BGMGOEHOLNL_proto_rawDescOnce sync.Once
	file_BGMGOEHOLNL_proto_rawDescData = file_BGMGOEHOLNL_proto_rawDesc
)

func file_BGMGOEHOLNL_proto_rawDescGZIP() []byte {
	file_BGMGOEHOLNL_proto_rawDescOnce.Do(func() {
		file_BGMGOEHOLNL_proto_rawDescData = protoimpl.X.CompressGZIP(file_BGMGOEHOLNL_proto_rawDescData)
	})
	return file_BGMGOEHOLNL_proto_rawDescData
}

var file_BGMGOEHOLNL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BGMGOEHOLNL_proto_goTypes = []interface{}{
	(*BGMGOEHOLNL)(nil), // 0: BGMGOEHOLNL
	(*IIPPNNPPDEG)(nil), // 1: IIPPNNPPDEG
}
var file_BGMGOEHOLNL_proto_depIdxs = []int32{
	1, // 0: BGMGOEHOLNL.MEGGIEFEENH:type_name -> IIPPNNPPDEG
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BGMGOEHOLNL_proto_init() }
func file_BGMGOEHOLNL_proto_init() {
	if File_BGMGOEHOLNL_proto != nil {
		return
	}
	file_IIPPNNPPDEG_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BGMGOEHOLNL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BGMGOEHOLNL); i {
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
			RawDescriptor: file_BGMGOEHOLNL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BGMGOEHOLNL_proto_goTypes,
		DependencyIndexes: file_BGMGOEHOLNL_proto_depIdxs,
		MessageInfos:      file_BGMGOEHOLNL_proto_msgTypes,
	}.Build()
	File_BGMGOEHOLNL_proto = out.File
	file_BGMGOEHOLNL_proto_rawDesc = nil
	file_BGMGOEHOLNL_proto_goTypes = nil
	file_BGMGOEHOLNL_proto_depIdxs = nil
}