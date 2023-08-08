package mysql

import (
	"database/sql"
	"fmt"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

type UserTmp struct {
	gorm.Model
	Name string `gorm:"size:255;index:idx_name,unique"`
}

func TestDB(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:32786)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	db.SetConnMaxLifetime(time.Second * 300)
	defer db.Close()
	fmt.Println(dsn)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("create database %s", "test_ddd"))
	if err != nil {
		log.Panicf("create db error:%v", err)
	}
}

func TestCreate(t *testing.T) {
	//err := initClientWithDocker(DefaultClient, "mysql")
	err := InitMysqlClient(DefaultClient, "root", "123456", "127.0.0.1:3306", "test")
	if err != nil {
		t.Errorf("init mysql error: %v", err)
	}
	defer CloseMysqlClient(DefaultClient)

	ormDB := GetMysqlClient(DefaultClient).DB

	// 创建数据表
	if err = ormDB.AutoMigrate(&UserTmp{}); err != nil {
		t.Errorf("create table error: %v", err)
	}

	// 添加数据
	user := UserTmp{
		Name: "test",
	}
	if err = ormDB.Create(&user).Error; err != nil {
		t.Errorf("insert error %v", err)
	}

	// 校验插入的数据
	queryUser := &UserTmp{}
	ormDB.First(queryUser)
	assert.Equal(t, "test", queryUser.Name)

	// 删除数据表
	ormDB.Exec("drop table user_tmp")
}

/*func TestMain(m *testing.M) {
	os.Exit(mysqltesting.RunWithMongoInDocker(m))
}*/

// 从 docker 创建客户端
func initClientWithDocker(clientName, dbname string, options ...Option) error {
	// todo 如果创建 docker 的同时，连接 mysql 会报错：driver: bad connection
	// 创建数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", mysqltesting.Username, mysqltesting.Password, mysqltesting.HostPort)
	db, err := sql.Open("mysql", dsn)
	db.SetConnMaxLifetime(time.Second * 300)
	defer db.Close()
	fmt.Println(dsn)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("create database %s", dbname))
	if err != nil {
		log.Panicf("create db error:%v", err)
	}

	// 初始化连接
	return InitMysqlClient(
		clientName,
		mysqltesting.Username,
		mysqltesting.Password,
		mysqltesting.HostPort,
		dbname,
		options...)
}
