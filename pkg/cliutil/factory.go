package cliutil

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/labels"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/common"
	"go.indent.com/indent-go/pkg/oauthutil"
	"go.indent.com/indent-go/pkg/petitioncfg"
)

// Ensure Factory is implemented.
var _ Factory = new(factoryImpl)

// Factory provides shared configuration used across the CLI.
type Factory interface {
	// Logger returns a zap.Logger for logging messages in the CLI.
	Logger() *zap.Logger

	// Config returns the global configuration of the CLI.
	Config() *Config

	// Store returns a token Store.
	Store() oauthutil.Store

	// API returns an APIClient configured to use the Platform API.
	API(context.Context) APIClient

	// CurrentUser returns a Resource representing the currently logged-in user.
	CurrentUser(context.Context) *auditv1.Resource

	// IsLoggedIn returns whether credentials and are active or not
	IsLoggedIn(context.Context) bool

	// AppConfigName returns the ConfigName for the space.
	AppConfigName(context.Context) string
}

// New returns a new Factory using defaults to authenticate against the Platform API.
func New() (Factory, *Config) {
	logger, err := zap.NewProduction(zap.IncreaseLevel(zapcore.InfoLevel))
	if err != nil {
		panic(fmt.Sprintf("failed to setup logger: %v", err))
	}

	config := NewConfig()
	return &factoryImpl{
		logger: logger,
		config: config,
	}, config
}

// factoryImpl is the runtime implementation of Factory.
type factoryImpl struct {
	logger *zap.Logger
	config *Config
	store  oauthutil.Store
}

func (f *factoryImpl) Logger() *zap.Logger {
	return f.logger
}

func (f *factoryImpl) Config() *Config {
	f.setupEnv()
	return f.config
}

func (f *factoryImpl) Store() oauthutil.Store {
	f.setupEnv()
	if f.store == nil {
		f.store = oauthutil.NewStore(f.config.Environment.Name)
	}
	return f.store
}

func (f *factoryImpl) API(ctx context.Context) APIClient {
	f.setupEnv()
	logger := f.Logger()
	client, err := NewAPIClient(ctx, f)
	if err != nil {
		logger.Fatal("Failed to setup API client", zap.Error(err))
	}
	return client
}

func (f *factoryImpl) CurrentUser(ctx context.Context) (user *auditv1.Resource) {
	logger := f.Logger()
	claims, err := f.Store().Claims()
	if err != nil {
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

func (f *factoryImpl) setupEnv() {
	if f.config.Staging {
		f.config.Environment = envStaging
	}
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
