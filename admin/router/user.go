package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/admin/api/user"
	"github.com/lwzphper/go-mall/admin/middleware"
)

func InitUser(r *gin.RouterGroup) {
	// 授权
	authGroup := r.Group("auth") //.Use(middleware.Trace())
	{
		authGroup.POST("login", user.Login)
		authGroup.POST("register", user.Register)
	}

	// 会员信息
	memberGroup := r.Group("user").Use(middleware.JwtAuth())
	{
		memberGroup.GET("", user.Detail)
		memberGroup.PUT("", user.Update)
	}

}
