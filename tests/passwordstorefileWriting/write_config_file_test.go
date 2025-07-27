package passwordstorefileWriting

import (
	"os"
	"testing"
	"vb-password-store-base/environment"
	"vb-password-store-base/passwordstore/io"
	"vb-password-store-base/passwordstore/passwordstoreConfig"
)

func TestConfigFileWrite(t *testing.T) {
	os.Setenv("VB_PASSWORD_STORE_DIR", "/home/carl-moritz/")
	path, _ := environment.LookUpAndGetEnvValue("VB_PASSWORD_STORE_DIR")
	var owner passwordstoreConfig.Owner = *passwordstoreConfig.NewOwner(path, "carl-moritz")
	io.WriteFileContents(&owner)
}
