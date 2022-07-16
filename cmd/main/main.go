package main

import (
	"fmt"
	"log"
	"net/http"

	"fahadmdkamal.com/booksapi/pkg/routes"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
	fmt.Printf("Server is Listening at localhost:80000")
}