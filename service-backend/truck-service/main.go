package main

import (
	"fleet-backend/truck-service/impl"
	truckService "fleet-backend/truck-service/proto"
	"fmt"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	fmt.Println("Starting truck service.")
	service := micro.NewService(
		micro.Name("truck-service"),
	)
	service.Init()

	repo, err := impl.NewTruckRepository()
	if err != nil {
		log.Fatal(err)
	}
	svc := impl.Service{
		Service: repo,
	}
	handler := impl.Handler{
		Service: svc,
	}

	err = truckService.RegisterTruckServiceHandler(
		service.Server(),
		handler)

	if err != nil {
		panic("Could not register the handler.")
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
