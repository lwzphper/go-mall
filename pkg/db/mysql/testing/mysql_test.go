package mysqltesting

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type UserTmp struct {
	gorm.Model
	Name string `gorm:"size:255;index:idx_name,unique"`
}

var err error

func TestCreate(t *testing.T) {
	// 创建数据表
	if err = GormDB.AutoMigrate(&UserTmp{}); err != nil {
		t.Errorf("create table error: %v", err)
	}

	// 添加数据
	user := UserTmp{
		Name: "test",
	}
	if err = GormDB.Create(&user).Error; err != nil {
		t.Errorf("insert error %v", err)
	}

	// 校验插入的数据
	queryUser := &UserTmp{}
	GormDB.First(queryUser)
	assert.Equal(t, "test", queryUser.Name)

	// 删除数据表
	GormDB.Exec("drop table user_tmp")
}

func TestMain(m *testing.M) {
	RunMysqlInDocker(m)
}
