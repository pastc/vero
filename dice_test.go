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
			value int
		}
	}{
		{"1c5cff3922c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d", "5b60f37f764fdb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71", 32139, 0, struct {
			value int
		}{97}},
		{"737bef126f149cdc10b40b29c2fcbbc3aae465fa506b171cba91661560a7a1fc", "868af726a9b00af3771c32b25db86ec7a281f721c150ff04a4adf97a059b40c5", 43289, 0, struct {
			value int
		}{3853}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%d,%d", tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration), func(t *testing.T) {
			value, err := Dice(tt.serverSeed, tt.clientSeed, tt.nonce, tt.iteration)
			if err != nil {
				t.Fatalf("got %v", err)
			}
			if value != tt.want.value {
				t.Errorf("got %d, want %d", value, tt.want.value)
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
