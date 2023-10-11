package response

import (
	"github.com/lwzphper/go-mall/pkg/until"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"
)

func NewMemberResponse() *MemberResponse {
	return &MemberResponse{}
}

type MemberResponse struct {
	Id            uint64 `json:"id"`
	MemberLevelId uint64 `json:"member_level_id"`
	Username      string `json:"username"`
	Phone         string `json:"phone"`
	Icon          string `json:"icon"`
	Status        int32  `json:"status"`
	Gender        int32  `json:"gender"`
	Birthday      string `json:"birthday"`
	City          string `json:"city"`
	Job           string `json:"job"`
	Growth        int32  `json:"growth"`
	CreatedAt     string `json:"created_at"`
}

func (r *MemberResponse) Marshal(p *memberpb.MemberEntity) {
	r.Id = p.Id
	r.MemberLevelId = p.MemberLevelId
	r.Username = p.Username
	r.Phone = p.Phone
	r.Icon = p.Icon
	r.Status = int32(p.Status)
	r.Gender = int32(p.Gender)
	r.City = p.Job
	r.Growth = p.Growth
	r.CreatedAt = p.CreatedAt

	if p.Birthday != nil {
		r.Birthday = until.TimeToDate(p.Birthday.AsTime())
	} else {
		r.Birthday = ""
	}
}
