package rule

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
)

// RegisterOnOff 开头
func RegisterOnOff(v *validator.Validate, t ut.Translator) {
	// 验证规则
	err := v.RegisterValidation("on-off", func(fl validator.FieldLevel) bool {
		if val, ok := fl.Field().Interface().(uint32); ok {
			return val == 0 || val == 1
		}
		return false
	})
	if err != nil {
		logger.Errorf("Register on-off validation error")
	}

	// 自定义错误信息
	_ = v.RegisterTranslation("on-off", t, func(ut ut.Translator) error {
		return ut.Add("on-off", "{0}值有误!", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("on-off", fe.Field())
		return t
	})
}
