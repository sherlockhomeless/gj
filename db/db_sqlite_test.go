package db

import (
	"math/rand"
	"testing"
)

func make_printable(rando []byte) string{
	for i := 0; i < len(rando); i++ {
		rando[i] = (rando[i] % 25) + 97 // wrapping all random bytes to printable ascii characters
	}
	return string(rando)
}

// Test_Add_Read adds random entries to a test-db then tries to read those entries
func Test_Add_Read(t *testing.T) {

	var test_shorthand = make([]byte, 20)
	var test_fullpath = make([]byte, 20)
	prio := -1

	rand.Read(test_shorthand)
	rand.Read(test_fullpath)

	short := make_printable(test_shorthand)
	full := make_printable(test_fullpath)

	e := Entry{
		shorthand:   short,
		full_path: full,
		prio:  prio,
		extra: "TESTING",
	}

	AddEntry(&e)

	res := GetShort(short)

	if len(res) < 1{
		t.Fail()
	}

}

