package db

import (
	"math/rand"
	"testing"
)

func Test_Add_Read(t *testing.T) {

	var test_shorthand []byte
	var test_fullpath []byte
	prio := -1

	rand.Read(test_shorthand)
	rand.Read(test_fullpath)

	short := string(test_shorthand)
	full := string(test_fullpath)

	e := Entry{
		key:   short,
		value: full,
		prio:  prio,
		extra: "",
	}


}

