package mysqltesting

import (
	"database/sql"
	"fmt"
	"github.com/lwzphper/go-mall/pkg/common/config/db"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
	"time"
)

// https://cloud.tencent.com/developer/article/2221128
// https://github.com/ory/dockertest

type TestDBSetting struct {
	Driver       string
	ImageName    string
	ImageVersion string
	Database     string
	ENV          []string
	PortID       string
}

var GormDB *gorm.DB
var connDB *sql.DB
var dockerPort string

var (
	password  = "123456"
	dbSetting = TestDBSetting{
		Driver: "mysql",
		//ImageName:    "mariadb",
		//ImageVersion: "10.4.7",
		ImageName:    "mysql/mysql-server",
		ImageVersion: "8.0.28",
		Database:     "go_mall_test",
		ENV:          []string{fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", password), "MYSQL_ROOT_HOST=%"},
		PortID:       "3306/tcp",
	}
)

// RunMysqlInDocker 在容器中运行 mysql
func RunMysqlInDocker(m *testing.M) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 5
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: dbSetting.ImageName,
		Tag:        dbSetting.ImageVersion,
		Env:        dbSetting.ENV,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("could not pull resource: %s", err)
	}

	var runCode int

	defer func() {
		//// You can't defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}

		os.Exit(runCode)
	}()

	dockerPort = resource.GetPort(dbSetting.PortID)
	if err := pool.Retry(func() error {
		// 创建数据库
		dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:%s)/?charset=utf8mb4&multiStatements=true", password, dockerPort)
		connDB, err = sql.Open("mysql", dsn)
		if err != nil {
			return err
		}
		if err = connDB.Ping(); err != nil {
			return err
		}
		return connDB.Ping()
	}); err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	// 创建数据库
	_, err = connDB.Exec(fmt.Sprintf("create database %s", dbSetting.Database))
	if err != nil {
		log.Fatalf("create database error: %s", err)
	}
	_ = connDB.Close()

	// 初始化 gorm 数据库
	initMysql()

	runCode = m.Run()
}

// 初始化 mysql
func initMysql() {
	mysqlCfg := db.NewDefaultMysql()
	mysqlCfg.Port = dockerPort
	mysqlCfg.Database = dbSetting.Database
	mysqlCfg.Password = password

	err := mysqlCfg.InitDB()
	if err != nil {
		log.Fatalf("init mysql db error: %s", err)
	}

	GormDB = mysqlCfg.GetDB()
}
