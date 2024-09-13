package auth

import (
	"kibzrael/workouttracker/cmd/data"
	"kibzrael/workouttracker/cmd/utils"
	"net/http"
)

func Logout(res http.ResponseWriter, req *http.Request){
	userId, _ := req.Context().Value(AuthUserId).(int)

	db := data.DB()
	result := db.Model(&data.User{}).Where("id = ?", userId).Update("logged_at", nil)
	if result.Error == nil{
		response := map[string]any{"message": "Logged out successfully", "status": http.StatusOK}
		utils.JsonResponse(&res, response)
	} else{
		utils.ApiPanic(&res, &result.Error)
	}

}