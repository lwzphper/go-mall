package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/admin/api"
	"github.com/lwzphper/go-mall/admin/common"
	"github.com/lwzphper/go-mall/pkg/response"
)

// Detail 获取会员详情
func Detail(c *gin.Context) {
	userId, exists := common.UserIDFromContext(c)
	if !exists {
		api.HandleUserIdNotExistError(c)
		return
	}
	response.Success(c.Writer, userId)
}

// Update 更新会员信息
func Update(c *gin.Context) {
	userId, exists := common.UserIDFromContext(c)
	if !exists {
		api.HandleUserIdNotExistError(c)
		return
	}

	response.Success(c.Writer, userId)
}
