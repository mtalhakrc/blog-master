package config

var cfg *AppConfig

type AppConfig struct {
	Db     DatabaseConfig
	Server ServerConfig
}

type DatabaseConfig struct {
	Name     string
	Host     string
	User     string
	Password string
	Port     string
}

type ServerConfig struct {
	JwtSecret string
}

func Setup() *AppConfig {
	cfg = &AppConfig{
		Db: DatabaseConfig{
			Name:     "blog-master",
			Host:     "localhost",
			User:     "postgres",
			Password: "postgres",
			Port:     "5432",
		},
		Server: ServerConfig{
			JwtSecret: "talha-gizli-jwt",
		},
	}
	return cfg
}
