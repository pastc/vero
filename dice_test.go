package vero

import (
	"fmt"
	"github.com/pastc/vero/internal"
	"strconv"
	"testing"
)

func TestDice(t *testing.T) {
	tests := []struct {
		serverSeed, clientSeed string
		nonce                  int
		iteration              int
		want                   struct {
			value float64
		}
	}{
		{"964cd1665174434d3b82b0a7e9dd5b8bbbc58056a4c3d411d89afcdc2141fa81",
			"20244eaefdc61956bde0c835b0698d8bf5eeddaa7a27e9fae53db5da5a3a0967", 32139, 0, struct {
				value float64
			}{37.50},
		},
		{"ad0111b329b54e2947d1ee14c7b40242019bae11114775932b7865c227636a3a",
			"2cf0c2cb4476cdf70ce68a42bc86b17814dafead8a8c30c128d203dc7270816f", 43289, 0, struct {
				value float64
			}{35.61},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%d,%d", tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration), func(t *testing.T) {
			value, err := Dice(tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration)
			if err != nil {
				t.Fatalf("got %v", err)
			}
			if value != tt.want.value {
				t.Errorf("got %f, want %f", value, tt.want.value)
			}
		})
	}
}

func FuzzDice(f *testing.F) {
	f.Add(0, 1, 0, 0)
	f.Fuzz(func(t *testing.T, serverSeedNum int, clientSeedNum int, nonce int, iteration int) {
		serverSeed := internal.Hash256(strconv.Itoa(serverSeedNum))
		clientSeed := internal.Hash256(strconv.Itoa(clientSeedNum))
		_, err := Dice(serverSeed, clientSeed, nonce, iteration)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
