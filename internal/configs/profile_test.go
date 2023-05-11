package configs_test

// import (
// 	"os"
// 	"testing"

// 	"mgmt/internal/configs"

// 	"github.com/stretchr/testify/assert"
// )

// func TestActiveProfile(t *testing.T) {
// 	type testActiveProfile struct {
// 		testName              string
// 		setEnvKey             string
// 		setEnvValue           string
// 		expectedActiveProfile string
// 	}

// 	testCases := []testActiveProfile{
// 		{
// 			testName:              "set pipeline profile",
// 			setEnvKey:             "GO_ENV",
// 			setEnvValue:           "pipeline",
// 			expectedActiveProfile: "pipeline",
// 		},

// 		{
// 			testName:              "set capitalized profile name",
// 			setEnvKey:             "GO_ENV",
// 			setEnvValue:           "PIPELINE",
// 			expectedActiveProfile: "pipeline",
// 		},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc

// 		t.Run(tc.testName, func(t *testing.T) {
// 			t.Setenv(tc.setEnvKey, tc.setEnvValue)
// 			err := configs.GetActiveProfile()
// 			if err != nil {
// 				t.Fatalf("testActiveProfile failed when calling func GetActiveProfile(), with error: %v", err)
// 			}

// 			actualActiveProfile := os.Getenv("ACTIVE_PROFILE")

// 			assert.Equal(t, tc.expectedActiveProfile, actualActiveProfile)
// 			t.Logf("GO_ENV is set to: %v, expected ACTIVE_PROFILE to be: %v, got ACTIVE_PROFILE: %v", tc.setEnvValue, tc.expectedActiveProfile, actualActiveProfile)
// 		})
// 	}
// }

// func TestErrorHandling(t *testing.T) {
// 	type testActiveProfile struct {
// 		testName    string
// 		setEnvKey   string
// 		setEnvValue string
// 	}

// 	testCases := []testActiveProfile{
// 		{testName: "set non-existent profile", setEnvKey: "GO_ENV", setEnvValue: "fake_profile"},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc

// 		t.Run(tc.testName, func(t *testing.T) {
// 			t.Setenv(tc.setEnvKey, tc.setEnvValue)

// 			t.Logf("GO_ENV is set to: %v, expected to catch an error", tc.setEnvValue)

// 			err := configs.GetActiveProfile()
// 			if err == nil {
// 				t.Errorf("Expected error from function GetActiveProfile() for Env Key: %v with value: %v, Error was not cached", tc.setEnvKey, tc.setEnvValue)
// 			} else {
// 				t.Logf("Error was successfully catched! catched error has message: %v", err)
// 			}
// 		})
// 	}
// }
