package handlers 

import (
	"net/http"
	"goSeatut/services"
)

// Users provides API access to the user service
type Users struct {
	Service services.UserService
}

// NewUsers creates a users handler
func NewUsers(s services.UserService) *Users {
	return &Users{s}
}

// Handler handles requests to /users
func (u *Users) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":
		case "POST":
		case "PUT":
		case "DELETE":
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
} 
