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

	for _, claim := range claims {
		if claim.hasOverlap {
			continue
		}
		for _, otherClaim := range claims {
			claim.overlapsWith(otherClaim)
		}
	}

	for _, c := range claims {
		if !c.hasOverlap {
			log.Printf("answer: %v", c.id)
			log.Printf("took:  %v", time.Since(start))
			return
		}
	}
}

type Claim struct {
	id         int
	x          int
	y          int
	x2         int
	y2         int
	hasOverlap bool
}

func (c *Claim) overlapsWith(o *Claim) {
	if c.id == o.id {
		return
	}
	if c.x < o.x && c.x2 > o.x {
		if c.y < o.y && c.y2 > o.y {
			c.hasOverlap = true
			o.hasOverlap = true
		}
	}

	if c.x < o.x2 && c.x2 > o.x {
		if c.y < o.y2 && c.y2 > o.y {
			c.hasOverlap = true
			o.hasOverlap = true
		}
	}
}

func parseInput() ([]*Claim, error) {
	claims := []*Claim{}
	file, _ := os.Open("input.txt")
	input := bufio.NewReader(file)

	for {
		claim := &Claim{}
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
