package main

import (
	"log"
	"net/http"

	"github.com/rest-api/golang/middlewares"

	"github.com/rest-api/golang/controllers/authcontroller"
	"github.com/rest-api/golang/controllers/productcontroller"
	"github.com/rest-api/golang/controllers/categorycontroller"
	"github.com/rest-api/golang/controllers/catheringcontroller"
	"github.com/rest-api/golang/controllers/reviewcontroller"
	"github.com/rest-api/golang/controllers/cartcontroller"
	"github.com/rest-api/golang/controllers/transactioncontroller"
	"github.com/rest-api/golang/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	route := mux.NewRouter()

	route.HandleFunc("/login", authcontroller.Login).Methods("POST")
	route.HandleFunc("/register", authcontroller.Register).Methods("POST")
	route.HandleFunc("/resetPassword", authcontroller.ResetPassword).Methods("POST")
	route.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	r := route.PathPrefix("").Subrouter()
	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/{cathering_id}/getAllDailyProducts", productcontroller.GetAllDailyProduct).Methods("GET")
	r.HandleFunc("/getAllProductsWithCartChecker", productcontroller.GetAllProductsWithCartChecker).Methods("GET")
	r.HandleFunc("/getProduct/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/createTransaction", transactioncontroller.CreateTransaction).Methods("POST")
	r.HandleFunc("/searchAll/", catheringcontroller.SearchAll).Methods("GET")
	r.HandleFunc("/searchAll/{search}", catheringcontroller.SearchAll).Methods("GET")
	r.HandleFunc("/createProduct", productcontroller.Create).Methods("POST")
	r.HandleFunc("/updateProduct/{product_id}", productcontroller.Update).Methods("POST")
	r.HandleFunc("/deleteProduct", productcontroller.Delete).Methods("DELETE")
	r.HandleFunc("/getAllCategories", categorycontroller.Index).Methods("GET")
	r.HandleFunc("/getCategory/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/setTransaction/success/{id_transaction}", transactioncontroller.SetToSuccess).Methods("GET")
	r.HandleFunc("/createCategory", productcontroller.Create).Methods("POST")
	r.HandleFunc("/updateCategory/{id}", categorycontroller.Update).Methods("POST")
	r.HandleFunc("/deleteCategory", productcontroller.Delete).Methods("DELETE")
	r.HandleFunc("/getAllCarts", cartcontroller.Index).Methods("GET")
	r.HandleFunc("/getCart/{id}", cartcontroller.Show).Methods("GET")
	r.HandleFunc("/getCart", cartcontroller.GetCartBasedOnCathering).Methods("GET")
	r.HandleFunc("/getAllCartProduct", cartcontroller.GetAddedProductByCathering).Methods("GET")
	r.HandleFunc("/getAllTransactionGroup/{user_id}", transactioncontroller.GetAllTransactionGroups).Methods("GET") 
	r.HandleFunc("/getTransactionGroup/{transaction_id}", transactioncontroller.GetTransactionGroupById).Methods("GET")
	r.HandleFunc("/getCartByUserId", cartcontroller.GetCartBasedOnUserId).Methods("GET")
	r.HandleFunc("/createCart", cartcontroller.Create).Methods("POST")
	r.HandleFunc("/updateCart", cartcontroller.Update).Methods("POST")
	r.HandleFunc("/removeCart", cartcontroller.RemoveCart).Methods("DELETE")
	r.HandleFunc("/setSnapToken/{id_transaction}", transactioncontroller.SetSnapToken).Methods("GET")
	r.HandleFunc("/getAllProductByUserId/{user_id}", productcontroller.GetAllProductByUserId).Methods("GET")
	r.HandleFunc("/deleteProduct/{product_id}", productcontroller.DeleteProductByProductId).Methods("DELETE")
	r.HandleFunc("/getUserReview", reviewcontroller.GetUserReview).Methods("GET")
	r.HandleFunc("/searchProduct/{search}", productcontroller.SearchProduct).Methods("GET")
	r.HandleFunc("/createUserReviews", reviewcontroller.CreateUserReviews).Methods("POST")
	r.HandleFunc("/getAllCatheringByPopularity", catheringcontroller.GetAllCatheringByPopularity).Methods("GET")
	r.HandleFunc("/resetIdTransaction/{id_transaction}", transactioncontroller.SetNewId).Methods("GET")


	rCathering := route.PathPrefix("").Subrouter()
	rCathering.HandleFunc("/catherings", catheringcontroller.Index).Methods("GET")
	rCathering.HandleFunc("/getCathering/{id}", catheringcontroller.Show).Methods("GET")
	rCathering.HandleFunc("/getCatheringByGenre/{genre}", catheringcontroller.GetAllCatheringByGenre).Methods("GET")
	rCathering.HandleFunc("/createCathering", catheringcontroller.Create).Methods("POST")
	rCathering.HandleFunc("/updateCathering/{id}", catheringcontroller.Update).Methods("POST")
	rCathering.HandleFunc("/deleteCathering", catheringcontroller.Delete).Methods("DELETE")
	// r.Use(middlewares.JWTMiddleware)
	api := route.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", route))

}
