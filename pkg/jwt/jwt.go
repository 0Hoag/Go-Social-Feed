package jwt

import (
	"log"

	"github.com/golang-jwt/jwt"
)

type Manager interface {
	Verify(token string) (Payload, error)
}

type Payload struct {
	jwt.StandardClaims
	UserID  string `json:"sub"`
	Type    string `json:"type"`
	Refresh bool   `json:"refresh"`
}

type implManager struct {
	secretKey string
}

func NewManager(secretKey string) Manager {
	return &implManager{
		secretKey: secretKey,
	}
}

// Verify verifies the token and returns the payload
func (m implManager) Verify(token string) (Payload, error) {
	if token == "" {
		return Payload{}, ErrInvalidToken
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Printf("jwt.ParseWithClaims: %v", ErrInvalidToken)
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		log.Printf("jwt.ParseWithClaims: %v", err)
		return Payload{}, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		log.Printf("Parsing to Payload: %v", ErrInvalidToken)
		return Payload{}, ErrInvalidToken
	}

	return *payload, nil
}
