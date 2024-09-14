package workouttracker

import (
	"fmt"
	"kibzrael/workouttracker/cmd/auth"
	"net/http"
)

func CreateWorkout(res http.ResponseWriter, req *http.Request){
	userId, _ := req.Context().Value(auth.AuthUserId).(uint)

	fmt.Fprintln(res, "Create Workout userId:", userId)
}