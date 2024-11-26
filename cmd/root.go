package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logger-go",
	Short: "This is a log tool.",
	Long:  "This is a log tool for simply create info/error/warn/test log.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
