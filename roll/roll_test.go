package roll

import (
	"fmt"
	"testing"
)

func TestRoll(t *testing.T) {
	Maximum = 15
	tests := []struct {
		serverSeed, publicSeed string
		nonce                  int
		wantColor              string
		wantValue              int
	}{
		{"1c5cff3922c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d", "1c064b20e2ed52a5c4db0361a2523e8901db2342f95bd0dd1d9a68a46b8cc483", 5345510, "Red-Bait", 4},
		{"5b60f37f764fcb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71", "2bbff841c1af824d8d6ae228f8204ae08453202883c2d8ccad947581ce1414a2", 5327980, "Green", 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s,%d", tt.serverSeed, tt.publicSeed, tt.nonce)
		t.Run(testName, func(t *testing.T) {
			color, value := Roll(tt.serverSeed, tt.publicSeed, tt.nonce)
			if color != tt.wantColor || int(value) != tt.wantValue {
				t.Errorf("got %s, want %s", color, tt.wantColor)
			}
		})
	}
}
