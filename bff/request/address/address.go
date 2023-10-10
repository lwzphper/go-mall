package address

type Update struct {
	Id        uint64 `form:"username" json:"username" xml:"username"  binding:"required" label:"id"`
	Name      string `form:"name" json:"name" xml:"name"  binding:"required" label:"用户名称"`
	Phone     string `form:"phone" json:"phone" xml:"phone"  binding:"required" label:"手机号码"`
	IsDefault uint32 `form:"is_default" json:"is_default" xml:"is_default"  binding:"on-off" label:"是否默认"`
}
