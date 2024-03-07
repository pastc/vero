package vero

import (
	"fmt"
	"github.com/pastc/vero/internal"
	"strconv"
	"testing"
)

func TestRoll(t *testing.T) {
	maximum := 15
	colorMap := map[int]string{
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
	baitMap := map[int]string{
		4:  "Bait",
		11: "Bait",
	}

	tests := []struct {
		serverSeed, publicSeed string
		nonce                  int
		want                   struct {
			color string
			value int
		}
	}{
		{"964cd1665174434d3b82b0a7e9dd5b8bbbc58056a4c3d411d89afcdc2141fa81",
			"20244eaefdc61956bde0c835b0698d8bf5eeddaa7a27e9fae53db5da5a3a0967", 5345510, struct {
				color string
				value int
			}{"Black", 14},
		},
		{"ad0111b329b54e2947d1ee14c7b40242019bae11114775932b7865c227636a3a",
			"2cf0c2cb4476cdf70ce68a42bc86b17814dafead8a8c30c128d203dc7270816f", 5327980, struct {
				color string
				value int
			}{"Black", 10},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%d", tt.serverSeed, tt.publicSeed, tt.nonce), func(t *testing.T) {
			color, value, err := Roll(tt.serverSeed, tt.publicSeed, tt.nonce, maximum, colorMap, baitMap)
			if err != nil {
				t.Fatalf("got %v", err)
			}
			if color != tt.want.color {
				t.Errorf("got %s, want %s", color, tt.want.color)
			}
			if value != tt.want.value {
				t.Errorf("got %d, want %d", value, tt.want.value)
			}
		})
	}
}

func FuzzRoll(f *testing.F) {
	colorMap := map[int]string{
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
	baitMap := map[int]string{
		4:  "Bait",
		11: "Bait",
	}

	f.Add(0, 1, 0, 0)
	f.Fuzz(func(t *testing.T, serverSeedNum int, clientSeedNum int, nonce int, maximum int) {
		_, _, err := Roll(internal.Hash256(strconv.Itoa(serverSeedNum)), internal.Hash256(strconv.Itoa(clientSeedNum)), nonce, maximum, colorMap, baitMap)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
