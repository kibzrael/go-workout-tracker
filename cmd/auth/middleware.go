package auth

import (
	"context"
	"errors"
	"kibzrael/workouttracker/cmd/data"
	"kibzrael/workouttracker/cmd/utils"
	"net/http"
)

type userId string

const AuthUserId userId = "auth.userId"

func JWTAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		claims, err := utils.DecodeJWT(req.Header.Get("Authorization"))
		if err != nil{
			utils.ApiPanic(&res, &err)
			return
		}
		db := data.DB()
		var user data.User
		if err := db.Where("id = ?", claims["id"]).First(&user).Error; err != nil{
			utils.ApiPanic(&res, &err)
			return
		}else if float64(user.LoggedAt.Unix()) != claims["loggedAt"]{
			err := errors.New("invalid token")
			utils.ApiPanic(&res, &err)
			return
		}
		ctx := context.WithValue(req.Context(), AuthUserId, user.ID)
		r := req.WithContext(ctx)
		next.ServeHTTP(res, r)
	})
}