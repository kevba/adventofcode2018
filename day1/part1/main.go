package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	freq := 0
	f, _ := os.Open("input.txt")
	input := bufio.NewReader(f)

	for {
		l, _, err := input.ReadLine()
		if err != nil {
			break
		}
		adj, _ := strconv.Atoi(string(l))
		freq = freq + adj
	}

	log.Println(freq)
}
