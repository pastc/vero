package vero

import (
	"fmt"
	"github.com/pastc/vero/internal"
	"strconv"
	"testing"
)

func TestPlinko(t *testing.T) {
	tests := []struct {
		serverSeed, clientSeed string
		nonce                  int
		iteration              int
		want                   struct {
			value int
		}
	}{
		{"1c5cff3925c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d", "5b60f37f764fdb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71", 493587, 0, struct {
			value int
		}{-1}},
		{"737bef126f149cdc10b40b2922fcbbc3aae465fa506b171cba91661560a7a1fc", "868af726a9b00af3771c32b25db86ec7a281f721c150ff04a4adf97a059b40c5", 364597, 0, struct {
			value int
		}{1}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%d,%d", tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration), func(t *testing.T) {
			value, err := Plinko(tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration)
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
	f.Add(0, 1, 0, 0)
	f.Fuzz(func(t *testing.T, serverSeedNum int, clientSeedNum int, nonce int, iteration int) {
		serverSeed := internal.Hash512(strconv.Itoa(serverSeedNum))
		clientSeed := internal.Hash512(strconv.Itoa(clientSeedNum))
		_, err := Plinko(serverSeed, clientSeed, nonce, iteration)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
