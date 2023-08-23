package dao

import (
	"context"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/lwzphper/go-mall/server/member/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试创建会员
func TestCreateMember(t *testing.T) {
	// 创建数据表
	/*if err := mysqltesting.GormDB.AutoMigrate(&entity.Member{}); err != nil {
		t.Errorf("create table error: %v", err)
	}*/

	dao := NewMember(mysqltesting.GormDB)
	ctx := context.Background()

	var username = "张三"

	member := &entity.Member{
		Username: username,
		Password: "123456",
	}
	err := dao.CreateMember(ctx, member)
	if err != nil {
		t.Errorf("create member error:%v", err)
	}

	memberRecord, err := dao.GetMemberByUsername(ctx, "张三")
	if err != nil {
		t.Errorf("get member info error: %v", memberRecord)
	}

	assert.Equal(t, username, memberRecord.Username)
}

// 测试获取会员
func TestGetMemberByUsername(t *testing.T) {

}

func initTable() {
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
