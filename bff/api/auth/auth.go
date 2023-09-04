package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/request/auth"
	"github.com/lwzphper/go-mall/bff/utils/response"
)

// Login 登录
func Login(c *gin.Context) {
	var req auth.Login
	if err := c.ShouldBind(&req); err != nil {
		return
	}
	response.Success(c, nil)
}

// Register 注册
func Register(c *gin.Context) {
	var req auth.Register
	if err := c.ShouldBind(&req); err != nil {
		return
	}
	response.Success(c, nil)
}
