package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	inputs := getInputSlice()
	twos := 0
	threes := 0

	for _, box := range inputs {
		hasTwo, hasThree := findMatch([]rune(box))
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	log.Printf("answer: %v", twos*threes)
	log.Printf("took:  %v", time.Since(start))
}

func getInputSlice() []string {
	boxes := []string{}
	file, _ := os.Open("input.txt")
	input := bufio.NewReader(file)

	for {
		l, _, err := input.ReadLine()
		if err != nil {
			break
		}
		boxes = append(boxes, string(l))
	}

	return boxes
}

func findMatch(box []rune) (bool, bool) {
	two := false
	three := false

	sort.Slice(box, func(a, b int) bool {
		return box[a] < box[b]
	})

	for i := 0; i < len(box); i++ {
		if i+1 > len(box)-1 {
			continue
		}
		if box[i] == box[i+1] {
			if i+2 > len(box)-1 {
				two = true
				i = i + 1
				continue
			}

			if box[i] != box[i+2] {
				two = true
				i = i + 1
				continue
			}

			if box[i] == box[i+2] {
				if i+3 > len(box)-1 {
					three = true
					i = i + 2
					continue
				}
				if box[i] != box[i+3] {
					three = true
					i = i + 2
					continue
				}
			}
		}
	}

	return two, three
}
