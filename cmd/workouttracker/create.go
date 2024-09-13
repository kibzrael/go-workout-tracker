package workouttracker

import (
	"fmt"
	"kibzrael/workouttracker/cmd/auth"
	"log"
	"net/http"
)

func CreateWorkout(res http.ResponseWriter, req *http.Request){
	userId, ok := req.Context().Value(auth.AuthUserId).(string)
	if !ok{
		log.Println("Invalid user token")
		res.WriteHeader(http.StatusBadRequest)
	}

	fmt.Fprintln(res, "Create Workout userId:", userId)
}