package config

import (
	"github.com/lwzphper/go-mall/pkg/common/config"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
)

type Config struct {
	App     *app.App        `toml:"app" yaml:"app" mapstructure:"app"`
	Logging *config.Logging `toml:"logging" yaml:"logging" mapstructure:"logging"`
	Jaeger  *config.Jaeger  `toml:"jaeger" yaml:"jaeger" mapstructure:"jaeger"`
	Jwt     *config.Jwt     `toml:"jwt" yaml:"jwt" mapstructure:"jwt"`
}

func NewDefaultConfig() *Config {
	return &Config{
		App:     app.NewDefaultApp(),
		Logging: config.NewDefaultLogging(),
		Jaeger:  config.NewJaeger(),
		Jwt:     config.NewJwt(),
	}
}
