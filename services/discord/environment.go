package main

import (
	"log"
	"os"
)

var TOKEN, ok = os.LookupEnv("TOKEN")


func validateEnvironment() {
	if !ok {
		log.Fatal("TOKEN env variable not set!")
	}
}
