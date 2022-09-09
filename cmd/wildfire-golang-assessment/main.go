package main

import (
	"github.com/mr1hm/wildfire-golang-assessment/internal/api"
	"log"

	"github.com/mr1hm/wildfire-golang-assessment/internal/server"
)

func main() {
	// Initialize New `ResponseData`
	api.RespData = api.RespData.NewData()

	// Fire up server
	err := server.Serve()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
