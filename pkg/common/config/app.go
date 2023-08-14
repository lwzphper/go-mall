package config

func NewDefaultApp() *App {
	return &App{
		Name:    "go-mall",
		Address: ":8081",
	}
}

type App struct {
	Name    string `toml:"name" yaml:"name" mapstructure:"name" env:"APP_NAME"`
	Address string `toml:"address" yaml:"address" mapstructure:"address" env:"APP_ADDRESS"`
}
