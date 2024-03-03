package crash

import (
	"fmt"
	"testing"
)

func TestCrash(t *testing.T) {
	HouseEdge = 6.66

	tests := []struct {
		seed string
		want struct {
			value float64
		}
	}{
		{"2826d440b0fcad643e3008693c3a93ef81b31675ca00d686e44c40d5e83d7bb6", struct {
			value float64
		}{2.11}},
		{"5b60f37f764fcb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71", struct {
			value float64
		}{1.27}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.seed), func(t *testing.T) {
			value, err := Crash(tt.seed)
			if err != nil {
				t.Errorf("got %d", err)
			}
			if value != tt.want.value {
				t.Errorf("got %f, want %f", value, tt.want.value)
			}
		})
	}
}
