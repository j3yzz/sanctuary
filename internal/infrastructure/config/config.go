package config

type Config struct {
	PostgresHost     string `koanf:"POSTGRES_HOST"`
	PostgresPort     string `koanf:"POSTGRES_PORT"`
	PostgresDatabase string `koanf:"POSTGRES_DB"`
	PostgresUser     string `koanf:"POSTGRES_USER"`
	PostgresPassword string `koanf:"POSTGRES_PASSWORD"`
}
