package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	defaultTokenDuration = 24 * time.Hour
)

type Manager interface {
	Verify(token string) (Payload, error)
	Generate(userID string, roles []string, permissions []string) (string, error)
}

type Payload struct {
	UserID      string   `json:"sub"`
	Permissions []string `json:"permissions,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	jwt.StandardClaims
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

func (m implManager) Generate(userID string, roles []string, permissions []string) (string, error) {
	now := time.Now()

	claims := Payload{
		UserID:      userID,
		Roles:       roles,
		Permissions: permissions,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(defaultTokenDuration).Unix(),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}
