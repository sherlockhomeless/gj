// parser reads from the history file and generates entries
package parser

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Entry struct {
	timestamp int64
	command string
	arguments string
	path string
}

// ReadHistory parses the history file given by path into a slice of Entries
func  ReadHistory(path string) ([]Entry, error) {

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	c := make(chan Entry, 5000)
	var wg sync.WaitGroup
	var results []Entry

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go parseLine(line, c, &wg)

		go func() {
			results = append(results, <-c)
		}()
	}

	wg.Wait()

	return results, nil
}

// parseLine parses the line and parses the corresponding entry struct into c
// format is like this: ": 1633975492:0;cp big_five.md psychology_introduction.md"
func parseLine(line string, c chan Entry, wg *sync.WaitGroup){
	defer wg.Done()
	log.Printf("reading line \"%s\"\n", line)

	timestamp_s := strings.Split(line, ":")[1]
	timestamp_s = strings.Replace(timestamp_s, " ", "",1)
	timestamp, err := strconv.ParseInt(timestamp_s, 10, 64)

	if err != nil{
		log.Fatal("Could not parse line: " + line)
	}

	cmd := strings.Split(line, ";")[1]
	c <- Entry{timestamp: timestamp, command: cmd}
}
