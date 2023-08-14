package db

func NewDefaultMongoDB() *Mongodb {
	return &Mongodb{
		Host: []string{"127.0.0.1:27017"},
	}
}

type Mongodb struct {
	Database string   `toml:"database" yaml:"database" mapstructure:"database" env:"MONGODB_DATABASE"`
	UserName string   `toml:"username" yaml:"username" mapstructure:"username" env:"MONGODB_USERNAME"`
	Password string   `toml:"password" yaml:"password" mapstructure:"password" env:"MONGODB_PASSWORD"`
	Host     []string `toml:"host" yaml:"host" mapstructure:"host" env:"MONGODB_HOST"`
}
