package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const multipleClosest = -1

func main() {
	start := time.Now()
	defer func() {
		log.Printf("took:  %v", time.Since(start))
	}()

	input := getInput()
	a := newAreaHelper(input)

	input = applyOffset(a.xMin, a.yMin, input)

	a.markAllWithingZone(input, 10000)

	zoneSize := 0

	for _, ySlice := range a.area {
		for _, count := range ySlice {
			if count == 1 {
				zoneSize++
			}
		}
	}

	log.Printf("answer: %v", zoneSize)
}

func getInput() [][]int {
	var inputs [][]int
	file, _ := os.Open("input.txt")
	input := bufio.NewReader(file)

	for {
		l, _, err := input.ReadLine()
		if err != nil {
			break
		}
		coords := strings.Split(string(l), ", ")
		y, _ := strconv.Atoi(coords[0])
		x, _ := strconv.Atoi(coords[1])
		inputs = append(inputs, []int{x, y})
	}

	return inputs
}

func getMinMax(coords [][]int, index int) (int, int) {
	min := -1
	max := -1

	for _, j := range coords {
		if j[index] < min || min == -1 {
			min = j[index]
		} else if j[index] > max || max == -1 {
			max = j[index]
		}
	}

	return min, max
}

func applyOffset(xOff, yOff int, input [][]int) [][]int {
	var newInput [][]int
	for i := 0; i < len(input); i++ {
		newInput = append(newInput, []int{
			input[i][0] - xOff,
			input[i][1] - yOff,
		})
	}

	return newInput
}

type areaHelper struct {
	xMin int
	xMax int
	yMin int
	yMax int
	area [][]int
}

func newAreaHelper(input [][]int) *areaHelper {
	var areas [][]int
	x1, x2 := getMinMax(input, 0)
	y1, y2 := getMinMax(input, 1)

	for i := x1; i <= x2; i++ {
		areaY := make([]int, y2-y1+1)
		areas = append(areas, areaY)
	}

	return &areaHelper{
		x1,
		x2,
		y1,
		y2,
		areas,
	}
}

func (a *areaHelper) markAllWithingZone(coords [][]int, maxDist int) {
	for x, ySlice := range a.area {
		for y := range ySlice {
			totalDist := 0
			allCoords := true
			for _, c := range coords {
				totalDist = totalDist + getDistance(x, c[0]) + getDistance(y, c[1])
				if totalDist >= maxDist {
					allCoords = false
					break
				}
			}
			if allCoords {
				ySlice[y] = 1
			}
		}
	}
}

func getDistance(from, to int) int {
	if from > to {
		return from - to
	}
	return to - from
}

func isIn(i int, s []int) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}
	return false
}
