package vero

import (
	"github.com/pastc/vero/internal"
	"math"
	"strconv"
)

// Maximum is the maximum value that can be rolled, counting from 0

// ColorMap is colors mapped to values

// BaitMap is baits mapped to values

// Roll generates a random number from 0 to Maximum and returns the color and bait from ColorMap and BaitMap respectively.
func Roll(serverSeed string, publicSeed string, nonce int, maximum int, colorMap map[int]string, baitMap map[int]string) (string, int, error) {
	game := "ROLL"
	seed := internal.GetCombinedSeed(game, publicSeed, strconv.Itoa(nonce))

	hash := internal.Hmac256(serverSeed, seed)

	rollValue, err := internal.GetRandomInt(maximum, hash)
	if err != nil {
		return "", 0, err
	}
	rollColor := internal.GetRollColor(int(rollValue), colorMap, baitMap)

	return rollColor, int(math.Floor(rollValue)), nil
}
