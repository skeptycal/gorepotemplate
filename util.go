package gorepo

import (
	"fmt"
	"log"
	"os"
)

// Optional is an optional argument
type Optional interface{}

// checkVariable - use os package to set the env variable and then get it
// if value == "" ... skip the 'set' step and simply return the value
//
// LookupEnv retrieves the value of the environment variable named by the key. If the variable is present in the environment the value (which may be empty) is returned and the boolean is true. Otherwise the returned value will be empty and the boolean will be false.
//
func checkVariable(key string, valueOptional Optional) string {
	value := ""
	if valueOptional != nil {
		// Set value string
		value = fmt.Sprintf("%s", value)

		// set env variable using os package
		os.Setenv(key, value)
	}
	// return the env variable using os package
	value, ok := os.LookupEnv(key)
	if ok == false {
		return ""
	}
	return value
}

// CheckEnvStrings performs some initial environment checks
func CheckEnvStrings() {
	key := "name"
	value := checkVariable(key, "test_var")

	log.Printf("%s = %s \n", key, value)

	key = "COLUMNS"
	value = checkVariable(key, nil)

	log.Printf("%s = %s \n", key, value)

	key = "SHELL"
	value = checkVariable(key, nil)

	log.Printf("%s = %s \n", key, value)
}
