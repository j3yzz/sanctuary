package server

import (
	"github.com/j3yzz/sanctuary/internal/infrastructure/config"
	"github.com/j3yzz/sanctuary/internal/infrastructure/db"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run Server",
			Run: func(_ *cobra.Command, _ []string) {
				cfg := config.Provide()

				db.New(db.Config{
					PostgresHost:     cfg.PostgresHost,
					PostgresPort:     cfg.PostgresPort,
					PostgresDatabase: cfg.PostgresDatabase,
					PostgresUser:     cfg.PostgresUser,
					PostgresPassword: cfg.PostgresPassword,
				})

			},
		},
	)
}
