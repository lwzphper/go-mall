package config

import "time"

type Jwt struct {
	Secret string        `toml:"secret" yaml:"secret" mapstructure:"secret" env:"JWT_SECRET"`
	TTL    time.Duration `toml:"ttl" yaml:"ttl" mapstructure:"ttl" env:"JWT_TTL"`
}

func NewJwt() *Jwt {
	return &Jwt{
		Secret: "go-mail",
		TTL:    24 * time.Hour,
	}
}

func (j *Jwt) GetSecret() []byte {
	return []byte(j.Secret)
}
