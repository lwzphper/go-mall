package mysql

import (
	"database/sql"
	"fmt"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
	"time"
)

type User struct {
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
	err := initClientWithDocker(DefaultClient, "tmp_database")
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

// 从 docker 创建客户端(数据库创建失败，待排查)
func initClientWithDocker(clientName, dbname string, options ...Option) error {
	// 创建数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", mysqltesting.Username,
		mysqltesting.Password, mysqltesting.HostPort)
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
