package handlers

import (
	"net/http"
	"goSeatut/services"
)

// Token expors an API to the token service
type Tokens struct {
	Service services.TokenService
}

// NewTokens creates new handle for tokens
func NewTokens(s services.TokenService) *Tokens {
	return &Tokens{s}
}

// Handler will return tokens
func (t *Tokens) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			user := &services.User{
				ID: 1,
				FirstName: "Big",
				LastName: "Boss",
				Roles: []string{services.AdministratorRole},
			}
			token, err := t.Service.Get(user)
			if err != nil {
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			}
			w.Write([]byte(token))
		
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}