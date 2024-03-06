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
		{"62472aade71d19f24f145306f5755fca07498ce90823b223db734568e4665dedce7fd8d33a6fdcdbd1a5e9a8d2bcfce53ef757048fac6a987d55fc064bdcd0b8", struct {
			value int
		}{205}},
		{"8b13c8014a7704bbccec153354259eba7f8cdfab47caf51e6701e60727f5500f75e9f506fc61c3e6f5063775c17c70b5af476000fadf04ca44399ef465be352a", struct {
			value int
		}{2203}},
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
	f.Add(0)
	f.Fuzz(func(t *testing.T, seed int) {
		_, err := Crash(internal.Hash512(strconv.Itoa(seed)))
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
