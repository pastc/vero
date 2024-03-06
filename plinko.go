package vero

import (
	"github.com/pastc/vero/internal"
	"math"
	"strconv"
)

// Plinko generates a number with a 1/2 chance between -1 and 1 where -1 is left and 1 is right.
func Plinko(serverSeed string, clientSeed string, nonce int, iteration int) (int, error) {
	game := "PLINKO"
	seed := internal.GetCombinedSeed(game, clientSeed, strconv.Itoa(nonce), strconv.Itoa(iteration))

	hash := internal.Hmac512(serverSeed, seed)

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
			return Plinko(serverSeed, clientSeed, nonce, iteration+1)
		}
	}

	luckyNumber := int(math.Floor(math.Mod(float64(lucky), math.Pow(10, 4))))
	if luckyNumber < 5000 {
		return -1, nil
	} else {
		return 1, nil
	}
}
