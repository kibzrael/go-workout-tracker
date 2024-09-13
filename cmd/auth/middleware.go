package auth

import (
	"context"
	"log"
	"net/http"
)

type userId string

const AuthUserId userId = "auth.userId"

func JWTAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		// TODO: Handle JWT Authentication
		log.Println("Token user id", 2)
		ctx := context.WithValue(req.Context(), AuthUserId, 2)
		r := req.WithContext(ctx)
		next.ServeHTTP(res, r)
	})
}