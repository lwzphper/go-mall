package address

import (
	"github.com/lwzphper/go-mall/pkg/common/id"
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
)

type Entity struct {
	Id uint64 `form:"id" json:"id" xml:"id"  binding:"required" label:"id"`
	Address
}

func (r *Entity) Marshal(mId id.MemberID) *addresspb.Entity {
	return &addresspb.Entity{
		Id:        r.Id,
		Name:      r.Name,
		Phone:     r.Phone,
		Province:  r.Province,
		City:      r.City,
		Region:    r.Region,
		Detail:    r.Detail,
		IsDefault: r.IsDefault,
		PostCode:  r.PostCode,
		MemberId:  mId.Uint64(),
	}
}

type Address struct {
	Name      string `form:"name" json:"name" xml:"name"  binding:"required" label:"收货人"`
	Phone     string `form:"phone" json:"phone" xml:"phone"  binding:"required" label:"手机号码"`
	Province  string `form:"province" json:"province" xml:"province"  binding:"required" label:"省份"`
	City      string `form:"city" json:"city" xml:"city"  binding:"required" label:"城市"`
	Region    string `form:"region" json:"region" xml:"region"  binding:"required" label:"地区"`
	Detail    string `form:"detail" json:"detail" xml:"detail"  binding:"required" label:"详细地址"`
	IsDefault uint32 `form:"is_default" json:"is_default" xml:"is_default"  binding:"on-off" label:"是否默认地址"`
	PostCode  string `form:"post_code" json:"post_code" xml:"post_code"  label:"邮政编码"`
}

func (r *Address) Marshal(mId id.MemberID) *addresspb.CreateRequest {
	return &addresspb.CreateRequest{
		Name:      r.Name,
		Phone:     r.Phone,
		Province:  r.Province,
		City:      r.City,
		Region:    r.Region,
		Detail:    r.Detail,
		IsDefault: r.IsDefault,
		PostCode:  r.PostCode,
		MemberId:  mId.Uint64(),
	}
}

type Delete struct {
	Id uint64 `form:"id" json:"id" xml:"id"  binding:"required" label:"id"`
}

func (r *Delete) Marshal(mId id.MemberID) *addresspb.DeleteRequest {
	return &addresspb.DeleteRequest{
		Id:       r.Id,
		MemberId: mId.Uint64(),
	}
}
