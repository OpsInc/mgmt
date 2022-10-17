package env_test

import (
	"errors"
	"forms/env"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

type testEnv struct {
	testName                   string
	envKey                     string
	envValue                   string
	defaultValue               string
	expectedValueGetEnv        string
	expectedValueGetEnvDefault string
	expectedErrorGetEnv        error
	expectedErrorGetEnvDefault error
}

var testEnvValues = []testEnv{ //nolint: gochecknoglobals
	{
		testName:                   "1 envValue takes precedence",
		envKey:                     "env1",
		envValue:                   "value1",
		defaultValue:               "notEmptyDefault1",
		expectedValueGetEnv:        "value1",
		expectedValueGetEnvDefault: "value1",
		expectedErrorGetEnv:        nil,
		expectedErrorGetEnvDefault: nil,
	},
	{
		testName:                   "1.1 envKey capitalized (subtest)",
		envKey:                     "ENV1.1",
		envValue:                   "value1.1",
		defaultValue:               "notEmptyDefault1.1",
		expectedValueGetEnv:        "value1.1",
		expectedValueGetEnvDefault: "value1.1",
		expectedErrorGetEnv:        nil,
		expectedErrorGetEnvDefault: nil,
	},
	{
		testName:                   "2 only envValue is defined",
		envKey:                     "env2",
		envValue:                   "value2",
		defaultValue:               "",
		expectedValueGetEnvDefault: "value2",
		expectedValueGetEnv:        "value2",
		expectedErrorGetEnvDefault: nil,
		expectedErrorGetEnv:        nil,
	},
	{
		testName:                   "2.1 envKey capitialized (subtest)",
		envKey:                     "ENV2.1",
		envValue:                   "value2.1",
		defaultValue:               "",
		expectedValueGetEnvDefault: "value2.1",
		expectedValueGetEnv:        "value2.1",
		expectedErrorGetEnvDefault: nil,
		expectedErrorGetEnv:        nil,
	},
	{
		testName:                   "3 only defaultValue is defined",
		envKey:                     "env3",
		envValue:                   "",
		defaultValue:               "notEmptyDefault3",
		expectedValueGetEnvDefault: "notEmptyDefault3",
		expectedValueGetEnv:        "",
		expectedErrorGetEnvDefault: nil,
		expectedErrorGetEnv:        errTest,
	},
	{
		testName:                   "3.1 envKey capitialized (subtest)",
		envKey:                     "ENV3.1",
		envValue:                   "",
		defaultValue:               "notEmptyDefault3.1",
		expectedValueGetEnvDefault: "notEmptyDefault3.1",
		expectedValueGetEnv:        "",
		expectedErrorGetEnvDefault: nil,
		expectedErrorGetEnv:        errTest,
	},
	{
		testName:                   "4 is Error when only envKey defined",
		envKey:                     "env4",
		envValue:                   "",
		defaultValue:               "",
		expectedErrorGetEnvDefault: errTest,
		expectedErrorGetEnv:        errTest,
	},
	{
		testName:                   "4.1 is Error envKey capitalized (subtest)",
		envKey:                     "ENV4.1",
		envValue:                   "",
		defaultValue:               "",
		expectedErrorGetEnvDefault: errTest,
		expectedErrorGetEnv:        errTest,
	},
	{
		testName:                   "5 is Error when all fields are empty",
		envKey:                     "",
		envValue:                   "",
		defaultValue:               "",
		expectedErrorGetEnvDefault: errTest,
		expectedErrorGetEnv:        errTest,
	},
	{
		testName:                   "6 is Error when only envValue is defined",
		envKey:                     "",
		envValue:                   "value6",
		defaultValue:               "",
		expectedErrorGetEnvDefault: errTest,
		expectedErrorGetEnv:        errTest,
	},
	{
		testName:                   "7 is Error when only envValue & defaultValue are defined",
		envKey:                     "",
		envValue:                   "value7",
		defaultValue:               "notEmptyDefault7",
		expectedErrorGetEnvDefault: errTest,
		expectedErrorGetEnv:        errTest,
	},
	{
		testName:                   "8 is Error when only defaultValue is defined",
		envKey:                     "",
		envValue:                   "",
		defaultValue:               "notEmptyDefault8",
		expectedErrorGetEnvDefault: errTest,
		expectedErrorGetEnv:        errTest,
	},
}

var errTest = errors.New("")

func TestGetEnvDefault(t *testing.T) {
	for _, envTest := range testEnvValues {
		t.Run(envTest.testName, func(t *testing.T) {
			// Set the ENV var if envKey is defined
			if envTest.envKey != "" {
				t.Setenv(envTest.envKey, envTest.envValue)
			}

			// Fetch envValue will be returned from the function if defined
			// If not, defaultValue will be defined instead
			got, _ := env.GetDefault(envTest.envKey, envTest.defaultValue)
			want := envTest.expectedValueGetEnvDefault

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}

func TestGetDefaultError(t *testing.T) {
	for _, envTest := range testEnvValues {
		t.Run(envTest.testName, func(t *testing.T) {
			var (
				errIsNil           bool // Default value is false
				expectedErrorIsNil bool // Default value is false
			)

			// Set the ENV var if envKey is defined
			if envTest.envKey != "" {
				t.Setenv(envTest.envKey, envTest.envValue)
			}

			// Fetch envValue will be returned from the function if defined
			// If not, defaultValue will be defined instead
			_, err := env.GetDefault(envTest.envKey, envTest.defaultValue)

			// Define err into "true" value if nil
			if err == nil {
				errIsNil = true
			}

			// Define envTest.expectedErrorGetEnvDefault into "true" value if nil
			if envTest.expectedErrorGetEnvDefault == nil {
				expectedErrorIsNil = true
			}

			// Both bool vars should have the same value
			// If not, it will FAIL
			if !assert.Equal(t, expectedErrorIsNil, errIsNil) {
				t.Errorf("Error not catched, error msg is: %v", err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	for _, envTest := range testEnvValues {
		t.Run(envTest.testName, func(t *testing.T) {
			// Set the ENV var if envKey is defined
			if envTest.envKey != "" {
				t.Setenv(envTest.envKey, envTest.envValue)
			}

			// Fetch envValue will be returned from the function if defined
			got, _ := env.Get(envTest.envKey)
			want := envTest.expectedValueGetEnv

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}

func TestGetError(t *testing.T) {
	for _, envTest := range testEnvValues {
		t.Run(envTest.testName, func(t *testing.T) {
			var (
				errIsNil           bool // Default value is false
				expectedErrorIsNil bool // Default value is false
			)

			// Set the ENV var if envKey is defined
			if envTest.envKey != "" {
				t.Setenv(envTest.envKey, envTest.envValue)
			}

			// Fetch envValue will be returned from the function if defined
			_, err := env.Get(envTest.envKey)

			// Define err into "true" value if nil
			if err == nil {
				errIsNil = true
			}

			// Define envTest.expectedErrorGetEnv into "true" value if nil
			if envTest.expectedErrorGetEnv == nil {
				expectedErrorIsNil = true
			}

			// Both bool vars should have the same value
			// If not, it will FAIL
			if !assert.Equal(t, expectedErrorIsNil, errIsNil) {
				t.Errorf("Error not catched, error msg is: %v", err)
			}
		})
	}
}
