package configs_test

import (
	"mgmt/configs"
	"testing"
)

type testEnv struct {
	testName      string
	envValue      string
	expectedValue string
}

const (
	envKey       string = "TEMPLATE_DIR"
	defaultValue string = "/var/www/html/"
)

func TestFetchVars(t *testing.T) {
	testEnvValues := []testEnv{
		{testName: "1. Testing if envValue takes precedence over defaultValue", envValue: "value1", expectedValue: "value1"},
		{testName: "2. Testing when defaultValue is defined but not envValue", envValue: "", expectedValue: defaultValue},
	}

	for _, envTest := range testEnvValues {
		t.Run(envTest.testName, func(t *testing.T) {
			// Set the ENV var if envKey is defined
			t.Setenv(envKey, envTest.envValue)

			// FetchVars will return a string
			// envValue if defined or else defaultValue value will be defined
			got := configs.FetchVars().TemplateDir
			want := envTest.expectedValue

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}
