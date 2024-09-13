package workouttracker

import (
	"fmt"
	"net/http"
)

func WorkoutDetails(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Workout Details id: %v\n", req.PathValue("id"))
}