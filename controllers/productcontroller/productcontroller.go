package productcontroller

import (
	"encoding/json"
	"fmt"

	// "fmt"

	
	"net/http"

	// "strconv"

	"github.com/gorilla/mux"
	"github.com/rest-api/golang/models"
)

func GetAllProductsWithCartChecker(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	user_id := r.FormValue("user_id")
	cathering_id := r.FormValue("cathering_id")
	priceOrder := r.FormValue("price_order")
	productType := r.FormValue("product_type")
	var Product []models.Product
	var products_id []int
	sqlStatement := "SELECT * FROM products WHERE cathering_id = ? AND ((CAST( NOW() AS Date ) BETWEEN start_date and due_date) OR start_date IS NULL OR due_date IS NULL)"
	// SELECT * FROM products WHERE (CAST( NOW() AS Date ) BETWEEN start_date and due_date OR start_date IS NULL ; 
	if(productType == "d"){
		sqlStatement = sqlStatement + "AND type = 'daily'"
		
	}else if(productType == "r"){
		sqlStatement = sqlStatement + "AND type = 'recurring'"
	}



	if(priceOrder == "h"){
		sqlStatement = sqlStatement + "ORDER BY harga DESC"
	}else if(priceOrder == "l"){
		sqlStatement = sqlStatement + "ORDER BY harga ASC"
	}

	

	models.DB.Raw(sqlStatement, cathering_id).Find(&Product)

	
	
	models.DB.Raw("SELECT c.product_id  FROM products as p LEFT JOIN carts AS c ON p.id = c.product_id WHERE c.user_id = ? AND c.cathering_id = ? ", user_id, cathering_id).Scan(&products_id)
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
	var Product models.Product
	result := models.DB.First(&Product, id)

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
	decoder := json.NewDecoder(r.Body)
	product_id := mux.Vars(r)["product_id"]
	
		
		var Product1 models.Product
		decoder.Decode(&Product1)
		
		fmt.Println(Product1.Harga)
		result := models.DB.Model(&Product1).Where("id", product_id).Updates(Product1)
		
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

func GetAllProductByUserId(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["user_id"]
	var Product []models.Product
	var Cathering models.Cathering
	models.DB.Where("user_id", id).First(&Cathering)
	models.DB.Where("type=?", "daily").Where("cathering_id", Cathering.Id).Find(&Product)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Product, "statusCode":200})
	if err := models.DB.Find(&Product).Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
}

func DeleteProductByProductId(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["product_id"]
	var Product models.Product

	models.DB.Delete(&Product, id)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Product, "statusCode":200})
	if err := models.DB.Find(&Product).Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
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