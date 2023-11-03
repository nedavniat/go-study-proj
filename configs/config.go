package config

var Conf struct {
	App struct {
		Port string `envconfig:"APP_PORT"`
	}
	DB struct {
		URL string `envconfig:"DATABASE_URL"`
	}
}


