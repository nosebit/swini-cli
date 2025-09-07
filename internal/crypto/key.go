package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateAndStoreKeyPair(privatePath string) (string, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}

	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return "", err
	}

	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privBytes,
	}

	f, err := os.OpenFile(privatePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if err := pem.Encode(f, pemBlock); err != nil {
		return "", err
	}

	pubkey := elliptic.MarshalCompressed(priv.PublicKey.Curve, priv.PublicKey.X, priv.PublicKey.Y)
	return encodeHex(pubkey), nil
}

func encodeHex(b []byte) string {
	const hex = "0123456789abcdef"
	out := make([]byte, len(b)*2)
	for i, v := range b {
		out[i*2] = hex[v>>4]
		out[i*2+1] = hex[v&0x0f]
	}
	return string(out)
}
