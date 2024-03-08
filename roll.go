package vero

import (
	"github.com/pastc/vero/internal"
	"strconv"
)

// maximum is the maximum value that can be rolled, counting from 0

// colorMap is colors mapped to values

// baitMap is baits mapped to values

// Roll generates a random integer from 0 to maximum and returns the color and bait from colorMap and baitMap
// respectively.
func Roll(serverSeed string, publicSeed string, nonce int, maximum int, colorMap map[int]string, baitMap map[int]string) (string, int, error) {
	game := "ROLL"
	seed := internal.GetCombinedSeed(game, publicSeed, strconv.Itoa(nonce))

	hash := internal.Hmac256(serverSeed, seed)

	rollValue, err := internal.GetRandomInt(maximum, hash)
	if err != nil {
		return "", 0, err
	}
	rollColor := internal.GetRollColor(int(rollValue), colorMap, baitMap)

	return rollColor, int(rollValue), nil
}
