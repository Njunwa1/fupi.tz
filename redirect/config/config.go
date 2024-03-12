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
	portStr := "50053"
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func GetShortenerURL() string {
	//return getEnvironmentValue("KEYGEN_SERVICE_URL")
	return "localhost:50051"
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return os.Getenv(key)
}

func GetRedisURL() string {
	//return getEnvironmentValue("REDIS_URL")
	return "localhost:6379"
}

func GetRabbitMQURL() string {
	//return getEnvironmentValue("RABBITMQ_URL")
	return "amqp://guest:guest@localhost:5672/"
}
