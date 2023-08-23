package db

import (
	"context"
	"database/sql"
	"fmt"
	mysqlDB "github.com/lwzphper/go-mall/pkg/db/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

type Mysql struct {
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
	lock        sync.Mutex
}

func NewDefaultMysql() *Mysql {
	return &Mysql{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Database:    "default_db",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
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
	if err := mysqlDB.InitMysqlClient(mysqlDB.DefaultClient, m.UserName, m.Password, m.Host, m.Database, options...); err != nil {
		return err
	}
	db = mysqlDB.GetMysqlClient(mysqlDB.DefaultClient).DB
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
