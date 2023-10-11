package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/admin/common"
	"github.com/lwzphper/go-mall/admin/global"
	"github.com/lwzphper/go-mall/pkg/jwt"
	"github.com/lwzphper/go-mall/pkg/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.UnauthorizedError(c.Writer)
			c.Abort()
			return
		}

		validator := jwt.NewTokenValidator(global.JwtSecret)
		userId, err := validator.Validator(token)
		if err != nil {
			if err == jwt.ErrExpiredOrNotValid {
				response.UnauthorizedError(c.Writer, response.WithMsg("授权已过期"))
			} else {
				response.UnauthorizedError(c.Writer, response.WithMsg("未登陆"))
			}
			c.Abort()
			return
		}

		// 设置 member_id
		common.ContextWithUserID(c, userId)
		c.Next()
	}
}
