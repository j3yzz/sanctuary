package migration

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/j3yzz/sanctuary/internal/infrastructure/config"
	"github.com/j3yzz/sanctuary/internal/infrastructure/db"
	"github.com/j3yzz/sanctuary/internal/infrastructure/db/migration"
	"github.com/spf13/cobra"
	"log"
)

const migrationPath = "migrations"

func Register(root *cobra.Command) {
	cfg := config.Provide()
	gormDB := db.New(db.Config{
		PostgresHost:     cfg.PostgresHost,
		PostgresPort:     cfg.PostgresPort,
		PostgresDatabase: cfg.PostgresDatabase,
		PostgresUser:     cfg.PostgresUser,
		PostgresPassword: cfg.PostgresPassword,
	})

	m, err := migration.New(gormDB, migrationPath)
	if err != nil {
		log.Fatalln(err)
	}

	root.AddCommand(
		&cobra.Command{
			Use:   "migrate",
			Short: "Run migration files",
			Run: func(_ *cobra.Command, _ []string) {
				err := m.Up()
				if err != nil {
					log.Fatalln(err)
				}
			},
		},
	)

	root.AddCommand(
		&cobra.Command{
			Use:   "migrate:rollback",
			Short: "Rollback migration files",
			Run: func(_ *cobra.Command, _ []string) {
				err := m.Down()
				if err != nil {
					log.Fatalln(err)
				}
			},
		},
	)
}
