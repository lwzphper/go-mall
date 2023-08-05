package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
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
		Database:    "default_db",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

type Mongodb struct {
	Database string   `toml:"database" yaml:"database" mapstructure:"database" env:"MONGODB_DATABASE"`
	UserName string   `toml:"username" yaml:"username" mapstructure:"username" env:"MONGODB_USERNAME"`
	Password string   `toml:"password" yaml:"password" mapstructure:"password" env:"MONGODB_PASSWORD"`
	Host     []string `toml:"host" yaml:"host" mapstructure:"host" env:"MONGODB_HOST"`
}

func NewDefaultMongoDB() *Mongodb {
	return &Mongodb{
		Host: []string{"127.0.0.1:27017"}, // 设置默认值
	}
}

type Config struct {
	Mysql   *Mysql   `toml:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Mongodb *Mongodb `toml:"Mongodb" yaml:"mongodb" mapstructure:"mongodb"`
}

func NewConfig() *Config {
	return &Config{
		Mysql:   NewDefaultMysql(),
		Mongodb: NewDefaultMongoDB(),
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	_ = os.Setenv("MYSQL_DATABASE", "go-mall")
	_ = os.Setenv("MYSQL_HOST", "127.0.0.1")
	_ = os.Setenv("MYSQL_PORT", "3306")
	_ = os.Setenv("MYSQL_USERNAME", "root")
	_ = os.Setenv("MYSQL_PASSWORD", "123456")
	_ = os.Setenv("MYSQL_MAX_OPEN_CONN", "1000")
	_ = os.Setenv("MYSQL_MAX_IDLE_CONN", "100")
	_ = os.Setenv("MYSQL_MAX_LIFE_SECOND", "30")
	_ = os.Setenv("MYSQL_TABLE_PREFIX", "")
	_ = os.Setenv("MONGODB_DATABASE", "shop")
	_ = os.Setenv("MONGODB_USERNAME", "root")
	_ = os.Setenv("MONGODB_PASSWORD", "123456")
	_ = os.Setenv("MONGODB_HOST", "127.0.0.1:27017,127.0.0.1:27018") // 数组逗号分割

	cfg := NewConfig()
	err := LoadConfigFromEnv(cfg)
	if err != nil {
		t.Errorf("load config from env error:%v", err)
	}

	assert.Equal(t, "go-mall", cfg.Mysql.Database)
	assert.Equal(t, "127.0.0.1:27018", cfg.Mongodb.Host[1])
}

func TestLoadConfigFromToml(t *testing.T) {
	cfg := NewConfig()
	err := LoadConfigFromToml("./test/config.toml", cfg)
	if err != nil {
		t.Errorf("load config from env error:%v", err)
	}

	assert.Equal(t, "go-mall", cfg.Mysql.Database)
	assert.Equal(t, "127.0.0.1:27018", cfg.Mongodb.Host[1])
}

func TestLoadConfigFromYml(t *testing.T) {
	cfg := NewConfig()
	err := LoadConfigFromYml("./test/config.yml", cfg)
	if err != nil {
		t.Errorf("load config from env error:%v", err)
	}

	assert.Equal(t, "go-mall", cfg.Mysql.Database)
	assert.Equal(t, "127.0.0.1:27018", cfg.Mongodb.Host[1])
}
