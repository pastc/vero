package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
)

func GetCombinedSeed(seedParameters ...string) string {
	return strings.Join(seedParameters, "-")
}

func GetLucky(hash string, index int) (float64, error) {
	hashLucky := hash[index*5 : (index*5)+5]
	luckyNumber, err := strconv.ParseInt(hashLucky, 16, 64)
	return float64(luckyNumber), err
}

func GetRandomInt(max int, hash string) (float64, error) {
	// Value from hash
	subHash := hash[0:13]
	valueFromHash, err := strconv.ParseInt(subHash, 16, 64)
	if err != nil {
		return 0, err
	}

	// Dynamic result for this roll
	e := math.Pow(2, 52)
	result := math.Floor((float64(valueFromHash) / e) * float64(max))

	return result, nil
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

func Hash256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}

func Hmac256(key string, s string) string {
	hmacHash := hmac.New(sha256.New, []byte(key))
	hmacHash.Write([]byte(s))
	return hex.EncodeToString(hmacHash.Sum(nil))
}

func Hmac512(key string, s string) string {
	hmacHash := hmac.New(sha512.New, []byte(key))
	hmacHash.Write([]byte(s))
	return hex.EncodeToString(hmacHash.Sum(nil))
}
