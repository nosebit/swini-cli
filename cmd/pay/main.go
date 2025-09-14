package pay

import (
	"github.com/spf13/cobra"

	"swini-cli/cmd/pay/account"
	"swini-cli/cmd/pay/method"
)

func init() {
	MainCmd.AddCommand(account.MainCmd)
	MainCmd.AddCommand(method.MainCmd)
}

var MainCmd = &cobra.Command{
	Use:   "pay",
	Short: "Pay management",
}
