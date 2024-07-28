package utils

import (
	"log"
	"os"
)

type Env struct {
	RABBIT_MQ_URL        string
	SMTP_SENDER_EMAIL    string
	SMTP_SENDER_PASSWORD string
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

	envVariables["SMTP_SENDER_EMAIL"] = os.Getenv("SMTP_SENDER_EMAIL")
	envVariables["SMTP_SENDER_PASSWORD"] = os.Getenv("SMTP_SENDER_PASSWORD")
	envVariables["RABBIT_MQ_URL"] = os.Getenv("RABBIT_MQ_URL")

	for varName, value := range envVariables {
		if value == "" {
			log.Panicf("%s is not defined", varName)
		}
	}

	Environment = &Env{
		RABBIT_MQ_URL:        envVariables["RABBIT_MQ_URL"],
		SMTP_SENDER_EMAIL:    envVariables["SMTP_SENDER_EMAIL"],
		SMTP_SENDER_PASSWORD: envVariables["SMTP_SENDER_PASSWORD"],
	}
}
