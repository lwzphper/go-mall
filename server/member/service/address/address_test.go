package address

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/common/id"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/lwzphper/go-mall/pkg/db/mysql/testing/init_table"
	"github.com/lwzphper/go-mall/pkg/until"
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
	"github.com/lwzphper/go-mall/server/member/global"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

var srv *Service
var mutex sync.Mutex
var hasCreateDB = false

// 创建
func TestService_Create(t *testing.T) {
	initDB()

	req := addresspb.CreateRequest{
		Name:      "王五",
		Phone:     until.RandomString(11),
		IsDefault: uint32(1),
		PostCode:  "511450",
		Province:  "广东省",
		City:      "广州市",
		Region:    "天河区",
		Address:   "xxx小区3单元1号",
		MemberId:  37,
	}
	resp, err := srv.Create(context.Background(), &req)
	if err != nil {
		t.Errorf("create error:%s", err)
	}
	assert.Equal(t, id.AddressID(1).Uint64(), resp.Id)

}

// 更新（涉及到：更新、获取列表接口）
func TestService_Update(t *testing.T) {
	initDB()

	memberId := uint64(61)
	ctx := context.Background()

	// 创建
	req := addresspb.CreateRequest{
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
	resp, err := srv.Create(ctx, &req)
	if err != nil {
		t.Errorf("create error:%s", err)
	}

	// 更新
	entity := addresspb.Entity{
		Id:        resp.Id,
		MemberId:  req.MemberId,
		Name:      "张三",
		Phone:     until.RandomString(11),
		IsDefault: uint32(1),
		PostCode:  "412000",
		Province:  "长沙市",
		City:      "株洲市",
		Region:    "芦淞区",
		Address:   "xxx小区8单元3号",
	}
	_, err = srv.Update(ctx, &entity)
	if err != nil {
		t.Errorf("update error:%s", err)
	}

	// 校验
	list, err := srv.GetList(ctx, &addresspb.ListRequest{MemberId: memberId})
	if err != nil {
		t.Errorf("get address list error:%s", err)
	}
	if len(list.GetList()) == 0 {
		t.Errorf("get list data empty")
	}

	item := list.GetList()[0]
	assert.Equal(t, entity.Id, item.Id)
	assert.Equal(t, entity.Phone, item.Phone)
	assert.Equal(t, entity.Name, item.Name)
	assert.Equal(t, entity.IsDefault, item.IsDefault)
	assert.Equal(t, entity.PostCode, item.PostCode)
	assert.Equal(t, entity.Province, item.Province)
	assert.Equal(t, entity.City, item.City)
	assert.Equal(t, entity.Region, item.Region)
	assert.Equal(t, entity.Address, item.Address)
}

// 数据库初始化
func initDB() {
	mutex.Lock()
	if hasCreateDB == false {
		// 创建数据表
		if err := init_table.Address(); err != nil {
			log.Panicf("create table error: %v", err)
		}
		hasCreateDB = true
	}
	mutex.Unlock()

	global.DB = mysqltesting.GormDB

	// 本地调试
	/*cfg := db.NewDefaultMysql()
	cfg.Database = "go_mall"
	cfg.Host = "127.0.0.1"
	cfg.Password = "123456"
	_ = cfg.InitDB()
	global.DB = cfg.GetDB()

	srv = &Service{
		AddressDao: address.NewAddress(),
		Logger:     logger.NewDefaultLogger(),
	}*/
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
