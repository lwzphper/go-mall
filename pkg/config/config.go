package config

import (
	"errors"
	"github.com/caarlos0/env/v9"
	fsnotify2 "github.com/fsnotify/fsnotify"
	"github.com/lwzphper/go-mall/pkg/logger"
	"github.com/spf13/viper"
)

type FileType uint8

const (
	TomlFileType FileType = iota
	YamlFIleType
)

// LoadConfigFromToml 从 toml 中添加配置文件
func LoadConfigFromToml(filePath string, cfg interface{}) error {
	return loadConfig(filePath, TomlFileType, cfg)
}

// LoadConfigFromYml 从 yaml 中添加配置文件
func LoadConfigFromYml(filePath string, cfg interface{}) error {
	return loadConfig(filePath, YamlFIleType, cfg)
}

// LoadConfigFromEnv 从 env 中添加配置文件
func LoadConfigFromEnv(cfg interface{}) error {
	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}

func loadConfig(filePath string, fileType FileType, cfg interface{}) error {
	v := viper.New()
	// 设置配置文件路径
	v.SetConfigFile(filePath)
	// 设置文件类型
	switch fileType {
	case TomlFileType:
		v.SetConfigType("toml")
	case YamlFIleType:
		v.SetConfigType("yaml")
	default:
		return errors.New("file type not define")
	}
	// 读取配置
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// 将配置映射结构体
	if err := v.Unmarshal(cfg); err != nil {
		return err
	}

	// 监听配置文件变动,重新解析配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify2.Event) {
		if err := v.Unmarshal(cfg); err != nil {
			logger.Error("yml config watch error" + err.Error())
		}
	})
	return nil
}
