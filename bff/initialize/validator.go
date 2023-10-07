package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/bff/rule"
	"reflect"
)

const (
	EnLocale = "en"
	ZhLocale = "zh"
)

// InitValidator 初始验证器
func InitValidator(locale string) {
	// 修改gin框架中的validator引擎属性, 实现定制
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	// 注册一个获取 label 的tag的自定义方法
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		if name == "" {
			return field.Name
		}
		return name
	})

	// 设置语言
	setLocale(v, locale)

	// 注册自定义验证类
	rule.RegisterGender(v)
}

// 设置语言
func setLocale(v *validator.Validate, locale string) {
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
	uni := ut.New(enT, zhT, enT)
	var ok bool
	global.T, ok = uni.GetTranslator(locale)
	if !ok {
		global.L.Errorf("uni.GetTranslator error:%s", locale)
		return
	}

	var err error
	switch locale {
	case ZhLocale:
		err = zh_translations.RegisterDefaultTranslations(v, global.T)
	default:
		err = en_translations.RegisterDefaultTranslations(v, global.T)
	}
	if err != nil {
		global.L.Errorf("Register Default transition error:%v", err)
	}
}
