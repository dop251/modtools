package cmd

import (
	"os"

	"github.com/muesli/coral"
)

var (
	rootCmd = &coral.Command{
		Use: "modtools",
		PersistentPreRun: func(cmd *coral.Command, args []string) {
			cmd.SilenceUsage = true
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
