package auth

import (
	"encoding/json"
	"errors"
	"io"
	"kibzrael/workouttracker/cmd/data"
	"kibzrael/workouttracker/cmd/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Login(res http.ResponseWriter, req *http.Request){
	body, err := io.ReadAll(req.Body)
	if err != nil{
		utils.ApiPanic(&res, &err)
		return
	}

	var details data.User
	var user data.User
	json.Unmarshal(body, &details)

	db := data.DB()
	result := db.Where("email = ? AND password = ?", details.Email, details.Password).First(&user)
	if result.Error == nil{
		loggedAt := time.Now()
		db.Model(&data.User{}).Where("id = ?", user.ID).Update("logged_at", loggedAt)
		payload := jwt.MapClaims{"id": user.ID, "email": user.Email, "loggedAt": loggedAt.Unix()}
		token, err := utils.EncodeJWT(payload)
		if err != nil {
			utils.ApiPanic(&res, &err)
			return
		}
		user.LoggedAt = &loggedAt
		response := map[string]any{"message": "Login successfully", "user": user, "token": token}
		res.WriteHeader(http.StatusCreated)
		utils.JsonResponse(&res, response)
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound){
		res.WriteHeader(http.StatusNotFound)
		response := map[string]any{"message": "Invalid email or password", "status": http.StatusNotFound}
		utils.JsonResponse(&res, response)
	} else{
		utils.ApiPanic(&res, &result.Error)
	}
}