package main

import (
	"log"
	"net/http"

	"github.com/rest-api/golang/middlewares"

	"github.com/rest-api/golang/controllers/authcontroller"
	"github.com/rest-api/golang/controllers/productcontroller"
	"github.com/rest-api/golang/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	route := mux.NewRouter()

	route.HandleFunc("/login", authcontroller.Login).Methods("POST")
	route.HandleFunc("/register", authcontroller.Register).Methods("POST")
	route.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	r := route.PathPrefix("").Subrouter()
	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/getProduct/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/createProduct", productcontroller.Create).Methods("POST")
	r.HandleFunc("/updateProduct", productcontroller.Update).Methods("POST")
	r.HandleFunc("/deleteProduct", productcontroller.Delete).Methods("DELETE")
	// r.Use(middlewares.JWTMiddleware)
	api := route.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", route))

}
