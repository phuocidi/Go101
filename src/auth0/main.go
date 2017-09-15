package main

// All credit goes to
// https://auth0.com/blog/authentication-in-golang/
// Use only for practicing purpose and nothing else
import (
	"encoding/json"
	"fmt"
	"github.com/auth0-community/go-auth0"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"os"
	"strings"
)

// --------------------------------------------------------
// Object meta
// --------------------------------------------------------

type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

type Response struct {
	Message string `json:"message"`
}

// --------------------------------------------------------
// Global vars
// --------------------------------------------------------

/* We will create our catalog of VR experiences and store them in a slice. */
var products = []Product{
	Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	Product{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	Product{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

var AUTH0_CLIENT_SECRET = os.Getenv("AUTH0_CLIENT_SECRET")
var AUTH0_API_AUDIENCE = os.Getenv("AUTH0_API_AUDIENCE") //os.Getenv("AUTH0_IDENTIFIER")
var AUTH0_ISSUER_URL = os.Getenv("AUTH0_ISSUER_URL")

var JWKS_URI = os.Getenv("JWKS_URI")
var AUTH0_API_ISSUER = os.Getenv("AUTH0_API_ISSUER")

// --------------------------------------------------------
// Handlers
// --------------------------------------------------------

/* The feedback handler will add either positive or negative feedback to the product
   We would normally save this data to the database - but for this demo we'll fake it
   so that as long as the request is successful and we can match a product to our catalog of products
   we'll return an OK status. */
var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})

// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

/* The products handler will be called when the user makes a GET request to the /products endpoint.
   This handler will return a list of products available for users to review */
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Here we are converting the slice of products to json
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

/* The status handler will be invoked when the user calls the /status route
   It will simply return a string with the message "API is up and running" */
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

// --------------------------------------------------------
// Middleware
// --------------------------------------------------------

// Secure our API endpoint with Auth0
func authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKS_URI})
		audience := []string{AUTH0_API_AUDIENCE}

		configuration := auth0.NewConfiguration(client, audience, AUTH0_API_ISSUER, jose.RS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)
		fmt.Println("\n\n ", r, "\n\n")

		if err != nil {
			fmt.Println("Token is not valid or missing token")
			fmt.Println(err)
			fmt.Println("Token is not valid: ", token)

			response := Response{
				Message: "Missing or invalid token",
			}

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
		} else {
			next.ServeHTTP(w, r)
		}

	})

	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// 	AUTH0_CLIENT_SECRET := os.Getenv("AUTH0_CLIENT_SECRET")
	// 	AUTH0_API_AUDIENCE := os.Getenv("AUTH0_API_AUDIENCE") //os.Getenv("AUTH0_IDENTIFIER")
	// 	AUTH0_ISSUER_URL := os.Getenv("AUTH0_ISSUER_URL")

	// 	secret := []byte(AUTH0_CLIENT_SECRET)
	// 	secretProvider := auth0.NewKeyProvider(secret)
	// 	audience := []string(AUTH0_API_AUDIENCE)

	// 	configuration := auth0.NewConfiguration(secretProvider, audience, AUTH0_ISSUER_URL, jose.RS256)
	// 	validator := auth0.NewValidator(configuration)

	// 	token, err := validator.ValidateRequest(r)

	// 	if err != nil {
	// 		fmt.Println(r.URL)
	// 		fmt.Println(err)
	// 		fmt.Println("Token is not valid: ", token)
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		w.Write([]byte("Unauthorized"))
	// 	} else {
	// 		next.ServeHTTP(w, r)
	// 	}
	// })
}

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Our API is going to consist of three routes
	// /status - which we will call to make sure that our API is up and running
	// /products - which will retrieve a list of products that the user can leave feedback on
	// /products/{slug}/feedback - which will capture user feedback on products
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", authMiddleware(ProductsHandler)).Methods("GET")
	r.Handle("/products/{slug}/feedback", authMiddleware(AddFeedbackHandler)).Methods("POST")

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

func checkScope(r *http.Request, validator *auth0.JWTValidator, token *jwt.JSONWebToken) bool {
	claims := map[string]interface{}{}
	err := validator.Claims(r, token, &claims)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if strings.Contains(claims["scope"].(string), "read:messages") {
		return true
	} else {
		return false
	}
}
