package migration

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migrateDriver "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/j3yzz/sanctuary/internal/infrastructure/db"
)

type Migration struct {
	db            *db.GormDatabase
	migrationPath string
	migrate       *migrate.Migrate
}

func New(db *db.GormDatabase, migrationPath string) (*Migration, error) {
	migrationStruct := &Migration{
		db:            db,
		migrationPath: migrationPath,
	}

	sqlDB, _ := migrationStruct.db.DB.DB()
	instance, err := migrateDriver.WithInstance(sqlDB, &migrateDriver.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in connection to database: %v", err))
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationStruct.migrationPath),
		"postgres",
		instance,
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in creating migration database instance: %v", err))
	}

	migrationStruct.migrate = m

	return migrationStruct, nil
}

func (m *Migration) Up() error {
	err := m.migrate.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.New(fmt.Sprintf("apply migrations: %v", err))
	}

	err, _ = m.migrate.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("close migration instance: %v", err))
	}

	return nil
}

func (m *Migration) Down() error {
	err := m.migrate.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.New(fmt.Sprintf("dowm migrations: %v", err))
	}

	err, _ = m.migrate.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("close migration instance: %v", err))
	}

	return nil
}
