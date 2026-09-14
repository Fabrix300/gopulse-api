package main

import (
	"log"
	"net/http"

	"github.com/Fabrix300/gopulse-api/internal/application/monitor"
	"github.com/Fabrix300/gopulse-api/internal/infrastructure/persistence/memory"
	httptransport "github.com/Fabrix300/gopulse-api/internal/transport/http"
)

func main() {
	repository := memory.NewMonitorRepository()

	createMonitorService := monitor.NewCreateMonitorService(
		repository,
	)

	handler := httptransport.NewHandler(createMonitorService)

	server := &http.Server{
		Addr:    ":8088",
		Handler: handler,
	}

	log.Println("GoPulse API running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
