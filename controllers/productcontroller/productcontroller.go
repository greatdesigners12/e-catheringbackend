package productcontroller

import (
	"encoding/json"
	"fmt"

	// "fmt"

	"io"
	"net/http"

	// "strconv"

	"github.com/gorilla/mux"
	"github.com/rest-api/golang/models"
)

func GetAllProductsWithCartChecker(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	user_id := r.FormValue("user_id")
	cathering_id := r.FormValue("cathering_id")
	var Product []models.Product
	var products_id []int
	models.DB.Where("cathering_id", cathering_id).Find(&Product)
	models.DB.Raw("SELECT c.product_id  FROM products as p LEFT JOIN carts AS c ON p.id = c.product_id WHERE c.user_id = ? AND c.cathering_id = ?  ", user_id, cathering_id).Scan(&products_id)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Product, "carts" : products_id, "statusCode":200})
	if err := models.DB.Find(&Product).Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
}

func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var Product []models.Product
	models.DB.Model(&Product).Preload("Cathering").Find(&Product)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Product, "statusCode":200})
	if err := models.DB.Find(&Product).Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
}

func GetAllDailyProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["cathering_id"]
	var Product []models.Product
	models.DB.Where("type=?", "daily").Where("cathering_id", id).Find(&Product)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Product, "statusCode":200})
	if err := models.DB.Find(&Product).Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
}

func Show(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var Product []models.Product
	result := models.DB.First(&Product, id)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.RowsAffected == 0)
	if result.Error != nil {
		response, _  := json.Marshal(map[string]any{"status": "failed", "message": result.Error})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}else if result.RowsAffected == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(map[string]any{"status": "failed", "message": "No Products Found"})
		w.Write(status)
	}else{
		response, _  := json.Marshal(map[string]any{"status": "success","data":Product, "statusCode":200})
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}


	
}

func Create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
    var Product models.Product
    decoder.Decode(&Product)
	result := models.DB.Create(&Product)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else{
		w.WriteHeader(http.StatusOK)
		status,_ := json.Marshal(map[string]any{"status": "success","data":Product, "statusCode":200})
		w.Write(status)
	}
	
}

func Update(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	data, _ := io.ReadAll(r.Body)
	
	if len(data) == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal("Please insert some value first")
		w.Write(status)
	}else{
		
		var Product1 models.Product
		json.Unmarshal(data, &Product1)
		
		
		result := models.DB.Save(&Product1)
		
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			
			status,_ := json.Marshal(result.Error)
			w.Write(status)
		}else{
			w.WriteHeader(http.StatusOK)
			status,_ := json.Marshal(map[string]any{"data": Product1, "success" : true, "message": "Data has been updated"})
			w.Write(status)
		}
	}
	
	
	
}

func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	decoder := json.NewDecoder(r.Body)
	
	var idData models.IdData
	var Product models.Product
	decoder.Decode(&idData)

	
	result := models.DB.Delete(&Product, idData.Id)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else if result.RowsAffected == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(map[string]any{"status": "failed", "message": "No Products Found", "statusCode":100})
		w.Write(status)
	}else{
		w.WriteHeader(http.StatusOK)
		status,_ := json.Marshal(map[string]any{"status": "success", "statusCode":200})
		w.Write(status)
	}

}