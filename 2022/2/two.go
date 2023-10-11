package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Row struct {
	opponent string
	player   string
}

func main() {
	content, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	contentRows := strings.Split(string(content), "\n")

	s1 := 0
    s2 := 0

	for _, contentRow := range contentRows {
		if len(contentRow) == 0 {
			continue
		}
		opponent := string(contentRow[0])
		player := string(contentRow[2])

        s1 = s1 + calculateRound(opponent, player)
        s2 = s2 + calculateRoundTwo(opponent, player)
	}

    fmt.Printf("Answer one: %d\n", s1)
    fmt.Printf("Answer two: %d\n", s2)
}

func calculateRound(opponent string, player string) (score int) {
	if opponent == "A" {
        if player == "X" {
            score = 1 + 3
        } else if player == "Y" {
            score = 2 + 6
        } else {
            score = 3 
        }
	} else if opponent == "B" {
        if player == "X" {
            score = 1 
        } else if player == "Y" {
            score = 2 + 3
        } else {
            score = 3 + 6 
        }
	} else {
        if player == "X" {
            score = 1 + 6
        } else if player == "Y" {
            score = 2 
        } else {
            score = 3 + 3 
        }
	}

    return score
}
func calculateRoundTwo(opponent string, player string) (score int) {
	if opponent == "A" {
        if player == "X" {
            score = 3 
        } else if player == "Y" {
            score = 1 + 3
        } else {
            score = 2 + 6 
        }
	} else if opponent == "B" {
        if player == "X" {
            score = 1 
        } else if player == "Y" {
            score = 2 + 3
        } else {
            score = 3 + 6 
        }
	} else {
        if player == "X" {
            score = 2 
        } else if player == "Y" {
            score = 3 + 3 
        } else {
            score = 1 + 6 
        }
	}

    return score
}
