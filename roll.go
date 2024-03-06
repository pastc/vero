package vero

import (
	"github.com/pastc/vero/internal"
	"math"
	"strconv"
)

// Maximum is the maximum value that can be rolled
var Maximum = 15

// ColorMap is colors mapped to values
var ColorMap = map[int]string{
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

// BaitMap is baits mapped to values
var BaitMap = map[int]string{
	4:  "Bait",
	11: "Bait",
}

// Roll generates a random number from 0 to Maximum and returns the color and bait from ColorMap and BaitMap respectively.
func Roll(serverSeed string, publicSeed string, nonce int) (string, int, error) {
	game := "ROLL"
	seed := internal.GetCombinedSeed(game, publicSeed, strconv.Itoa(nonce))

	hash := internal.Hmac512(serverSeed, seed)

	rollValue, err := internal.GetRandomInt(Maximum, hash)
	if err != nil {
		return "", 0, err
	}
	rollColor := internal.GetRollColor(int(rollValue), ColorMap, BaitMap)

	return rollColor, int(math.Floor(rollValue)), nil
}
