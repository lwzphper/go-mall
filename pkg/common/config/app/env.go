package app

type Env string

const (
	EnvDevelopment Env = "development"
	EnvTest        Env = "test"
	EnvProduction  Env = "production"
)

func (e Env) String() string {
	return string(e)
}
