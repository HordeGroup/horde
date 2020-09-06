package config

type Config struct {
	Debug bool `json:"debug" default:"true"`

	Logger struct {
		Level uint32 `env:"LOGGER_LEVEL" default:"1"`
	}

	Server struct {
		Host    string `env:"SERVER_HOST"`
		Port    int    `env:"SERVER_PORT"`
		Verbose bool   `env:"SERVER_VERBOSE" default:"false"`
	}

	Database struct {
		Host     string `env:"DB_HOST"`
		Port     int64  `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PWD"`
		Name     string `env:"DB_NAME"`
	}

	Redis struct {
		Host     string `env:"REDIS_HOST"`
		Port     int64  `env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD"`
	}
	Session struct {
		TTL int64 `env:"SESSION_TTL"`
	}
}
