package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lwzphper/go-mall/admin/global"
	"github.com/lwzphper/go-mall/pkg/response"
)

// HandleUserIdNotExistError 处理会员信息不存在错误
func HandleUserIdNotExistError(c *gin.Context) {
	response.InternalError(c.Writer, response.WithMsg("[u err] 服务器异常，请稍后再试"))
}

// HandleValidatorError 处理表单验证错误
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		global.L.Errorf("Validator error：%v", err)
		response.InternalError(c.Writer)
		return
	}

	// 多条错误信息，只显示第一条
	for _, val := range errs.Translate(global.T) {
		response.FormValidError(c.Writer, val)
		return
	}
}
