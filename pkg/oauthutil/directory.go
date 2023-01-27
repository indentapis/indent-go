package oauthutil

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// configDirName is the name of the Indent CLI configuration directory.
	configDirName = "indentcli"

	// credentialDirName is name of the directory containing Platform Credentials.
	credentialDirName = "credentials"
)

func configDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Sprintf("Failed to detect user config dir: %v", err))
	}

	return filepath.Join(dir, configDirName)
}

func credentialDir() string {
	return filepath.Join(configDir(), credentialDirName)
}
