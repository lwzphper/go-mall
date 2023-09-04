package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
)

// https://seepine.com/go/error/

func Exception() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if global.C.App.Env == app.EnvProduction {
					global.L.Errorf("gin error: %s\n", err)
					// 简单返回友好提示，具体可自定义发生错误后处理逻辑
					c.JSON(500, gin.H{"msg": "服务器发生错误"})
					c.Abort()
				}
				// todo 这里错误处理
				c.JSON(500, gin.H{"msg": "服务器发生错误"})
				fmt.Println(err)
			}
		}()
		c.Next()
	}
}
