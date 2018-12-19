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
	defer func() {
		log.Printf("took:  %v", time.Since(start))
	}()

	var answer []byte
	steps := getInput()
	first := getFirstStep(steps)

	answer = append(answer, first.required)
	answer = append(answer, first.next)

	rest := getNextSteps(first, steps)
	for _, s := range rest {
		answer = append(answer, s.next)
	}

	log.Printf("answer is: %v", string(answer))
}

type step struct {
	required byte
	next     byte
}

type stepSlice []*step

// Len is part of sort.Interface.
func (d stepSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d stepSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d stepSlice) Less(i, j int) bool {
	return d[i].required < d[j].required
}

func getInput() stepSlice {
	var inputs []*step
	file, _ := os.Open("input.txt")
	input := bufio.NewReader(file)

	for {
		l, _, err := input.ReadLine()
		if err != nil {
			break
		}
		inputs = append(inputs, &step{l[5], l[36]})
	}

	return inputs
}

func getFirstStep(steps []*step) *step {
	var stepsWithARequirement []byte
	for _, s := range steps {
		stepsWithARequirement = append(stepsWithARequirement, s.next)
	}

	for _, s := range steps {
		if !isIn(s.required, stepsWithARequirement) {
			return s
		}
	}

	return &step{0, 0}
}

func getNextSteps(s *step, steps stepSlice) stepSlice {
	var nextSteps stepSlice

	log.Println(string(s.next))
	for _, ss := range steps {
		if s.next == ss.required {
			nextSteps = append(nextSteps, ss)
		}
	}

	sort.Sort(nextSteps)

	for _, sss := range nextSteps {
		nextSteps = append(nextSteps, getNextSteps(sss, steps)...)
	}

	return nextSteps
}

func isIn(i byte, s []byte) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}
	return false
}
