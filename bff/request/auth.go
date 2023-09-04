package request

type Auth struct {
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	SmsCode  string `form:"sms_code" json:"sms_code" xml:"sms_code"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}
