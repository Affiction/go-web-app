package main

import (
	"fmt"
	"net/http"

	"github.com/Affiction/go-web-app/cmd/pkg/handlers"
)

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server starting on port http://localhost%s\n", port)
	_ = http.ListenAndServe(port, nil)
}
