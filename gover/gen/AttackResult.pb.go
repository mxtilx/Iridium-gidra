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
// source: AttackResult.proto

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

type AttackResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsResistText                 bool                   `protobuf:"varint,1858,opt,name=is_resist_text,json=isResistText,proto3" json:"is_resist_text,omitempty"`
	Unk2700_GBANCFEPPIM          uint32                 `protobuf:"varint,1011,opt,name=Unk2700_GBANCFEPPIM,json=Unk2700GBANCFEPPIM,proto3" json:"Unk2700_GBANCFEPPIM,omitempty"`
	AmplifyReactionType          uint32                 `protobuf:"varint,2005,opt,name=amplify_reaction_type,json=amplifyReactionType,proto3" json:"amplify_reaction_type,omitempty"`
	EndureBreak                  uint32                 `protobuf:"varint,7,opt,name=endure_break,json=endureBreak,proto3" json:"endure_break,omitempty"`
	ElementType                  uint32                 `protobuf:"varint,5,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	ElementDurabilityAttenuation float32                `protobuf:"fixed32,425,opt,name=element_durability_attenuation,json=elementDurabilityAttenuation,proto3" json:"element_durability_attenuation,omitempty"`
	DefenseId                    uint32                 `protobuf:"varint,15,opt,name=defense_id,json=defenseId,proto3" json:"defense_id,omitempty"`
	AttackTimestampMs            uint32                 `protobuf:"varint,1188,opt,name=attack_timestamp_ms,json=attackTimestampMs,proto3" json:"attack_timestamp_ms,omitempty"`
	BulletFlyTimeMs              uint32                 `protobuf:"varint,91,opt,name=bullet_fly_time_ms,json=bulletFlyTimeMs,proto3" json:"bullet_fly_time_ms,omitempty"`
	IsCrit                       bool                   `protobuf:"varint,13,opt,name=is_crit,json=isCrit,proto3" json:"is_crit,omitempty"`
	ElementAmplifyRate           float32                `protobuf:"fixed32,900,opt,name=element_amplify_rate,json=elementAmplifyRate,proto3" json:"element_amplify_rate,omitempty"`
	AttackCount                  uint32                 `protobuf:"varint,1564,opt,name=attack_count,json=attackCount,proto3" json:"attack_count,omitempty"`
	CriticalRand                 uint32                 `protobuf:"varint,1664,opt,name=critical_rand,json=criticalRand,proto3" json:"critical_rand,omitempty"`
	HitPosType                   uint32                 `protobuf:"varint,2,opt,name=hit_pos_type,json=hitPosType,proto3" json:"hit_pos_type,omitempty"`
	AnimEventId                  string                 `protobuf:"bytes,4,opt,name=anim_event_id,json=animEventId,proto3" json:"anim_event_id,omitempty"`
	HitEffResult                 *AttackHitEffectResult `protobuf:"bytes,8,opt,name=hit_eff_result,json=hitEffResult,proto3" json:"hit_eff_result,omitempty"`
	DamageShield                 float32                `protobuf:"fixed32,1202,opt,name=damage_shield,json=damageShield,proto3" json:"damage_shield,omitempty"`
	EndureDelta                  float32                `protobuf:"fixed32,430,opt,name=endure_delta,json=endureDelta,proto3" json:"endure_delta,omitempty"`
	ResolvedDir                  *Vector                `protobuf:"bytes,1,opt,name=resolved_dir,json=resolvedDir,proto3" json:"resolved_dir,omitempty"`
	Damage                       float32                `protobuf:"fixed32,6,opt,name=damage,proto3" json:"damage,omitempty"`
	AddhurtReactionType          uint32                 `protobuf:"varint,1887,opt,name=addhurt_reaction_type,json=addhurtReactionType,proto3" json:"addhurt_reaction_type,omitempty"`
	HashedAnimEventId            uint32                 `protobuf:"varint,278,opt,name=hashed_anim_event_id,json=hashedAnimEventId,proto3" json:"hashed_anim_event_id,omitempty"`
	UseGadgetDamageAction        bool                   `protobuf:"varint,1418,opt,name=use_gadget_damage_action,json=useGadgetDamageAction,proto3" json:"use_gadget_damage_action,omitempty"`
	HitRetreatAngleCompat        int32                  `protobuf:"varint,9,opt,name=hit_retreat_angle_compat,json=hitRetreatAngleCompat,proto3" json:"hit_retreat_angle_compat,omitempty"`
	AbilityIdentifier            *AbilityIdentifier     `protobuf:"bytes,14,opt,name=ability_identifier,json=abilityIdentifier,proto3" json:"ability_identifier,omitempty"`
	AttackerId                   uint32                 `protobuf:"varint,11,opt,name=attacker_id,json=attackerId,proto3" json:"attacker_id,omitempty"`
	MuteElementHurt              bool                   `protobuf:"varint,1530,opt,name=mute_element_hurt,json=muteElementHurt,proto3" json:"mute_element_hurt,omitempty"`
	TargetType                   uint32                 `protobuf:"varint,1366,opt,name=target_type,json=targetType,proto3" json:"target_type,omitempty"`
	HitCollision                 *HitCollision          `protobuf:"bytes,10,opt,name=hit_collision,json=hitCollision,proto3" json:"hit_collision,omitempty"`
	GadgetDamageActionIdx        uint32                 `protobuf:"varint,1110,opt,name=gadget_damage_action_idx,json=gadgetDamageActionIdx,proto3" json:"gadget_damage_action_idx,omitempty"`
}

func (x *AttackResult) Reset() {
	*x = AttackResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AttackResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackResult) ProtoMessage() {}

func (x *AttackResult) ProtoReflect() protoreflect.Message {
	mi := &file_AttackResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackResult.ProtoReflect.Descriptor instead.
func (*AttackResult) Descriptor() ([]byte, []int) {
	return file_AttackResult_proto_rawDescGZIP(), []int{0}
}

func (x *AttackResult) GetIsResistText() bool {
	if x != nil {
		return x.IsResistText
	}
	return false
}

func (x *AttackResult) GetUnk2700_GBANCFEPPIM() uint32 {
	if x != nil {
		return x.Unk2700_GBANCFEPPIM
	}
	return 0
}

func (x *AttackResult) GetAmplifyReactionType() uint32 {
	if x != nil {
		return x.AmplifyReactionType
	}
	return 0
}

func (x *AttackResult) GetEndureBreak() uint32 {
	if x != nil {
		return x.EndureBreak
	}
	return 0
}

func (x *AttackResult) GetElementType() uint32 {
	if x != nil {
		return x.ElementType
	}
	return 0
}

func (x *AttackResult) GetElementDurabilityAttenuation() float32 {
	if x != nil {
		return x.ElementDurabilityAttenuation
	}
	return 0
}

func (x *AttackResult) GetDefenseId() uint32 {
	if x != nil {
		return x.DefenseId
	}
	return 0
}

func (x *AttackResult) GetAttackTimestampMs() uint32 {
	if x != nil {
		return x.AttackTimestampMs
	}
	return 0
}

func (x *AttackResult) GetBulletFlyTimeMs() uint32 {
	if x != nil {
		return x.BulletFlyTimeMs
	}
	return 0
}

func (x *AttackResult) GetIsCrit() bool {
	if x != nil {
		return x.IsCrit
	}
	return false
}

func (x *AttackResult) GetElementAmplifyRate() float32 {
	if x != nil {
		return x.ElementAmplifyRate
	}
	return 0
}

func (x *AttackResult) GetAttackCount() uint32 {
	if x != nil {
		return x.AttackCount
	}
	return 0
}

func (x *AttackResult) GetCriticalRand() uint32 {
	if x != nil {
		return x.CriticalRand
	}
	return 0
}

func (x *AttackResult) GetHitPosType() uint32 {
	if x != nil {
		return x.HitPosType
	}
	return 0
}

func (x *AttackResult) GetAnimEventId() string {
	if x != nil {
		return x.AnimEventId
	}
	return ""
}

func (x *AttackResult) GetHitEffResult() *AttackHitEffectResult {
	if x != nil {
		return x.HitEffResult
	}
	return nil
}

func (x *AttackResult) GetDamageShield() float32 {
	if x != nil {
		return x.DamageShield
	}
	return 0
}

func (x *AttackResult) GetEndureDelta() float32 {
	if x != nil {
		return x.EndureDelta
	}
	return 0
}

func (x *AttackResult) GetResolvedDir() *Vector {
	if x != nil {
		return x.ResolvedDir
	}
	return nil
}

func (x *AttackResult) GetDamage() float32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *AttackResult) GetAddhurtReactionType() uint32 {
	if x != nil {
		return x.AddhurtReactionType
	}
	return 0
}

func (x *AttackResult) GetHashedAnimEventId() uint32 {
	if x != nil {
		return x.HashedAnimEventId
	}
	return 0
}

func (x *AttackResult) GetUseGadgetDamageAction() bool {
	if x != nil {
		return x.UseGadgetDamageAction
	}
	return false
}

func (x *AttackResult) GetHitRetreatAngleCompat() int32 {
	if x != nil {
		return x.HitRetreatAngleCompat
	}
	return 0
}

func (x *AttackResult) GetAbilityIdentifier() *AbilityIdentifier {
	if x != nil {
		return x.AbilityIdentifier
	}
	return nil
}

func (x *AttackResult) GetAttackerId() uint32 {
	if x != nil {
		return x.AttackerId
	}
	return 0
}

func (x *AttackResult) GetMuteElementHurt() bool {
	if x != nil {
		return x.MuteElementHurt
	}
	return false
}

func (x *AttackResult) GetTargetType() uint32 {
	if x != nil {
		return x.TargetType
	}
	return 0
}

func (x *AttackResult) GetHitCollision() *HitCollision {
	if x != nil {
		return x.HitCollision
	}
	return nil
}

func (x *AttackResult) GetGadgetDamageActionIdx() uint32 {
	if x != nil {
		return x.GadgetDamageActionIdx
	}
	return 0
}

var File_AttackResult_proto protoreflect.FileDescriptor

var file_AttackResult_proto_rawDesc = []byte{
	0x0a, 0x12, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x48, 0x69, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x0a, 0x0a,
	0x0c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0xc2, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x47, 0x42, 0x41, 0x4e, 0x43, 0x46, 0x45, 0x50, 0x50, 0x49, 0x4d, 0x18, 0xf3, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x47, 0x42, 0x41, 0x4e, 0x43,
	0x46, 0x45, 0x50, 0x50, 0x49, 0x4d, 0x12, 0x33, 0x0a, 0x15, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66,
	0x79, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0xd5, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65,
	0x6e, 0x64, 0x75, 0x72, 0x65, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x75, 0x72, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x21,
	0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x45, 0x0a, 0x1e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0xa9, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1c, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74,
	0x65, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x65,
	0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x65,
	0x66, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0xa4,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x62, 0x75, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x66, 0x6c, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x5b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x46, 0x6c, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x63, 0x72, 0x69, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x43, 0x72, 0x69, 0x74, 0x12, 0x31,
	0x0a, 0x14, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66,
	0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x84, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x79, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x9c, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x80, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63,
	0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x68,
	0x69, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x68, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x3c, 0x0a, 0x0e, 0x68, 0x69, 0x74, 0x5f, 0x65, 0x66, 0x66, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x0c, 0x68, 0x69, 0x74, 0x45, 0x66, 0x66, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0xb2, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x75, 0x72, 0x65, 0x5f,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0xae, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x65, 0x6e,
	0x64, 0x75, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x44, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a,
	0x15, 0x61, 0x64, 0x64, 0x68, 0x75, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xdf, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x61,
	0x64, 0x64, 0x68, 0x75, 0x72, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x6e, 0x69,
	0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x96, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x11, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x41, 0x6e, 0x69, 0x6d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x75, 0x73, 0x65, 0x5f, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x8a, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x75, 0x73, 0x65, 0x47, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37,
	0x0a, 0x18, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x61, 0x6e,
	0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x15, 0x68, 0x69, 0x74, 0x52, 0x65, 0x74, 0x72, 0x65, 0x61, 0x74, 0x41, 0x6e, 0x67, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x12, 0x41, 0x0a, 0x12, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x11, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x6d,
	0x75, 0x74, 0x65, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x75, 0x72, 0x74,
	0x18, 0xfa, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6d, 0x75, 0x74, 0x65, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x48, 0x75, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xd6, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x68, 0x69,
	0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x68, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x18, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x78, 0x18, 0xd6, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x15, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x78, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AttackResult_proto_rawDescOnce sync.Once
	file_AttackResult_proto_rawDescData = file_AttackResult_proto_rawDesc
)

func file_AttackResult_proto_rawDescGZIP() []byte {
	file_AttackResult_proto_rawDescOnce.Do(func() {
		file_AttackResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_AttackResult_proto_rawDescData)
	})
	return file_AttackResult_proto_rawDescData
}

var file_AttackResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AttackResult_proto_goTypes = []interface{}{
	(*AttackResult)(nil),          // 0: AttackResult
	(*AttackHitEffectResult)(nil), // 1: AttackHitEffectResult
	(*Vector)(nil),                // 2: Vector
	(*AbilityIdentifier)(nil),     // 3: AbilityIdentifier
	(*HitCollision)(nil),          // 4: HitCollision
}
var file_AttackResult_proto_depIdxs = []int32{
	1, // 0: AttackResult.hit_eff_result:type_name -> AttackHitEffectResult
	2, // 1: AttackResult.resolved_dir:type_name -> Vector
	3, // 2: AttackResult.ability_identifier:type_name -> AbilityIdentifier
	4, // 3: AttackResult.hit_collision:type_name -> HitCollision
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_AttackResult_proto_init() }
func file_AttackResult_proto_init() {
	if File_AttackResult_proto != nil {
		return
	}
	file_AbilityIdentifier_proto_init()
	file_AttackHitEffectResult_proto_init()
	file_HitCollision_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AttackResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackResult); i {
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
			RawDescriptor: file_AttackResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AttackResult_proto_goTypes,
		DependencyIndexes: file_AttackResult_proto_depIdxs,
		MessageInfos:      file_AttackResult_proto_msgTypes,
	}.Build()
	File_AttackResult_proto = out.File
	file_AttackResult_proto_rawDesc = nil
	file_AttackResult_proto_goTypes = nil
	file_AttackResult_proto_depIdxs = nil
}