package auth

import (
	"net/http"
	"swini-cli/internal/crypto"
	"swini-cli/internal/localstore"

	"time"
)

// GenerateAuthHeaders loads the private key, signs the current datetime,
// and builds HTTP headers (to attach in a GraphQL request).
func HttpHeaders() (http.Header, error) {
	store, err := localstore.Load()

	if err != nil {
		return nil, err
	}

	headers := http.Header{}

	// Acount is absent
	if store.Account.ID == "" || store.Account.PvtKey == "" {
		return headers, nil
	}

	// Create the signature
	now := time.Now().UTC().Format(time.RFC3339)
	msg := store.Account.ID + now
	sig, err := crypto.SignMessage(store.Account.PvtKey, msg)

	if err != nil {
		return nil, err
	}

	// Create headers
	headers.Set("X-Swini-AccountID", store.Account.ID)
	headers.Set("X-Swini-Datetime", now)
	headers.Set("X-Swini-Signature", sig)

	return headers, nil
}
