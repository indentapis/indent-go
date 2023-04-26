package cliutil

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

const (
	// configDirName is the name of the Indent CLI configuration directory.
	configDirName = "access"

	// credentialDirName is name of the directory containing Platform Credentials.
	credentialDirName = "credentials"
)

func userConfigDir(logger *zap.Logger) string {
	dir, err := os.UserConfigDir()
	if err != nil {
		logger.Fatal("Failed to get user config directory", zap.Error(err))
	}
	return filepath.Join(dir, configDirName)
}

func credentialDir(configDir string) string {
	return filepath.Join(configDir, credentialDirName)
}
