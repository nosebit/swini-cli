package cmd

import (
	"fmt"
	"swini-cli/cmd/account"
	"swini-cli/internal/graphql"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "swini",
	Short: "Swini CLI",
}

func Execute() {
	graphql.InitSharedClient("http://localhost:8080/gql", &clientv2.Options{})

	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	mainCmd.AddCommand(account.AccountCmd)
}
