package vero

import (
	"fmt"
	"github.com/pastc/vero/internal"
	"strconv"
	"testing"
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
		serverSeed, publicSeed string
		nonce                  int
		want                   struct {
			color string
			value int
		}
	}{
		{"62476aade71d19f24f145306f5755fca07498ce90823b223db734568e4665dedce7fd8d33a6fdcdbd1a5e9a8d2bcfce53ef757048fac6a987d55fc064bdcd0b8", "8b13c8014a7704bbccec153354259eba7f8cdfab47caf51e6701e60727f5500f75e9f506fc61c3e6f5063775c17c70b5af476000fadf04ca44399ef465be352a", 5345510, struct {
			color string
			value int
		}{"Red", 2}},
		{"5ad4bacaf3ec34a4a0102a402924610ac70705f3d8bbea1051f0f0d57651c6904fb749a75b4d375f7f7554610bef1c0e93357462c12356734d4cbe4902c35e8e", "3f1b14ac3e6a2e00eca700dc8393ede8e12b3197d5d7a658eb677e86d367104e4924584c90643d055cbab76b209060f9d69a1d70e3c11523b40835999296423c", 5327980, struct {
			color string
			value int
		}{"Black", 9}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%d", tt.serverSeed, tt.publicSeed, tt.nonce), func(t *testing.T) {
			color, value, err := Roll(tt.serverSeed, tt.publicSeed, tt.nonce)
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
	f.Add(0, 1, 0)
	f.Fuzz(func(t *testing.T, serverSeedNum int, clientSeedNum int, nonce int) {
		_, _, err := Roll(internal.Hash512(strconv.Itoa(serverSeedNum)), internal.Hash512(strconv.Itoa(clientSeedNum)), nonce)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
