package crash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math"
	"strconv"
)

func Crash(seed string, houseEdgePercent float64) float64 {
	// Cannot be a sha256 hash since that will be the previous seed
	hmacHash := hmac.New(sha256.New, []byte(seed)).Sum(nil)
	hash := hex.EncodeToString(hmacHash)

	// Use the most significant 52-bit from the hash to calculate the crash point
	h, err := strconv.ParseInt(hash[:52/4], 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	e := math.Pow(2, 52)

	// Cool equation that determines whether you will live a luxurious life or on the streets
	result := (100*e - float64(h)) / (e - float64(h))

	// The house always wins
	// houseEdgePercent of 5 will result in modifier of 0.95 = 5% house edge with the lowest crash point of 100
	houseEdgeModifier := 1 - houseEdgePercent/100
	endResult := math.Max(100, result*houseEdgeModifier)

	return math.Floor(endResult)
}
