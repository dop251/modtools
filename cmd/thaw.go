package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "thaw modpath",
		Short: "Un-freeze a dependency",
		Long:  "Removes the specified module path from the list of exceptions.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return thaw(args[0])
		},
	})
}

func thaw(p string) error {
	e, err := loadExceptions()
	if err != nil {
		return err
	}
	e.Remove(p)
	return e.Save()
}
