package config

type Jaeger struct {
	Host string `toml:"host" yaml:"host" mapstructure:"host" env:"JAEGER_HOST"`
	Port string `toml:"port" yaml:"port" mapstructure:"port" env:"JAEGER_PORT"`
	Name string `toml:"name" yaml:"name" mapstructure:"name" env:"JAEGER_NAME"`
}

func NewJaeger() *Jaeger {
	return &Jaeger{
		Host: "127.0.0.1",
		Port: "6831",
		Name: "go-mail-jaeger",
	}
}
