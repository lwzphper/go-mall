package mysql

import (
	"fmt"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:255;index:idx_name,unique"`
}

func TestCreate(t *testing.T) {
	err := initClient(DefaultClient)
	if err != nil {
		t.Errorf("init mysql error: %v", err)
	}
	defer CloseMysqlClient(DefaultClient)

	ormDB := GetMysqlClient(DefaultClient).DB

	// 创建数据表
	if err = ormDB.AutoMigrate(&User{}); err != nil {
		t.Errorf("create table error: %v", err)
	}

	// 添加数据
	user := User{
		Name: "test",
	}
	if err = ormDB.Create(&user).Error; err != nil {
		t.Errorf("insert error %v", err)
	}

	// 校验插入的数据
	queryUser := &User{}
	ormDB.First(queryUser)

	assert.Equal(t, "test", queryUser.Name)
}

func TestMain(m *testing.M) {
	os.Exit(mysqltesting.RunWithMongoInDocker(m))
}

// initClient 创建客户端
func initClient(clientName string, options ...Option) error {
	return InitMysqlClient(
		clientName,
		mysqltesting.Username,
		mysqltesting.Password,
		fmt.Sprintf("%s:%s", mysqltesting.ContainerHostPort.HostIP, mysqltesting.ContainerHostPort.HostPort),
		mysqltesting.Database,
		options...)
}
