package method

import (
	"github.com/spf13/cobra"
)

func init() {
	MainCmd.AddCommand(SetupCmd)
}

var MainCmd = &cobra.Command{
	Use:   "method",
	Short: "Pay method management",
}
