// Package cliutil provides helpers for building Command Lines.
package cliutil

import (
	"errors"
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	// envPrefix is the prefix used for environment variables.
	envPrefix = "INDENT"

	// configFilename is the name of the configuration file.
	configFilename = configName + "." + configType

	// configName is the name of the configuration file without extension.
	configName = "access"

	// configType is the extension of the configuration file.
	configType = "yaml"
)

// NewConfig configures viper and returns a Config defaulting to the production environment.
func NewConfig(logger *zap.Logger) *Config {
	defaultConfigDir := userConfigDir(logger)

	// setup config file
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(defaultConfigDir)
	viper.AddConfigPath(".")

	// configure environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	if err := viper.BindEnv("SPACE"); err != nil {
		logger.Warn("Failed to bind environment variables", zap.Error(err))
	}

	// read config and setup for writing
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, new(viper.ConfigFileNotFoundError)) {
			logger.Fatal("Failed to read config file", zap.Error(err))
		}
	}

	// if no config file was found, use the default config file
	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		configFile = filepath.Join(defaultConfigDir, configFilename)
		viper.SetConfigFile(configFile)
	}

	// populate config
	config := new(Config)
	config.configDir = filepath.Dir(configFile)
	config.refresh(logger)
	return config
}

// Config shared by all commands of the CLI.
type Config struct {
	// configDir is the directory where the configuration file is stored.
	configDir string

	// Environment is the environment to use with the CLI.
	*Environment `mapstructure:"ENVIRONMENT" yaml:"environment"`

	// Space is the name of the space being used with the CLI.
	Space string `mapstructure:"SPACE" yaml:"space"`

	// Staging specifies that the staging environment should be used.
	Staging bool `mapstructure:"STAGING" yaml:"staging"`

	// JSONKeyFile specifies a path to a Google Developers service account
	// JSON key file. If none, then a standard oauth token is expected.
	JSONKeyFile string `mapstructure:"JSON_KEY_FILE" yaml:"jsonKeyFile"`

	// Verbose specifies that the logger should include debug messages and additional context.
	Verbose bool `mapstructure:"VERBOSE" yaml:"verbose"`

	// Headless specifies that the CLI should not open a browser login window.
	Headless bool `mapstructure:"HEADLESS" yaml:"headless"`
}

func (c *Config) refresh(logger *zap.Logger) {
	if err := viper.Unmarshal(c); err != nil {
		logger.Fatal("Failed to read config file", zap.Error(err))
	} else if c.Staging {
		c.Environment = envStaging
	} else if c.Environment == nil {
		c.Environment = envProduction
	}
}
