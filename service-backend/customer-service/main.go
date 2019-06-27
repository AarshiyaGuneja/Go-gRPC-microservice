package main

import (
	"fleet-backend/customer-service/impl"
	authService "fleet-backend/customer-service/proto"
	"fmt"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	fmt.Println("Starting customer service.")
	service := micro.NewService(
		micro.Name("customer-service"),
	)
	service.Init()

	repo, err := impl.NewCustomerRepository()
	if err != nil {
		log.Fatal(err)
	}
	svc := impl.Service{
		Service: repo,
	}
	handler := impl.Handler{
		Service: svc,
	}

	err = authService.RegisterCustomerServiceHandler(
		service.Server(),
		handler)

	if err != nil {
		panic("Could not register the handler.")
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
