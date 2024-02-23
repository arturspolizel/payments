package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JwtProcessor interface {
	NewToken(context TokenContext) (string, error)
	Validate(tokenString string) (TokenContext, error)
}

// TODO: Add group
type TokenContext struct {
	jwt.RegisteredClaims
	Email      string
	MerchantId uint
}

type JwtProcessorWithKey struct {
	key        []byte
	signingKey []byte
	algorithm  jwt.SigningMethod
}

func NewJwtProcessor(key []byte, algorithm jwt.SigningMethod) *JwtProcessorWithKey {
	return &JwtProcessorWithKey{
		key:       key,
		algorithm: algorithm,
	}
}

func NewJwtProcessorWithPrivate(key, signingKey []byte, algorithm jwt.SigningMethod) *JwtProcessorWithKey {
	return &JwtProcessorWithKey{
		key:       key,
		algorithm: algorithm,
	}
}

func (j *JwtProcessorWithKey) NewToken(context TokenContext) (string, error) {
	if len(j.signingKey) == 0 {
		return "", fmt.Errorf("Must provide a signing key for token creation")
	}
	token := jwt.NewWithClaims(j.algorithm, context)
	return token.SignedString(j.signingKey)
}

func (j *JwtProcessorWithKey) Validate(tokenString string) (TokenContext, error) {
	context := TokenContext{}
	token, err := jwt.ParseWithClaims(tokenString, &context, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if err != nil {
		return TokenContext{}, err
	}

	if !token.Valid {
		return TokenContext{}, fmt.Errorf("invalid jwt token")
	}

	return context, nil
}
