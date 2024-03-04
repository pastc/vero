package vero

import (
	"github.com/pastc/vero/internal"
	"math"
)

func Dice(serverSeed string, clientSeed string, nonce int, iteration int) (float64, error) {
	game := "DICE"
	seed := internal.GetCombinedSeed(game, serverSeed, clientSeed, nonce, iteration)

	hash := internal.Hmac(serverSeed, seed)

	index := 0
	lucky, err := internal.GetLucky(hash, index)
	if err != nil {
		return 0, err
	}

	for float64(lucky) >= math.Pow(10, 6) {
		index++
		lucky, err = internal.GetLucky(hash, index)
		if err != nil {
			return 0, err
		}

		if (index*5)+5 > 128 {
			return Dice(serverSeed, clientSeed, nonce, iteration+1)
		}
	}

	luckyNumber := math.Mod(float64(lucky), math.Floor(math.Pow(10, 4)))

	return luckyNumber, nil
}
