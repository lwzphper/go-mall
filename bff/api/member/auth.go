package member

import (
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/api"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/bff/request/auth"
	jwt2 "github.com/lwzphper/go-mall/pkg/jwt"
	"github.com/lwzphper/go-mall/pkg/response"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
)

// Login 登录
func Login(c *gin.Context) {
	var req auth.Login
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	// 获取用户信息
	member, err := global.MemberSrvClient.GetMemberByPhone(c, &memberpb.PhoneRequest{Phone: req.Phone})
	if err != nil {
		api.HandleGrpcErrorToHttp(c, err)
		return
	}

	// 检查密码是否正确
	checkRet, _ := global.MemberSrvClient.CheckPassWord(c, &memberpb.PasswordCheckInfo{
		Password:          req.Password,
		EncryptedPassword: member.Password,
	})
	if checkRet.Success == false {
		response.FormValidError(c.Writer, "密码不正确")
		return
	}

	// 生成 token
	tokenGen := jwt2.NewJwtTokenGen(global.C.App.Name, global.JwtSecret)
	token, err := tokenGen.GenerateToken(member.Id, global.C.Jwt.TTL)
	if err != nil {
		global.L.Errorf("create token error:%v", err)
		response.InternalError(c.Writer)
		return
	}

	result := member
	result.Password = "" // 删除密码
	response.Success(c.Writer, member, response.WithAuthHeader(token))
}

// Register 注册
func Register(c *gin.Context) {
	var req auth.Register
	if err := c.ShouldBind(&req); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	// 校验短信
	if req.SmsCode != "123456" {
		response.FormValidError(c.Writer, "短信验证码有误")
		return
	}

	createReq := &memberpb.CreateRequest{
		Phone:    req.Phone,
		Password: req.Password,
	}
	member, err := global.MemberSrvClient.CreateMember(c, createReq)
	if err != nil {
		api.HandleGrpcErrorToHttp(c, err)
		return
	}

	result := make(map[string]uint64, 1)
	result["id"] = member.Id
	response.Success(c.Writer, result)
}
