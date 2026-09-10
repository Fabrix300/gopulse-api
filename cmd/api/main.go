package main

import (
	"log"
	"net/http"

	httptransport "github.com/Fabrix300/gopulse-api/internal/transport/http"
)

func main() {
	handler := httptransport.NewHandler()

	server := &http.Server{
		Addr:    ":8088",
		Handler: handler,
	}

	log.Println("GoPulse API running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
