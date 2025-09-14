package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("swini version: 0.0.1")
	},
}
