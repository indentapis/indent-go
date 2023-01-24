package cliutil

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
)

// Ensure APIClient is implemented.
var _ APIClient = new(apiClientImpl)

// APIClient returns type specific clients.
type APIClient interface {
	Blocks() indentv1.BlockAPIClient
	Petitions() indentv1.PetitionAPIClient
	Resources() indentv1.ResourceAPIClient
	Webhooks() indentv1.WebhookAPIClient
}

// NewAPIClient creates a client configured with f.
func NewAPIClient(ctx context.Context, f Factory) (APIClient, error) {
	config := f.Config()
	conn, err := grpc.DialContext(ctx,
		config.Target,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: f.Store(),
		}))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Platform API: %w", err)
	}

	return apiClientImpl{
		conn: conn,
	}, nil
}

// apiClientImpl is the runtime implementation of APIClient.
type apiClientImpl struct {
	conn *grpc.ClientConn
}

func (c apiClientImpl) Blocks() indentv1.BlockAPIClient {
	return indentv1.NewBlockAPIClient(c.conn)
}

func (c apiClientImpl) Petitions() indentv1.PetitionAPIClient {
	return indentv1.NewPetitionAPIClient(c.conn)
}

func (c apiClientImpl) Resources() indentv1.ResourceAPIClient {
	return indentv1.NewResourceAPIClient(c.conn)
}

func (c apiClientImpl) Webhooks() indentv1.WebhookAPIClient {
	return indentv1.NewWebhookAPIClient(c.conn)
}