package common

import (
	"os"
	"testing"
)

// Testing environment variables.
const (
	// EnvPrefix is the prefix used of all Indent variables.
	EnvPrefix = "INDENT_"

	// EnvTestPrefix is the prefix used for env variables in testing.
	EnvTestPrefix = "TEST_"
)

// Names for environment variables used as cmd/indent-api configuration.
const (
	// EnvIndentSpace is the name of the space affected.
	EnvIndentSpace = "INDENT_SPACE"

	// EnvIndentDatabase is the Indent database name.
	EnvIndentDatabase = "INDENT_DATABASE"

	// EnvAuditDSN is used as the destination for Indent system audit events.
	EnvAuditDSN = "AUDIT_DSN"

	// EnvTelemetryAuditDSN is used as the destination for Indent telemetry audit events.
	EnvTelemetryAuditDSN = "TELEMETRY_AUDIT_DSN"

	// EnvSentryEnvironment is the environment that's included in Sentry exceptions.
	EnvSentryEnvironment = "SENTRY_ENVIRONMENT"

	// EnvTemporalFrontendEndpoint is the host and port used to connect to the Temporal frontend.
	EnvTemporalFrontendEndpoint = "TEMPORAL_FRONTEND_ENDPOINT"

	// EnvTemporalNamespace is the domain used to run workers.
	EnvTemporalNamespace = "TEMPORAL_NAMESPACE"

	// EnvProviderServiceAccountPath is the path to provider service account credentials.
	EnvProviderServiceAccountPath = "PROVIDER_SERVICE_ACCOUNT_PATH"
	// EnvGCSProject is the name of the provider Google Cloud Storage project.
	EnvGCSProject = "GCS_PROJECT"

	EnvTelemetryEndpoint = "TELEMETRY_ENDPOINT"

	// EnvNotificationsEmail is the email address associated with Indent Notifications.
	EnvNotificationsEmail = "INDENT_NOTIFICATIONS_EMAIL"

	// EnvSendgridKey is the Sendgrid API Key used to send Indent Notification emails.
	EnvSendgridKey = "SENDGRID_KEY"

	// EnvSendGridTemplatesPath is the path of the file containing SendGrid template to Protobuf message pairings.
	EnvSendGridTemplatesPath = "SENDGRID_TEMPLATE_PATH"

	// EnvVerboseLog specifies that a verbose log level be used.
	EnvVerboseLog = EnvPrefix + "VERBOSE"

	StorageCredentialsSA     = "/opt/storage-credentials/creds.json"      // #nosec G101 // loaded at runtime
	StorageSyncCredentialsSA = "/opt/storage-sync-credentials/creds.json" // #nosec G101 // loaded at runtime
)

// GetTestEnv returns the value of the test equivalent of the environment variable. Skips the test if unset or empty.
func GetTestEnv(t *testing.T, env string) string {
	t.Helper()
	testEnv := EnvTestPrefix + env
	val := os.Getenv(testEnv)
	if val == "" {
		t.Skipf("set %s to run tests", testEnv)
	}
	return val
}
