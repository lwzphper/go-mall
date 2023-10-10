package member

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
	"sync"
	"testing"
	"time"
)

var (
	memberDao    *Member
	ctx          context.Context
	hasInitTable bool // 是否初始化 table
	initLock     sync.Mutex
)

// 测试会员创建和查询
func TestCreateAndQueryMember(t *testing.T) {
	initTable()

	testCase := []struct {
		caseName  string // 测试名称
		username  string // 会员名称
		phone     string
		wantName  string // 期望值
		wantEqual bool   // 是否期望相等
	}{
		{
			caseName:  "test equal",
			username:  "张三",
			phone:     until.RandomString(11),
			wantName:  "张三",
			wantEqual: true,
		},
		{
			caseName:  "test not equal",
			username:  "李四",
			phone:     until.RandomString(11),
			wantName:  "张三",
			wantEqual: false,
		},
	}

	for _, c := range testCase {
		t.Run(c.caseName, func(t *testing.T) {
			member := entity.Member{
				Username: c.username,
				Phone:    c.phone,
				Password: until.RandomString(32),
			}
			err := memberDao.CreateMember(ctx, &member)
			if err != nil {
				t.Errorf("[%s]:create member error:%v", c.caseName, err)
			}

			memberRecord, err := memberDao.GetItemByWhere(ctx, &entity.Member{Username: c.username})
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
		Phone:    until.RandomString(11),
		Password: "123456",
	}
	err := memberDao.CreateMember(ctx, save)
	if err != nil {
		t.Errorf("create member error:%v", err)
	}

	member, err := memberDao.GetItemById(ctx, id.MemberID(save.Id))
	if err != nil {
		t.Errorf("get member by id error:%v", err)
	}
	assert.Equal(t, save.Username, member.Username)
}

func TestUpdate(t *testing.T) {
	initTable()

	save := &entity.Member{
		Username: "张三",
		Phone:    until.RandomString(11),
		Password: "123456",
	}
	err := memberDao.CreateMember(ctx, save)
	if err != nil {
		t.Errorf("create member error:%v", err)
	}

	uData := map[string]interface{}{
		"username":        "王五",
		"birthday":        until.TimeToDate(time.Now()),
		"member_level_id": uint64(1),
		"icon":            "https://www.baidu.com",
		"status":          entity.StatusDisable,
		"gender":          entity.GenderMan,
		"city":            "广州",
		"job":             "go开发工程师",
		"growth":          int32(100),
	}
	err = memberDao.UpdateById(ctx, id.MemberID(save.Id), uData)
	if err != nil {
		t.Errorf("update member error:%v", err)
	}

	member, err := memberDao.GetItemById(ctx, id.MemberID(save.Id))
	if err != nil {
		t.Errorf("get member by id error:%v", err)
	}
	assert.Equal(t, uData["username"], member.Username)
	assert.Equal(t, uData["birthday"], until.TimeToDate(*member.Birthday))
	assert.Equal(t, uData["member_level_id"], member.MemberLevelId)
	assert.Equal(t, uData["icon"], member.Icon)
	assert.Equal(t, uData["status"], member.Status)
	assert.Equal(t, uData["gender"], member.Gender)
	assert.Equal(t, uData["city"], member.City)
	assert.Equal(t, uData["job"], member.Job)
	assert.Equal(t, uData["growth"], member.Growth)
}

func TestMember_UpdateByEntity(t *testing.T) {
	initTable()

	save := &entity.Member{
		Username: "张三",
		Phone:    until.RandomString(11),
		Password: "123456",
	}
	err := memberDao.CreateMember(ctx, save)
	if err != nil {
		t.Errorf("create member error:%v", err)
	}

	save.Username = "王五"
	err = memberDao.UpdateByEntity(ctx, save)
	if err != nil {
		t.Errorf("update member error:%v", err)
	}

	member, err := memberDao.GetItemById(ctx, id.MemberID(save.Id))
	if err != nil {
		t.Errorf("get member by id error:%v", err)
	}
	assert.Equal(t, save.Username, member.Username)
}

func initTable() {
	initLock.Lock()
	defer initLock.Unlock()

	if hasInitTable {
		return
	}
	// 创建数据表
	if err := init_table.Member(); err != nil {
		log.Panicf("create table error: %v", err)
	}

	global.DB = mysqltesting.GormDB
	memberDao = NewMember(context.Background())
	ctx = context.Background()
	hasInitTable = true
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
