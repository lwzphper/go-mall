package service

import (
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/lwzphper/go-mall/pkg/db/mysql/testing/init_table"
	"testing"
)

func TestMemberService_CreateMember(t *testing.T) {
	// 创建数据表
	if err := init_table.Member(); err != nil {
		t.Errorf("create table error: %v", err)
	}

	//srv := &MemberService{
	//	MemberDao: nil,
	//	Logger:    nil,
	//}
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
