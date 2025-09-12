package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func KeyPairCreate() (privHex string, pubHex string, err error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}

	privHex = hex.EncodeToString(priv)
	pubHex = hex.EncodeToString(pub)

	return privHex, pubHex, nil
}

// SignMessage takes a hex-encoded private key and a message string,
// and returns the hex-encoded Ed25519 signature.
func SignMessage(privHex string, message string) (string, error) {
	// Decode hex private key
	privBytes, err := hex.DecodeString(privHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key hex: %w", err)
	}

	// Ensure the length is valid (ed25519.PrivateKey should be 64 bytes)
	if len(privBytes) != ed25519.PrivateKeySize {
		return "", fmt.Errorf("invalid private key length: got %d, want %d", len(privBytes), ed25519.PrivateKeySize)
	}

	// Sign the message
	sig := ed25519.Sign(privBytes, []byte(message))

	// Encode signature to hex
	return hex.EncodeToString(sig), nil
}

// PublicKeyFromPrivate takes a hex-encoded private key (64 bytes)
// and returns the associated public key (hex-encoded, 32 bytes).
func PubKeyFrom(privHex string) (string, error) {
	privBytes, err := hex.DecodeString(privHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key hex: %w", err)
	}

	if len(privBytes) != ed25519.PrivateKeySize {
		return "", fmt.Errorf("invalid private key length: got %d, want %d", len(privBytes), ed25519.PrivateKeySize)
	}

	// Extract the last 32 bytes (public key)
	pubBytes := privBytes[32:]

	return hex.EncodeToString(pubBytes), nil
}
