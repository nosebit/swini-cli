package account

import (
	"github.com/spf13/cobra"
)

func init() {
	MainCmd.AddCommand(SetupCmd)
}

var MainCmd = &cobra.Command{
	Use:   "account",
	Short: "Pay account management",
}
