package auth

// 使用参考：https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/
// 支持的规则：https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme

type Register struct {
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required,len:11"`
	SmsCode  uint32 `form:"sms_code" json:"sms_code" xml:"sms_code"  binding:"required,len:6"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required,min=6"`
}

type Login struct {
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required,len:11"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required,min=6"`
}
