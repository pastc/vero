package roll

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	Maximum  = 15
	ColorMap = map[int]string{
		0:  "Green",
		1:  "Red",
		2:  "Red",
		3:  "Red",
		4:  "Red",
		5:  "Red",
		6:  "Red",
		7:  "Red",
		8:  "Black",
		9:  "Black",
		10: "Black",
		11: "Black",
		12: "Black",
		13: "Black",
		14: "Black",
	}
	BaitMap = map[int]string{
		4:  "Bait",
		11: "Bait",
	}
)

func Roll(serverSeed string, publicSeed string, nonce int) (string, float64) {
	game := "ROULETTE"
	seed := getCombinedSeed(game, serverSeed, publicSeed, nonce)

	rollValue := getRandomInt(Maximum, seed)
	rollColor := getRollColor(int(rollValue), ColorMap, BaitMap)

	return rollColor, rollValue
}

func getRandomInt(max int, seed string) float64 {
	// Generate a hmac hash
	hmacHash := hmac.New(sha256.New, []byte(seed)).Sum(nil)
	hash := hex.EncodeToString(hmacHash)

	// Value from hash
	subHash := hash[0:13]
	valueFromHash, err := strconv.ParseInt(subHash, 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Dynamic result for this roll
	e := math.Pow(2, 52)
	result := float64(valueFromHash) / e

	return math.Floor(result * float64(max))
}

func getCombinedSeed(game string, serverSeed string, clientSeed string, nonce int) string {
	seedParameters := []string{serverSeed, clientSeed, strconv.Itoa(nonce)}
	if game != "" {
		seedParameters = append(seedParameters, "")
		copy(seedParameters[1:], seedParameters)
		seedParameters[0] = game
	}

	return strings.Join(seedParameters, "-")
}

func getRollColor(rollValue int, colorMap map[int]string, baitMap map[int]string) string {
	// May the odds be ever in your favor
	color, _ := colorMap[rollValue]
	bait, found := baitMap[rollValue]
	if found {
		return color + "-" + bait
	} else {
		return color
	}
}
