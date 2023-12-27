package server

import (
	"github.com/j3yzz/sanctuary/internal/infrastructure/config"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run Server",
			Run: func(_ *cobra.Command, _ []string) {
				config.Provide()
			},
		},
	)
}
