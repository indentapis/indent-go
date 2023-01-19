package auditv1

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"

	"go.indent.com/indent-go/pkg/rand"
)

const (
	testTarget = "4b5be66077a24e57b19eb20012fa1eb2"
	testDSN    = "https://write.indentapis.com/v1/4b5be66077a24e57b19eb20012fa1eb2"
)

func testClient(t *testing.T) *Client {
	t.Helper()

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	// setup Input target
	target := &Target{
		Name: testTarget,
	}

	client, err := NewClient(logger, target)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func testClientFromDSN(t *testing.T) *Client {
	t.Helper()

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	client, err := NewClientFromDSN(logger, testDSN)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func testEvents(num int) []*Event {
	events := make([]*Event, num)
	now := time.Now().UTC()
	for i := 0; i < num; i++ {
		events[i] = &Event{
			Event:      rand.VariableStr(4, 8),
			Timestamp:  rand.Timestamp(now, 14*24*time.Hour),
			Id:         rand.VariableStr(8, 22),
			ExternalId: rand.VariableStr(8, 22),
			SessionId:  rand.VariableStr(8, 22),
			Actor: &Resource{
				Id:          rand.VariableStr(4, 8),
				DisplayName: rand.VariableStr(4, 16),
				Email:       fmt.Sprintf("%s@%s.com", rand.VariableStr(3, 6), rand.VariableStr(4, 9)),
			},
		}
	}
	return events
}
