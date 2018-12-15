package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

const units = "abcdefghijklmnopqrstuvwxyz"

func main() {
	start := time.Now()
	defer func() {
		log.Printf("took:  %v", time.Since(start))
	}()

	input := reduce(getInput())

	resChan := make(chan int)
	for _, c := range units {
		oc := c

		go func() {
			stripped := stripBytes(input, byte(oc))
			output := reduce(stripped)
			resChan <- len(output)
		}()
	}

	results := waitForResults(resChan)
	shortest := len(input)
	for _, r := range results {
		if r < shortest {
			shortest = r
		}
	}

	log.Printf("answer %v", shortest)
}

func stripBytes(bs []byte, b byte) []byte {
	nbs := []byte{}
	for _, ob := range bs {
		if ob == b || ob+32 == b || ob == b+32 {
			continue
		}
		nbs = append(nbs, ob)
	}

	return nbs
}

func waitForResults(resChan chan int) []int {
	results := []int{}
	for {
		select {
		case r := <-resChan:
			results = append(results, r)
			if len(results) == len(units) {
				return results
			}
		}
	}
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
			i++

			continue
		}
		new = append(new, c)
	}

	if reactions != 0 {
		return reduce(new)
	}
	return new
}
