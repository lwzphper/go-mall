package aws_s3

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lwzphper/go-mall/pkg/logger"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Service struct {
	Client   *s3.S3
	EndPoint string
	Region   string
}

const DefaultClientName = "default-s3-client"

var clients map[string]*Service
var log *logger.Logger

func init() {
	log = logger.NewDefaultLogger()
}

// InitService 初始化客户端
func InitService(clientID, sk, token, region, point string) error {
	credential := credentials.NewStaticCredentials(clientID, sk, token)
	cfg := aws.NewConfig().WithCredentials(credential).WithRegion(region).
		WithEndpoint(point).WithS3ForcePathStyle(true).WithDisableSSL(true)
	sess, err := session.NewSession(cfg)
	if err != nil {
		return err
	}
	if len(clients) == 0 {
		clients = make(map[string]*Service, 0)
	}
	clients[clientID] = &Service{
		Client:   s3.New(sess),
		EndPoint: point,
		Region:   region,
	}
	return nil
}

// GetS3Client 获取客户端
func GetS3Client(clientName string) (*Service, error) {
	if service, ok := clients[clientName]; ok {
		return service, nil
	}
	return nil, errors.New(fmt.Sprintf("client %s has not initial", clientName))
}

// GetObj 获取对象
func (s *Service) GetObj(key, bucket string) ([]byte, error) {
	inputObject := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	out, err := s.Client.GetObject(inputObject)
	if err != nil && !IsNotFoundErr(err) {
		return nil, err
	}

	if out.Body != nil {
		res, err := ioutil.ReadAll(out.Body)
		_ = out.Body.Close()
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	err = awserr.New(s3.ErrCodeNoSuchKey, "empty body", nil)
	return nil, err
}

// PutObj 添加对象
func (s *Service) PutObj(key, bucket string, data []byte) error {
	inputObject := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String("application/octet-stream"),
		Body:        bytes.NewReader(data),
	}
	out, err := s.Client.PutObject(inputObject)
	if err != nil {
		outStr := ""
		if out != nil {
			outStr = out.String()
		}
		log.Errorf("PutS3Object error:%s  outStr:%s", err, outStr)
		return err
	}
	return nil
}

// PutFile 上传文件
func (s *Service) PutFile(key, bucket, fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return s.PutObj(key, bucket, content)
}

// IsNotFoundErr 判断资源是否不存在
func IsNotFoundErr(err error) bool {
	if aswErr, ok := err.(awserr.Error); ok {
		switch aswErr.Code() {
		case s3.ErrCodeNoSuchKey:
			return true
		}
	}
	return false
}
