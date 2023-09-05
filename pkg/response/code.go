package response

const (
	// 全局

	CodeSuccess       int = 200
	CodeError         int = 500
	CodeInvalidParams int = 400
	CodeFailAddData   int = 800

	// 其他模块自定义
)

var CodeMsgMap = map[int]string{
	CodeSuccess:       "成功",
	CodeError:         "服务器繁忙，请稍后再试。",
	CodeInvalidParams: "表单验证有误",
	CodeFailAddData:   "添加数据失败",
}

func GetMsg(c int) string {
	if msg, ok := CodeMsgMap[c]; ok {
		return msg
	}
	return "未知"
}
