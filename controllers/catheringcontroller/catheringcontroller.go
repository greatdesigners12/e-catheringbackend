package catheringcontroller

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

func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var Cathering []models.Cathering
	models.DB.Find(&Cathering)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Cathering, "statusCode":200})
	if err := models.DB.Find(&Cathering).Error; err != nil {
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
	var Cathering []models.Cathering
	result := models.DB.First(&Cathering, id)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.RowsAffected == 0)
	if result.Error != nil {
		response, _  := json.Marshal(map[string]any{"status": "failed", "message": result.Error})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}else if result.RowsAffected == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(map[string]any{"status": "failed", "message": "No Catherings Found"})
		w.Write(status)
	}else{
		response, _  := json.Marshal(map[string]any{"status": "success","data":Cathering, "statusCode":200})
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}


	
}

func SearchAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	search := mux.Vars(r)["search"]
	var Product []models.Product
	models.DB.Raw("SELECT * FROM products WHERE nama LIKE ?", "%" + search + "%").Scan(&Product)
	var Category []models.Category
	models.DB.Raw("SELECT * FROM categories WHERE nama_kategori LIKE ?", "%" + search + "%").Scan(&Category)
	var Cathering []models.Cathering
	models.DB.Raw("SELECT * FROM catherings WHERE nama LIKE ?", "%" + search + "%").Scan(&Cathering)
	response, _  := json.Marshal(map[string]any{"status": "success","products":Product, "categories" : Category, "catherings" : Cathering,  "statusCode":200})
	if err := models.DB.Find(&Product).Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	w.Write(response)
}

func GetAllCatheringByGenre(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	genre := mux.Vars(r)["genre"]
	var Cathering []models.Cathering
	result := models.DB
	if(genre == "cathering"){
		result = result.Raw("SELECT c.*, AVG(r.rating) as average_rating FROM catherings as c LEFT JOIN reviews as r ON c.id = r.cathering_id GROUP BY c.id").Find(&Cathering)
	}else if(genre == "PopularCathering"){
		result = result.Raw("SELECT c.*, AVG(r.rating) as average_rating FROM catherings as c LEFT JOIN reviews as r ON c.id = r.cathering_id GROUP BY c.id ORDER BY AVG(r.rating) DESC").Scan(&Cathering)
	}else{
		result = result.Raw("SELECT c.*, AVG(r.rating) as average_rating FROM catherings as c LEFT JOIN reviews as r ON c.id = r.cathering_id WHERE c.need_preorder = 0 GROUP BY c.id ORDER BY AVG(r.rating) DESC").Find(&Cathering)
	}
	
	if result.Error != nil {
		response, _  := json.Marshal(map[string]any{"status": "failed", "message": result.Error})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}else if result.RowsAffected == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(map[string]any{"status": "failed", "message": "No Catherings Found"})
		w.Write(status)
	}else{
		response, _  := json.Marshal(map[string]any{"status": "success","data":Cathering, "statusCode":200})
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}


	
}

func Create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
    var Cathering models.Cathering
    decoder.Decode(&Cathering)
	result := models.DB.Create(&Cathering)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else{
		w.WriteHeader(http.StatusOK)
		status,_ := json.Marshal(map[string]any{"status": "success","data":Cathering, "statusCode":200})
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
		
		var Cathering1 models.Cathering
		json.Unmarshal(data, &Cathering1)
		
		
		result := models.DB.Save(&Cathering1)
		
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			
			status,_ := json.Marshal(result.Error)
			w.Write(status)
		}else{
			w.WriteHeader(http.StatusOK)
			status,_ := json.Marshal(map[string]any{"data": Cathering1, "success" : true, "message": "Data has been updated"})
			w.Write(status)
		}
	}
	
	
	
}

func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	decoder := json.NewDecoder(r.Body)
	
	var idData models.IdData
	var Cathering models.Cathering
	decoder.Decode(&idData)

	
	result := models.DB.Delete(&Cathering, idData.Id)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else if result.RowsAffected == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(map[string]any{"status": "failed", "message": "No Catherings Found", "statusCode":100})
		w.Write(status)
	}else{
		w.WriteHeader(http.StatusOK)
		status,_ := json.Marshal(map[string]any{"status": "success", "statusCode":200})
		w.Write(status)
	}

}