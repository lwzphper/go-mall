package member

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api"
	"github.com/lwzphper/go-mall/bff/common"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/bff/reponse"
	"github.com/lwzphper/go-mall/pkg/response"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
)

// GetMemberDetail 获取会员详情
func GetMemberDetail(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	detail, err := global.MemberSrvClient.GetMemberById(c, &memberpb.IdRequest{Id: memberId})
	if err != nil {
		api.HandleGrpcErrorToHttp(c, err)
		return
	}

	// 这里空值会将字段过滤掉
	result := new(reponse.MemberResponse)
	result.Marshal(detail)
	response.Success(c.Writer, result)
}
