package dice

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
)

func Dice(serverSeed string, clientSeed string, nonce int, iteration int) (float64, error) {
	game := "DICE"
	seed := getCombinedSeed(game, serverSeed, clientSeed, nonce, iteration)

	hmacHash := hmac.New(sha256.New, []byte(serverSeed))
	hmacHash.Write([]byte(seed))
	hash := hex.EncodeToString(hmacHash.Sum(nil))

	index := 0
	lucky, err := getLucky(hash, index)
	if err != nil {
		return 0, err
	}

	for float64(lucky) >= math.Pow(10, 6) {
		index++
		lucky, err = getLucky(hash, index)
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

func getLucky(hash string, index int) (int, error) {
	hashLucky := hash[index*5 : (index*5)+5]
	luckyNumber, err := strconv.ParseInt(hashLucky, 16, 64)
	return int(luckyNumber), err
}

func getCombinedSeed(game string, serverSeed string, clientSeed string, nonce int, iteration int) string {
	seedParameters := []string{serverSeed, clientSeed, strconv.Itoa(nonce), strconv.Itoa(iteration)}
	if game != "" {
		seedParameters = append(seedParameters, "")
		copy(seedParameters[1:], seedParameters)
		seedParameters[0] = game
	}

	return strings.Join(seedParameters, "-")
}
