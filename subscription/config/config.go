package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	//return getEnvironmentValue("ENV")
	return "development"
}

func GetDataSourceURL() string {
	//return getEnvironmentValue("DATA_SOURCE_URL")
	return "mongodb://localhost:27017"
}

func GetApplicationPort() int {
	//portStr := getEnvironmentValue("APPLICATION_PORT")
	portStr := "50057"
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func GetPaymentServiceUrl() string{
	return "http://localhost:50058"
}

func GetPlanServiceUrl() string{
	return "http://localhost:50058"
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}

func GetJaegerURL() string {
	//return getEnvironmentValue("JAEGER_URL")
	return "http://localhost:14278/api/traces"
}
