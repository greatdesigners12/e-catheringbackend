package foodcontroller

import (
	"encoding/json"
	
	// "fmt"

	"net/http"
	"io"
	// "strconv"

	"github.com/gorilla/mux"
	"github.com/rest-api/golang/models"
)

func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var Food []models.Food
	models.DB.Limit(10).Find(&Food)
	response, _  := json.Marshal(Food)
	if err := models.DB.Find(&Food).Error; err != nil {
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
	var Food []models.Food
	result := models.DB.First(&Food, id)
	if result.Error != nil {
		response, _  := json.Marshal(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}else{
		response, _  := json.Marshal(Food)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
	
}

func Create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
    var food models.Food
    decoder.Decode(&food)
	result := models.DB.Create(&food)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else{
		w.WriteHeader(http.StatusOK)
		status,_ := json.Marshal("It's worked")
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
		
		var food1 models.Food
		json.Unmarshal(data, &food1)
		
		
		result := models.DB.Save(&food1)
		
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			
			status,_ := json.Marshal(result.Error)
			w.Write(status)
		}else{
			w.WriteHeader(http.StatusOK)
			status,_ := json.Marshal(map[string]any{"data": food1, "success" : true, "message": "Data has been updated"})
			w.Write(status)
		}
	}
	
	
	
}

func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var food models.Food
	id := mux.Vars(r)["id"]
	result := models.DB.Delete(&food, id)
	if result.Error != nil{
		status,_ := json.Marshal(result.Error)
		w.Write(status)
	}else{
		status,_ := json.Marshal("It's worked")
		w.Write(status)
	}

}