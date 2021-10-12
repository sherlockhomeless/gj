package parser

import (
	"fmt"
	"sync"
	"testing"
)

func Test_parseLine(t *testing.T) {
	line1:= ": 1633975468:0;ls"
	line2:= ": 1633975477:0;cp big_five.md"
	line3:= ": 1633975492:0;cp big_five.md psychology_introduction.md"

	var wg sync.WaitGroup
	wg.Add(3)
	var c = make(chan Entry, 3)

	go parseLine(line1, c, &wg)
	go parseLine(line2, c, &wg)
	go parseLine(line3, c, &wg)

	wg.Wait()
	close(c)

	var results []Entry
	for e := range c{
		results = append(results, e)
		fmt.Println(results)
	}


	if len(results) != 3 {
		t.Fail()
	}
}
