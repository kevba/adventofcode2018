package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	freq := 0
	reachedFreqs := []int{0}

	start := time.Now()
	inputs := getInputSlice()

	iterations := 0
	for {
		for i := 0; i < len(inputs); i++ {
			freq = freq + inputs[i]
			if isDuplicate(reachedFreqs, freq) {
				log.Printf("found duplicate (%v) after iterations: %v", freq, iterations)
				log.Printf("took:  %v", time.Since(start))
				return
			}
			reachedFreqs = append(reachedFreqs, freq)
		}
		iterations++
	}

}

func getInputSlice() []int {
	freqs := []int{}
	file, _ := os.Open("input.txt")
	input := bufio.NewReader(file)

	for {
		l, _, err := input.ReadLine()
		if err != nil {
			break
		}
		f, _ := strconv.Atoi(string(l))
		freqs = append(freqs, f)
	}

	return freqs
}

func isDuplicate(rfs []int, f int) bool {
	for _, rf := range rfs {
		if rf == f {
			return true
		}
	}
	return false
}
