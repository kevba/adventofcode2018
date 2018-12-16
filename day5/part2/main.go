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
		cc := c

		go func() {
			stripped := stripBytes(input, byte(cc))
			resChan <- len(reduce(stripped))
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

	return r[:len(r)-1]
}

func reduce(s []byte) []byte {
	new := []byte{}

	for {
		if len(s) == 0 {
			break
		}
		next := s[0]
		s = s[1:]

		if len(new) == 0 {
			new = append(new, next)
			continue
		}
		prev := new[len(new)-1]

		if prev+32 == next || prev-32 == next {
			if len(new) == 1 {
				new = []byte{}
			} else {
				new = new[:len(new)-1]
			}
		} else {
			new = append(new, next)
		}
	}

	return new
}
