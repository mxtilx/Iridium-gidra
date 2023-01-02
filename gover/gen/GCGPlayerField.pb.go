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
// source: GCGPlayerField.proto

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

type GCGPlayerField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3300_IKJMGAHCFPM uint32                 `protobuf:"varint,5,opt,name=Unk3300_IKJMGAHCFPM,json=Unk3300IKJMGAHCFPM,proto3" json:"Unk3300_IKJMGAHCFPM,omitempty"`
	ModifyZoneMap       map[uint32]*GCGZone    `protobuf:"bytes,7,rep,name=modify_zone_map,json=modifyZoneMap,proto3" json:"modify_zone_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Unk3300_GGHKFFADEAL uint32                 `protobuf:"varint,731,opt,name=Unk3300_GGHKFFADEAL,json=Unk3300GGHKFFADEAL,proto3" json:"Unk3300_GGHKFFADEAL,omitempty"`
	Unk3300_AOPJIOHMPOF *GCGZone               `protobuf:"bytes,10,opt,name=Unk3300_AOPJIOHMPOF,json=Unk3300AOPJIOHMPOF,proto3" json:"Unk3300_AOPJIOHMPOF,omitempty"`
	Unk3300_FDFPHNDOJML uint32                 `protobuf:"varint,12,opt,name=Unk3300_FDFPHNDOJML,json=Unk3300FDFPHNDOJML,proto3" json:"Unk3300_FDFPHNDOJML,omitempty"`
	Unk3300_IPLMHKCNDLE *GCGZone               `protobuf:"bytes,1,opt,name=Unk3300_IPLMHKCNDLE,json=Unk3300IPLMHKCNDLE,proto3" json:"Unk3300_IPLMHKCNDLE,omitempty"`
	Unk3300_EIHOMDLENMK *GCGZone               `protobuf:"bytes,9,opt,name=Unk3300_EIHOMDLENMK,json=Unk3300EIHOMDLENMK,proto3" json:"Unk3300_EIHOMDLENMK,omitempty"`
	WaitingList         []*GCGWaitingCharacter `protobuf:"bytes,2,rep,name=waiting_list,json=waitingList,proto3" json:"waiting_list,omitempty"`
	Unk3300_PBECINKKHND uint32                 `protobuf:"varint,15,opt,name=Unk3300_PBECINKKHND,json=Unk3300PBECINKKHND,proto3" json:"Unk3300_PBECINKKHND,omitempty"`
	ControllerId        uint32                 `protobuf:"varint,6,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	Unk3300_INDJNJJJNKL *GCGZone               `protobuf:"bytes,11,opt,name=Unk3300_INDJNJJJNKL,json=Unk3300INDJNJJJNKL,proto3" json:"Unk3300_INDJNJJJNKL,omitempty"`
	Unk3300_EFNAEFBECHD *GCGZone               `protobuf:"bytes,4,opt,name=Unk3300_EFNAEFBECHD,json=Unk3300EFNAEFBECHD,proto3" json:"Unk3300_EFNAEFBECHD,omitempty"`
	IsPassed            bool                   `protobuf:"varint,8,opt,name=is_passed,json=isPassed,proto3" json:"is_passed,omitempty"`
	IntentionList       []*GCGPVEIntention     `protobuf:"bytes,304,rep,name=intention_list,json=intentionList,proto3" json:"intention_list,omitempty"`
	DiceSideList        []GCGDiceSideType      `protobuf:"varint,13,rep,packed,name=dice_side_list,json=diceSideList,proto3,enum=GCGDiceSideType" json:"dice_side_list,omitempty"`
	DeckCardNum         uint32                 `protobuf:"varint,3,opt,name=deck_card_num,json=deckCardNum,proto3" json:"deck_card_num,omitempty"`
	Unk3300_GLNIFLOKBPM uint32                 `protobuf:"varint,14,opt,name=Unk3300_GLNIFLOKBPM,json=Unk3300GLNIFLOKBPM,proto3" json:"Unk3300_GLNIFLOKBPM,omitempty"`
}

func (x *GCGPlayerField) Reset() {
	*x = GCGPlayerField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGPlayerField_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGPlayerField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGPlayerField) ProtoMessage() {}

func (x *GCGPlayerField) ProtoReflect() protoreflect.Message {
	mi := &file_GCGPlayerField_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGPlayerField.ProtoReflect.Descriptor instead.
func (*GCGPlayerField) Descriptor() ([]byte, []int) {
	return file_GCGPlayerField_proto_rawDescGZIP(), []int{0}
}

func (x *GCGPlayerField) GetUnk3300_IKJMGAHCFPM() uint32 {
	if x != nil {
		return x.Unk3300_IKJMGAHCFPM
	}
	return 0
}

func (x *GCGPlayerField) GetModifyZoneMap() map[uint32]*GCGZone {
	if x != nil {
		return x.ModifyZoneMap
	}
	return nil
}

func (x *GCGPlayerField) GetUnk3300_GGHKFFADEAL() uint32 {
	if x != nil {
		return x.Unk3300_GGHKFFADEAL
	}
	return 0
}

func (x *GCGPlayerField) GetUnk3300_AOPJIOHMPOF() *GCGZone {
	if x != nil {
		return x.Unk3300_AOPJIOHMPOF
	}
	return nil
}

func (x *GCGPlayerField) GetUnk3300_FDFPHNDOJML() uint32 {
	if x != nil {
		return x.Unk3300_FDFPHNDOJML
	}
	return 0
}

func (x *GCGPlayerField) GetUnk3300_IPLMHKCNDLE() *GCGZone {
	if x != nil {
		return x.Unk3300_IPLMHKCNDLE
	}
	return nil
}

func (x *GCGPlayerField) GetUnk3300_EIHOMDLENMK() *GCGZone {
	if x != nil {
		return x.Unk3300_EIHOMDLENMK
	}
	return nil
}

func (x *GCGPlayerField) GetWaitingList() []*GCGWaitingCharacter {
	if x != nil {
		return x.WaitingList
	}
	return nil
}

func (x *GCGPlayerField) GetUnk3300_PBECINKKHND() uint32 {
	if x != nil {
		return x.Unk3300_PBECINKKHND
	}
	return 0
}

func (x *GCGPlayerField) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGPlayerField) GetUnk3300_INDJNJJJNKL() *GCGZone {
	if x != nil {
		return x.Unk3300_INDJNJJJNKL
	}
	return nil
}

func (x *GCGPlayerField) GetUnk3300_EFNAEFBECHD() *GCGZone {
	if x != nil {
		return x.Unk3300_EFNAEFBECHD
	}
	return nil
}

func (x *GCGPlayerField) GetIsPassed() bool {
	if x != nil {
		return x.IsPassed
	}
	return false
}

func (x *GCGPlayerField) GetIntentionList() []*GCGPVEIntention {
	if x != nil {
		return x.IntentionList
	}
	return nil
}

func (x *GCGPlayerField) GetDiceSideList() []GCGDiceSideType {
	if x != nil {
		return x.DiceSideList
	}
	return nil
}

func (x *GCGPlayerField) GetDeckCardNum() uint32 {
	if x != nil {
		return x.DeckCardNum
	}
	return 0
}

func (x *GCGPlayerField) GetUnk3300_GLNIFLOKBPM() uint32 {
	if x != nil {
		return x.Unk3300_GLNIFLOKBPM
	}
	return 0
}

var File_GCGPlayerField_proto protoreflect.FileDescriptor

var file_GCGPlayerField_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x43, 0x47, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x47, 0x43, 0x47, 0x44, 0x69, 0x63, 0x65, 0x53,
	0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x47,
	0x43, 0x47, 0x50, 0x56, 0x45, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x47, 0x43, 0x47, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x07, 0x0a, 0x0e, 0x47, 0x43, 0x47, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x49, 0x4b, 0x4a,
	0x4d, 0x47, 0x41, 0x48, 0x43, 0x46, 0x50, 0x4d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x49, 0x4b, 0x4a, 0x4d, 0x47, 0x41, 0x48, 0x43, 0x46,
	0x50, 0x4d, 0x12, 0x4a, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x47, 0x43,
	0x47, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x30,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x47, 0x47, 0x48, 0x4b, 0x46, 0x46,
	0x41, 0x44, 0x45, 0x41, 0x4c, 0x18, 0xdb, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x33, 0x33, 0x30, 0x30, 0x47, 0x47, 0x48, 0x4b, 0x46, 0x46, 0x41, 0x44, 0x45, 0x41, 0x4c,
	0x12, 0x39, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x41, 0x4f, 0x50, 0x4a,
	0x49, 0x4f, 0x48, 0x4d, 0x50, 0x4f, 0x46, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30,
	0x41, 0x4f, 0x50, 0x4a, 0x49, 0x4f, 0x48, 0x4d, 0x50, 0x4f, 0x46, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x46, 0x44, 0x46, 0x50, 0x48, 0x4e, 0x44, 0x4f, 0x4a,
	0x4d, 0x4c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x46, 0x44, 0x46, 0x50, 0x48, 0x4e, 0x44, 0x4f, 0x4a, 0x4d, 0x4c, 0x12, 0x39, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x49, 0x50, 0x4c, 0x4d, 0x48, 0x4b, 0x43, 0x4e,
	0x44, 0x4c, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x49, 0x50, 0x4c, 0x4d,
	0x48, 0x4b, 0x43, 0x4e, 0x44, 0x4c, 0x45, 0x12, 0x39, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x5f, 0x45, 0x49, 0x48, 0x4f, 0x4d, 0x44, 0x4c, 0x45, 0x4e, 0x4d, 0x4b, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x45, 0x49, 0x48, 0x4f, 0x4d, 0x44, 0x4c, 0x45, 0x4e,
	0x4d, 0x4b, 0x12, 0x37, 0x0a, 0x0c, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x47, 0x43, 0x47, 0x57, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0b,
	0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x50, 0x42, 0x45, 0x43, 0x49, 0x4e, 0x4b, 0x4b, 0x48,
	0x4e, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x50, 0x42, 0x45, 0x43, 0x49, 0x4e, 0x4b, 0x4b, 0x48, 0x4e, 0x44, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x49, 0x4e, 0x44,
	0x4a, 0x4e, 0x4a, 0x4a, 0x4a, 0x4e, 0x4b, 0x4c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x49, 0x4e, 0x44, 0x4a, 0x4e, 0x4a, 0x4a, 0x4a, 0x4e, 0x4b, 0x4c, 0x12, 0x39, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x45, 0x46, 0x4e, 0x41, 0x45, 0x46, 0x42, 0x45,
	0x43, 0x48, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x45, 0x46, 0x4e, 0x41,
	0x45, 0x46, 0x42, 0x45, 0x43, 0x48, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x61,
	0x73, 0x73, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xb0, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x47, 0x43, 0x47, 0x50, 0x56, 0x45, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x0e, 0x64, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x47, 0x43, 0x47, 0x44, 0x69, 0x63, 0x65,
	0x53, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x64, 0x69, 0x63, 0x65, 0x53, 0x69,
	0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64,
	0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x47, 0x4c, 0x4e, 0x49, 0x46, 0x4c, 0x4f, 0x4b, 0x42, 0x50,
	0x4d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30,
	0x47, 0x4c, 0x4e, 0x49, 0x46, 0x4c, 0x4f, 0x4b, 0x42, 0x50, 0x4d, 0x1a, 0x4a, 0x0a, 0x12, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGPlayerField_proto_rawDescOnce sync.Once
	file_GCGPlayerField_proto_rawDescData = file_GCGPlayerField_proto_rawDesc
)

func file_GCGPlayerField_proto_rawDescGZIP() []byte {
	file_GCGPlayerField_proto_rawDescOnce.Do(func() {
		file_GCGPlayerField_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGPlayerField_proto_rawDescData)
	})
	return file_GCGPlayerField_proto_rawDescData
}

var file_GCGPlayerField_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GCGPlayerField_proto_goTypes = []interface{}{
	(*GCGPlayerField)(nil),      // 0: GCGPlayerField
	nil,                         // 1: GCGPlayerField.ModifyZoneMapEntry
	(*GCGZone)(nil),             // 2: GCGZone
	(*GCGWaitingCharacter)(nil), // 3: GCGWaitingCharacter
	(*GCGPVEIntention)(nil),     // 4: GCGPVEIntention
	(GCGDiceSideType)(0),        // 5: GCGDiceSideType
}
var file_GCGPlayerField_proto_depIdxs = []int32{
	1,  // 0: GCGPlayerField.modify_zone_map:type_name -> GCGPlayerField.ModifyZoneMapEntry
	2,  // 1: GCGPlayerField.Unk3300_AOPJIOHMPOF:type_name -> GCGZone
	2,  // 2: GCGPlayerField.Unk3300_IPLMHKCNDLE:type_name -> GCGZone
	2,  // 3: GCGPlayerField.Unk3300_EIHOMDLENMK:type_name -> GCGZone
	3,  // 4: GCGPlayerField.waiting_list:type_name -> GCGWaitingCharacter
	2,  // 5: GCGPlayerField.Unk3300_INDJNJJJNKL:type_name -> GCGZone
	2,  // 6: GCGPlayerField.Unk3300_EFNAEFBECHD:type_name -> GCGZone
	4,  // 7: GCGPlayerField.intention_list:type_name -> GCGPVEIntention
	5,  // 8: GCGPlayerField.dice_side_list:type_name -> GCGDiceSideType
	2,  // 9: GCGPlayerField.ModifyZoneMapEntry.value:type_name -> GCGZone
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_GCGPlayerField_proto_init() }
func file_GCGPlayerField_proto_init() {
	if File_GCGPlayerField_proto != nil {
		return
	}
	file_GCGDiceSideType_proto_init()
	file_GCGPVEIntention_proto_init()
	file_GCGWaitingCharacter_proto_init()
	file_GCGZone_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGPlayerField_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGPlayerField); i {
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
			RawDescriptor: file_GCGPlayerField_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGPlayerField_proto_goTypes,
		DependencyIndexes: file_GCGPlayerField_proto_depIdxs,
		MessageInfos:      file_GCGPlayerField_proto_msgTypes,
	}.Build()
	File_GCGPlayerField_proto = out.File
	file_GCGPlayerField_proto_rawDesc = nil
	file_GCGPlayerField_proto_goTypes = nil
	file_GCGPlayerField_proto_depIdxs = nil
}