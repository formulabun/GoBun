package main

import (
	"log"
	"os"
)

var TARGET, ok = os.LookupEnv("TARGET")


func validateEnvironment() {
	if !ok {
		log.Fatal("TARGET env variable not set!")
	}
}
