package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	inputs := getInputSlice()

	for i, input := range inputs {
		box := []rune(input)
		match, index := compareBox(box, inputs[i+1:])
		if match {
			newID := string(append(box[:index], box[index+1:]...))
			log.Printf("answer: %v", newID)
			log.Printf("took:  %v", time.Since(start))
			return
		}
	}
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

func compareBox(box []rune, boxes []string) (bool, int) {
	for _, b := range boxes {
		diffCounter := 0
		diffIndex := 0
		for i, r := range b {
			if r != box[i] {
				diffCounter++
				diffIndex = i
			}
			if diffCounter > 1 {
				break
			}
		}

		if diffCounter == 1 {
			return true, diffIndex
		}
	}
	return false, 0
}
