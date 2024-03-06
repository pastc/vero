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
		{"62476aade71d19f24f145306f5755fca07498ce90823b223db734568e4665dedce7fd8d33a6fdcdbd1a5e9a8d2bcfce53ef757048fac6a987d55fc064bdcd0b8", "8b13c8014a7704bbccec153354259eba7f8cdfab47caf51e6701e60727f5500f75e9f506fc61c3e6f5063775c17c70b5af476000fadf04ca44399ef465be352a", 32139, 0, struct {
			value int
		}{232}},
		{"5ad4bacaf3ec34a4a0102a402924610ac70705f3d8bbea1051f0f0d57651c6904fb749a75b4d375f7f7554610bef1c0e93357462c12356734d4cbe4902c35e8e", "3f1b14ac3e6a2e00eca700dc8393ede8e12b3197d5d7a658eb677e86d367104e4924584c90643d055cbab76b209060f9d69a1d70e3c11523b40835999296423c", 43289, 0, struct {
			value int
		}{2457}},
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
		serverSeed := internal.Hash512(strconv.Itoa(serverSeedNum))
		clientSeed := internal.Hash512(strconv.Itoa(clientSeedNum))
		_, err := Dice(serverSeed, clientSeed, nonce, iteration)
		if err != nil {
			t.Fatalf("got %v", err)
		}
	})
}
