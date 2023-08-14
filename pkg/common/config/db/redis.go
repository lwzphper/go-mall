package db

func NewDefaultRedis() *Redis {
	return &Redis{
		Host: "127.0.0.1:6379",
	}
}

type Redis struct {
	Host        string `toml:"host" yaml:"host" mapstructure:"host" env:"REDIS_HOST"`
	Password    string `toml:"password" yaml:"password" mapstructure:"password" env:"REDIS_PASSWORD"`
	Database    int    `toml:"database" yaml:"database" mapstructure:"database" env:"REDIS_DATABASE"`
	MinIdleConn int    `toml:"min_idle_conn" yaml:"min_idle_conn" mapstructure:"min_idle_conn" env:"REDIS_MIN_IDLE_CONN"`
	PoolSize    int    `toml:"pool_size" yaml:"pool_size" mapstructure:"pool_size" env:"REDIS_POOL_SIZE"`
	MaxRetries  int    `toml:"max_retries" yaml:"max_retries" mapstructure:"max_retries" env:"REDIS_MAX_RETRIES"`
}
