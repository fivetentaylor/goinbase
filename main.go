package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/btree"
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
	tree := btree.NewWithStringComparator(16)
	tree.Put("hello", "world")
	tree.Put("goodbye", "friends")
	tree.Put("yes", "no")
	tree.Put("bye", "hello")
	tree.Put("bye", 12)
	fmt.Println(tree.Keys())
	fmt.Println(tree.Values())

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
