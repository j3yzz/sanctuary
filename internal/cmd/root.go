package cmd

import (
	"github.com/j3yzz/sanctuary/internal/cmd/migration"
	"github.com/j3yzz/sanctuary/internal/cmd/server"
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	cmd := &cobra.Command{
		Use:   "Sanctuary",
		Short: "A simple book store search engine",
	}

	migration.Register(cmd)
	server.Register(cmd)

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
