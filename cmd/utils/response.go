package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(res *http.ResponseWriter, response any){
	(*res).Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(response)
	if err != nil{
		ApiPanic(res, &err)
		return
	}
	(*res).Write(data)
}