package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func httpHandler(w http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s Inside the handler using the WITH: \n", r.RemoteAddr, r.Method)
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.With(httpHandler).Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		_, err := w.Write([]byte("Hello World."))
		fmt.Println("A response came in")
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
