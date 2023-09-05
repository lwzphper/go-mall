package auth

type Register struct {
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required,len=11" label:"手机号码"`
	SmsCode  uint32 `form:"sms_code" json:"sms_code" xml:"sms_code"  binding:"required,len=6" label:"短信验证码"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required,min=6" label:"密码"`
}

type Login struct {
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required,len:11" label:"手机号码"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required,min=6" label:"密码"`
}
