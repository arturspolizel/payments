package utils

import (
	"crypto/ed25519"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ed25519PrivKey struct {
	Version          int
	ObjectIdentifier struct {
		ObjectIdentifier asn1.ObjectIdentifier
	}
	PrivateKey []byte
}

type ed25519PubKey struct {
	OBjectIdentifier struct {
		ObjectIdentifier asn1.ObjectIdentifier
	}
	PublicKey asn1.BitString
}

type JwtProcessor interface {
	NewToken(context TokenContext) (string, error)
	Validate(tokenString string) (TokenContext, error)
}

// TODO: Add group
type TokenContext struct {
	jwt.RegisteredClaims
	Email      string `json:"email"`
	MerchantId uint   `json:"merchantId"`
}

type JwtProcessorWithKey struct {
	key        ed25519.PublicKey
	signingKey ed25519.PrivateKey
	algorithm  jwt.SigningMethod
}

// NewJwtProcessor creates a processor with a public key for validating tokens. Key must be
// in PEM format
func NewJwtProcessor(key []byte, algorithm jwt.SigningMethod) *JwtProcessorWithKey {
	pubKey := decodePublicKey(key)
	return &JwtProcessorWithKey{
		key:       pubKey,
		algorithm: algorithm,
	}
}

// NewJwtProcessorWithPrivate creates a processor with a public and private key for validating and
// signing tokens. Keys must be in PEM format.
func NewJwtProcessorWithPrivate(key, signingKey []byte, algorithm jwt.SigningMethod) *JwtProcessorWithKey {
	pubKey := decodePublicKey(key)
	privKey := decodePrivateKey(signingKey)
	return &JwtProcessorWithKey{
		key:        pubKey,
		signingKey: privKey,
		algorithm:  algorithm,
	}
}

func (j *JwtProcessorWithKey) NewToken(context TokenContext) (string, error) {
	if len(j.signingKey) == 0 {
		return "", fmt.Errorf("must provide a signing key for token creation")
	}
	token := jwt.NewWithClaims(j.algorithm, context)
	return token.SignedString(j.signingKey)
}

func (j *JwtProcessorWithKey) Validate(tokenString string) (TokenContext, error) {
	context := TokenContext{}
	token, err := jwt.ParseWithClaims(tokenString, &context, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	}, jwt.WithValidMethods([]string{j.algorithm.Alg()}))

	if err != nil {
		return TokenContext{}, err
	}

	if !token.Valid {
		return TokenContext{}, fmt.Errorf("invalid jwt token")
	}

	expirationDate, err := token.Claims.GetExpirationTime()
	if err != nil || time.Now().After(expirationDate.Time) {
		return TokenContext{}, fmt.Errorf("token has expired, please login or refresh")
	}

	return context, nil
}

func Logger(jwtProcessor JwtProcessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := strings.Split(c.Request.Header.Get("Authorization"), " ")

		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			// invalid token format
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format. Please use a Bearer schema."})
		}
		reqToken := bearerToken[1]

		token, err := jwtProcessor.Validate(reqToken)
		if err != nil {
			// unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature."})
		}
		// Set example variable
		c.Set("token", token)

		// before request
		c.Next()
	}
}

func decodePrivateKey(key []byte) ed25519.PrivateKey {
	// TODO: Support other algorithms
	var block *pem.Block
	block, _ = pem.Decode(key)

	var asn1PrivKey ed25519PrivKey
	asn1.Unmarshal(block.Bytes, &asn1PrivKey)

	privateKey := ed25519.NewKeyFromSeed(asn1PrivKey.PrivateKey[2:])
	return privateKey
}

func decodePublicKey(key []byte) ed25519.PublicKey {
	// TODO: Support other algorithms
	var block *pem.Block
	block, _ = pem.Decode(key)

	var asn1PubKey ed25519PubKey
	asn1.Unmarshal(block.Bytes, &asn1PubKey)

	publicKey := ed25519.PublicKey(asn1PubKey.PublicKey.Bytes)
	return publicKey
}
