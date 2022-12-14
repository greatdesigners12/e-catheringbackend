package transactioncontroller

import (
	"encoding/json"
	"fmt"
	"time"
	"net/http"
	

	"github.com/rest-api/golang/models"
	
	"math/rand"
	"github.com/gorilla/mux"
)

type TransactionRequest struct{
	Products []models.TransactionProduct
	Carts_id []int
	Cathering_id int
	User_id int
	Shipping_price int
	Total_price int
	Full_address string 
	Status string
}


func SetToSuccess(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	id_transaction := mux.Vars(r)["id_transaction"]
	
	
		
	
		var transactionGroup models.TransactionGroup
		

		result := models.DB.Model(&transactionGroup).Where("id_transaction", id_transaction).Update("status", "Terbayar")
		
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			
			status,_ := json.Marshal(result.Error)
			w.Write(status)
		}else{
			w.WriteHeader(http.StatusOK)
			status,_ := json.Marshal(map[string]any{"data": transactionGroup, "success" : true, "message": "Data has been updated"})
			w.Write(status)
		}
	
}

func SetSnapToken(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	snap_token := r.FormValue("snap_token")
	id_transaction := mux.Vars(r)["id_transaction"]
	
		var transactionGroup models.TransactionGroup
		

		result := models.DB.Model(&transactionGroup).Where("id_transaction", id_transaction).Update("snap_token", snap_token)
		
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			
			status,_ := json.Marshal(result.Error)
			w.Write(status)
		}else{
			w.WriteHeader(http.StatusOK)
			status,_ := json.Marshal(map[string]any{"data": transactionGroup, "success" : true, "message": "Data has been updated"})
			w.Write(status)
		}
}



func SetNewId(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id_transaction := mux.Vars(r)["id_transaction"]
	
		var transactionGroup models.TransactionGroup
		
		rand.Seed(time.Now().UnixNano())

		var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
		str := make([]rune, 10)
		for i := range str {
			str[i] = chars[rand.Intn(len(chars))]
		}
		fmt.Println(id_transaction)

		result := models.DB.Model(&transactionGroup).Where("id_transaction", id_transaction).Update("id_transaction", string(str))
		
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			
			status,_ := json.Marshal(result.Error)
			w.Write(status)
		}else{
			w.WriteHeader(http.StatusOK)
			status,_ := json.Marshal(map[string]any{"data": string(str), "success" : true, "message": "Data has been updated"})
			w.Write(status)
		}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var result TransactionRequest
    
	// var TransactionProduct models.TransactionProduct
	err := decoder.Decode(&result)
	fmt.Println(err)
	rand.Seed(time.Now().UnixNano())

	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
    str := make([]rune, 10)
    for i := range str {
        str[i] = chars[rand.Intn(len(chars))]
    }
	var carts []models.Cart
	var TransactionGroup models.TransactionGroup
	TransactionGroup.Shipping_price = int64(result.Shipping_price)
	TransactionGroup.TotalPrice = int64(result.Total_price)
	TransactionGroup.User_id = int64(result.User_id)
	TransactionGroup.Cathering_id = int64(result.Cathering_id)
	TransactionGroup.DateTransaction = time.Now()
	TransactionGroup.Full_address = result.Full_address
	TransactionGroup.Id_transaction = string(str)
	TransactionGroup.Status = "Belum terbayar"
	
	models.DB.Create(&TransactionGroup)
	models.DB.Create(&result.Products)
	for _, product := range result.Products {
		var TransactionRelation models.TransactionGroupRelation
		
		TransactionRelation.Transaction_group_id = TransactionGroup.Id
		TransactionRelation.Transaction_product_id = product.Id
		models.DB.Create(&TransactionRelation)
	}
	fmt.Println(result.Carts_id)
	models.DB.Delete(&carts, result.Carts_id)
	
	
	w.WriteHeader(http.StatusOK)
	status,_ := json.Marshal(map[string]any{"status": "success", "data" : result	, "statusCode":200})
	w.Write(status)
	
}

func GetAllTransactionGroups(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	user_id := mux.Vars(r)["user_id"]
	
	var Transactions []models.TransactionGroup

	result := models.DB.Where("user_id", user_id).Preload("TransactionGroupRelation.TransactionProduct").Preload("Cathering").Order("id DESC").Find(&Transactions)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Transactions, "statusCode":200})
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
	
}

func GetTransactionGroupById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	
	transaction_id := mux.Vars(r)["transaction_id"]
	
	var Transactions models.TransactionGroup
	
	result := models.DB.Preload("TransactionGroupRelation.TransactionProduct").Preload("User.UserInformation").Find(&Transactions, transaction_id)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Transactions, "statusCode":200})
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
	
}

