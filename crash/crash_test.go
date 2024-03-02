package crash

import (
	"fmt"
	"testing"
)

func TestCrash(t *testing.T) {
	crash := Crash("cadaaef371bc977aae209dc9be1a30665550adf89fa40fc17771051914d1f9fc", 6.66)
	fmt.Println(crash / 100)
}
