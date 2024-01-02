package rand

import (
	"math"
)

// IDNum generates a non-zero random numeric ID.
func IDNum() (id uint64) {
	for id == 0 {
		id = 1 + uint64(Int64n(math.MaxInt64-1))
	}
	return
}
