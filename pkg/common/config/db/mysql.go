package db

import (
	"context"
	"database/sql"
	"fmt"
	mysqlDB "github.com/lwzphper/go-mall/pkg/db/mysql"
	"github.com/lwzphper/go-mall/pkg/file"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

type Mysql struct {
	ClientName  string `toml:"client_name" yaml:"client_name" mapstructure:"client_name" env:"MYSQL_NAME"`
	Host        string `toml:"host" yaml:"host" mapstructure:"host" env:"MYSQL_HOST"`
	Port        string `toml:"port" yaml:"port" mapstructure:"port" env:"MYSQL_PORT"`
	UserName    string `toml:"username" yaml:"username" mapstructure:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" yaml:"password" mapstructure:"password" env:"MYSQL_PASSWORD"`
	Database    string `toml:"database" yaml:"database" mapstructure:"database" env:"MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" yaml:"max_open_conn" mapstructure:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" yaml:"max_idle_conn" mapstructure:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" yaml:"max_life_time" mapstructure:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" yaml:"max_idle_time" mapstructure:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
	TablePrefix string `toml:"table_prefix" yaml:"table_prefix" mapstructure:"table_prefix" env:"MYSQL_TABLE_PREFIX"`
	LogFileName string `toml:"log_file_name" yaml:"log_file_name" mapstructure:"log_file_name" env:"MYSQL_LOG_FILE_NAME"`
	lock        sync.Mutex
	LogLevel    logger.LogLevel
}

func NewDefaultMysql() *Mysql {
	return &Mysql{
		ClientName:  mysqlDB.DefaultClient,
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Database:    "default_db",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
		MaxLifeTime: 1800,
		LogLevel:    logger.Info,
	}
}

// GetDB 获取 gorm 对象
func (m *Mysql) GetDB() *gorm.DB {
	if db == nil {
		panic("数据库未初始化")
	}
	return db
}

func (m *Mysql) InitDB() error {
	options := []mysqlDB.Option{
		mysqlDB.WithMaxOpenConn(m.MaxOpenConn),
		mysqlDB.WithMaxIdleConn(m.MaxIdleConn),
		mysqlDB.WithConnMaxLifeSecond(time.Duration(m.MaxLifeTime)),
		mysqlDB.WithMaxIdleTime(time.Duration(m.MaxIdleTime)),
	}

	// 设置日志
	logConf := mysqlDB.NewDefaultLoggerConf()
	logConf.LogLevel = m.LogLevel
	var logWriter io.Writer
	if m.LogFileName == "" {
		logWriter = os.Stdout
	} else {
		err := file.IsNotExistMkDir(filepath.Dir(m.LogFileName))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot create dir：%s", m.LogFileName))
		}

		open, err := os.OpenFile(m.LogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return errors.Wrap(err, "cannot open mysql log file")
		}
		logWriter = open
	}
	options = append(options, mysqlDB.WithLogger(logWriter, logConf))

	// 初始化客户端
	if err := mysqlDB.InitMysqlClient(m.ClientName, m.UserName, m.Password, m.Host, m.Port, m.Database, options...); err != nil {
		return err
	}
	db = mysqlDB.GetMysqlClient(m.ClientName).DB
	return nil
}

// GetSqlDB 获取 sql 对象
func (m *Mysql) GetSqlDB() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if sqlDB == nil {
		conn, err := m.getSqlDBConn()
		if err != nil {
			return nil, err
		}
		sqlDB = conn
	}
	return sqlDB, nil
}

// 用于初始化操作
func (m *Mysql) getSqlDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	sqlDB.SetMaxOpenConns(m.MaxOpenConn)
	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime != 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleTime != 0 {
		sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return sqlDB, nil
}
