package address

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
	dao          *Address
	ctx          context.Context
	hasInitTable bool // 是否初始化 table
	initLock     sync.Mutex
)

// 测试创建和查询
func TestCreateAndQuery(t *testing.T) {
	initTable()

	memberId := uint64(29)
	testCase := []struct {
		caseName  string
		name      string
		phone     string
		isDefault uint32
		postCode  string
		province  string
		city      string
		region    string
		address   string
		memberId  uint64
	}{
		{
			caseName:  "test equal",
			name:      "张三",
			phone:     until.RandomString(11),
			isDefault: 0,
			postCode:  "511450",
			province:  "广东省",
			city:      "广州市",
			region:    "天河区",
			address:   "xxx小区3单元1号",
			memberId:  memberId,
		},
		{
			caseName:  "test equal",
			name:      "王五",
			phone:     until.RandomString(11),
			isDefault: 0,
			postCode:  "511450",
			province:  "广东省",
			city:      "广州市",
			region:    "天河区",
			address:   "xxx小区3单元1号",
			memberId:  memberId,
		},
	}

	for i, c := range testCase {
		t.Run(c.caseName, func(t *testing.T) {
			item := entity.Address{
				Name:      c.name,
				Phone:     c.phone,
				IsDefault: c.isDefault,
				PostCode:  c.postCode,
				Province:  c.province,
				City:      c.city,
				Region:    c.region,
				Address:   c.address,
				MemberId:  c.memberId,
			}
			err := dao.Create(ctx, &item)
			if err != nil {
				t.Errorf("[%s]:create member error:%v", c.caseName, err)
			}

			record, err := dao.GetList(ctx, id.MemberID(memberId))
			if err != nil {
				t.Errorf("[%s]:get list error: %v", c.caseName, err)
			}
			if len(record) == 0 {
				t.Errorf("[%s]:get list data empty", c.caseName)
			}

			assert.Equal(t, c.name, record[i].Name)
		})
	}
}

func TestUpdate(t *testing.T) {
	initTable()

	memberId := uint64(107)
	save := &entity.Address{
		Name:      "王五",
		Phone:     until.RandomString(11),
		IsDefault: uint32(1),
		PostCode:  "511450",
		Province:  "广东省",
		City:      "广州市",
		Region:    "天河区",
		Address:   "xxx小区3单元1号",
		MemberId:  memberId,
	}
	err := dao.Create(ctx, save)
	if err != nil {
		t.Errorf("create member error:%v", err)
	}

	uData := map[string]interface{}{
		"name":       "张三",
		"phone":      until.TimeToDate(time.Now()),
		"is_default": uint32(1),
		"post_code":  "412000",
		"province":   "长沙市",
		"city":       "株洲市",
		"region":     "芦淞区",
		"address":    "xxx小区8单元3号",
	}
	err = dao.UpdateById(ctx, id.AddressID(save.Id), uData)
	if err != nil {
		t.Errorf("update member error:%v", err)
	}

	member, err := dao.GetList(ctx, id.MemberID(memberId))
	if err != nil {
		t.Errorf("get list error:%v", err)
	}
	if len(member) == 0 {
		t.Errorf("cannot get list by member id：%d", memberId)
	}
	assert.Equal(t, uData["name"], member[0].Name)
	assert.Equal(t, uData["phone"], member[0].Phone)
	assert.Equal(t, uData["is_default"], member[0].IsDefault)
	assert.Equal(t, uData["post_code"], member[0].PostCode)
	assert.Equal(t, uData["province"], member[0].Province)
	assert.Equal(t, uData["city"], member[0].City)
	assert.Equal(t, uData["region"], member[0].Region)
	assert.Equal(t, uData["address"], member[0].Address)
}

func initTable() {
	initLock.Lock()
	defer initLock.Unlock()

	if hasInitTable {
		return
	}
	// 创建数据表
	if err := init_table.Address(); err != nil {
		log.Panicf("create table error: %v", err)
	}

	global.DB = mysqltesting.GormDB
	dao = NewAddress()
	ctx = context.Background()
	hasInitTable = true
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
