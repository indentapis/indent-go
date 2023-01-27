// Package auditv1 sends events to the Audit API.
package auditv1

import (
	"crypto/x509"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	auditEndpoint    = "write.indentapis.com:443"
	defaultAuditHost = "write.indentapis.com"
)

// NewClientFromDSN returns a client that sends to input by DSN.
func NewClientFromDSN(logger *zap.Logger, dsn string) (*Client, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	tlsCreds := credentials.NewClientTLSFromCert(pool, "")

	target, err := targetFromDSN(dsn)
	if err != nil {
		return nil, err
	}
	endpoint, err := endpointFromDSN(dsn)
	if err != nil {
		return nil, err
	}

	con, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		return nil, err
	}

	return &Client{
		Target: target,
		Audit:  NewAuditAPIClient(con),
		Log:    logger,
	}, nil
}

// NewClient returns a client that sends to input.
func NewClient(logger *zap.Logger, target *Target) (*Client, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
	}
	tlsCreds := credentials.NewClientTLSFromCert(pool, "")

	con, err := grpc.Dial(auditEndpoint, grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		return nil, err
	}

	return &Client{
		Target: target,
		Audit:  NewAuditAPIClient(con),
		Log:    logger,
	}, nil
}

// Client sends events to the Audit API.
type Client struct {
	// Target to send Events to.
	Target *Target

	// Audit communicates with the AuditAPI.
	Audit AuditAPIClient

	// Log prints informational messages.
	Log *zap.Logger

	// DebugEvents are events written with WriteDebug for validation in testing.
	DebugEvents []*Event

	// Debug determines if all Write method calls are marked for debugging. Disables sending to Audit API.
	Debug bool
}
