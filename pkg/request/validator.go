package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// Validator 验证器接口
type Validator interface {
	// GetMessages 获取验证器自定义错误信息
	GetMessages() ValidatorMessages
}

// ValidatorMessages 验证器自定义错误信息字典
type ValidatorMessages map[string]string

// GetErrorMsg 获取自定义错误信息
func GetErrorMsg(request Validator, err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, v := range err.(validator.ValidationErrors) {
			if message, exist := request.GetMessages()[v.Field()+"."+v.Tag()]; exist {
				return message
			}
			return v.Error()
		}
		return "Parameter error"
	}
	// 其他错误，如：json请求字符串不是json格式等
	return err.Error()
}
