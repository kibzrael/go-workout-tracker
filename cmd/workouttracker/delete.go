package workouttracker

import (
	"fmt"
	"net/http"
)

func DeleteWorkout(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Delete Workout id: %v\n", req.PathValue("id"))
}