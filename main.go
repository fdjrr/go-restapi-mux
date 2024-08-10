package main

import (
	"github/fdjrr/go-restapi-mux/controllers/authcontroller"
	"github/fdjrr/go-restapi-mux/controllers/productcontroller"
	"github/fdjrr/go-restapi-mux/middlewares"
	"github/fdjrr/go-restapi-mux/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/login", authcontroller.Login).Methods("POST")
	api.HandleFunc("/register", authcontroller.Register).Methods("POST")

	apim := r.PathPrefix("/api").Subrouter()
	apim.HandleFunc("/logout", authcontroller.Logout).Methods("POST")
	apim.HandleFunc("/products", productcontroller.Index).Methods("GET")
	apim.HandleFunc("/products", productcontroller.Create).Methods("POST")
	apim.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	apim.HandleFunc("/products/{id}", productcontroller.Update).Methods("PUT")
	apim.HandleFunc("/products/{id}", productcontroller.Delete).Methods("DELETE")

	apim.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":3000", r))
}
