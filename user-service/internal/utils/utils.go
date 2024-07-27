package utils

import (
	"log"
	"os"
)

type Env struct {
	RABBIT_MQ_URL string
	REDIS_URL     string
	PORT          string
}

var (
	Environment *Env
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ParseEnvsOrPanic() {
	envVariables := make(map[string]string)

	envVariables["REDIS_URL"] = os.Getenv("REDIS_URL")
	envVariables["RABBIT_MQ_URL"] = os.Getenv("RABBIT_MQ_URL")
	envVariables["PORT"] = os.Getenv("PORT")

	for varName, value := range envVariables {
		if value == "" {
			log.Panicf("%s is not defined", varName)
		}
	}

	Environment = &Env{
		RABBIT_MQ_URL: envVariables["RABBIT_MQ_URL"],
		REDIS_URL:     envVariables["REDIS_URL"],
		PORT:          envVariables["PORT"],
	}
}
