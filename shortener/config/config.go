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
	portStr := "50051"
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func GetKeyGenServiceUrl() string {
	//return getEnvironmentValue("KEYGEN_SERVICE_URL")
	return "localhost:50052"
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}

func GetSymmetricKey() []byte {
	//return getEnvironmentValue("SYMMETRIC_KEY")
	return []byte("12345678901234567890123456789012")
}
