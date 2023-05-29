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
// source: EGEBOGIIBGB.proto

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

type EGEBOGIIBGB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KKHMFNJJMHB      uint32         `protobuf:"varint,15,opt,name=KKHMFNJJMHB,proto3" json:"KKHMFNJJMHB,omitempty"`
	DungeonGuid      uint64         `protobuf:"varint,11,opt,name=dungeon_guid,json=dungeonGuid,proto3" json:"dungeon_guid,omitempty"`
	NLDAPLLDGFO      uint32         `protobuf:"varint,1751,opt,name=NLDAPLLDGFO,proto3" json:"NLDAPLLDGFO,omitempty"`
	ShareCode        string         `protobuf:"bytes,8,opt,name=share_code,json=shareCode,proto3" json:"share_code,omitempty"`
	FBPHOMIGIAB      *FHFAAHDDDOE   `protobuf:"bytes,5,opt,name=FBPHOMIGIAB,proto3" json:"FBPHOMIGIAB,omitempty"`
	IsPsnPlatform    bool           `protobuf:"varint,489,opt,name=is_psn_platform,json=isPsnPlatform,proto3" json:"is_psn_platform,omitempty"`
	JFGIJOALDLD      []*DEKEKNMCKPB `protobuf:"bytes,12,rep,name=JFGIJOALDLD,proto3" json:"JFGIJOALDLD,omitempty"`
	DFFHACPAKKN      bool           `protobuf:"varint,14,opt,name=DFFHACPAKKN,proto3" json:"DFFHACPAKKN,omitempty"`
	DungeonId        uint32         `protobuf:"varint,7,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	BCGJBPBHCIH      uint32         `protobuf:"varint,1,opt,name=BCGJBPBHCIH,proto3" json:"BCGJBPBHCIH,omitempty"`
	GEHEFABOGJD      uint32         `protobuf:"varint,9,opt,name=GEHEFABOGJD,proto3" json:"GEHEFABOGJD,omitempty"`
	CreatorNickname  string         `protobuf:"bytes,3,opt,name=creator_nickname,json=creatorNickname,proto3" json:"creator_nickname,omitempty"`
	TagList          []uint32       `protobuf:"varint,6,rep,packed,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	BPPDFFHHANO      *AFCLGCNDDBH   `protobuf:"bytes,10,opt,name=BPPDFFHHANO,proto3" json:"BPPDFFHHANO,omitempty"`
	DNOFGBKGANP      bool           `protobuf:"varint,2,opt,name=DNOFGBKGANP,proto3" json:"DNOFGBKGANP,omitempty"`
	HLDJKLKMMCC      GLNOHEBLMJP    `protobuf:"varint,13,opt,name=HLDJKLKMMCC,proto3,enum=GLNOHEBLMJP" json:"HLDJKLKMMCC,omitempty"`
	FirstPublishTime uint32         `protobuf:"varint,4,opt,name=first_publish_time,json=firstPublishTime,proto3" json:"first_publish_time,omitempty"`
	HHBNJPAEHMO      bool           `protobuf:"varint,1188,opt,name=HHBNJPAEHMO,proto3" json:"HHBNJPAEHMO,omitempty"`
}

func (x *EGEBOGIIBGB) Reset() {
	*x = EGEBOGIIBGB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EGEBOGIIBGB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EGEBOGIIBGB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EGEBOGIIBGB) ProtoMessage() {}

func (x *EGEBOGIIBGB) ProtoReflect() protoreflect.Message {
	mi := &file_EGEBOGIIBGB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EGEBOGIIBGB.ProtoReflect.Descriptor instead.
func (*EGEBOGIIBGB) Descriptor() ([]byte, []int) {
	return file_EGEBOGIIBGB_proto_rawDescGZIP(), []int{0}
}

func (x *EGEBOGIIBGB) GetKKHMFNJJMHB() uint32 {
	if x != nil {
		return x.KKHMFNJJMHB
	}
	return 0
}

func (x *EGEBOGIIBGB) GetDungeonGuid() uint64 {
	if x != nil {
		return x.DungeonGuid
	}
	return 0
}

func (x *EGEBOGIIBGB) GetNLDAPLLDGFO() uint32 {
	if x != nil {
		return x.NLDAPLLDGFO
	}
	return 0
}

func (x *EGEBOGIIBGB) GetShareCode() string {
	if x != nil {
		return x.ShareCode
	}
	return ""
}

func (x *EGEBOGIIBGB) GetFBPHOMIGIAB() *FHFAAHDDDOE {
	if x != nil {
		return x.FBPHOMIGIAB
	}
	return nil
}

func (x *EGEBOGIIBGB) GetIsPsnPlatform() bool {
	if x != nil {
		return x.IsPsnPlatform
	}
	return false
}

func (x *EGEBOGIIBGB) GetJFGIJOALDLD() []*DEKEKNMCKPB {
	if x != nil {
		return x.JFGIJOALDLD
	}
	return nil
}

func (x *EGEBOGIIBGB) GetDFFHACPAKKN() bool {
	if x != nil {
		return x.DFFHACPAKKN
	}
	return false
}

func (x *EGEBOGIIBGB) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *EGEBOGIIBGB) GetBCGJBPBHCIH() uint32 {
	if x != nil {
		return x.BCGJBPBHCIH
	}
	return 0
}

func (x *EGEBOGIIBGB) GetGEHEFABOGJD() uint32 {
	if x != nil {
		return x.GEHEFABOGJD
	}
	return 0
}

func (x *EGEBOGIIBGB) GetCreatorNickname() string {
	if x != nil {
		return x.CreatorNickname
	}
	return ""
}

func (x *EGEBOGIIBGB) GetTagList() []uint32 {
	if x != nil {
		return x.TagList
	}
	return nil
}

func (x *EGEBOGIIBGB) GetBPPDFFHHANO() *AFCLGCNDDBH {
	if x != nil {
		return x.BPPDFFHHANO
	}
	return nil
}

func (x *EGEBOGIIBGB) GetDNOFGBKGANP() bool {
	if x != nil {
		return x.DNOFGBKGANP
	}
	return false
}

func (x *EGEBOGIIBGB) GetHLDJKLKMMCC() GLNOHEBLMJP {
	if x != nil {
		return x.HLDJKLKMMCC
	}
	return GLNOHEBLMJP_GLNOHEBLMJP_EditUgcDungeonEdit
}

func (x *EGEBOGIIBGB) GetFirstPublishTime() uint32 {
	if x != nil {
		return x.FirstPublishTime
	}
	return 0
}

func (x *EGEBOGIIBGB) GetHHBNJPAEHMO() bool {
	if x != nil {
		return x.HHBNJPAEHMO
	}
	return false
}

var File_EGEBOGIIBGB_proto protoreflect.FileDescriptor

var file_EGEBOGIIBGB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x47, 0x45, 0x42, 0x4f, 0x47, 0x49, 0x49, 0x42, 0x47, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x48, 0x46, 0x41, 0x41, 0x48, 0x44, 0x44, 0x44, 0x4f, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x45, 0x4b, 0x45, 0x4b, 0x4e, 0x4d, 0x43,
	0x4b, 0x50, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x46, 0x43, 0x4c, 0x47,
	0x43, 0x4e, 0x44, 0x44, 0x42, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4c,
	0x4e, 0x4f, 0x48, 0x45, 0x42, 0x4c, 0x4d, 0x4a, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbb, 0x05, 0x0a, 0x0b, 0x45, 0x47, 0x45, 0x42, 0x4f, 0x47, 0x49, 0x49, 0x42, 0x47, 0x42, 0x12,
	0x20, 0x0a, 0x0b, 0x4b, 0x4b, 0x48, 0x4d, 0x46, 0x4e, 0x4a, 0x4a, 0x4d, 0x48, 0x42, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x4b, 0x48, 0x4d, 0x46, 0x4e, 0x4a, 0x4a, 0x4d, 0x48,
	0x42, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x47, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x4e, 0x4c, 0x44, 0x41, 0x50, 0x4c, 0x4c, 0x44,
	0x47, 0x46, 0x4f, 0x18, 0xd7, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x4c, 0x44, 0x41,
	0x50, 0x4c, 0x4c, 0x44, 0x47, 0x46, 0x4f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x42, 0x50, 0x48, 0x4f, 0x4d,
	0x49, 0x47, 0x49, 0x41, 0x42, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x48,
	0x46, 0x41, 0x41, 0x48, 0x44, 0x44, 0x44, 0x4f, 0x45, 0x52, 0x0b, 0x46, 0x42, 0x50, 0x48, 0x4f,
	0x4d, 0x49, 0x47, 0x49, 0x41, 0x42, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x70, 0x73, 0x6e,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0xe9, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x50, 0x73, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x2e, 0x0a, 0x0b, 0x4a, 0x46, 0x47, 0x49, 0x4a, 0x4f, 0x41, 0x4c, 0x44, 0x4c, 0x44, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x45, 0x4b, 0x45, 0x4b, 0x4e, 0x4d, 0x43, 0x4b,
	0x50, 0x42, 0x52, 0x0b, 0x4a, 0x46, 0x47, 0x49, 0x4a, 0x4f, 0x41, 0x4c, 0x44, 0x4c, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x46, 0x46, 0x48, 0x41, 0x43, 0x50, 0x41, 0x4b, 0x4b, 0x4e, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x46, 0x46, 0x48, 0x41, 0x43, 0x50, 0x41, 0x4b, 0x4b,
	0x4e, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x42, 0x43, 0x47, 0x4a, 0x42, 0x50, 0x42, 0x48, 0x43, 0x49, 0x48, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x43, 0x47, 0x4a, 0x42, 0x50, 0x42, 0x48, 0x43,
	0x49, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x45, 0x48, 0x45, 0x46, 0x41, 0x42, 0x4f, 0x47, 0x4a,
	0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x45, 0x48, 0x45, 0x46, 0x41, 0x42,
	0x4f, 0x47, 0x4a, 0x44, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x42, 0x50,
	0x50, 0x44, 0x46, 0x46, 0x48, 0x48, 0x41, 0x4e, 0x4f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x41, 0x46, 0x43, 0x4c, 0x47, 0x43, 0x4e, 0x44, 0x44, 0x42, 0x48, 0x52, 0x0b, 0x42,
	0x50, 0x50, 0x44, 0x46, 0x46, 0x48, 0x48, 0x41, 0x4e, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4e,
	0x4f, 0x46, 0x47, 0x42, 0x4b, 0x47, 0x41, 0x4e, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x44, 0x4e, 0x4f, 0x46, 0x47, 0x42, 0x4b, 0x47, 0x41, 0x4e, 0x50, 0x12, 0x2e, 0x0a, 0x0b,
	0x48, 0x4c, 0x44, 0x4a, 0x4b, 0x4c, 0x4b, 0x4d, 0x4d, 0x43, 0x43, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x47, 0x4c, 0x4e, 0x4f, 0x48, 0x45, 0x42, 0x4c, 0x4d, 0x4a, 0x50, 0x52,
	0x0b, 0x48, 0x4c, 0x44, 0x4a, 0x4b, 0x4c, 0x4b, 0x4d, 0x4d, 0x43, 0x43, 0x12, 0x2c, 0x0a, 0x12,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x48, 0x48,
	0x42, 0x4e, 0x4a, 0x50, 0x41, 0x45, 0x48, 0x4d, 0x4f, 0x18, 0xa4, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x48, 0x48, 0x42, 0x4e, 0x4a, 0x50, 0x41, 0x45, 0x48, 0x4d, 0x4f, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EGEBOGIIBGB_proto_rawDescOnce sync.Once
	file_EGEBOGIIBGB_proto_rawDescData = file_EGEBOGIIBGB_proto_rawDesc
)

func file_EGEBOGIIBGB_proto_rawDescGZIP() []byte {
	file_EGEBOGIIBGB_proto_rawDescOnce.Do(func() {
		file_EGEBOGIIBGB_proto_rawDescData = protoimpl.X.CompressGZIP(file_EGEBOGIIBGB_proto_rawDescData)
	})
	return file_EGEBOGIIBGB_proto_rawDescData
}

var file_EGEBOGIIBGB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EGEBOGIIBGB_proto_goTypes = []interface{}{
	(*EGEBOGIIBGB)(nil), // 0: EGEBOGIIBGB
	(*FHFAAHDDDOE)(nil), // 1: FHFAAHDDDOE
	(*DEKEKNMCKPB)(nil), // 2: DEKEKNMCKPB
	(*AFCLGCNDDBH)(nil), // 3: AFCLGCNDDBH
	(GLNOHEBLMJP)(0),    // 4: GLNOHEBLMJP
}
var file_EGEBOGIIBGB_proto_depIdxs = []int32{
	1, // 0: EGEBOGIIBGB.FBPHOMIGIAB:type_name -> FHFAAHDDDOE
	2, // 1: EGEBOGIIBGB.JFGIJOALDLD:type_name -> DEKEKNMCKPB
	3, // 2: EGEBOGIIBGB.BPPDFFHHANO:type_name -> AFCLGCNDDBH
	4, // 3: EGEBOGIIBGB.HLDJKLKMMCC:type_name -> GLNOHEBLMJP
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_EGEBOGIIBGB_proto_init() }
func file_EGEBOGIIBGB_proto_init() {
	if File_EGEBOGIIBGB_proto != nil {
		return
	}
	file_FHFAAHDDDOE_proto_init()
	file_DEKEKNMCKPB_proto_init()
	file_AFCLGCNDDBH_proto_init()
	file_GLNOHEBLMJP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EGEBOGIIBGB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EGEBOGIIBGB); i {
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
			RawDescriptor: file_EGEBOGIIBGB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EGEBOGIIBGB_proto_goTypes,
		DependencyIndexes: file_EGEBOGIIBGB_proto_depIdxs,
		MessageInfos:      file_EGEBOGIIBGB_proto_msgTypes,
	}.Build()
	File_EGEBOGIIBGB_proto = out.File
	file_EGEBOGIIBGB_proto_rawDesc = nil
	file_EGEBOGIIBGB_proto_goTypes = nil
	file_EGEBOGIIBGB_proto_depIdxs = nil
}