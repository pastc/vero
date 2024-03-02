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
	hash := hex.EncodeToString(hmac.New(sha256.New, []byte(seed)).Sum(nil))

	// Use the most significant 52-bit from the hash to calculate the crash point
	h, err := strconv.ParseInt(hash[0:52/4], 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	e := math.Pow(2, 52)
	result := (100*e - float64(h)) / (e - float64(h))

	// houseEdgePercent of 5 will result in modifier of 0.95 = 5% house edge with a lowest crashpoint of 100
	houseEdgeModifier := 1 - houseEdgePercent/100
	endResult := math.Max(100, result*houseEdgeModifier)

	return math.Floor(endResult)
}
