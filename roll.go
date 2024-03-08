package vero

import (
	"strconv"

	"github.com/pastc/vero/v2/internal"
)

// maximum is the maximum value that can be rolled, counting from 0

// Roll generates a random integer from 0 to maximum and returns the color and bait from colorMap and baitMap
// respectively.
func Roll(serverSeed string, publicSeed string, nonce int, maximum int) (int, error) {
	game := "ROLL"
	seed := internal.GetCombinedSeed(game, publicSeed, strconv.Itoa(nonce))

	hash := internal.Hmac256(serverSeed, seed)

	rollValue, err := internal.GetRandomInt(maximum, hash)
	if err != nil {
		return 0, err
	}

	return int(rollValue), nil
}
