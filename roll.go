package vero

import (
	"github.com/pastc/vero/internal"
)

var (
	Maximum  = 15 // The maximum value that can be rolled
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
	} // Colors mapped to values
	BaitMap = map[int]string{
		4:  "Bait",
		11: "Bait",
	} // Bait mapped to values
)

// Roll Generates a random number from 0 to Maximum and gets the color and bait from ColorMap and BaitMap respectively.
func Roll(serverSeed string, clientSeed string, nonce int) (string, float64, error) {
	game := "ROLL"
	seed := internal.GetCombinedSeed(game, serverSeed, clientSeed, nonce, -1)

	rollValue, err := internal.GetRandomInt(Maximum, seed)
	if err != nil {
		return "", 0, err
	}
	rollColor := internal.GetRollColor(int(rollValue), ColorMap, BaitMap)

	return rollColor, rollValue, nil
}
