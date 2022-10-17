package configs

import (
	"fmt"
	"forms/env"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

var (
	err              error
	errVoidEnvValue  = fmt.Errorf("function GetActiveProfile() failed with error: environment variable GO_ENV is empty or set to an unrecognized value")
	errRuntimeCaller = fmt.Errorf("function GetActiveProfile() failed while calling function runtime.Caller(), error: failed to get the current execution filename")
	errOsChdir       = fmt.Errorf("function GetActiveProfile() failed while calling function os.Chdir(), with error: %w", err)
	errGetEnv        = fmt.Errorf("function GetActiveProfile() failed while calling function env.Get(), with error: %w", err)
	errGoDotEnv      = fmt.Errorf("function GetActiveProfile() failed while calling function godotenv(), with error: %w", err)
)

func GetActiveProfile() error {
	_, goFileExecuted, _, ok := runtime.Caller(0)
	if !ok {
		return errRuntimeCaller
	}

	projectRootDir := path.Join(path.Dir(goFileExecuted), "..")

	err := os.Chdir(projectRootDir)
	if err != nil {
		return errOsChdir
	}

	activeProfile, err := env.Get("GO_ENV")
	if err != nil {
		return errGetEnv
	}

	activeProfile = strings.ToLower(activeProfile)

	switch activeProfile {
	case "local":
		err := godotenv.Overload(projectRootDir + "/.env-local")
		if err != nil {
			return errGoDotEnv
		}

		return nil

	case "pipeline":
		err := godotenv.Overload(projectRootDir + "/.env-pipeline")
		if err != nil {
			return errGoDotEnv
		}

		return nil

	default:
		return errVoidEnvValue
	}
}
