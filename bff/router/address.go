package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api/address"
	"github.com/lwzphper/go-mall/bff/middleware"
)

func InitAddress(r *gin.RouterGroup) {
	addressGroup := r.Group("address").Use(middleware.JwtAuth())
	{
		addressGroup.GET("", address.List)
		addressGroup.POST("", address.Create)
		addressGroup.PUT("", address.Update)
		addressGroup.DELETE("", address.Delete)
	}
}
