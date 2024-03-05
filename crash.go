package vero

import (
	"github.com/pastc/vero/internal"
	"math"
	"strconv"
)

var (
	HouseEdge = 6.66 // The house edge, i.e, how much the house gets on average.
)

func Crash(serverSeed string) (int, error) {
	// Cannot be a sha256 hash since that will be the previous seed
	hmac := internal.Hmac(serverSeed, "")

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
	houseEdgeModifier := 1 - HouseEdge/100
	endResult := math.Max(100, result*houseEdgeModifier)

	return int(math.Floor(endResult)), nil
}