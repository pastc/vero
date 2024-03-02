package crash

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"vero"
)

func TestCrash(t *testing.T) {
	crash := Crash(vero.Hash(strconv.Itoa(rand.Int())), 6.66)
	fmt.Println(crash / 100)
}
