package auth

import (
	"fmt"
	"net/http"
)

func Logout(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "Logout")
}