package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(config Config) *GormDatabase {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDatabase,
		config.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error in connection to mysql: %v", err)
	}

	database := new(GormDatabase)
	database.DB = db

	return database
}
