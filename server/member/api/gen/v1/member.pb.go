// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.9.2
// source: member.proto

package memberpb

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

type MemberStatus int32

const (
	MemberStatus_DISABLED MemberStatus = 0
	MemberStatus_ENABLED  MemberStatus = 1
)

// Enum value maps for MemberStatus.
var (
	MemberStatus_name = map[int32]string{
		0: "DISABLED",
		1: "ENABLED",
	}
	MemberStatus_value = map[string]int32{
		"DISABLED": 0,
		"ENABLED":  1,
	}
)

func (x MemberStatus) Enum() *MemberStatus {
	p := new(MemberStatus)
	*p = x
	return p
}

func (x MemberStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_member_proto_enumTypes[0].Descriptor()
}

func (MemberStatus) Type() protoreflect.EnumType {
	return &file_member_proto_enumTypes[0]
}

func (x MemberStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberStatus.Descriptor instead.
func (MemberStatus) EnumDescriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{0}
}

type MemberGender int32

const (
	MemberGender_MAN   MemberGender = 0
	MemberGender_WOMAN MemberGender = 1
)

// Enum value maps for MemberGender.
var (
	MemberGender_name = map[int32]string{
		0: "MAN",
		1: "WOMAN",
	}
	MemberGender_value = map[string]int32{
		"MAN":   0,
		"WOMAN": 1,
	}
)

func (x MemberGender) Enum() *MemberGender {
	p := new(MemberGender)
	*p = x
	return p
}

func (x MemberGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberGender) Descriptor() protoreflect.EnumDescriptor {
	return file_member_proto_enumTypes[1].Descriptor()
}

func (MemberGender) Type() protoreflect.EnumType {
	return &file_member_proto_enumTypes[1]
}

func (x MemberGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberGender.Descriptor instead.
func (MemberGender) EnumDescriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{1}
}

type GetMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetMemberRequest) Reset() {
	*x = GetMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberRequest) ProtoMessage() {}

func (x *GetMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberRequest.ProtoReflect.Descriptor instead.
func (*GetMemberRequest) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{0}
}

func (x *GetMemberRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetMemberRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type MemberEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Member *Member `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *MemberEntity) Reset() {
	*x = MemberEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberEntity) ProtoMessage() {}

func (x *MemberEntity) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberEntity.ProtoReflect.Descriptor instead.
func (*MemberEntity) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{1}
}

func (x *MemberEntity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberEntity) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberLevelId int64        `protobuf:"varint,1,opt,name=member_level_id,json=memberLevelId,proto3" json:"member_level_id,omitempty"`
	Username      string       `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string       `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Nickname      string       `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Phone         string       `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Icon          string       `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	Status        MemberStatus `protobuf:"varint,7,opt,name=status,proto3,enum=member.v1.MemberStatus" json:"status,omitempty"`
	Gender        MemberGender `protobuf:"varint,8,opt,name=gender,proto3,enum=member.v1.MemberGender" json:"gender,omitempty"`
	Birthday      string       `protobuf:"bytes,9,opt,name=birthday,proto3" json:"birthday,omitempty"`
	City          string       `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	Job           string       `protobuf:"bytes,11,opt,name=job,proto3" json:"job,omitempty"`
	Growth        int32        `protobuf:"varint,12,opt,name=growth,proto3" json:"growth,omitempty"`
	CratedAt      string       `protobuf:"bytes,13,opt,name=crated_at,json=cratedAt,proto3" json:"crated_at,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{2}
}

func (x *Member) GetMemberLevelId() int64 {
	if x != nil {
		return x.MemberLevelId
	}
	return 0
}

func (x *Member) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Member) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Member) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Member) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Member) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Member) GetStatus() MemberStatus {
	if x != nil {
		return x.Status
	}
	return MemberStatus_DISABLED
}

func (x *Member) GetGender() MemberGender {
	if x != nil {
		return x.Gender
	}
	return MemberGender_MAN
}

func (x *Member) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *Member) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Member) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Member) GetGrowth() int32 {
	if x != nil {
		return x.Growth
	}
	return 0
}

func (x *Member) GetCratedAt() string {
	if x != nil {
		return x.CratedAt
	}
	return ""
}

type BasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *BasicInfo) Reset() {
	*x = BasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicInfo) ProtoMessage() {}

func (x *BasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicInfo.ProtoReflect.Descriptor instead.
func (*BasicInfo) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{3}
}

func (x *BasicInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BasicInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BasicInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member   *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	Password string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRequest) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *CreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_member_proto protoreflect.FileDescriptor

var file_member_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0x49, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x29, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x87, 0x03, 0x0a, 0x06, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6a, 0x6f, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x53, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x2a, 0x29, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x22, 0x0a, 0x0c,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03,
	0x4d, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x01,
	0x32, 0x8c, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_member_proto_rawDescOnce sync.Once
	file_member_proto_rawDescData = file_member_proto_rawDesc
)

func file_member_proto_rawDescGZIP() []byte {
	file_member_proto_rawDescOnce.Do(func() {
		file_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_proto_rawDescData)
	})
	return file_member_proto_rawDescData
}

var file_member_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_member_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_member_proto_goTypes = []interface{}{
	(MemberStatus)(0),        // 0: member.v1.MemberStatus
	(MemberGender)(0),        // 1: member.v1.MemberGender
	(*GetMemberRequest)(nil), // 2: member.v1.GetMemberRequest
	(*MemberEntity)(nil),     // 3: member.v1.MemberEntity
	(*Member)(nil),           // 4: member.v1.Member
	(*BasicInfo)(nil),        // 5: member.v1.BasicInfo
	(*CreateRequest)(nil),    // 6: member.v1.CreateRequest
}
var file_member_proto_depIdxs = []int32{
	4, // 0: member.v1.MemberEntity.member:type_name -> member.v1.Member
	0, // 1: member.v1.Member.status:type_name -> member.v1.MemberStatus
	1, // 2: member.v1.Member.gender:type_name -> member.v1.MemberGender
	4, // 3: member.v1.CreateRequest.member:type_name -> member.v1.Member
	2, // 4: member.v1.MemberService.GetMember:input_type -> member.v1.GetMemberRequest
	6, // 5: member.v1.MemberService.CreateMember:input_type -> member.v1.CreateRequest
	4, // 6: member.v1.MemberService.GetMember:output_type -> member.v1.Member
	5, // 7: member.v1.MemberService.CreateMember:output_type -> member.v1.BasicInfo
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_member_proto_init() }
func file_member_proto_init() {
	if File_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberRequest); i {
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
		file_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberEntity); i {
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
		file_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicInfo); i {
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
		file_member_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
			RawDescriptor: file_member_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_member_proto_goTypes,
		DependencyIndexes: file_member_proto_depIdxs,
		EnumInfos:         file_member_proto_enumTypes,
		MessageInfos:      file_member_proto_msgTypes,
	}.Build()
	File_member_proto = out.File
	file_member_proto_rawDesc = nil
	file_member_proto_goTypes = nil
	file_member_proto_depIdxs = nil
}