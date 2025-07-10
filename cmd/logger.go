package cmd

import (
	loggergo "github.com/Alonza0314/logger-go/v2"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Create a info log.",
	Long:    "Use this command to create a info log.",
	Example: `logger-go info [tag] [message]`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loggergo.Info(args[0], args[1])
	},
}

var errorCmd = &cobra.Command{
	Use:     "error",
	Short:   "Create a error log.",
	Long:    "Use this command to create a error log.",
	Example: `logger-go error [tag] [message]`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loggergo.Error(args[0], args[1])
	},
}

var warnCmd = &cobra.Command{
	Use:     "warn",
	Short:   "Create a warn log.",
	Long:    "Use this command to create a warn log.",
	Example: `logger-go warn [tag] [message]`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loggergo.Warn(args[0], args[1])
	},
}

var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "Create a test log.",
	Long:    "Use this command to create a test log.",
	Example: `logger-go test [tag] [message]`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loggergo.Test(args[0], args[1])
	},
}

var debugCmd = &cobra.Command{
	Use:     "debug",
	Short:   "Create a debug log.",
	Long:    "Use this command to create a debug log.",
	Example: `logger-go debug [tag] [message]`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loggergo.Debug(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(errorCmd)
	rootCmd.AddCommand(warnCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(debugCmd)
}
