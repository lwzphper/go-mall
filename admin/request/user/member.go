package user

import "time"

type Update struct {
	Username string    `form:"username" json:"username" xml:"username"  binding:"required" label:"用户名"`
	Icon     string    `form:"icon" json:"icon" xml:"icon"  binding:"required" label:"头像"`
	Birthday time.Time `form:"birthday" json:"birthday" xml:"birthday" label:"生日日期" time_format:"2006-01-02"`
	Gender   int32     `form:"gender" json:"gender" binding:"gender" label:"性别"`
}
