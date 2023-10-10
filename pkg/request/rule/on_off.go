package rule

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/lwzphper/go-mall/bff/global"
)

// RegisterOnOff 开头
func RegisterOnOff(v *validator.Validate) {
	// 验证规则
	err := v.RegisterValidation("on-off", func(fl validator.FieldLevel) bool {
		if val, ok := fl.Field().Interface().(uint32); ok {
			return val == 0 || val == 1
		}
		return false
	})
	if err != nil {
		global.L.Errorf("Register on-off validation error")
	}

	// 自定义错误信息
	_ = v.RegisterTranslation("on-off", global.T, func(ut ut.Translator) error {
		return ut.Add("on-off", "{0}值有误!", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("on-off", fe.Field())
		return t
	})
}
