package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rest-api/golang/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rest-api/golang/helper"
	"gorm.io/gorm"

	"github.com/rest-api/golang/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error(), "token" : "", "status" : "400", "userId" : ""}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan email
	var user models.User
	if err := models.DB.Where("email = ?", userInput.Email).Preload("Role").First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email atau password salah", "token" : "", "status" : "400", "userId" : ""}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error(), "token" : "", "status" : "400", "userId" : ""}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Password salah", "token" : "", "status" : "400", "userId" : ""}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// medeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error(), "token" : "", "status" : "400", "userId" : ""}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]any{"message": "login berhasil", "role" : user.Role.Role,"token" : token, "status" : "200",  "userId" : user.Id}
	helper.ResponseJSON(w, http.StatusOK, response)
}

type ResetPasswordRequest struct{
	User_id  int;
	OldPassword string;
	NewPassword string;

}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var data ResetPasswordRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	var user models.User
	if err := models.DB.Where("id = ?", data.User_id).Preload("Role").First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email atau password salah", "token" : "", "status" : "400", "userId" : ""}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error(), "token" : "", "status" : "400", "userId" : ""}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.OldPassword)); err != nil {
		response := map[string]string{"message": "Password salah", "token" : "", "status" : "400", "userId" : ""}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	newHashPassword, _ := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	data.NewPassword = string(newHashPassword)

	models.DB.Model(&user).Update("password", newHashPassword)

	response := map[string]any{"data" : user,"message": "Reset password berhasil", "status" : "200"}
	helper.ResponseJSON(w, http.StatusOK, response)


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
