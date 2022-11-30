package main

import (
	"log"
	"net/http"

	"github.com/rest-api/golang/middlewares"

	"github.com/rest-api/golang/controllers/authcontroller"
	"github.com/rest-api/golang/controllers/productcontroller"
	"github.com/rest-api/golang/controllers/foodcontroller"
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
	r.HandleFunc("/foods", foodcontroller.Index).Methods("GET")
	r.HandleFunc("/getFood/{id}", foodcontroller.Show).Methods("GET")
	r.HandleFunc("/createFood", foodcontroller.Create).Methods("POST")
	r.HandleFunc("/updateFood", foodcontroller.Update).Methods("POST")
	r.HandleFunc("/deleteFood/{id}", foodcontroller.Delete).Methods("POST")
	r.Use(middlewares.JWTMiddleware)
	api := route.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")


	log.Fatal(http.ListenAndServe(":8080", route))

}
