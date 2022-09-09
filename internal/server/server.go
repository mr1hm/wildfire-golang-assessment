package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mr1hm/wildfire-golang-assessment/internal/config"
)

type ServerError struct {
	Err error
}

func (se *ServerError) Error() string {
	return fmt.Sprintf("[ API ] Error: %v", se.Err)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func Serve() error {
	// Set config variables
	err := config.ReadConfig()
	if err != nil {
		log.Println(err.Error())
		return &ServerError{
			Err: errors.New(fmt.Sprintf("%v", err)),
		}
	}

	// Set router, middleware and handler
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/joke", getMessage).Methods("GET")

	// Fire up server
	log.Printf("Server starting on Port %v...", config.Port)
	err = http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Println(err.Error())
		return &ServerError{
			Err: errors.New(fmt.Sprintf("Error starting server: %v", err)),
		}
	}

	fmt.Println("Server is running!")

	return nil
}
