package method

import (
	"context"
	"fmt"

	"swini-cli/internal/graphql"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup swini payment method",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for setting up payment would go here
		res, err := graphql.SharedClient.PayMethodSetup(context.TODO())

		if err != nil {
			fmt.Println("Error setting up payment:", err)
			return
		}

		setupURL := res.PayMethodSetup.GetURL()

		err = browser.OpenURL(setupURL)

		if err != nil {
			fmt.Println("Failed to open browser:", err)
		}
	},
}
