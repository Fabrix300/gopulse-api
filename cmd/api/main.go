package main

import (
	"log"
	"net/http"

	"github.com/Fabrix300/gopulse-api/internal/application/monitor"
	"github.com/Fabrix300/gopulse-api/internal/infrastructure/persistence/memory"
	httptransport "github.com/Fabrix300/gopulse-api/internal/transport/http"
)

var PORT = ":8088"

func main() {

	repository := memory.NewMonitorRepository()

	createMonitorService := monitor.NewCreateMonitorService(
		repository,
	)

	handler := httptransport.NewHandler(createMonitorService)
	router := httptransport.NewRouter(handler)

	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Println("GoPulse API running on " + PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
