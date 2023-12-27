package config

type Config struct {
	PostgresDatabase string `koanf:"POSTGRES_DB"`
	PostgresUser     string `koanf:"POSTGRES_USER"`
	PostgresPassword string `koanf:"POSTGRES_PASSWORD"`
}
