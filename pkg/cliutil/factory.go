package cliutil

import (
	"context"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/labels"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/common"
	"go.indent.com/indent-go/pkg/fileutil"
	"go.indent.com/indent-go/pkg/oauthutil"
	"go.indent.com/indent-go/pkg/petitioncfg"
)

const (
	// OutputJSON encodes the output as JSON.
	OutputJSON = "json"
)

const (
	// configDirPerm is the permission bits for the configuration directory.
	configDirPerm = os.FileMode(0o750)
)

// Ensure Factory is implemented.
var _ Factory = new(factoryImpl)

// Factory provides shared configuration used across the CLI.
type Factory interface {
	// Setup readies the factory for use.
	Setup()

	// Logger returns a zap.Logger for logging messages in the CLI.
	Logger() *zap.Logger

	// Config returns the global configuration of the CLI.
	Config() *Config

	// WriteConfig writes the current configuration to disk.
	WriteConfig()

	// Store returns a token Store.
	Store() oauthutil.Store

	// OutputJSON writes the given proto.Message to stdout as JSON.
	OutputJSON(proto.Message)

	// API returns an APIClient configured to use the Platform API.
	API(context.Context) APIClient

	// Resources returns all the Resources in the space.
	Resources(ctx context.Context, view string) *indentv1.ListResourcesResponse

	// SelectResource uses an interactive prompt to select a resource.
	SelectResource(ctx context.Context, view string) (selected *auditv1.Resource)

	// CurrentUser returns a Resource representing the currently logged-in user.
	CurrentUser(context.Context) *auditv1.Resource

	// IsLoggedIn returns whether credentials and are active or not
	IsLoggedIn(context.Context) bool

	// AppConfigName returns the ConfigName for the space.
	AppConfigName(context.Context) string
}

// New returns a new Factory using defaults to authenticate against the Platform API.
func New(logger *zap.Logger, rootCmd *cobra.Command) (Factory, *Config) {
	config := NewConfig(logger)
	return &factoryImpl{
		logger:  logger,
		rootCmd: rootCmd,
		config:  config,
	}, config
}

// factoryImpl is the runtime implementation of Factory.
type factoryImpl struct {
	logger  *zap.Logger
	rootCmd *cobra.Command
	config  *Config
	store   oauthutil.Store
	m       jsonpb.Marshaler
}

func (f *factoryImpl) Setup() {
	if err := viper.BindPFlags(f.rootCmd.Flags()); err != nil {
		f.logger.Fatal("Failed to bind flags", zap.Error(err))
	}
	f.config.refresh(f.logger)
}

func (f *factoryImpl) Logger() *zap.Logger {
	if !f.config.Verbose {
		return f.logger
	}

	logger := f.logger.WithOptions(zap.IncreaseLevel(zap.DebugLevel))
	return logger.With(
		zap.String("config", viper.ConfigFileUsed()),
		zap.String("space", f.config.Space),
		zap.String("environment", f.config.Environment.Name),
	)
}

func (f *factoryImpl) Config() *Config {
	return f.config
}

func (f *factoryImpl) WriteConfig() {
	logger := f.Logger()

	// resolve creation / permission issues
	if dirErr, parentErr := f.writeStatus(); dirErr != nil {
		if parentErr != nil {
			f.printFixPermissionsAndExit()
		} else if dirErr = os.MkdirAll(f.config.configDir, configDirPerm); dirErr != nil {
			logger.Fatal("Failed to create config directory", zap.Error(dirErr))
		}
	}

	logger.Debug("Writing configuration file")
	if err := viper.WriteConfig(); err != nil {
		logger.Fatal("Failed to write config file", zap.Error(err))
	}
	f.config.refresh(logger)
}

func (f *factoryImpl) Store() oauthutil.Store {
	if f.store == nil {
		f.store = oauthutil.NewStore(f.rootCmd.Context(), f.config.Headless, credentialDir(f.config.configDir), f.config.Environment.Name)
	}
	return f.store
}

func (f *factoryImpl) OutputJSON(msg proto.Message) {
	if err := f.m.Marshal(f.rootCmd.OutOrStdout(), msg); err != nil {
		f.Logger().Fatal("Failed to encode as JSON", zap.Error(err))
	}
}

func (f *factoryImpl) API(ctx context.Context) APIClient {
	logger := f.Logger()
	client, err := NewAPIClient(ctx, f)
	if err != nil {
		logger.Fatal("Failed to setup API client", zap.Error(err))
	}
	return client
}

func (f *factoryImpl) Resources(ctx context.Context, view string) *indentv1.ListResourcesResponse {
	logger := f.Logger()
	client := f.API(ctx).Resources()

	resp, err := client.ListResources(ctx, &indentv1.ListResourcesRequest{
		SpaceName: f.Config().Space,
		View:      view,
	})
	if err != nil {
		logger.Fatal("Failed to list Resources", zap.Error(err))
	}
	return resp
}

// SelectResource provides an interactive prompt for selecting resources.
func (f *factoryImpl) SelectResource(ctx context.Context, view string) (selected *auditv1.Resource) {
	logger := f.Logger()
	resources := f.Resources(ctx, view)
	if s, err := NewSelect(resources.GetResources()); err != nil {
		logger.Fatal("Failed to create select", zap.Error(err))
	} else if selected, err = s.Run(); err != nil {
		logger.Fatal("Failed to run select", zap.Error(err))
	}
	logger.Debug("Selected", zap.Object("resource", selected))
	return selected
}

func (f *factoryImpl) CurrentUser(ctx context.Context) (user *auditv1.Resource) {
	logger := f.Logger()
	store := f.Store()
	if err := store.UpdateUserInfo(); err != nil {
		logger.Fatal("Failed to update user info", zap.Error(err))
	} else if claims, err := store.Claims(); err != nil {
		logger.Fatal("Failed to determine current user", zap.Error(err))
	} else if user, err = claims.Resource(); err != nil {
		logger.Fatal("Failed loading resource from claims", zap.Error(err))
	}
	return lookupUser(ctx, f, user)
}

func (f *factoryImpl) IsLoggedIn(ctx context.Context) bool {
	if claims, err := f.Store().Claims(); err != nil {
		return false
	} else if _, err = claims.Resource(); err != nil {
		return false
	}
	return true
}

// ConfigName returns the ConfigName for the space.
func (f *factoryImpl) AppConfigName(ctx context.Context) string {
	logger := f.Logger()
	cfgLabels := petitioncfg.ConfigLabels()
	delete(cfgLabels, common.LabelAppConfigID)
	cfgSel, err := labels.Set(cfgLabels).AsValidatedSelector()
	if err != nil {
		logger.Fatal("Failed to determine app config selector", zap.Error(err))
	}

	client := f.API(ctx).Blocks()
	resp, err := client.ListBlocks(ctx, &indentv1.ListBlocksRequest{
		SpaceName:     f.Config().Space,
		LabelSelector: cfgSel.String(),
	})
	if err != nil {
		logger.Fatal("Failed to list app config blocks", zap.Error(err))
	}

	cfgBlock := resp.GetBlocks()[0]

	var ok bool
	var name string
	if name, ok = cfgBlock.Labels[common.LabelAppConfigID]; !ok {
		logger.Fatal("Config ID is not set on the block", zap.Error(err), zap.Object("block", cfgBlock))
	}
	return name
}

// writeStatus checks if the config file, and it's parent directory are writable.
func (f *factoryImpl) writeStatus() (dirErr, parentErr error) {
	parent := filepath.Dir(f.config.configDir)
	dirErr = fileutil.IsWritable(f.config.configDir)
	parentErr = fileutil.IsWritable(parent)
	return
}

// printFixPermissionsAndExit prints instructions for setting up the config directory.
func (f *factoryImpl) printFixPermissionsAndExit() {
	logger := f.Logger()
	if err := fileutil.PrintPermissionsFix(f.rootCmd.OutOrStderr(), f.config.configDir); err != nil {
		logger.Fatal("Failed to print config directory setup", zap.Error(err))
	}
	os.Exit(0)
}
