package auth

import (
	"fmt"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "Login")
}