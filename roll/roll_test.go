package roll

import (
	"fmt"
	"strconv"
	"testing"
	"vero"
)

func TestRoll(t *testing.T) {
	Maximum = 15
	ColorMap = map[int]string{
		0:  "Green",
		1:  "Red",
		2:  "Red",
		3:  "Red",
		4:  "Red",
		5:  "Red",
		6:  "Red",
		7:  "Red",
		8:  "Black",
		9:  "Black",
		10: "Black",
		11: "Black",
		12: "Black",
		13: "Black",
		14: "Black",
	}
	BaitMap = map[int]string{
		4:  "Bait",
		11: "Bait",
	}

	tests := []struct {
		serverSeed, clientSeed string
		nonce                  int
		want                   struct {
			color string
			value int
		}
	}{
		{"1c5cff3922c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d", "1c064b20e2ed52a5c4db0361a2523e8901db2342f95bd0dd1d9a68a46b8cc483", 5345510, struct {
			color string
			value int
		}{"Red", 7}},
		{"5b60f37f764fdb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71", "2bbff841c1af824d8d6ae228f8204ae08453202883c2d8ccad947581ce1414a2", 5327980, struct {
			color string
			value int
		}{"Red", 6}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%d", tt.serverSeed, tt.clientSeed, tt.nonce), func(t *testing.T) {
			color, value, err := Roll(tt.serverSeed, tt.clientSeed, tt.nonce)
			if err != nil {
				t.Fatalf("got %v", err)
			}
			if color != tt.want.color {
				t.Errorf("got %s, want %s", color, tt.want.color)
			}
			if int(value) != tt.want.value {
				t.Errorf("got %d, want %d", int(value), tt.want.value)
			}
		})
	}
}

func FuzzRoll(f *testing.F) {
	f.Add(0, 1, 0)
	f.Fuzz(func(t *testing.T, serverSeedNum int, clientSeedNum int, nonce int) {
		_, _, err := Roll(vero.Hash(strconv.Itoa(serverSeedNum)), vero.Hash(strconv.Itoa(clientSeedNum)), nonce)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
