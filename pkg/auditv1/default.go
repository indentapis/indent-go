package auditv1

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

var (
	// DefaultClient for sending Events.
	DefaultClient *Client
)

// Init configures the DefaultClient for the provided dsn. Errors will be logged..
func Init(dsn string) {
	l := zap.L()
	target, err := TargetFromDSN(dsn)
	if err == nil {
		err = SetupClient(l, target)
	}

	if err != nil {
		if errMsg := fmt.Sprintf("Init for DefaultClient failed: %v", err); l != nil {
			l.Error(errMsg)
		} else {
			fmt.Println(errMsg)
		}
	}
}

// SetupClient configures the DefaultClient for t.
func SetupClient(l *zap.Logger, t *Target) (err error) {
	if DefaultClient == nil {
		if DefaultClient, err = NewClient(l, t); err != nil {
			return
		}
	} else {
		l.Warn("Audit's DefaultClient has already been configured")
	}
	return nil
}

// Write event asynchronously to DefaultClient. Errors are logged including event.
func Write(event *Event) {
	if DefaultClient != nil {
		DefaultClient.Write(event)
	} else {
		fmt.Println("Init must be called before using DefaultClient")
	}
}

// WriteEvents to DefaultClient. Failures contained in err will contain event.
func WriteEvents(ctx context.Context, events []*Event) (err error) {
	if DefaultClient != nil {
		if err := DefaultClient.WriteEvents(ctx, events); err != nil {
			return err
		}
	} else {
		return errors.New("the func Setup must be called before using DefaultClient")
	}
	return nil
}
