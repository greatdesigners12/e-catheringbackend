package reviewcontroller

import (
	"encoding/json"

	"net/http"
	

	"github.com/rest-api/golang/models"
	

)

func CreateUserReviews(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	

	decoder := json.NewDecoder(r.Body)
	var review models.Review
	decoder.Decode(&review)
    result := models.DB.Create(&review)

	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	status,_ := json.Marshal(map[string]any{"data": review, "success" : true, "message": "Data has been updated"})
	w.Write(status)
	

}

func GetUserReview(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	user_id := r.FormValue("user_id")
	cathering_id := r.FormValue("cathering_id")
	decoder := json.NewDecoder(r.Body)
	var review []models.Review
	decoder.Decode(&review)
    result := models.DB.Where("user_id", user_id).Where("cathering_id", cathering_id).Find(&review)

	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	status,_ := json.Marshal(map[string]any{"data": review, "success" : true, "message": "Data has been updated"})
	w.Write(status)
	

}