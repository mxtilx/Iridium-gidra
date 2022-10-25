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
// source: TreasureMapCurrencyNotify.proto

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

// CmdId: 2171
// EnetChannelId: 0
// EnetIsReliable: true
type TreasureMapCurrencyNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyNum uint32 `protobuf:"varint,8,opt,name=currency_num,json=currencyNum,proto3" json:"currency_num,omitempty"`
}

func (x *TreasureMapCurrencyNotify) Reset() {
	*x = TreasureMapCurrencyNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureMapCurrencyNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureMapCurrencyNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureMapCurrencyNotify) ProtoMessage() {}

func (x *TreasureMapCurrencyNotify) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureMapCurrencyNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureMapCurrencyNotify.ProtoReflect.Descriptor instead.
func (*TreasureMapCurrencyNotify) Descriptor() ([]byte, []int) {
	return file_TreasureMapCurrencyNotify_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureMapCurrencyNotify) GetCurrencyNum() uint32 {
	if x != nil {
		return x.CurrencyNum
	}
	return 0
}

var File_TreasureMapCurrencyNotify_proto protoreflect.FileDescriptor

var file_TreasureMapCurrencyNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3e, 0x0a, 0x19, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x75,
	0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TreasureMapCurrencyNotify_proto_rawDescOnce sync.Once
	file_TreasureMapCurrencyNotify_proto_rawDescData = file_TreasureMapCurrencyNotify_proto_rawDesc
)

func file_TreasureMapCurrencyNotify_proto_rawDescGZIP() []byte {
	file_TreasureMapCurrencyNotify_proto_rawDescOnce.Do(func() {
		file_TreasureMapCurrencyNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureMapCurrencyNotify_proto_rawDescData)
	})
	return file_TreasureMapCurrencyNotify_proto_rawDescData
}

var file_TreasureMapCurrencyNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureMapCurrencyNotify_proto_goTypes = []interface{}{
	(*TreasureMapCurrencyNotify)(nil), // 0: TreasureMapCurrencyNotify
}
var file_TreasureMapCurrencyNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TreasureMapCurrencyNotify_proto_init() }
func file_TreasureMapCurrencyNotify_proto_init() {
	if File_TreasureMapCurrencyNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TreasureMapCurrencyNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureMapCurrencyNotify); i {
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
			RawDescriptor: file_TreasureMapCurrencyNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureMapCurrencyNotify_proto_goTypes,
		DependencyIndexes: file_TreasureMapCurrencyNotify_proto_depIdxs,
		MessageInfos:      file_TreasureMapCurrencyNotify_proto_msgTypes,
	}.Build()
	File_TreasureMapCurrencyNotify_proto = out.File
	file_TreasureMapCurrencyNotify_proto_rawDesc = nil
	file_TreasureMapCurrencyNotify_proto_goTypes = nil
	file_TreasureMapCurrencyNotify_proto_depIdxs = nil
}