package dao

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/common/id"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/lwzphper/go-mall/pkg/db/mysql/testing/init_table"
	"github.com/lwzphper/go-mall/pkg/until"
	"github.com/lwzphper/go-mall/server/member/entity"
	"github.com/lwzphper/go-mall/server/member/global"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	dao *Member
	ctx context.Context
)

// 测试会员创建和查询
func TestCreateAndQueryMember(t *testing.T) {
	initTable()

	testCase := []struct {
		caseName  string // 测试名称
		username  string // 会员名称
		wantName  string // 期望值
		wantEqual bool   // 是否期望相等
	}{
		{
			caseName:  "test equal",
			username:  "张三",
			wantName:  "张三",
			wantEqual: true,
		},
		{
			caseName:  "test not equal",
			username:  "李四",
			wantName:  "张三",
			wantEqual: false,
		},
	}

	for _, c := range testCase {
		t.Run(c.caseName, func(t *testing.T) {
			member := entity.Member{
				Username: c.username,
				Password: until.RandomString(32),
			}
			err := dao.CreateMember(ctx, &member)
			if err != nil {
				t.Errorf("[%s]:create member error:%v", c.caseName, err)
			}

			memberRecord, err := dao.GetItemByWhere(ctx, &entity.Member{Username: c.username})
			if err != nil {
				t.Errorf("[%s]:get member info error: %v", c.caseName, err)
			}

			if c.wantEqual && c.wantName != memberRecord.Username {
				t.Errorf("[%s]:want %s but get %s", c.caseName, c.wantName, memberRecord.Username)
			}

			if c.wantEqual == false && c.wantName == memberRecord.Username {
				t.Errorf("[%s]:value equal %s", c.caseName, c.wantName)
			}
		})
	}
}

func TestGetItemById(t *testing.T) {
	initTable()

	save := &entity.Member{
		Username: "张三",
		Phone:    "15800000001",
		Password: "123456",
	}
	err := dao.CreateMember(ctx, save)
	if err != nil {
		t.Errorf("create member error:%v", err)
	}

	member, err := dao.GetItemById(ctx, id.MemberID(save.Id))
	if err != nil {
		t.Errorf("get member by id error:%v", err)
	}
	assert.Equal(t, save.Username, member.Username)
}

func TestUpdate(t *testing.T) {
	initTable()

	save := &entity.Member{
		Username: "张三",
		Phone:    "15800000001",
		Password: "123456",
	}
	err := dao.CreateMember(ctx, save)
	if err != nil {
		t.Errorf("create member error:%v", err)
	}

	save.Username = "王五"
	err = dao.Update(ctx, save)
	if err != nil {
		t.Errorf("update member error:%v", err)
	}

	member, err := dao.GetItemById(ctx, id.MemberID(save.Id))
	if err != nil {
		t.Errorf("get member by id error:%v", err)
	}
	assert.Equal(t, save.Username, member.Username)
}

func initTable() {
	// 创建数据表
	if err := init_table.Member(); err != nil {
		log.Panicf("create table error: %v", err)
	}

	global.DB = mysqltesting.GormDB
	dao = NewMember()
	ctx = context.Background()
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
