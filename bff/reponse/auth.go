package reponse

import memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"

type LoginResponse struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}

func (r *LoginResponse) Marshal(p *memberpb.MemberEntity) {
	r.Id = p.Id
	r.Username = p.Username
	r.Icon = p.Icon
}
