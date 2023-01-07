package transactioncontroller

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"net/http"

	"github.com/rest-api/golang/models"

	"math/rand"

	"github.com/gorilla/mux"
)

type TransactionRequest struct {
	Products       []models.TransactionProduct
	Carts_id       []int
	Cathering_id   int
	User_id        int
	Shipping_price int
	Total_price    int
	Full_address   string
	Status         string
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
	TransactionGroup.Status = "Belum dibayar"
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
	status, _ := json.Marshal(map[string]any{"status": "success", "data": result, "statusCode": 200})
	w.Write(status)

}

func GetAllPaidGroups(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")
	cathering_id := mux.Vars(r)["cathering_id"]

	var Transactions []models.TransactionGroup

	result := models.DB.Where(map[string]interface{}{"cathering_id": cathering_id, "status": "Terbayar"}).Or(map[string]interface{}{"cathering_id": cathering_id, "status": "Dalam Pengantaran"}).Preload("TransactionGroupRelation.TransactionProduct").Preload("Cathering").Find(&Transactions)
	response, _ := json.Marshal(map[string]any{"status": "success", "data": Transactions, "statusCode": 200})
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(response)

}

func UpdatePaidGroups(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	data, _ := io.ReadAll(r.Body)
	id := mux.Vars(r)["id"]

	if len(data) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		status, _ := json.Marshal("Please insert some value first")
		w.Write(status)
	} else {

		var TransactionGroup models.TransactionGroup
		json.Unmarshal(data, &TransactionGroup)

		result := models.DB.Model(&TransactionGroup).Where("id", id).Updates(&TransactionGroup)

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)

			status, _ := json.Marshal(result.Error)
			w.Write(status)
		} else {
			w.WriteHeader(http.StatusOK)
			status, _ := json.Marshal(map[string]any{"data": TransactionGroup, "success": true, "message": "Data has been updated"})
			w.Write(status)
		}
	}
}

func DetailPaidGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["transaction_group_id"]
	w.Header().Set("Content-Type", "application/json")

	var transactionProducts []models.TransactionProduct
	result := models.DB.Debug().Raw(`
    SELECT tp.name, tp.price, tp.time
    FROM transaction_groups AS tg
    INNER JOIN transaction_group_relations AS tgr ON tgr.transaction_group_id = tg.id
    INNER JOIN transaction_products AS tp ON tp.id = tgr.transaction_product_id WHERE tg.id = ?;
    `,id).Find(&transactionProducts)
	response, _ := json.Marshal(map[string]any{"status": "success", "data": transactionProducts, "statusCode": 200})
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(response)
}

func GetAllTransactionGroups(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")
	user_id := mux.Vars(r)["user_id"]

	var Transactions []models.TransactionGroup

	result := models.DB.Where("user_id", user_id).Preload("TransactionGroupRelation.TransactionProduct").Preload("Cathering").Find(&Transactions)
	response, _ := json.Marshal(map[string]any{"status": "success", "data": Transactions, "statusCode": 200})
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
	response, _ := json.Marshal(map[string]any{"status": "success", "data": Transactions, "statusCode": 200})
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(response)

}
