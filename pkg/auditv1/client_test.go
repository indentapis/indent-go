package auditv1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	testTarget = "4b5be66077a24e57b19eb20012fa1eb2"
	testDSN    = "https://write.indentapis.com/v1/4b5be66077a24e57b19eb20012fa1eb2"
)

func testClient(t *testing.T) *Client {
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
			Event:      randStr(4, 8),
			Timestamp:  randTimestamp(now, 14*24*time.Hour),
			Id:         randStr(8, 22),
			ExternalId: randStr(8, 22),
			SessionId:  randStr(8, 22),
			Actor: &Resource{
				Id:          randStr(4, 8),
				DisplayName: randStr(4, 16),
				Email:       fmt.Sprintf("%s@%s.com", randStr(3, 6), randStr(4, 9)),
			},
		}
	}
	return events
}

func randStr(min, max int) (str string) {
	length := max
	if min != max {
		length = min + rand.Intn(max-min)
	}

	chars := "0123456789abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < length; i++ {
		c := string(chars[rand.Intn(len(chars))])
		str += c
	}
	return
}

func randTimestamp(center time.Time, variance time.Duration) *timestamp.Timestamp {
	varInt := int64(variance)
	diff := time.Duration(rand.Int63n(varInt) - (varInt / 2))

	ts, err := ptypes.TimestampProto(center.Add(diff))
	if err != nil {
		panic(err)
	}
	return ts
}
