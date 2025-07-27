package environment

import (
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	value, exists := LookUpAndGetEnvValue("VB_PASSWORD_STORE_ENVIRONMENT_PATH")
	if exists {
		godotenv.Overload(value)
	}
}

func LookUpAndGetEnvValue(key string) (string, bool) {
	return os.LookupEnv(key)
}
