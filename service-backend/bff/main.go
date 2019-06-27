package main

import (
	"encoding/json"
	"fleet-backend/bff/impl"
	"fleet-backend/common/middleware"
	common "fleet-backend/common/util"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro"
	"github.com/rs/cors"
	"net/http"
)

type genericHandler struct {
	ApplicationConfig common.ApplicationConfig
}

type TestRequest struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type TestResponse struct {
	Name    string
	Email   string
	Message string
	Id      int
}

func (g genericHandler) testFunc(res http.ResponseWriter, req *http.Request) {
	request := &TestRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response := &TestResponse{
		Name:    request.Name,
		Email:   request.Email,
		Message: "User successfully created.",
		Id:      1,
	}

	json.NewEncoder(res).Encode(response)
}

func main() {
	fmt.Println("Starting bff-service.")
	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	applicationConfig := common.GetApplicationConfig()
	config := applicationConfig.BffConfig()
	logger.WithField("service-name", config.ServiceName).Infoln("Starting service...")

	service := micro.NewService(
		micro.Name(config.ServiceName),
	)

	service.Init()

	r := mux.NewRouter()

	handler := genericHandler{
		ApplicationConfig: applicationConfig,
	}

	r.Path("/testFromMain").HandlerFunc(handler.testFunc)
	// generic-api
	impl.RegisterGenericEndpoints(r, service.Client(), applicationConfig)

	go func() {

		port := config.Port
		logger.WithField("Port", port).Infoln("Starting Rest services now.")
		corsHandler := cors.AllowAll().Handler(r)
		handler := handlers.RecoveryHandler()(corsHandler)

		err := http.ListenAndServe(config.Port, handler)
		if err != nil {
			logger.WithField("Error", err.Error()).Panicln("Could not start the server.")
		}

	}()

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
