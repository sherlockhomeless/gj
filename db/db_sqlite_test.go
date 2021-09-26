package db

import (
	"fmt"
	"math/rand"
	"testing"
)

func make_printable(rando []byte) string{
	for i := 0; i < len(rando); i++ {
		rando[i] = (rando[i] % 25) + 97 // wrapping all random bytes to printable ascii characters
	}
	return string(rando)
}

func Test_Add_Read(t *testing.T) {

	var test_shorthand = make([]byte, 20)
	var test_fullpath = make([]byte, 20)
	prio := -1

	rand.Read(test_shorthand)
	rand.Read(test_fullpath)

	short := make_printable(test_shorthand)
	full := make_printable(test_fullpath)

	e := Entry{
		key:   short,
		value: full,
		prio:  prio,
		extra: "",
	}

	fmt.Printf("content: %s, %s", short, full)

	if e.prio != -1 {
		t.Fail()
	}

}

