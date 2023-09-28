package config

import (
	"github.com/lwzphper/go-mall/pkg/common/config"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	configDB "github.com/lwzphper/go-mall/pkg/common/config/db"
)

type Config struct {
	App     *app.App        `toml:"app" yaml:"app" mapstructure:"app"`
	Mysql   *configDB.Mysql `toml:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Logging *config.Logging `toml:"logging" yaml:"logging" mapstructure:"logging"`
}

func NewDefaultConfig() *Config {
	return &Config{
		App:     app.NewDefaultApp(),
		Mysql:   configDB.NewDefaultMysql(),
		Logging: config.NewDefaultLogging(),
	}
}
