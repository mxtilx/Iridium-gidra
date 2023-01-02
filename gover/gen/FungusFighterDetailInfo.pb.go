// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
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
// 	protoc        v3.11.3
// source: FungusFighterDetailInfo.proto

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

type FungusFighterDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FungusDetailList                  []*FungusDetail                 `protobuf:"bytes,6,rep,name=fungus_detail_list,json=fungusDetailList,proto3" json:"fungus_detail_list,omitempty"`
	Unk3300_GIHAKKAJHDH               []uint32                        `protobuf:"varint,1,rep,packed,name=Unk3300_GIHAKKAJHDH,json=Unk3300GIHAKKAJHDH,proto3" json:"Unk3300_GIHAKKAJHDH,omitempty"`
	Unk3300_KGKFJJHFHAB               []uint32                        `protobuf:"varint,8,rep,packed,name=Unk3300_KGKFJJHFHAB,json=Unk3300KGKFJJHFHAB,proto3" json:"Unk3300_KGKFJJHFHAB,omitempty"`
	TrainingDungeonProgressDetailList []*FungusTrainingProgressDetail `protobuf:"bytes,3,rep,name=training_dungeon_progress_detail_list,json=trainingDungeonProgressDetailList,proto3" json:"training_dungeon_progress_detail_list,omitempty"`
	Unk3300_HPDOJOBPFMJ               []uint32                        `protobuf:"varint,9,rep,packed,name=Unk3300_HPDOJOBPFMJ,json=Unk3300HPDOJOBPFMJ,proto3" json:"Unk3300_HPDOJOBPFMJ,omitempty"`
	PlotStageDetailList               []*FungusPlotStageDetail        `protobuf:"bytes,11,rep,name=plot_stage_detail_list,json=plotStageDetailList,proto3" json:"plot_stage_detail_list,omitempty"`
	TrainingDungeonDetailList         []*FungusTrainingDungeonDetail  `protobuf:"bytes,2,rep,name=training_dungeon_detail_list,json=trainingDungeonDetailList,proto3" json:"training_dungeon_detail_list,omitempty"`
}

func (x *FungusFighterDetailInfo) Reset() {
	*x = FungusFighterDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FungusFighterDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungusFighterDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungusFighterDetailInfo) ProtoMessage() {}

func (x *FungusFighterDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FungusFighterDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungusFighterDetailInfo.ProtoReflect.Descriptor instead.
func (*FungusFighterDetailInfo) Descriptor() ([]byte, []int) {
	return file_FungusFighterDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FungusFighterDetailInfo) GetFungusDetailList() []*FungusDetail {
	if x != nil {
		return x.FungusDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetUnk3300_GIHAKKAJHDH() []uint32 {
	if x != nil {
		return x.Unk3300_GIHAKKAJHDH
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetUnk3300_KGKFJJHFHAB() []uint32 {
	if x != nil {
		return x.Unk3300_KGKFJJHFHAB
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonProgressDetailList() []*FungusTrainingProgressDetail {
	if x != nil {
		return x.TrainingDungeonProgressDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetUnk3300_HPDOJOBPFMJ() []uint32 {
	if x != nil {
		return x.Unk3300_HPDOJOBPFMJ
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetPlotStageDetailList() []*FungusPlotStageDetail {
	if x != nil {
		return x.PlotStageDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonDetailList() []*FungusTrainingDungeonDetail {
	if x != nil {
		return x.TrainingDungeonDetailList
	}
	return nil
}

var File_FungusFighterDetailInfo_proto protoreflect.FileDescriptor

var file_FungusFighterDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x50, 0x6c, 0x6f, 0x74, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x17, 0x46, 0x75, 0x6e, 0x67,
	0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x12, 0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x10,
	0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x47, 0x49, 0x48, 0x41,
	0x4b, 0x4b, 0x41, 0x4a, 0x48, 0x44, 0x48, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x47, 0x49, 0x48, 0x41, 0x4b, 0x4b, 0x41, 0x4a, 0x48, 0x44,
	0x48, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4b, 0x47, 0x4b,
	0x46, 0x4a, 0x4a, 0x48, 0x46, 0x48, 0x41, 0x42, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4b, 0x47, 0x4b, 0x46, 0x4a, 0x4a, 0x48, 0x46, 0x48,
	0x41, 0x42, 0x12, 0x6f, 0x0a, 0x25, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x21, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x48,
	0x50, 0x44, 0x4f, 0x4a, 0x4f, 0x42, 0x50, 0x46, 0x4d, 0x4a, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x48, 0x50, 0x44, 0x4f, 0x4a, 0x4f, 0x42,
	0x50, 0x46, 0x4d, 0x4a, 0x12, 0x4b, 0x0a, 0x16, 0x70, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x50, 0x6c, 0x6f,
	0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x13, 0x70, 0x6c,
	0x6f, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x5d, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x19, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FungusFighterDetailInfo_proto_rawDescOnce sync.Once
	file_FungusFighterDetailInfo_proto_rawDescData = file_FungusFighterDetailInfo_proto_rawDesc
)

func file_FungusFighterDetailInfo_proto_rawDescGZIP() []byte {
	file_FungusFighterDetailInfo_proto_rawDescOnce.Do(func() {
		file_FungusFighterDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FungusFighterDetailInfo_proto_rawDescData)
	})
	return file_FungusFighterDetailInfo_proto_rawDescData
}

var file_FungusFighterDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FungusFighterDetailInfo_proto_goTypes = []interface{}{
	(*FungusFighterDetailInfo)(nil),      // 0: FungusFighterDetailInfo
	(*FungusDetail)(nil),                 // 1: FungusDetail
	(*FungusTrainingProgressDetail)(nil), // 2: FungusTrainingProgressDetail
	(*FungusPlotStageDetail)(nil),        // 3: FungusPlotStageDetail
	(*FungusTrainingDungeonDetail)(nil),  // 4: FungusTrainingDungeonDetail
}
var file_FungusFighterDetailInfo_proto_depIdxs = []int32{
	1, // 0: FungusFighterDetailInfo.fungus_detail_list:type_name -> FungusDetail
	2, // 1: FungusFighterDetailInfo.training_dungeon_progress_detail_list:type_name -> FungusTrainingProgressDetail
	3, // 2: FungusFighterDetailInfo.plot_stage_detail_list:type_name -> FungusPlotStageDetail
	4, // 3: FungusFighterDetailInfo.training_dungeon_detail_list:type_name -> FungusTrainingDungeonDetail
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_FungusFighterDetailInfo_proto_init() }
func file_FungusFighterDetailInfo_proto_init() {
	if File_FungusFighterDetailInfo_proto != nil {
		return
	}
	file_FungusDetail_proto_init()
	file_FungusPlotStageDetail_proto_init()
	file_FungusTrainingDungeonDetail_proto_init()
	file_FungusTrainingProgressDetail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FungusFighterDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungusFighterDetailInfo); i {
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
			RawDescriptor: file_FungusFighterDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FungusFighterDetailInfo_proto_goTypes,
		DependencyIndexes: file_FungusFighterDetailInfo_proto_depIdxs,
		MessageInfos:      file_FungusFighterDetailInfo_proto_msgTypes,
	}.Build()
	File_FungusFighterDetailInfo_proto = out.File
	file_FungusFighterDetailInfo_proto_rawDesc = nil
	file_FungusFighterDetailInfo_proto_goTypes = nil
	file_FungusFighterDetailInfo_proto_depIdxs = nil
}