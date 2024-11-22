package config

type Config struct {
	Env *Enviroment
}

type Enviroment struct {
	SERVER_PORT       int    `mapstructure:"SERVER_PORT" validate:"required"`
	DATABASE_HOST     string `mapstructre:"DATABASE_HOST" validate:"required"`
	DATABASE_PORT     int    `mapstructre:"DATABASE_HOST" validate:"required"`
	DATABASE_SSL_MODE string `mapstructre:"DATABASE_HOST" validate:"required"`
	DATABASE_NAME     string `mapstructre:"DATABASE_HOST" validate:"required"`
	DATABASE_PASSWORD string `mapstructre:"DATABASE_HOST" validate:"required"`
	DATABASE_USERNAME string `mapstructure:"DATABASE_USERNAME" validate:"required"`
}
