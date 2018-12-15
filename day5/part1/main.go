package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		log.Printf("took:  %v", time.Since(start))
	}()

	input := getInput()

	result := reduce(input)
	log.Printf("answer %v", len(result))
}

func getInput() []byte {
	file, _ := os.Open("input.txt")
	r, _ := ioutil.ReadAll(file)
	return r
}

func reduce(s []byte) []byte {
	new := []byte{}

	reactions := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if i+1 > len(s)-1 {
			new = append(new, c)
			// Stop the loop because there is no next value to compare against
			break
		}
		next := s[i+1]
		if c-32 == next || c+32 == next {
			reactions++
			i = i + 1

			continue
		}
		new = append(new, c)
	}

	if reactions != 0 {
		return reduce(new)
	}
	return new
}
