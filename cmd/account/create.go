package account

import (
	"context"
	"fmt"
	"swini-cli/internal/crypto"
	"swini-cli/internal/graphql"
	"swini-cli/internal/localstore"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Run: func(cmd *cobra.Command, args []string) {
		store, err := localstore.Load()
		if err != nil {
			fmt.Println("Error loading local store:", err)
			return
		}

		if store.Account.ID != "" {
			fmt.Println("Account already created with ID:", store.Account.ID)
			return
		}

		pvtkey, pubkey, err := crypto.KeyPairCreate()
		if err != nil {
			fmt.Println("Error generating key pair:", err)
			return
		}

		res, err := graphql.SharedClient.AccountCreate(context.TODO(), pubkey)

		if err != nil {
			fmt.Println("Error creating account:", err)
			return
		}

		accountID := res.AccountCreate.GetID()

		store.Account.ID = accountID
		store.Account.PvtKey = pvtkey

		err = store.Save()
		if err != nil {
			fmt.Println("Error saving account to local store:", err)
			return
		}

		fmt.Println("Account created. ID:", accountID)
	},
}
