package vero

import (
	"fmt"
	"github.com/pastc/vero/internal"
	"strconv"
	"testing"
)

func TestCrash(t *testing.T) {
	HouseEdge = 6.66

	tests := []struct {
		seed string
		want struct {
			value int
		}
	}{
		{"2826d440b0fcad643e3008693c3a93ef81b31675ca00d686e44c40d5e83d7bb6", struct {
			value int
		}{126}},
		{"5b60f37f764fcb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71", struct {
			value int
		}{288}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.seed), func(t *testing.T) {
			value, err := Crash(tt.seed)
			if err != nil {
				t.Fatalf("got %v", err)
			}
			if value != tt.want.value {
				t.Errorf("got %d, want %d", value, tt.want.value)
			}
		})
	}
}

func FuzzCrash(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int) {
		f.Add(0)
		_, err := Crash(internal.Hash(strconv.Itoa(seed)))
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
