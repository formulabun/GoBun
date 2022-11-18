package env

import (
	"log"
	"os"
)

var TOKEN, ok1 = os.LookupEnv("TOKEN")
var APPICATIONID, ok2 = os.LookupEnv("APPLICATIONID")
var TESTGUILD, ok3 = os.LookupEnv("TESTGUILD")

func ValidateEnvironment() {
	if !ok1 {
		log.Fatal("TOKEN env variable not set!")
	}
  if !ok2 {
    log.Fatal("TOKEN env not set!")
  }
  if !ok3 {
    log.Fatal("TESTGUILD env not set!")
  }
}
