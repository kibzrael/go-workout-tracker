package workouttracker

import (
	"fmt"
	"net/http"
)

func Root(){
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Hello World")
	})

	fmt.Println("App listening on http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil{
		panic(err)
	}
}