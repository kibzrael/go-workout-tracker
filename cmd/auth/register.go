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

func Register(res http.ResponseWriter, req *http.Request){
	body, err := io.ReadAll(req.Body)
	if err != nil{
		utils.ApiPanic(&res, &err)
		return
	}

	var user data.User
	if err := json.Unmarshal(body, &user); err != nil{
		utils.ApiPanic(&res, &err)
		return
	}
	if user.Email == "" || user.Password == "" {
		err := errors.New("email and password are required")
		utils.ApiPanic(&res, &err)
		return
	}
	user.Password = utils.HashPassword(user.Password)
	
	db := data.DB()
	if result := db.Create(&user); result.Error == nil{
		loggedAt := time.Now()
		db.Model(&data.User{}).Where("id = ?", user.ID).Update("logged_at", loggedAt)
		payload := jwt.MapClaims{"id": user.ID, "email": user.Email, "loggedAt": loggedAt.Unix()}
		token, err := utils.EncodeJWT(payload)
		if err != nil {
			utils.ApiPanic(&res, &err)
			return
		}
		user.LoggedAt = &loggedAt
		response := map[string]any{"message": "Registration successfully", "user": user, "token": token}
		res.WriteHeader(http.StatusCreated)
		utils.JsonResponse(&res, response)
	} else if errors.Is(result.Error, gorm.ErrDuplicatedKey){
		res.WriteHeader(http.StatusConflict)
		response := map[string]any{"message": "Email is already taken", "status": http.StatusConflict}
		utils.JsonResponse(&res, response)
	} else {
		utils.ApiPanic(&res, &result.Error)
	}

}