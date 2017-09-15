package main

import (
	"fmt"
	"github.com/markbates/goth/gothic"
	"log"
	"net/http"
	"strings"
)

// clientID
// 208960536631-ac8p6vrd49t9i9lhgmasm4c8eqp4vnqe.apps.googleusercontent.com

// client secrect
// qVX8UOu6a7qZTD0ml7WH1lJ-

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// sucess - call the next handler
	h.next.ServeHTTP(w, r)
}

// loginHandler handles the third-party login process
// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]

	switch action {
	case "login":
		if gothUser, err := gothic.CompleteUserAuth(w, r); err != nil {
			fmt.Println("gothUser:Login ", gothUser)
			fmt.Println("gothUser:URL \n", r.URL)
			gothic.BeginAuthHandler(w, r)
		} else {
			// try to get the user without re-authenticating
			log.Println("TODO handle login[re-authenticate] for", provider)
		}
	case "callback":
		gothUser, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Println("gothUser:Auth ", gothUser)
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
