package vero

import (
	"github.com/pastc/vero/internal"
	"math"
	"strconv"
)

// Plinko generates a list of coordinates that the ball went through
func Plinko(serverSeed string, clientSeed string, nonce int, iteration int, rows int) (int, float64, error) {
	game := "PLINKO"

	var coordinate int

	// repeat it the number of rows (n)
	for i := range rows {
		seed := internal.GetCombinedSeed(game, clientSeed, strconv.Itoa(nonce), strconv.Itoa(iteration), strconv.Itoa(i))

		hash := internal.Hmac256(serverSeed, seed)

		index := 0
		lucky, err := internal.GetLucky(hash, index)
		if err != nil {
			return 0, 0, err
		}

		for float64(lucky) >= math.Pow(10, 6) {
			index++
			lucky, err = internal.GetLucky(hash, index)
			if err != nil {
				return 0, 0, err
			}

			if (index*5)+5 > 128 {
				return Plinko(serverSeed, clientSeed, nonce, iteration+1, rows)
			}
		}

		luckyNumber := int(math.Floor(math.Mod(float64(lucky), math.Pow(10, 4))))
		if luckyNumber < 5000 {
			coordinate -= 1
		} else {
			coordinate += 1
		}
	}

	return (rows + coordinate) / 2, math.Trunc(internal.BinomialDistribution(rows, (rows+coordinate)/2)*math.Pow(10,
		6)) / math.Pow(10, 6), nil
}
