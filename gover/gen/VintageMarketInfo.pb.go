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
// source: VintageMarketInfo.proto

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

type VintageMarketInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BargainInfoMap      map[uint32]bool           `protobuf:"bytes,9,rep,name=bargain_info_map,json=bargainInfoMap,proto3" json:"bargain_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Unk3300_FJDIMCJEAOB bool                      `protobuf:"varint,8,opt,name=Unk3300_FJDIMCJEAOB,json=Unk3300FJDIMCJEAOB,proto3" json:"Unk3300_FJDIMCJEAOB,omitempty"`
	Unk3300_COKBOKAOCNJ []uint32                  `protobuf:"varint,1007,rep,packed,name=Unk3300_COKBOKAOCNJ,json=Unk3300COKBOKAOCNJ,proto3" json:"Unk3300_COKBOKAOCNJ,omitempty"`
	HelpSkillId         uint32                    `protobuf:"varint,760,opt,name=help_skill_id,json=helpSkillId,proto3" json:"help_skill_id,omitempty"`
	Unk3300_ABHHACBDCFI []uint32                  `protobuf:"varint,7,rep,packed,name=Unk3300_ABHHACBDCFI,json=Unk3300ABHHACBDCFI,proto3" json:"Unk3300_ABHHACBDCFI,omitempty"`
	OpenStoreList       []*VintageMarketStoreInfo `protobuf:"bytes,2,rep,name=open_store_list,json=openStoreList,proto3" json:"open_store_list,omitempty"`
	Unk3300_EDBJLBICEJF []uint32                  `protobuf:"varint,10,rep,packed,name=Unk3300_EDBJLBICEJF,json=Unk3300EDBJLBICEJF,proto3" json:"Unk3300_EDBJLBICEJF,omitempty"`
	Unk3300_PAMKPAOCJJJ bool                      `protobuf:"varint,11,opt,name=Unk3300_PAMKPAOCJJJ,json=Unk3300PAMKPAOCJJJ,proto3" json:"Unk3300_PAMKPAOCJJJ,omitempty"`
	Unk3300_FOPKAIIAMFM uint32                    `protobuf:"varint,1826,opt,name=Unk3300_FOPKAIIAMFM,json=Unk3300FOPKAIIAMFM,proto3" json:"Unk3300_FOPKAIIAMFM,omitempty"`
	Unk3300_NBNOJJBNJPB bool                      `protobuf:"varint,470,opt,name=Unk3300_NBNOJJBNJPB,json=Unk3300NBNOJJBNJPB,proto3" json:"Unk3300_NBNOJJBNJPB,omitempty"`
	DealInfo            *VintageMarketDealInfo    `protobuf:"bytes,12,opt,name=deal_info,json=dealInfo,proto3" json:"deal_info,omitempty"`
	StoreRound          uint32                    `protobuf:"varint,3,opt,name=store_round,json=storeRound,proto3" json:"store_round,omitempty"`
	UnlockStrategyList  []uint32                  `protobuf:"varint,13,rep,packed,name=unlock_strategy_list,json=unlockStrategyList,proto3" json:"unlock_strategy_list,omitempty"`
	Unk3300_AIAMBBODLBO []uint32                  `protobuf:"varint,14,rep,packed,name=Unk3300_AIAMBBODLBO,json=Unk3300AIAMBBODLBO,proto3" json:"Unk3300_AIAMBBODLBO,omitempty"`
	Unk3300_NDFHAGHBCFN bool                      `protobuf:"varint,6,opt,name=Unk3300_NDFHAGHBCFN,json=Unk3300NDFHAGHBCFN,proto3" json:"Unk3300_NDFHAGHBCFN,omitempty"`
	Unk3300_LMNLHHPMAMN uint32                    `protobuf:"varint,1,opt,name=Unk3300_LMNLHHPMAMN,json=Unk3300LMNLHHPMAMN,proto3" json:"Unk3300_LMNLHHPMAMN,omitempty"`
	Unk3300_ALHLENELIEO bool                      `protobuf:"varint,1594,opt,name=Unk3300_ALHLENELIEO,json=Unk3300ALHLENELIEO,proto3" json:"Unk3300_ALHLENELIEO,omitempty"`
	Unk3300_ACJLANLBCGK uint32                    `protobuf:"varint,1658,opt,name=Unk3300_ACJLANLBCGK,json=Unk3300ACJLANLBCGK,proto3" json:"Unk3300_ACJLANLBCGK,omitempty"`
	Unk3300_HHKKEDDIGLA uint32                    `protobuf:"varint,5,opt,name=Unk3300_HHKKEDDIGLA,json=Unk3300HHKKEDDIGLA,proto3" json:"Unk3300_HHKKEDDIGLA,omitempty"`
	Unk3300_KJDEFIBKBPE bool                      `protobuf:"varint,4,opt,name=Unk3300_KJDEFIBKBPE,json=Unk3300KJDEFIBKBPE,proto3" json:"Unk3300_KJDEFIBKBPE,omitempty"`
	Unk3300_ICJOCDGLFFD bool                      `protobuf:"varint,15,opt,name=Unk3300_ICJOCDGLFFD,json=Unk3300ICJOCDGLFFD,proto3" json:"Unk3300_ICJOCDGLFFD,omitempty"`
	Unk3300_NMHHGNPKBEL bool                      `protobuf:"varint,1721,opt,name=Unk3300_NMHHGNPKBEL,json=Unk3300NMHHGNPKBEL,proto3" json:"Unk3300_NMHHGNPKBEL,omitempty"`
}

func (x *VintageMarketInfo) Reset() {
	*x = VintageMarketInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VintageMarketInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VintageMarketInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VintageMarketInfo) ProtoMessage() {}

func (x *VintageMarketInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VintageMarketInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VintageMarketInfo.ProtoReflect.Descriptor instead.
func (*VintageMarketInfo) Descriptor() ([]byte, []int) {
	return file_VintageMarketInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VintageMarketInfo) GetBargainInfoMap() map[uint32]bool {
	if x != nil {
		return x.BargainInfoMap
	}
	return nil
}

func (x *VintageMarketInfo) GetUnk3300_FJDIMCJEAOB() bool {
	if x != nil {
		return x.Unk3300_FJDIMCJEAOB
	}
	return false
}

func (x *VintageMarketInfo) GetUnk3300_COKBOKAOCNJ() []uint32 {
	if x != nil {
		return x.Unk3300_COKBOKAOCNJ
	}
	return nil
}

func (x *VintageMarketInfo) GetHelpSkillId() uint32 {
	if x != nil {
		return x.HelpSkillId
	}
	return 0
}

func (x *VintageMarketInfo) GetUnk3300_ABHHACBDCFI() []uint32 {
	if x != nil {
		return x.Unk3300_ABHHACBDCFI
	}
	return nil
}

func (x *VintageMarketInfo) GetOpenStoreList() []*VintageMarketStoreInfo {
	if x != nil {
		return x.OpenStoreList
	}
	return nil
}

func (x *VintageMarketInfo) GetUnk3300_EDBJLBICEJF() []uint32 {
	if x != nil {
		return x.Unk3300_EDBJLBICEJF
	}
	return nil
}

func (x *VintageMarketInfo) GetUnk3300_PAMKPAOCJJJ() bool {
	if x != nil {
		return x.Unk3300_PAMKPAOCJJJ
	}
	return false
}

func (x *VintageMarketInfo) GetUnk3300_FOPKAIIAMFM() uint32 {
	if x != nil {
		return x.Unk3300_FOPKAIIAMFM
	}
	return 0
}

func (x *VintageMarketInfo) GetUnk3300_NBNOJJBNJPB() bool {
	if x != nil {
		return x.Unk3300_NBNOJJBNJPB
	}
	return false
}

func (x *VintageMarketInfo) GetDealInfo() *VintageMarketDealInfo {
	if x != nil {
		return x.DealInfo
	}
	return nil
}

func (x *VintageMarketInfo) GetStoreRound() uint32 {
	if x != nil {
		return x.StoreRound
	}
	return 0
}

func (x *VintageMarketInfo) GetUnlockStrategyList() []uint32 {
	if x != nil {
		return x.UnlockStrategyList
	}
	return nil
}

func (x *VintageMarketInfo) GetUnk3300_AIAMBBODLBO() []uint32 {
	if x != nil {
		return x.Unk3300_AIAMBBODLBO
	}
	return nil
}

func (x *VintageMarketInfo) GetUnk3300_NDFHAGHBCFN() bool {
	if x != nil {
		return x.Unk3300_NDFHAGHBCFN
	}
	return false
}

func (x *VintageMarketInfo) GetUnk3300_LMNLHHPMAMN() uint32 {
	if x != nil {
		return x.Unk3300_LMNLHHPMAMN
	}
	return 0
}

func (x *VintageMarketInfo) GetUnk3300_ALHLENELIEO() bool {
	if x != nil {
		return x.Unk3300_ALHLENELIEO
	}
	return false
}

func (x *VintageMarketInfo) GetUnk3300_ACJLANLBCGK() uint32 {
	if x != nil {
		return x.Unk3300_ACJLANLBCGK
	}
	return 0
}

func (x *VintageMarketInfo) GetUnk3300_HHKKEDDIGLA() uint32 {
	if x != nil {
		return x.Unk3300_HHKKEDDIGLA
	}
	return 0
}

func (x *VintageMarketInfo) GetUnk3300_KJDEFIBKBPE() bool {
	if x != nil {
		return x.Unk3300_KJDEFIBKBPE
	}
	return false
}

func (x *VintageMarketInfo) GetUnk3300_ICJOCDGLFFD() bool {
	if x != nil {
		return x.Unk3300_ICJOCDGLFFD
	}
	return false
}

func (x *VintageMarketInfo) GetUnk3300_NMHHGNPKBEL() bool {
	if x != nil {
		return x.Unk3300_NMHHGNPKBEL
	}
	return false
}

var File_VintageMarketInfo_proto protoreflect.FileDescriptor

var file_VintageMarketInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x56, 0x69, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x09, 0x0a, 0x11, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x10, 0x62, 0x61,
	0x72, 0x67, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x62, 0x61,
	0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12, 0x2f, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x46, 0x4a, 0x44, 0x49, 0x4d, 0x43, 0x4a, 0x45,
	0x41, 0x4f, 0x42, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x46, 0x4a, 0x44, 0x49, 0x4d, 0x43, 0x4a, 0x45, 0x41, 0x4f, 0x42, 0x12, 0x30, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x43, 0x4f, 0x4b, 0x42, 0x4f, 0x4b, 0x41,
	0x4f, 0x43, 0x4e, 0x4a, 0x18, 0xef, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x33, 0x33, 0x30, 0x30, 0x43, 0x4f, 0x4b, 0x42, 0x4f, 0x4b, 0x41, 0x4f, 0x43, 0x4e, 0x4a, 0x12,
	0x23, 0x0a, 0x0d, 0x68, 0x65, 0x6c, 0x70, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0xf8, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x68, 0x65, 0x6c, 0x70, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f,
	0x41, 0x42, 0x48, 0x48, 0x41, 0x43, 0x42, 0x44, 0x43, 0x46, 0x49, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x41, 0x42, 0x48, 0x48, 0x41, 0x43,
	0x42, 0x44, 0x43, 0x46, 0x49, 0x12, 0x3f, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x5f, 0x45, 0x44, 0x42, 0x4a, 0x4c, 0x42, 0x49, 0x43, 0x45, 0x4a, 0x46, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x45, 0x44, 0x42, 0x4a,
	0x4c, 0x42, 0x49, 0x43, 0x45, 0x4a, 0x46, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x5f, 0x50, 0x41, 0x4d, 0x4b, 0x50, 0x41, 0x4f, 0x43, 0x4a, 0x4a, 0x4a, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x50, 0x41, 0x4d,
	0x4b, 0x50, 0x41, 0x4f, 0x43, 0x4a, 0x4a, 0x4a, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33,
	0x33, 0x30, 0x30, 0x5f, 0x46, 0x4f, 0x50, 0x4b, 0x41, 0x49, 0x49, 0x41, 0x4d, 0x46, 0x4d, 0x18,
	0xa2, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x46,
	0x4f, 0x50, 0x4b, 0x41, 0x49, 0x49, 0x41, 0x4d, 0x46, 0x4d, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4e, 0x42, 0x4e, 0x4f, 0x4a, 0x4a, 0x42, 0x4e, 0x4a, 0x50,
	0x42, 0x18, 0xd6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x4e, 0x42, 0x4e, 0x4f, 0x4a, 0x4a, 0x42, 0x4e, 0x4a, 0x50, 0x42, 0x12, 0x33, 0x0a, 0x09,
	0x64, 0x65, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44,
	0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f,
	0x41, 0x49, 0x41, 0x4d, 0x42, 0x42, 0x4f, 0x44, 0x4c, 0x42, 0x4f, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x41, 0x49, 0x41, 0x4d, 0x42, 0x42,
	0x4f, 0x44, 0x4c, 0x42, 0x4f, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30,
	0x5f, 0x4e, 0x44, 0x46, 0x48, 0x41, 0x47, 0x48, 0x42, 0x43, 0x46, 0x4e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4e, 0x44, 0x46, 0x48, 0x41,
	0x47, 0x48, 0x42, 0x43, 0x46, 0x4e, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x5f, 0x4c, 0x4d, 0x4e, 0x4c, 0x48, 0x48, 0x50, 0x4d, 0x41, 0x4d, 0x4e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4c, 0x4d, 0x4e, 0x4c,
	0x48, 0x48, 0x50, 0x4d, 0x41, 0x4d, 0x4e, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x5f, 0x41, 0x4c, 0x48, 0x4c, 0x45, 0x4e, 0x45, 0x4c, 0x49, 0x45, 0x4f, 0x18, 0xba,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x41, 0x4c,
	0x48, 0x4c, 0x45, 0x4e, 0x45, 0x4c, 0x49, 0x45, 0x4f, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x33, 0x33, 0x30, 0x30, 0x5f, 0x41, 0x43, 0x4a, 0x4c, 0x41, 0x4e, 0x4c, 0x42, 0x43, 0x47, 0x4b,
	0x18, 0xfa, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30,
	0x41, 0x43, 0x4a, 0x4c, 0x41, 0x4e, 0x4c, 0x42, 0x43, 0x47, 0x4b, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x48, 0x48, 0x4b, 0x4b, 0x45, 0x44, 0x44, 0x49, 0x47,
	0x4c, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x48, 0x48, 0x4b, 0x4b, 0x45, 0x44, 0x44, 0x49, 0x47, 0x4c, 0x41, 0x12, 0x2f, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4b, 0x4a, 0x44, 0x45, 0x46, 0x49, 0x42, 0x4b,
	0x42, 0x50, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x4b, 0x4a, 0x44, 0x45, 0x46, 0x49, 0x42, 0x4b, 0x42, 0x50, 0x45, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x49, 0x43, 0x4a, 0x4f, 0x43, 0x44, 0x47,
	0x4c, 0x46, 0x46, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x33, 0x30, 0x30, 0x49, 0x43, 0x4a, 0x4f, 0x43, 0x44, 0x47, 0x4c, 0x46, 0x46, 0x44, 0x12, 0x30,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4e, 0x4d, 0x48, 0x48, 0x47, 0x4e,
	0x50, 0x4b, 0x42, 0x45, 0x4c, 0x18, 0xb9, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x33, 0x33, 0x30, 0x30, 0x4e, 0x4d, 0x48, 0x48, 0x47, 0x4e, 0x50, 0x4b, 0x42, 0x45, 0x4c,
	0x1a, 0x41, 0x0a, 0x13, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_VintageMarketInfo_proto_rawDescOnce sync.Once
	file_VintageMarketInfo_proto_rawDescData = file_VintageMarketInfo_proto_rawDesc
)

func file_VintageMarketInfo_proto_rawDescGZIP() []byte {
	file_VintageMarketInfo_proto_rawDescOnce.Do(func() {
		file_VintageMarketInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_VintageMarketInfo_proto_rawDescData)
	})
	return file_VintageMarketInfo_proto_rawDescData
}

var file_VintageMarketInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_VintageMarketInfo_proto_goTypes = []interface{}{
	(*VintageMarketInfo)(nil),      // 0: VintageMarketInfo
	nil,                            // 1: VintageMarketInfo.BargainInfoMapEntry
	(*VintageMarketStoreInfo)(nil), // 2: VintageMarketStoreInfo
	(*VintageMarketDealInfo)(nil),  // 3: VintageMarketDealInfo
}
var file_VintageMarketInfo_proto_depIdxs = []int32{
	1, // 0: VintageMarketInfo.bargain_info_map:type_name -> VintageMarketInfo.BargainInfoMapEntry
	2, // 1: VintageMarketInfo.open_store_list:type_name -> VintageMarketStoreInfo
	3, // 2: VintageMarketInfo.deal_info:type_name -> VintageMarketDealInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_VintageMarketInfo_proto_init() }
func file_VintageMarketInfo_proto_init() {
	if File_VintageMarketInfo_proto != nil {
		return
	}
	file_VintageMarketDealInfo_proto_init()
	file_VintageMarketStoreInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_VintageMarketInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VintageMarketInfo); i {
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
			RawDescriptor: file_VintageMarketInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VintageMarketInfo_proto_goTypes,
		DependencyIndexes: file_VintageMarketInfo_proto_depIdxs,
		MessageInfos:      file_VintageMarketInfo_proto_msgTypes,
	}.Build()
	File_VintageMarketInfo_proto = out.File
	file_VintageMarketInfo_proto_rawDesc = nil
	file_VintageMarketInfo_proto_goTypes = nil
	file_VintageMarketInfo_proto_depIdxs = nil
}