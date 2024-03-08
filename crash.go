package vero

import (
	"github.com/pastc/vero/internal"
	"math"
	"strconv"
)

// HouseEdge i.e, percentage that the house gets

// Crash generates a random number from 0 to lowest crash point that is calculated with the HouseEdge in mind.
func Crash(serverSeed string, houseEdge float64) (int, error) {
	game := "CRASH"
	seed := internal.GetCombinedSeed(game)

	hmac := internal.Hmac256(serverSeed, seed)

	// Use the most significant 52-bit from the hash to calculate the crash point
	h, err := strconv.ParseInt(hmac[:52/4], 16, 64)
	if err != nil {
		return 0, err
	}
	e := math.Pow(2, 52)

	// Cool equation that determines whether you will live a luxurious life or on the streets
	result := (100*e - float64(h)) / (e - float64(h))

	// The house always wins
	// houseEdgePercent of 5 will result in modifier of 0.95 = 5% house edge with the lowest crash point of 100
	houseEdgeModifier := 1 - houseEdge/100
	endResult := math.Floor(math.Max(100, result*houseEdgeModifier))

	return int(endResult), nil
}
