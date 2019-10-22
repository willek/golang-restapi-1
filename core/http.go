package core

import (
	"../route"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HTTP() {
	r := mux.NewRouter()
	r.HandleFunc("/", route.Handler)

	r.HandleFunc("/blog", route.Blogs).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Connected to port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
