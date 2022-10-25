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
// source: LanternRiteActivityDetailInfo.proto

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

type LanternRiteActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_ONOHODJPIGK *Unk2700_JCNIPOJMFMH   `protobuf:"bytes,13,opt,name=Unk2700_ONOHODJPIGK,json=Unk2700ONOHODJPIGK,proto3" json:"Unk2700_ONOHODJPIGK,omitempty"`
	Unk2700_PHKHIPLDOOA []*Unk2700_LLGDCAKMCKL `protobuf:"bytes,5,rep,name=Unk2700_PHKHIPLDOOA,json=Unk2700PHKHIPLDOOA,proto3" json:"Unk2700_PHKHIPLDOOA,omitempty"`
	Unk2700_MPOCLGBFNAK *Unk2700_MJGFEHOMKJE   `protobuf:"bytes,8,opt,name=Unk2700_MPOCLGBFNAK,json=Unk2700MPOCLGBFNAK,proto3" json:"Unk2700_MPOCLGBFNAK,omitempty"`
	Unk2700_KGGCKHBIOED bool                   `protobuf:"varint,2,opt,name=Unk2700_KGGCKHBIOED,json=Unk2700KGGCKHBIOED,proto3" json:"Unk2700_KGGCKHBIOED,omitempty"`
	IsContentClosed     bool                   `protobuf:"varint,14,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	Unk2700_EOGEAIHJPFD bool                   `protobuf:"varint,6,opt,name=Unk2700_EOGEAIHJPFD,json=Unk2700EOGEAIHJPFD,proto3" json:"Unk2700_EOGEAIHJPFD,omitempty"`
}

func (x *LanternRiteActivityDetailInfo) Reset() {
	*x = LanternRiteActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LanternRiteActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanternRiteActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanternRiteActivityDetailInfo) ProtoMessage() {}

func (x *LanternRiteActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_LanternRiteActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanternRiteActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*LanternRiteActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_LanternRiteActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *LanternRiteActivityDetailInfo) GetUnk2700_ONOHODJPIGK() *Unk2700_JCNIPOJMFMH {
	if x != nil {
		return x.Unk2700_ONOHODJPIGK
	}
	return nil
}

func (x *LanternRiteActivityDetailInfo) GetUnk2700_PHKHIPLDOOA() []*Unk2700_LLGDCAKMCKL {
	if x != nil {
		return x.Unk2700_PHKHIPLDOOA
	}
	return nil
}

func (x *LanternRiteActivityDetailInfo) GetUnk2700_MPOCLGBFNAK() *Unk2700_MJGFEHOMKJE {
	if x != nil {
		return x.Unk2700_MPOCLGBFNAK
	}
	return nil
}

func (x *LanternRiteActivityDetailInfo) GetUnk2700_KGGCKHBIOED() bool {
	if x != nil {
		return x.Unk2700_KGGCKHBIOED
	}
	return false
}

func (x *LanternRiteActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *LanternRiteActivityDetailInfo) GetUnk2700_EOGEAIHJPFD() bool {
	if x != nil {
		return x.Unk2700_EOGEAIHJPFD
	}
	return false
}

var File_LanternRiteActivityDetailInfo_proto protoreflect.FileDescriptor

var file_LanternRiteActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x23, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a,
	0x43, 0x4e, 0x49, 0x50, 0x4f, 0x4a, 0x4d, 0x46, 0x4d, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4c, 0x4c, 0x47, 0x44, 0x43, 0x41,
	0x4b, 0x4d, 0x43, 0x4b, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x4a, 0x47, 0x46, 0x45, 0x48, 0x4f, 0x4d, 0x4b, 0x4a, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x1d, 0x4c, 0x61, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x4f, 0x4e, 0x4f, 0x48, 0x4f, 0x44, 0x4a, 0x50, 0x49, 0x47, 0x4b, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x4a, 0x43, 0x4e, 0x49, 0x50, 0x4f, 0x4a, 0x4d, 0x46, 0x4d, 0x48, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x4f, 0x4e, 0x4f, 0x48, 0x4f, 0x44, 0x4a, 0x50, 0x49, 0x47, 0x4b, 0x12,
	0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x48, 0x4b, 0x48, 0x49,
	0x50, 0x4c, 0x44, 0x4f, 0x4f, 0x41, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4c, 0x4c, 0x47, 0x44, 0x43, 0x41, 0x4b, 0x4d, 0x43,
	0x4b, 0x4c, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x50, 0x48, 0x4b, 0x48, 0x49,
	0x50, 0x4c, 0x44, 0x4f, 0x4f, 0x41, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x5f, 0x4d, 0x50, 0x4f, 0x43, 0x4c, 0x47, 0x42, 0x46, 0x4e, 0x41, 0x4b, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x4a,
	0x47, 0x46, 0x45, 0x48, 0x4f, 0x4d, 0x4b, 0x4a, 0x45, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x4d, 0x50, 0x4f, 0x43, 0x4c, 0x47, 0x42, 0x46, 0x4e, 0x41, 0x4b, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x47, 0x47, 0x43, 0x4b, 0x48, 0x42,
	0x49, 0x4f, 0x45, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x4b, 0x47, 0x47, 0x43, 0x4b, 0x48, 0x42, 0x49, 0x4f, 0x45, 0x44, 0x12, 0x2a,
	0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x4f, 0x47, 0x45, 0x41, 0x49, 0x48, 0x4a, 0x50, 0x46,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x45, 0x4f, 0x47, 0x45, 0x41, 0x49, 0x48, 0x4a, 0x50, 0x46, 0x44, 0x42, 0x06, 0x5a, 0x04, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LanternRiteActivityDetailInfo_proto_rawDescOnce sync.Once
	file_LanternRiteActivityDetailInfo_proto_rawDescData = file_LanternRiteActivityDetailInfo_proto_rawDesc
)

func file_LanternRiteActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_LanternRiteActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_LanternRiteActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_LanternRiteActivityDetailInfo_proto_rawDescData)
	})
	return file_LanternRiteActivityDetailInfo_proto_rawDescData
}

var file_LanternRiteActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LanternRiteActivityDetailInfo_proto_goTypes = []interface{}{
	(*LanternRiteActivityDetailInfo)(nil), // 0: LanternRiteActivityDetailInfo
	(*Unk2700_JCNIPOJMFMH)(nil),           // 1: Unk2700_JCNIPOJMFMH
	(*Unk2700_LLGDCAKMCKL)(nil),           // 2: Unk2700_LLGDCAKMCKL
	(*Unk2700_MJGFEHOMKJE)(nil),           // 3: Unk2700_MJGFEHOMKJE
}
var file_LanternRiteActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: LanternRiteActivityDetailInfo.Unk2700_ONOHODJPIGK:type_name -> Unk2700_JCNIPOJMFMH
	2, // 1: LanternRiteActivityDetailInfo.Unk2700_PHKHIPLDOOA:type_name -> Unk2700_LLGDCAKMCKL
	3, // 2: LanternRiteActivityDetailInfo.Unk2700_MPOCLGBFNAK:type_name -> Unk2700_MJGFEHOMKJE
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_LanternRiteActivityDetailInfo_proto_init() }
func file_LanternRiteActivityDetailInfo_proto_init() {
	if File_LanternRiteActivityDetailInfo_proto != nil {
		return
	}
	file_Unk2700_JCNIPOJMFMH_proto_init()
	file_Unk2700_LLGDCAKMCKL_proto_init()
	file_Unk2700_MJGFEHOMKJE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LanternRiteActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanternRiteActivityDetailInfo); i {
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
			RawDescriptor: file_LanternRiteActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LanternRiteActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_LanternRiteActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_LanternRiteActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_LanternRiteActivityDetailInfo_proto = out.File
	file_LanternRiteActivityDetailInfo_proto_rawDesc = nil
	file_LanternRiteActivityDetailInfo_proto_goTypes = nil
	file_LanternRiteActivityDetailInfo_proto_depIdxs = nil
}