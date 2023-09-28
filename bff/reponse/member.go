package reponse

import "github.com/golang/protobuf/ptypes/timestamp"

type MemberResponse struct {
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberLevelId uint64 `protobuf:"varint,2,opt,name=member_level_id,json=memberLevelId,proto3" json:"member_level_id,omitempty"`
	Username      string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password      string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Phone         string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Icon          string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	//Status        MemberStatus         `protobuf:"varint,7,opt,name=status,proto3,enum=member.v1.MemberStatus" json:"status,omitempty"`
	//Gender        MemberGender         `protobuf:"varint,8,opt,name=gender,proto3,enum=member.v1.MemberGender" json:"gender,omitempty"`
	Birthday  *timestamp.Timestamp `protobuf:"bytes,9,opt,name=birthday,proto3" json:"birthday,omitempty"`
	City      string               `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	Job       string               `protobuf:"bytes,11,opt,name=job,proto3" json:"job,omitempty"`
	Growth    int32                `protobuf:"varint,12,opt,name=growth,proto3" json:"growth,omitempty"`
	CreatedAt string               `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}
