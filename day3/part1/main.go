package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	claims, err := parseInput()
	if err != nil {
		log.Fatalln(err)
	}
	countedClaims := getClaimCounted(claims)

	totalDupClaimed := 0
	for _, xClaims := range countedClaims {
		for _, claimCount := range xClaims {
			if claimCount > 1 {
				totalDupClaimed++
			}
		}
	}
	log.Printf("answer: %v", totalDupClaimed)
	log.Printf("took:  %v", time.Since(start))
}

type Claim struct {
	id int
	x  int
	y  int
	x2 int
	y2 int
}

func parseInput() ([]Claim, error) {
	claims := []Claim{}
	file, _ := os.Open("input.txt")
	input := bufio.NewReader(file)

	for {
		claim := Claim{}
		l, _, err := input.ReadLine()
		if err != nil {
			if err == io.EOF {
				return claims, nil
			}
			return claims, err
		}

		parts := strings.Split(string(l), "@")

		claim.id, err = strconv.Atoi(parts[0][1 : len(parts[0])-1])
		if err != nil {
			return claims, err
		}

		cParts := strings.Split(parts[1], ":")
		startParts := strings.Split(cParts[0], ",")
		endParts := strings.Split(cParts[1], "x")

		claim.x, err = strconv.Atoi(startParts[0][1:])
		if err != nil {
			return claims, err
		}
		claim.y, err = strconv.Atoi(startParts[1])
		if err != nil {
			return claims, err
		}

		endXCount, err := strconv.Atoi(endParts[0][1:])
		if err != nil {
			return claims, err
		}
		endYCount, err := strconv.Atoi(endParts[1])
		if err != nil {
			return claims, err
		}

		claim.x2 = claim.x + endXCount
		claim.y2 = claim.y + endYCount

		if err != nil {
			break
		}
		claims = append(claims, claim)
	}

	return claims, nil
}

func getMaxX(claims []Claim) int {
	largest := 0
	for _, c := range claims {
		if c.x2 > largest {
			largest = c.x2
		}
	}
	return largest
}

func getMaxY(claims []Claim) int {
	largest := 0
	for _, c := range claims {
		if c.y2 > largest {
			largest = c.y2
		}
	}
	return largest
}

func getClaimCounted(claims []Claim) [][]int {
	maxX := getMaxX(claims)
	maxY := getMaxY(claims)

	// x,y,claimcounter
	claimedCounters := make([][]int, maxX)
	for i := range claimedCounters {
		claimedCounters[i] = make([]int, maxY)
	}

	for _, claim := range claims {
		for i := claim.x; i < claim.x2; i++ {
			for j := claim.y; j < claim.y2; j++ {
				claimedCounters[i][j] = claimedCounters[i][j] + 1
			}

		}
	}
	return claimedCounters
}
