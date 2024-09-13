package workouttracker

import (
	"fmt"
	"net/http"
)

func WorkoutReports(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "Workout Reports")
}