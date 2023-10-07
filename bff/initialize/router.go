package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/middleware"
	"github.com/lwzphper/go-mall/bff/router"
	"github.com/lwzphper/go-mall/pkg/response"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	// 心跳检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 404 处理
	r.NoRoute(func(c *gin.Context) {
		response.NotFoundError(c.Writer)
	})

	r.Use(middleware.Cors())      //配置跨域
	r.Use(middleware.Exception()) // 错误处理

	v1Group := r.Group("/v1")
	router.InitMember(v1Group)
	return r
}
