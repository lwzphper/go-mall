package member

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api"
	"github.com/lwzphper/go-mall/bff/common"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/bff/request/member"
	member2 "github.com/lwzphper/go-mall/bff/response"
	ginhelper "github.com/lwzphper/go-mall/pkg/gin"
	"github.com/lwzphper/go-mall/pkg/response"
	"github.com/lwzphper/go-mall/pkg/until"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"
)

// Detail 获取会员详情
func Detail(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	detail, err := global.MemberSrvClient.GetMemberById(c, &memberpb.IdRequest{Id: memberId})
	if err != nil {
		ginhelper.HandleGrpcErrorToHttp(c, err)
		return
	}

	// 这里空值会将字段过滤掉
	result := member2.NewMemberResponse()
	result.Marshal(detail)
	response.Success(c.Writer, result)
}

// Update 更新会员信息
func Update(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	var req member.Update
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	_, err := global.MemberSrvClient.UpdateMember(c, &memberpb.MemberEntity{
		Id:       memberId,
		Username: req.Username,
		Icon:     req.Icon,
		Birthday: until.TimeToPb(req.Birthday),
		Gender:   memberpb.MemberGender(req.Gender),
	})
	if err != nil {
		ginhelper.HandleGrpcErrorToHttp(c, err)
		return
	}
	response.Success(c.Writer, nil)
}
