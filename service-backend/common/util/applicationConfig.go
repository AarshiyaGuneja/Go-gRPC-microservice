package util

import (
	"fleet-backend/common/middleware"
	"fmt"
	"os"
)

const (
	bffServiceName = "fleet-backend-bff"
)

func buildApplicationConfig(env string) ApplicationConfig {
	return ApplicationConfig{
		Environment: env,
	}
}

func LocalApplicationConfig() ApplicationConfig {
	return buildApplicationConfig(
		"local")
}

func (a ApplicationConfig) IsLocalEnvironment() bool {
	return a.Environment == "local"
}

func GetApplicationConfig() ApplicationConfig {
	//flag.Parse()
	env := os.Getenv("MICRO_ENV")
	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	logger.WithField("MICRO_ENV", env).Infoln("Loading configs.")
	switch env {
	case "local":
		return LocalApplicationConfig()
	default:
		panic(fmt.Sprintf("Could not source the application config for Environment = %s", env))
	}
}

type ApplicationConfig struct {
	Environment string
}

func (a ApplicationConfig) IsLocalEnv() bool {
	return a.Environment == "local"
}

func (a ApplicationConfig) RequiresSetup() bool {
	return a.Environment == "local" || a.Environment == "dev"
}

type ServiceConfig struct {
	ServiceName string
}

type BffConfig struct {
	ServiceName string
	Port        string
}

func (a ApplicationConfig) BffConfig() BffConfig {
	return BffConfig{
		ServiceName: bffServiceName,
		Port:        ":20022",
	}
}

func (a ApplicationConfig) buildServiceConfig(serviceName string) ServiceConfig {
	return ServiceConfig{
		ServiceName: serviceName,
	}
}
