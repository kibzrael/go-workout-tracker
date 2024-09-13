package auth

import (
	"fmt"
	"net/http"
)

func Register(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "Register")
}