package server

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run Server",
			Run: func(_ *cobra.Command, _ []string) {
				fmt.Println("Hello From Server")
			},
		},
	)
}
