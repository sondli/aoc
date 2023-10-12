package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	contentRows := strings.Split(string(content), "\n")

	itemTypes := []rune{}

	for _, contentRow := range contentRows {
		line := string(contentRow)
		splitOn := len(line) / 2
		lastIndex := len(line)
		first := line[0:splitOn]
		second := line[splitOn:lastIndex]

		firstRunes := []rune(first)

		for i := 0; i < len(first); i++ {
			if strings.ContainsRune(second, firstRunes[i]) {
			    itemTypes = append(itemTypes, firstRunes[i])	
				break
			}
		}
	}

    sum := 0
    for i := 0; i < len(itemTypes); i++ {
        sum += getPrio(itemTypes[i])  
    }

	fmt.Printf("Answer one: %d\n", sum)

    badges := []rune{}
    
    for i:= 0; i < len(contentRows) - 2; i += 3 {
        group := contentRows[i:i + 3]
        for _, letter := range group[0] {
           if strings.ContainsRune(group[1], letter) && strings.ContainsRune(group[2], letter) {
               badges = append(badges, letter) 
               break
           }
        }
    }

    sumTwo := 0
    for i := 0; i < len(badges); i++ {
        sumTwo += getPrio(badges[i])  
    }

	fmt.Printf("Answer two: %d\n", sumTwo)
}

func getPrio(char rune) (prio int) {
	if unicode.IsUpper(char) {
		return int(char) - 38
	} else {
		return int(char) - 96
	}
}
