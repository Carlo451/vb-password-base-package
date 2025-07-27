package environment

import (
	"os"
	"testing"
	"vb-password-store-base/environment"
)

func TestSimpleOsEnvironmentVar(t *testing.T) {
	os.Setenv("TESTCASE", "OS")
	val, exists := environment.LookUpAndGetEnvValue("TESTCASE")
	if exists {
		if val != "OS" {
			t.Errorf("Expected OS")
		}
	}
}

func TestOverLoadOfEnvFile(t *testing.T) {
	os.Setenv("TESTCASE", "OS")
	os.Setenv("VB_PASSWORD_STORE_ENVIRONMENT_PATH", "test.env")
	environment.Init()
	val, exists := environment.LookUpAndGetEnvValue("TESTCASE")
	if !exists {
		t.Errorf("Did not find env variable TESTCASE")
	}
	if val != "FILE" {
		t.Errorf("Expected FILE from the env file")
	}

}

// tests that if no env FIle exists the normal os Get env loads the vars
func TestOverLoadOfEnvFileNoExtraFileSet(t *testing.T) {
	os.Setenv("TESTCASE", "OS")
	os.Setenv("VB_PASSWORD_STORE_ENVIRONMENT_PATH", "")
	environment.Init()
	val, exists := environment.LookUpAndGetEnvValue("TESTCASE")
	if exists {
		if val != "OS" {
			t.Errorf("Expected OS")
		}
	}
}
