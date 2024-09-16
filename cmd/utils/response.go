package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeJsonBody(req *http.Request, target any) error {
	body, err := io.ReadAll(req.Body)
	if err != nil{
		return err
	}
	if err := json.Unmarshal(body, target); err != nil{
		return err
	}
	return nil
}

func JsonResponse(res *http.ResponseWriter, response any){
	(*res).Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(response)
	if err != nil{
		ApiPanic(res, &err)
		return
	}
	(*res).Write(data)
}