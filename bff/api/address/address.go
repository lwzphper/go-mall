package address

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api"
	"github.com/lwzphper/go-mall/bff/common"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/bff/response"
	untilResponse "github.com/lwzphper/go-mall/pkg/response"
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
)

// List 获取列表
func List(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	list, err := global.AddressSrvClient.GetList(c, &addresspb.ListRequest{MemberId: memberId})
	if err != nil {
		api.HandleGrpcErrorToHttp(c, err)
	}

	result := response.NewAddressList()
	result.Marshal(list)
	untilResponse.Success(c.Writer, result)
}

// Update 更新
func Update(c *gin.Context) {
	/*memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	global.AddressSrvClient.Update(c, &addresspb.Entity{})*/
}

// Delete 删除
func Delete(c *gin.Context) {
	/*memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}*/
}
