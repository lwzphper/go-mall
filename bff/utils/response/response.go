package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	"log"
	"net/http"
)

var successCode = 0

// Success 成功响应
func Success(c *gin.Context, data interface{}, opts ...Option) {
	resp := &Data{
		Code:    &successCode,
		Message: "OK",
		Data:    data,
	}

	for _, opt := range opts {
		opt(resp)
	}
	c.JSON(http.StatusOK, resp)

	// 开发环境，打印响应结果
	if global.C.App.Env == app.EnvDevelopment {
		respBytes, _ := json.Marshal(resp)
		log.Print(c.Request.Context(), "resp: %s", string(respBytes))
	}
}
