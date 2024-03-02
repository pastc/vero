package crash

import (
	"testing"
)

func TestCrash(t *testing.T) {
	hash := "2826d440b0fcad643e3008693c3a93ef81b31675ca00d686e44c40d5e83d7bb6"
	//randomHash := vero.Hash(strconv.Itoa(rand.Int()))
	crash, err := Crash(hash, 6.66)
	if err != nil {
		t.Fatal(err)
	}
	if crash/100 != 2.11 {
		t.Fatalf("%f; want 2.11", crash/100)
	}
}
