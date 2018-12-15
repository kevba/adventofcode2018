package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		log.Printf("took:  %v", time.Since(start))
	}()

	guards := getGuards()
	var mostSleepGuard *Guard

	for _, g := range guards {
		if mostSleepGuard == nil {
			mostSleepGuard = g
		} else if g.TotalSleepMins() > mostSleepGuard.TotalSleepMins() {
			mostSleepGuard = g
		}
	}

	minMost := mostSleepGuard.MostSleptMinute()

	log.Printf("Guard: %v, Minute slept the most: %v", mostSleepGuard.id, minMost)
	log.Printf("awnser: %v", mostSleepGuard.id*minMost)
}

func getInputSlice() []string {
	logs := []string{}
	file, _ := os.Open("input.txt")
	input := bufio.NewReader(file)

	for {
		l, _, err := input.ReadLine()
		if err != nil {
			break
		}
		logs = append(logs, string(l))
	}

	sort.Strings(logs)
	return logs
}

type Guard struct {
	id           int
	sleepMinutes []int
}

func (g *Guard) AddSleep(from, to int) {
	for i := from; i < to; i++ {
		g.sleepMinutes = append(g.sleepMinutes, i)
	}
}

func (g *Guard) TotalSleepMins() int {
	return len(g.sleepMinutes)
}

func (g *Guard) MostSleptMinute() int {
	occurances := map[int]int{}
	for _, m := range g.sleepMinutes[1:] {
		if o, ok := occurances[m]; ok {
			occurances[m] = o + 1
		} else {
			occurances[m] = 1
		}
	}

	var mostOcc int
	var mostMin int

	for m, occ := range occurances {
		if occ > mostOcc {
			mostOcc = occ
			mostMin = m
		}
	}

	return mostMin
}

func getGuards() map[int]*Guard {
	guards := map[int]*Guard{}

	var currentGuard *Guard
	inputStrings := getInputSlice()
	var fellAsleep time.Time

	for _, l := range inputStrings {
		// Contains a new guard taking watch
		if l[25:26] == "#" {
			guardID, _ := strconv.Atoi(strings.Split(l[26:], " ")[0])
			if g, ok := guards[guardID]; ok {
				currentGuard = g
			} else {
				guards[guardID] = &Guard{id: guardID}
				currentGuard = guards[guardID]
			}
		} else {
			t, _ := time.Parse("2006-01-02 15:04", l[1:17])
			if strings.Contains(l, "asleep") {
				fellAsleep = t
			} else {
				currentGuard.AddSleep(fellAsleep.Minute(), t.Minute())
			}
		}
	}

	return guards
}
