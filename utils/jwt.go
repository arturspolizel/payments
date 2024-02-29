package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
