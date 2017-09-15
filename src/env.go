package main

import (
	"fmt"
	"os"
)

func main() {
	auth_client_id := os.Getenv("AUTH_CLIENT_SECRET")
	auth_identifier := os.Getenv("AUTH_IDENTIFIER")
	fmt.Println(auth_client_id)
	fmt.Println(auth_identifier)
}
