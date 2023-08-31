package until

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	dateFormat     = "2006-01-02"
	timeFormat     = "15:04:05"
	dateTimeFormat = dateFormat + " " + timeFormat
)

// TimeToDateTime 将时间转换为 yyyy-mm-dd HH:ii:ss 格式
func TimeToDateTime(time time.Time) string {
	return time.Format(dateTimeFormat)
}

// TimeToDate 将时间转换为 yyyy-mm-dd 格式
func TimeToDate(time time.Time) string {
	return time.Format(dateFormat)
}

// PbTimeToDate 将 protobuf 时间转换成 yyyy-mm-dd HH:ii:ss 格式
func PbTimeToDate(ts *timestamppb.Timestamp) string {
	return time.Unix(ts.GetSeconds(), 0).Format(dateFormat)
}
