package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/middleware"
	"github.com/lwzphper/go-mall/bff/router"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	r.Use(middleware.Cors())
	ApiGroup := r.Group("/v1")
	router.InitAuth(ApiGroup)
	router.InitMemberRouter(ApiGroup)
	return r
}
