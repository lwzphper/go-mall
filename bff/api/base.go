package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/pkg/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// RemoveTopStruct 移除字字段表单前缀。如： LoginFrom.username 移除 LoginFrom. 前缀
func RemoveTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// HandleMemberIdNotExistError 处理会员信息不存在错误
func HandleMemberIdNotExistError(c *gin.Context) {
	response.InternalError(c.Writer, response.WithMsg("[m err] 服务器异常，请稍后再试"))
}

// HandleValidatorError 处理表单验证错误
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		global.L.Errorf("Validator error：%v", err)
		response.InternalError(c.Writer)
		return
	}

	// 多条错误信息，只显示第一条
	for _, val := range errs.Translate(global.T) {
		response.FormValidError(c.Writer, val)
		return
	}
}

// HandleGrpcErrorToHttp 将 grpc 错误转换成 http 错误信息
func HandleGrpcErrorToHttp(c *gin.Context, err error) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Unavailable:
				response.InternalError(c.Writer, response.WithMsg("服务不可用"))
			//case codes.NotFound:
			//	response.NotFoundError(c.Writer, e.Message())
			//case codes.Internal:
			//	response.InternalError(c.Writer)
			//case codes.InvalidArgument:
			//	response.FormValidError(c.Writer, e.Message())
			//case codes.AlreadyExists:
			//	response.InternalError(c.Writer, response.WithMsg("数据已存在"))
			default:
				response.InternalError(c.Writer, response.WithMsg(e.Message()))
			}
			return
		}
	}
}