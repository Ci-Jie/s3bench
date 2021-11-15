package cmd

import (
	"s3bench/controller"

	"github.com/spf13/cobra"
)

func newStartCmd() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "start",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			controller.Running()
		},
	}
}
