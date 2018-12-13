package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	freq := 0
	reachedFreqs := []int{0}

	inputs := getInputSlice()

	for {
		for i := 0; i < len(inputs); i++ {
			freq = freq + inputs[i]
			if isDuplicate(reachedFreqs, freq) {
				log.Println("found duplicate")
				log.Println(freq)
				return
			}
			reachedFreqs = append(reachedFreqs, freq)
		}
		log.Println("resetting loop")
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
