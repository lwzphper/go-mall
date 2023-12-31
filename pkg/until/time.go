package until

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

var (
	cst *time.Location
)

const (
	DateLayout = "2006-01-02"
	TimeLayout = "15:04:05"
	CSTLayout  = DateLayout + " " + TimeLayout
)

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	// 默认设置为中国时区
	time.Local = cst
}

// TimeToDateTime 将时间转换为 yyyy-mm-dd HH:ii:ss 格式
func TimeToDateTime(time time.Time) string {
	return time.Format(CSTLayout)
}

// TimeToDate 将时间转换为 yyyy-mm-dd 格式
func TimeToDate(time time.Time) string {
	return time.Format(DateLayout)
}

// PbTimeToDate 将 protobuf 时间转换成 yyyy-mm-dd 格式
func PbTimeToDate(ts *timestamppb.Timestamp) string {
	return time.Unix(ts.GetSeconds(), 0).Format(DateLayout)
}

// TimeToPb 将时间转换成 PbTime
func TimeToPb(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}

// RFC3339ToCSTLayout convert rfc3339 value to china standard time layout
// 2020-11-08T08:18:46+08:00 => 2020-11-08 08:18:46
func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}

// CSTLayoutString 格式化时间
// 返回 "2006-01-02 15:04:05" 格式的时间
func CSTLayoutString() string {
	ts := time.Now()
	return ts.In(cst).Format(CSTLayout)
}

// ParseCSTInLocation 格式化时间
func ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(CSTLayout, date, cst)
}

// CSTLayoutStringToUnix 返回 unix 时间戳
// 2020-01-24 21:11:11 => 1579871471
func CSTLayoutStringToUnix(cstLayoutString string) (int64, error) {
	stamp, err := time.ParseInLocation(CSTLayout, cstLayoutString, cst)
	if err != nil {
		return 0, err
	}
	return stamp.Unix(), nil
}

// GMTLayoutString 格式化时间
// 返回 "Mon, 02 Jan 2006 15:04:05 GMT" 格式的时间
func GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

// ParseGMTInLocation 格式化时间
func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}
