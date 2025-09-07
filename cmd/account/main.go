package account

import (
	"github.com/spf13/cobra"
)

func init() {
	AccountCmd.AddCommand(CreateCmd)
}

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account management",
}
