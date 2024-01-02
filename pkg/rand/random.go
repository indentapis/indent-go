// Package rand provides utilities for generating random values.
package rand

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// bools can be (1) true or (2) false
const countBoolOptions = 2

// Str generates a randomized string of the specified length.
func Str(length int) (str string) {
	chars := "0123456789abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < length; i++ {
		c := string(chars[Intn(len(chars))])
		str += c
	}
	return
}

// VariableStr returns a randomized string with a uniform random length between
// min and max.
func VariableStr(min, max int) string {
	length := min
	if d := max - min; d > 0 {
		length += Intn(d)
	}
	return Str(length)
}

// Intn returns a number in [0, max).
func Intn(max int) int {
	return int(Int64n(int64(max)))
}

// Int64n returns a number in [0, max).
func Int64n(max int64) int64 {
	num, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return num.Int64()
}

// Bool returns a bool
func Bool() bool {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(countBoolOptions)))
	if err != nil {
		panic(err)
	}
	return int(num.Int64()) == 1
}

// Timestamp returns a uniformly random timestamp between (center - variance/2)
// and (center + variance/2).
func Timestamp(center time.Time, variance time.Duration) *timestamp.Timestamp {
	variance64 := int64(variance)
	// Get a random scaledVariance in range [0, variance].
	scaledVariance, err := rand.Int(rand.Reader, big.NewInt(variance64))
	if err != nil {
		panic(err)
	}
	// Transpose scaled variance into range [-variance/2, variance/2].
	d := scaledVariance.Int64() - (variance64 / 2)
	ts, err := ptypes.TimestampProto(center.Add(time.Duration(d)))
	if err != nil {
		panic(fmt.Errorf("failed to convert time to timestamp: %w", err))
	}
	return ts
}
