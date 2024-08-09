package main

import (
	"github/fdjrr/go-restapi-mux/controllers/productcontroller"
	"github/fdjrr/go-restapi-mux/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/products", productcontroller.Create).Methods("POST")
	r.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/products/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/products/{id}", productcontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
