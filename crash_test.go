package vero

import (
	"strconv"
	"testing"

	"github.com/pastc/vero/internal"
)

func TestCrash(t *testing.T) {
	houseEdge := 6.66

	tests := []struct {
		serverSeed string
		want       struct {
			value int
		}
	}{
		{"964cd1665174434d3b82b0a7e9dd5b8bbbc58056a4c3d411d89afcdc2141fa81", struct {
			value int
		}{230}},
		{"ad0111b329b54e2947d1ee14c7b40242019bae11114775932b7865c227636a3a", struct {
			value int
		}{1034}},
	}

	for _, tt := range tests {
		t.Run(tt.serverSeed, func(t *testing.T) {
			value, err := Crash(tt.serverSeed, houseEdge)
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
	houseEdge := 6.66
	f.Add(0)
	f.Fuzz(func(t *testing.T, seed int) {
		_, err := Crash(internal.Hash256(strconv.Itoa(seed)), houseEdge)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
