package auth

import (
	"encoding/json"
	"errors"
	"io"
	"kibzrael/workouttracker/cmd/data"
	"kibzrael/workouttracker/cmd/utils"
	"net/http"

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
	
	db := data.DB()
	if result := db.Create(&user); result.Error == nil{
		response := map[string]any{"message": "Registration successfully", "user": user}
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