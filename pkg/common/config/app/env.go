package app

type Env string

const (
	ENV_DEVELOPMENT Env = "development"
	ENV_TEST        Env = "test"
	ENV_PRODUCTION  Env = "production"
)

func (e Env) String() string {
	return string(e)
}
