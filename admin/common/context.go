package common

import "github.com/gin-gonic/gin"

var ctxUserIdKey = "memberId"

// ContextWithUserID 将用户id设置到 context 中
func ContextWithUserID(c *gin.Context, id uint64) {
	c.Set(ctxUserIdKey, id)
}

// UserIDFromContext 从 context 中获取用户id
func UserIDFromContext(c *gin.Context) (uint64, bool) {
	value, ok := c.Get(ctxUserIdKey)
	if !ok {
		return 0, false
	}
	return value.(uint64), true
}
