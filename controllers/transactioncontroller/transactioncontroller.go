package transactioncontroller

import (
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/rest-api/golang/helper"

	"github.com/rest-api/golang/models"
	"golang.org/x/crypto/bcrypt"
)

type TransactionRequest struct{
	Products []models.TransactionProduct
	Cathering_id int
	User_id int
	Shipping_price int
	Total_price int
	Daily_time string 
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var result TransactionRequest
    
	
	// var TransactionProduct models.TransactionProduct
	err := decoder.Decode(&result)
	fmt.Println(err)
	var TransactionGroup models.TransactionGroup
	TransactionGroup.Shipping_price = int64(result.Shipping_price)
	TransactionGroup.TotalPrice = int64(result.Total_price)
	TransactionGroup.Status = "Belum dibayar"
	TransactionGroup.User_id = int64(result.User_id)
	TransactionGroup.Cathering_id = int64(result.Cathering_id)
	TransactionGroup.DateTransaction = time.Now()
	TransactionGroup.Daily_time = result.Daily_time
	fmt.Println(result)
	models.DB.Create(&TransactionGroup)
	models.DB.Create(&result.Products)
	for _, product := range result.Products {
		var TransactionRelation models.TransactionGroupRelation
		TransactionRelation.Transaction_group_id = TransactionGroup.Id
		TransactionRelation.Transaction_product_id = product.Id
		models.DB.Create(&TransactionRelation)
	  }
	// result := models.DB.Create(&Product)
	// if result.Error != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	status,_ := json.Marshal(result.Error)
	// 	w.Write(status)
	// }else{
	// 	w.WriteHeader(http.StatusOK)
	// 	status,_ := json.Marshal(map[string]any{"status": "success","data":Product, "statusCode":200})
	// 	w.Write(status)
	// }
}

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success", "status" : "200"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token yang ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil", "status" : "200"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
