package rule

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
)

const (
	Default int32 = iota
	Male
	Female
	end
)

func RegisterGender(v *validator.Validate, t ut.Translator) {
	// 验证规则
	err := v.RegisterValidation("gender", func(fl validator.FieldLevel) bool {
		if val, ok := fl.Field().Interface().(int32); ok {
			return Default <= val && val < end
		}
		return false
	})
	if err != nil {
		logger.Errorf("Register gender validation error")
	}

	// 自定义错误信息
	_ = v.RegisterTranslation("gender", t, func(ut ut.Translator) error {
		return ut.Add("gender", "{0}格式有误!", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gender", fe.Field())
		return t
	})
}
