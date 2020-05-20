package main

import (
	"fmt"
	"github.com/fivetentaylor/goinbase/market"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"time"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	market.HelloWorld()

	r := mux.NewRouter()
	r.HandleFunc("/products", ProductsHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
