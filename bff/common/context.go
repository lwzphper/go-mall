package common

import "github.com/gin-gonic/gin"

var ctxMemberIdKey = "memberId"

// ContextWithMemberID 将用户id设置到 context 中
func ContextWithMemberID(c *gin.Context, id uint64) {
	c.Set(ctxMemberIdKey, id)
}

// MemberIDFromContext 从 context 中获取用户id
func MemberIDFromContext(c *gin.Context) (uint64, bool) {
	value, ok := c.Get(ctxMemberIdKey)
	if !ok {
		return 0, false
	}
	return value.(uint64), true
}
