package config

import "time"

type Jwt struct {
	Secret string        `toml:"secret" yaml:"secret" mapstructure:"host" env:"JWT_SECRET"`
	TTL    time.Duration `toml:"ttl" yaml:"ttl" mapstructure:"host" env:"JWT_TTL"`
}

func NewJwt() *Jwt {
	return &Jwt{
		Secret: "go-mail",
		TTL:    24 * time.Hour,
	}
}
