package categorycontroller

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
	
	var Category []models.Category
	models.DB.Find(&Category)
	response, _  := json.Marshal(map[string]any{"status": "success","data":Category, "statusCode":200})
	if err := models.DB.Find(&Category).Error; err != nil {
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
	var Category []models.Category
	result := models.DB.First(&Category, id)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.RowsAffected == 0)
	if result.Error != nil {
		response, _  := json.Marshal(map[string]any{"status": "failed", "message": result.Error})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}else if result.RowsAffected == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(map[string]any{"status": "failed", "message": "No Categorys Found"})
		w.Write(status)
	}else{
		response, _  := json.Marshal(map[string]any{"status": "success","data":Category, "statusCode":200})
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}


	
}

func Create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
    var Category models.Category
    decoder.Decode(&Category)
	result := models.DB.Create(&Category)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else{
		w.WriteHeader(http.StatusOK)
		status,_ := json.Marshal(map[string]any{"status": "success","data":Category, "statusCode":200})
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
		
		var Category1 models.Category
		json.Unmarshal(data, &Category1)
		
		
		result := models.DB.Save(&Category1)
		
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			
			status,_ := json.Marshal(result.Error)
			w.Write(status)
		}else{
			w.WriteHeader(http.StatusOK)
			status,_ := json.Marshal(map[string]any{"data": Category1, "success" : true, "message": "Data has been updated"})
			w.Write(status)
		}
	}
	
	
	
}

func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	decoder := json.NewDecoder(r.Body)
	
	var idData models.IdData
	var Category models.Category
	decoder.Decode(&idData)

	
	result := models.DB.Delete(&Category, idData.Id)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else if result.RowsAffected == 0{
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(map[string]any{"status": "failed", "message": "No Categorys Found", "statusCode":100})
		w.Write(status)
	}else{
		w.WriteHeader(http.StatusOK)
		status,_ := json.Marshal(map[string]any{"status": "success", "statusCode":200})
		w.Write(status)
	}

}