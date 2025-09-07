package account

import (
	"context"
	"fmt"
	"os"
	"swini-cli/internal/crypto"
	"swini-cli/internal/graphql"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		swiniDir := home + "/.swini"
		if err := os.MkdirAll(swiniDir, 0700); err != nil {
			fmt.Println("Error creating .swini directory:", err)
			return
		}
		privPath := swiniDir + "/account.key"
		pubkey, err := crypto.GenerateAndStoreKeyPair(privPath)
		if err != nil {
			fmt.Println("Error generating key pair:", err)
			return
		}

		res, err := graphql.SharedClient.AccountCreate(context.TODO(), pubkey)

		if err != nil {
			fmt.Println("Error creating account:", err)
			return
		}

		userID := res.AccountCreate.GetUserID()

		idPath := swiniDir + "/user.id"
		if err := os.WriteFile(idPath, []byte(userID), 0600); err != nil {
			fmt.Println("Error saving user id:", err)
			return
		}
		fmt.Println("Account created. User ID:", userID)
	},
}
