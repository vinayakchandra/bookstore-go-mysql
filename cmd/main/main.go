package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vinayakchandra/bookstore-go-mysql/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("server started at 9010...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
