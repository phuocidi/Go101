package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"goSeatut/handlers"
	"goSeatut/services"
)

type API struct {
	encryptionKey []byte
	AclService services.ACLService
	
	Hello *handlers.Hello
	Tokens *handlers.Tokens
	Users *handlers.Users
}

// NewAPI creates a new API
func NewAPI(certPath, keyPath string) *API {
	aclService := services.NewACLService()
	tokenService := services.NewTokenService()
	userService := services.NewUserService()
	helloService := services.NewHelloService()

	return &API {
		encryptionKey: []byte("tranhuuphuoc"),
		AclService: aclService,
		Tokens: handlers.NewTokens(tokenService),
		Hello: handlers.NewHello(helloService),
		Users: handlers.NewUsers(userService),
	}
}

// Middleware
// Authenticate provides Authentication middleware for handlers
func (a *API) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var tokenString string
		// Get token from the Authorization header
		// format: Authorization: Bearer <token>
		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			tokenString = tokens[0]
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		// If the token is empty
		if tokenString == "" {
			http.Error(w,http.StatusText(http.StatusUnauthorized),http.StatusUnauthorized)
			return
		}
		// Now parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"] )
			}
			return a.encryptionKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["admin"], claims["user"], claims["exp"])
		} else {
			fmt.Println(err)
		}

		if token != nil && token.Valid {
			// Everything worked! Set the user in the context
			context.Set(r, "user", token)
			next.ServeHTTP(w,r)
			fmt.Println("Authenticated sucessfully")
			return
		}

		// Token is invalid
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	})
}

// Authorization provides authorization middleware for our handlers
func (a *API) Authorize(permissions ...services.Permission) func (next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
			// TODO: Get User information from Request
			user := &services.User{
				ID: 1,
				FirstName: "Admin",
				LastName: "User",
				Roles: []string{services.AdministratorRole},
			}

			for _, permission := range permissions {
				if err := a.AclService.CheckPermission(user, permission); err !=nil {
					http.Error(w, http.StatusText(http.StatusForbidden),http.StatusForbidden)
					return
				}
			}
			next.ServeHTTP(w,r)
		})
	}
}

// SecureHeaders add adds secure headers to the API
func (a *API) SecureHeaders( next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// TODOL add security headers here
		next.ServeHTTP(w,r)
	})
}

