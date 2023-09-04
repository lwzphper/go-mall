package response

const (
	//全局
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	FAIL_ADD_DATA  = 800

	// 其他模块自定义
)

var CodeMsgMap = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "内部错误",
	INVALID_PARAMS: "表单验证有误",
	FAIL_ADD_DATA:  "添加数据失败",
}

func GetMsg(code int) string {
	if msg, ok := CodeMsgMap[code]; ok {
		return msg
	}
	return "未知"
}
