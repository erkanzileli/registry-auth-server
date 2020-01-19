package auth

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/docker/libtrust"
)

var tokenExpiresIn = int64(90)

type header struct {
	Type      string `json:"typ"`
	Algorithm string `json:"alg"`
	KeyID     string `json:"kid"`
}

type payload struct {
	Issuer     string `json:"iss"`
	Subject    string `json:"sub"`
	Audience   string `json:"aud"`
	Expiration int64  `json:"exp"`
	NotBefore  int64  `json:"nbf"`
	IssuedAt   int64  `json:"iat"`
	JwtID      string `json:"jti"`

	// Private fields
	Access []access `json:"access"`
}

type access struct {
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

func readKeysFromCert(certFile, keyFile string) (libtrust.PublicKey, libtrust.PrivateKey) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		log.Fatal(err)
	}
	publicKey, err := libtrust.FromCryptoPublicKey(x509Cert.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := libtrust.FromCryptoPrivateKey(cert.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	return publicKey, privateKey
}

func resolveSignAlgFromPrivateKey(privateKey libtrust.PrivateKey) (sigAlg string) {
	_, sigAlg, err := privateKey.Sign(strings.NewReader("dummy"), 0)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// CreateToken func
func CreateToken(u *User, cert, key string) string {
	// Resolve keys from certificates
	publicKey, privateKey := readKeysFromCert(cert, key)
	sigAlg := resolveSignAlgFromPrivateKey(privateKey)

	// Header
	h := header{
		Type:      "JWT",
		Algorithm: sigAlg,
		KeyID:     publicKey.KeyID(),
	}
	headerJSON, err := json.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}

	// Get current time as seconds
	var now = time.Now()

	// Authorize user
	accesses := u.Authorize()

	// Token payload
	payload := payload{
		Issuer:     "auth.registry.ezileli.dev",
		Subject:    u.Username,
		Audience:   "registry.ezileli.dev",
		Expiration: now.Unix() + tokenExpiresIn,
		NotBefore:  now.Unix() - 5,
		IssuedAt:   now.Unix(),
		JwtID:      "Docker registry",
		Access:     accesses,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	headerAndPayload := fmt.Sprintf("%s%s%s", joseBase64UrlEncode(headerJSON), ".", joseBase64UrlEncode(payloadJSON))

	sig, sigAlg2, err := privateKey.Sign(strings.NewReader(headerAndPayload), 0)
	if err != nil || sigAlg2 != sigAlg {
		log.Fatal(err)
	}
	token := fmt.Sprintf("%s%s%s", headerAndPayload, ".", joseBase64UrlEncode(sig))

	return token
}

func joseBase64UrlEncode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}
