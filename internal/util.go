package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
)

func GetCombinedSeed(game string, serverSeed string, clientSeed string, nonce int, iteration int) string {
	seedParameters := []string{serverSeed, clientSeed, strconv.Itoa(nonce)}
	if iteration != -1 {
		seedParameters = append(seedParameters, strconv.Itoa(iteration))
	}
	if game != "" {
		seedParameters = append(seedParameters, "")
		copy(seedParameters[1:], seedParameters)
		seedParameters[0] = game
	}

	return strings.Join(seedParameters, "-")
}

func GetLucky(hash string, index int) (int, error) {
	hashLucky := hash[index*5 : (index*5)+5]
	luckyNumber, err := strconv.ParseInt(hashLucky, 16, 64)
	return int(luckyNumber), err
}

func GetRandomInt(max int, seed string) (float64, error) {
	// Generate a hmac hash
	hash := Hmac(seed, "")

	// Value from hash
	subHash := hash[0:13]
	valueFromHash, err := strconv.ParseInt(subHash, 16, 64)
	if err != nil {
		return 0, err
	}

	// Dynamic result for this roll
	e := math.Pow(2, 52)
	result := float64(valueFromHash) / e

	return math.Floor(result * float64(max)), nil
}

func GetRollColor(rollValue int, colorMap map[int]string, baitMap map[int]string) string {
	// May the odds be ever in your favor
	color := colorMap[rollValue]
	bait, found := baitMap[rollValue]
	if found {
		return color + "-" + bait
	} else {
		return color
	}
}

func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return string(bs)
}

func Hmac(key string, s string) string {
	hmacHash := hmac.New(sha256.New, []byte(key))
	if s != "" {
		hmacHash.Write([]byte(s))
	}
	hash := hex.EncodeToString(hmacHash.Sum(nil))
	return hash
}
