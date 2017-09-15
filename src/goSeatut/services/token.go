package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("tranhuuphuoc")

// Token defines a token for uor application
type Token string

// TokenService provides a token
type TokenService interface {
	Get(u *User)(string, error)
}

type tokenService struct {
	UserService UserService
}

func NewTokenService() TokenService {
	return &tokenService{
		UserService: NewUserService(),
	}
}

// Get retrieves a token for a user
func (s *tokenService) Get(u *User)(string, error) {
	claims := jwt.MapClaims{
			// Set token claims
		"admin": true,
		"user": u,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	user, err := s.UserService.Read(u.ID)
	if err != nil {
		return "", errors.New("Failed to retrieve user")
	}
	if user == nil {
		return "", errors.New("Failed to retrieve user")
	}

	// Sign token with key
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("Failed to retrieve user")
	}
	if user == nil {
		return "", errors.New("Failed to retrieve user")
	}

	return tokenString, nil
}
