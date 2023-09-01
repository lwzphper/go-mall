package aws_s3

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	Bucket = "test-bucket"
	SK     = "admin123"
	Region = "weed"
	Point  = "127.0.0.1:8333"
)

func TestPutAndGetObj(t *testing.T) {
	err := InitService(DefaultClientName, SK, "", Region, Point)
	if err != nil {
		t.Error("InitService error", err)
	}
	client, err := GetS3Client(DefaultClientName)
	if err != nil {
		t.Error("get client error", err)
	}

	key := "mall_content/test"
	putStr := "this is for test"
	err = client.PutObj(key, Bucket, []byte(putStr))
	if err != nil {
		t.Error("PutObj error", err)
	}
	res, err := client.GetObj(key, Bucket)
	if err != nil {
		t.Error("GetObj error", err, string(res))
	}
	assert.Equal(t, putStr, string(res))
}

// 测试文件上传
func TestPutAndGetFile(t *testing.T) {
	err := InitService(DefaultClientName, SK, "", Region, Point)
	if err != nil {
		t.Error("InitService error", err)
	}

	client, err := GetS3Client(DefaultClientName)
	if err != nil {
		t.Error("get client error", err)
	}

	key := "test.png"
	err = client.PutFile(key, Bucket, "./test/test.png")
	if err != nil {
		t.Errorf("put file error:%v", err)
	}

	res, err := client.GetObj(key, Bucket)
	if err != nil {
		t.Errorf("get obj error:%v", err)
	}
	err = ioutil.WriteFile("./test/get.png", res, 0644)
	if err != nil {
		t.Errorf("save result error:%v", err)
	}
}
