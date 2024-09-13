package workouttracker

import (
	"log"
	"net/http"
	"time"
)

type Middleware func (http.Handler) http.Handler

func MiddlewareStack(ms ...Middleware) Middleware{
	return func (next http.Handler) http.Handler{
		for i := len(ms) - 1; i >= 0; i--{
			middleware := ms[i]
			next = middleware(next)
		}
		return next
	}
}

type WrappedWriter struct{
	http.ResponseWriter
	statusCode int
}

func (w *WrappedWriter) WriteHeader(statusCode int){
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc( func(res http.ResponseWriter, req *http.Request){
		start := time.Now()

		wrapper := &WrappedWriter{
			ResponseWriter: res,
			statusCode: http.StatusOK,
		}
		next.ServeHTTP(wrapper, req)

		log.Println(wrapper.statusCode, req.Method, req.URL.Path, time.Since(start))
	})
}

func JsonType(next http.Handler) http.Handler{
	return http.HandlerFunc( func(res http.ResponseWriter, req *http.Request){
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}