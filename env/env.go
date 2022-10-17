package env

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	logPrefix = "\nLOG --->"
	logSufix  = "<---"
)

var (
	errEmptyEnvKey        = fmt.Errorf("\nfunction got envKey: \"\" --> Was expecting a non-empty string")
	errEmptyValuesDefault = fmt.Errorf("\ngot envValue: \"\" and defaultValue: \"\". Was expecting at least envValue or defaultValue to be a valid non-empty string parameter")
	errEmptyValues        = fmt.Errorf("\ngot envValue: \"\". Was expecting at least envValue to have a non-empty string parameter")
)

// GetDefault fetches the env variable envKey with a defaultValue specified
// If envKey is empty: err == errEmptyEnvKey
// If the env variable value returns empty, it will use the defaultValue instead
// If env variable returns empty and defaultValue is empty as well: err == errEmptyValuesDefault
func GetDefault(envKey string, defaultValue string) (string, error) {
	var envValue string

	if envKey == "" {
		return "", errEmptyEnvKey
	}

	upperEnv := os.Getenv(strings.ToUpper(envKey))
	lowerEnv := os.Getenv(strings.ToLower(envKey))

	switch {
	case upperEnv != "":
		envValue = upperEnv
	case lowerEnv != "":
		envValue = lowerEnv
	case defaultValue != "":
		log.Println(logPrefix, envKey, "is not defined. Default value", defaultValue, "has been used instead", logSufix)

		return defaultValue, nil
	default:
		return "", errEmptyValuesDefault
	}

	log.Printf("%v Value for %v has been defined as: %v %v", logPrefix, envKey, envValue, logSufix)

	return envValue, nil
}

// Get fetches the env variable envKey
// If envKey is empty: err == errEmptyEnvKey
// If the env variable value returns empty: err == errEmptyValues
func Get(envKey string) (string, error) {
	var envValue string

	if envKey == "" {
		return "", errEmptyEnvKey
	}

	upperEnv := os.Getenv(strings.ToUpper(envKey))
	lowerEnv := os.Getenv(strings.ToLower(envKey))

	switch {
	case upperEnv != "":
		envValue = upperEnv
	case lowerEnv != "":
		envValue = lowerEnv
	default:
		return "", errEmptyValues
	}

	log.Printf("Value for %v has been defined as: %v\n", envKey, envValue)

	return envValue, nil
}
