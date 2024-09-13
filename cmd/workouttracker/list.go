package workouttracker

import (
	"fmt"
	"net/http"
)

func ListWorkout(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "List Workout")
}