package response

import (
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
)

type AddressEntity struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	IsDefault uint32 `json:"is_default"`
	PostCode  string `json:"post_code"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Detail    string `json:"detail"`
}

func NewAddressList() *AddressList {
	return &AddressList{}
}

type AddressList struct {
	list []AddressEntity
}

func (r *AddressList) Marshal(p *addresspb.ListResponse) {
	for _, item := range p.GetList() {
		r.list = append(r.list, AddressEntity{
			Id:        item.Id,
			Name:      item.Name,
			Phone:     item.Phone,
			IsDefault: item.IsDefault,
			PostCode:  item.PostCode,
			Province:  item.Province,
			City:      item.City,
			Region:    item.Region,
			Detail:    item.Detail,
		})
	}
}
