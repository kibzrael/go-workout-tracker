package workouttracker

import (
	"fmt"
	"net/http"
)

func UpdateWorkout(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Update Workout id: %v\n", req.PathValue("id"))
}