package until

import "time"

// TimeToYmdHis 将时间转换为 yyyy-mm-dd HH:ii:ss 格式
func TimeToYmdHis(time *time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
