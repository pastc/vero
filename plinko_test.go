package vero

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/pastc/vero/v2/internal"
)

func TestPlinko(t *testing.T) {
	tests := []struct {
		serverSeed, clientSeed string
		nonce                  int
		iteration              int
		rows                   int
		want                   struct {
			value int
		}
	}{
		{"964cd1665174434d3b82b0a7e9dd5b8bbbc58056a4c3d411d89afcdc2141fa81",
			"20244eaefdc61956bde0c835b0698d8bf5eeddaa7a27e9fae53db5da5a3a0967", 493587, 0, 16, struct {
				value int
			}{9},
		},
		{"ad0111b329b54e2947d1ee14c7b40242019bae11114775932b7865c227636a3a",
			"2cf0c2cb4476cdf70ce68a42bc86b17814dafead8a8c30c128d203dc7270816f", 364597, 0, 12, struct {
				value int
			}{2},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%d,%d", tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration), func(t *testing.T) {
			value, err := Plinko(tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration, tt.rows)
			if err != nil {
				t.Fatalf("got %v", err)
			}
			if value != tt.want.value {
				t.Errorf("got %d, want %d", value, tt.want.value)
			}
		})
	}
}

func FuzzPlinko(f *testing.F) {
	f.Add(0, 1, 2, 0, 1)
	f.Fuzz(func(t *testing.T, serverSeedNum int, clientSeedNum int, nonce int, iteration int, rows int) {
		serverSeed := internal.Hash256(strconv.Itoa(serverSeedNum))
		clientSeed := internal.Hash256(strconv.Itoa(clientSeedNum))

		_, err := Plinko(serverSeed, clientSeed, nonce, iteration, rows)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
