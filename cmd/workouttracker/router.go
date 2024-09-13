package workouttracker

import (
	"fmt"
	"kibzrael/workouttracker/cmd/auth"
	"net/http"
)

func Router(){
	workoutRouter := http.NewServeMux()
	workoutRouter.HandleFunc("GET /list", ListWorkout)
	workoutRouter.HandleFunc("POST /create", CreateWorkout)
	workoutRouter.HandleFunc("GET /reports", WorkoutReports)
	workoutRouter.HandleFunc("GET /{id}", WorkoutDetails)
	workoutRouter.HandleFunc("PATCH /{id}", UpdateWorkout)
	workoutRouter.HandleFunc("DELETE /{id}", DeleteWorkout)
	workoutRouter.HandleFunc("POST /{id}/schedule", ScheduleWorkout)

	router := http.NewServeMux()
	router.HandleFunc("POST /login", auth.Login)
	router.HandleFunc("POST /register", auth.Register)
	// 
	logoutHandler := http.HandlerFunc(auth.Logout)
	router.Handle("POST /logout", auth.JWTAuth(logoutHandler))

	router.Handle("/", auth.JWTAuth(workoutRouter))

	middlewares := MiddlewareStack(Logger)
	server := http.Server{
		Addr: ":8000",
		Handler: middlewares(router),
	}

	fmt.Println("App listening on http://localhost:8000")
	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}