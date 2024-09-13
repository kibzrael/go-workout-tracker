package workouttracker

import (
	"fmt"
	"net/http"
)

func ScheduleWorkout(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Schedule Workout id: %v\n", req.PathValue("id"))
}