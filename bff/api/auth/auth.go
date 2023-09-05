package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api"
	"github.com/lwzphper/go-mall/bff/request/auth"
	"github.com/lwzphper/go-mall/pkg/response"
)

// Login 登录
func Login(c *gin.Context) {
	var req auth.Login
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}
	response.Success(c.Writer, nil)
}

// Register 注册
func Register(c *gin.Context) {
	var req auth.Register
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}
	response.Success(c.Writer, nil)
}
