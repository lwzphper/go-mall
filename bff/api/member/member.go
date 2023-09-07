package member

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/common"
)

func GetMemberDetail(c *gin.Context) {
	memberId, exists := common.MemberIDFromContext(c)
	if !exists {
		fmt.Println("未认证")
		return
	}
	fmt.Printf("userid %d", memberId)
}
