package rand

import (
	"testing"
)

func TestBytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		num := Intn(100)
		if len(Bytes(num)) != num {
			t.Errorf("expected %d bytes", num)
		}
	}
}

func TestBytesEmpty(t *testing.T) {
	t.Parallel()
	if len(Bytes(0)) != 0 {
		t.Error("expected 0 bytes")
	}
}

func TestHex(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		num := Intn(100)
		if len(Hex(num)) != num*2 {
			t.Errorf("expected %d characters", num*2)
		}
	}
}

func TestHexEmpty(t *testing.T) {
	t.Parallel()
	if Hex(0) != "" {
		t.Error("expected 0 characters")
	}
}
