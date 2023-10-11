package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/admin/global"
	"github.com/lwzphper/go-mall/pkg/response"
	"runtime/debug"
)

// 参考：https://seepine.com/go/error/

func Exception() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errStr := errorToString(err)
				global.L.Errorf("gin error: %s\n", errStr)
				response.InternalError(c.Writer)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func errorToString(err interface{}) string {
	switch v := err.(type) {
	//case WrapError: // 自定义异常
	//	// 符合预期的错误，可以直接返回给客户端
	//	return v.Msg
	case error:
		return fmt.Sprintf("panic: %v\n%s", v.Error(), debug.Stack())
	default:
		// 同上
		return err.(string)
	}
}
