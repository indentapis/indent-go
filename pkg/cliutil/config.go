// Package cliutil provides helpers for building Command Lines.
package cliutil

import (
	"os"

	"go.indent.com/indent-go/pkg/common"
)

// NewConfig returns a Config with defaults set.
func NewConfig() *Config {
	return &Config{
		Environment: envProduction,
		Space:       os.Getenv(common.EnvIndentSpace),
	}
}

// Config shared by all commands of the CLI.
type Config struct {
	*Environment

	// Space is the name of the space being used with the CLI.
	Space string

	// Staging specifies that the staging environment should be used.
	Staging bool
}
