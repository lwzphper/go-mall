package address

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api"
	"github.com/lwzphper/go-mall/bff/common"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/bff/request/address"
	"github.com/lwzphper/go-mall/bff/response"
	"github.com/lwzphper/go-mall/pkg/common/id"
	ginhelper "github.com/lwzphper/go-mall/pkg/gin"
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
		ginhelper.HandleGrpcErrorToHttp(c, err)
	}

	result := response.NewAddressList()
	result.Marshal(list)
	untilResponse.Success(c.Writer, result)
}

// Create 创建
func Create(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	var req address.Address
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	_, err := global.AddressSrvClient.Create(c, req.Marshal(id.MemberID(memberId)))
	if err != nil {
		ginhelper.HandleGrpcErrorToHttp(c, err)
		return
	}
	untilResponse.Success(c.Writer, nil)
}

// Update 更新
func Update(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	var req address.Entity
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	_, err := global.AddressSrvClient.Update(c, req.Marshal(id.MemberID(memberId)))
	if err != nil {
		ginhelper.HandleGrpcErrorToHttp(c, err)
		return
	}
	untilResponse.Success(c.Writer, nil)
}

// Delete 删除
func Delete(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		api.HandleMemberIdNotExistError(c)
		return
	}

	var req address.Delete
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	_, err := global.AddressSrvClient.Delete(c, req.Marshal(id.MemberID(memberId)))
	if err != nil {
		ginhelper.HandleGrpcErrorToHttp(c, err)
		return
	}
	untilResponse.Success(c.Writer, nil)
}
