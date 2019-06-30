package main

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run equator server",
	Long:  "serve initializes then starts the equator HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		initApp(cmd, args)
		app.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
