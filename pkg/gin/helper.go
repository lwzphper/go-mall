package gin

import (
	"github.com/gin-gonic/gin"
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

// HandleGrpcErrorToHttp 将 grpc 错误转换成 http 错误信息
func HandleGrpcErrorToHttp(c *gin.Context, err error) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			errMsg := "[g]" + e.Message()
			switch e.Code() {
			case codes.Unavailable:
				response.InternalError(c.Writer, response.WithMsg("[g] 服务不可用"))
			//case codes.NotFound:
			//	response.NotFoundError(c.Writer, e.Message())
			//case codes.Internal:
			//	response.InternalError(c.Writer)
			//case codes.InvalidArgument: // bff 层对数据校验了，grpc层就不校验
			//	response.FormValidError(c.Writer, errMsg)
			//case codes.AlreadyExists:
			//	response.InternalError(c.Writer, response.WithMsg("数据已存在"))
			default:
				response.InternalError(c.Writer, response.WithMsg(errMsg))
			}
			return
		}
	}
}
