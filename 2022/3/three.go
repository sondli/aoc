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

    score := 0

    for _, contentRow := range contentRows {
        line := string(contentRow) 
        splitOn := len(line) / 2 
        lastIndex := len(line)
        first := line[0:splitOn]
        second := line[splitOn:lastIndex]

        runes := []rune(first)
        for i := 0; i < len(first); i++ {
            rune := runes[i]
            toAdd := 0
            if (strings.ContainsRune(second, rune) && toAdd == 0) {
                fmt.Printf(string(runes) + " " + second + "\n")
                fmt.Println(string(rune))
                fmt.Println(getPrio(rune))
                toAdd = getPrio(rune)
            }
            score = score + toAdd
        }

	}

    fmt.Printf("Answer: %d\n", score) 

}

func getPrio(char rune) (prio int) {
    if unicode.IsUpper(char) {
        return int(char) - 38
    } else {
        return int(char) - 96
    }
}
