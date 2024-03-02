package roll

import (
	"testing"
)

func TestRoll(t *testing.T) {
	Maximum = 15

	color, value := Roll("1c064b20e2ed52a5c4db0361a2523e8901db2342f95bd0dd1d9a68a46b8cc483", "1c5cff3922c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d", 5345510)
	if color != "Red-Bait" || value != 4 {
		t.Fatalf("%s, %f", color, value)
	}
	color, value = Roll("2bbff841c1af824d8d6ae228f8204ae08453202883c2d8ccad947581ce1414a2", "5b60f37f764fcb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71", 5327980)
	if color != "Green" || value != 0 {
		t.Fatalf("%s, %f", color, value)
	}
}
