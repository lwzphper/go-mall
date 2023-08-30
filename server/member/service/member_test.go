package service

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/common/id"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/lwzphper/go-mall/pkg/db/mysql/testing/init_table"
	"github.com/lwzphper/go-mall/pkg/logger"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/global"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemberService_CreateMember(t *testing.T) {
	// 创建数据表
	if err := init_table.Member(); err != nil {
		t.Errorf("create table error: %v", err)
	}

	global.DB = mysqltesting.GormDB
	srv := &MemberService{
		MemberDao: dao.NewMember(),
		Logger:    logger.NewDefaultLogger(),
	}
	member, err := srv.CreateMember(context.Background(), &memberpb.CreateRequest{
		Username: "张三",
		Phone:    "15800000001",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, id.MemberID(1).Uint64(), member.Id)
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
