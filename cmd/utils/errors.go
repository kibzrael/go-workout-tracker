package utils

import (
	"fmt"
	"net/http"
)

func ApiPanic(res *http.ResponseWriter, err *error){
	(*res).WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(*res, "Error Occured:", *err)
}