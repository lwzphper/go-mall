package config

import (
	"github.com/lwzphper/go-mall/pkg/common/config"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
)

type Config struct {
	App     *app.App        `toml:"app" yaml:"app" mapstructure:"app"`
	Logging *config.Logging `toml:"logging" yaml:"logging" mapstructure:"logging"`
}

func NewDefaultConfig() *Config {
	return &Config{
		App:     app.NewDefaultApp(),
		Logging: config.NewDefaultLogging(),
	}
}
