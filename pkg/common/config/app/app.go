package app

func NewDefaultApp() *App {
	return &App{
		Name: "go-mall",
		Addr: ":8081",
		Env:  EnvDevelopment,
	}
}

type App struct {
	Name string `toml:"name" yaml:"name" mapstructure:"name" env:"APP_NAME"`
	Addr string `toml:"address" yaml:"address" mapstructure:"address" env:"APP_ADDRESS"`
	Env  Env    `toml:"env" yaml:"env" mapstructure:"env" env:"APP_ENV"`
}
