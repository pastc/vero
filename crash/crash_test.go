package crash

import (
	"fmt"
	"math"
	"testing"
)

func TestCrash(t *testing.T) {
	crash := Crash("2826d440b0fcad643e3008693c3a93ef81b31675ca00d686e44c40d5e83d7bb6", 6.66)
	fmt.Println(toFixed(crash/100, 2))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
