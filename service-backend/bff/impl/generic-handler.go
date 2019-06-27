package impl

import (
	"encoding/json"
	"fleet-backend/common/middleware"
	"fleet-backend/common/util"
	customerService "fleet-backend/customer-service/proto"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/client"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type genericHandler struct {
	ApplicationConfig util.ApplicationConfig
	CustomerService   customerService.CustomerService
}

func RegisterGenericEndpoints(router *mux.Router, client client.Client, appConfig util.ApplicationConfig) {
	handler := genericHandler{
		ApplicationConfig: appConfig,
		CustomerService:   customerService.NewCustomerService("customer-service", client),
	}

	router.PathPrefix("/api")

	router.Path("/testRequest/user").HandlerFunc(handler.testFunc)

	// SignUp Driver
	router.Path("/signUp").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := middleware.NewContext("main")
		request := new(customerService.SignUpRequest)
		params := r.URL.Query()
		request.Name = params.Get("Name")
		request.Email = params.Get("Email")
		request.Password = params.Get("Password")
		request.FleetCompanyName = params.Get("FleetCompanyName")

		_, _ = handler.CustomerService.SignUp(ctx, request)
	})

	// Create Driver
	router.Path("/createDriver").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := middleware.NewContext("main")
		request := new(customerService.Driver)
		params := r.URL.Query()
		request.Id = params.Get("Id")
		request.Name = params.Get("Name")
		request.Email = params.Get("Email")
		request.Password = params.Get("Password")
		request.FleetCompanyId= params.Get("FleetCompanyId")

		_, _ = handler.CustomerService.CreateDriver(ctx, request)
	})
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

func (g genericHandler) adminOnly(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if token, err := extractToken(req); err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		} else {
			if response, err := g.ValidateToken(token, "ADMIN"); err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			} else {
				if !response.Valid {
					res.WriteHeader(http.StatusUnauthorized)
					return
				} else {
					handlerFunc(res, req)
				}
			}
		}
	}
}

type tokenValidationResponse struct {
	Valid bool
}

func (g genericHandler) ValidateToken(token, tokenType string) (tokenValidationResponse, error) {
	response := tokenValidationResponse{}
	if token == tokenType {
		response.Valid = true
	} else {
		response.Valid = false
	}
	return response, nil
}

func extractToken(req *http.Request) (string, error) {
	authorizationHeader := req.Header.Get("authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			return bearerToken[1], nil
		} else {
			return "", errors.New("bearer token not found")
		}
	} else {
		return "", errors.New("authorization token not found in header")
	}
}
