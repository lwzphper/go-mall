package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api/member"
	"github.com/lwzphper/go-mall/bff/middleware"
)

func InitMember(r *gin.RouterGroup) {
	// 授权
	authGroup := r.Group("auth") //.Use(middleware.Trace())
	{
		authGroup.POST("login", member.Login)
		authGroup.POST("register", member.Register)
	}

	// 会员信息
	memberGroup := r.Group("user").Use(middleware.JwtAuth())
	{
		memberGroup.GET("", member.Detail)
		memberGroup.PUT("", member.Update)
	}

}
