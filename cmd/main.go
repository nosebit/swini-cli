package cmd

import (
	"fmt"
	"swini-cli/cmd/account"
	"swini-cli/cmd/pay"
	"swini-cli/internal/config"
	"swini-cli/internal/graphql"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "swini",
	Short: "Swini CLI",
}

func Execute() {
	cfg, err := config.Load()

	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	graphql.InitSharedClient(cfg.ApiUrl, &clientv2.Options{})

	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	mainCmd.AddCommand(account.MainCmd)
	mainCmd.AddCommand(pay.MainCmd)
	mainCmd.AddCommand(VersionCmd)
}
