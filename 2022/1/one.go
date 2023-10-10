package main

import (
	"fmt"
    "strconv"
	"log"
	"os"
	"strings"
    "slices"
)

func main() {
	content, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	contentString := string(content)

	items := strings.Split(contentString, "\n")

	calsPerElf := []int{}
    currentCals := 0

	for _, item := range items {
        if item == "" {
            calsPerElf = append(calsPerElf, currentCals)
            currentCals = 0
            continue
        }
        itemInt, err := strconv.Atoi(item)

        if err != nil {
            log.Fatal(err)

        }

        currentCals += itemInt
	}
    slices.Sort(calsPerElf)

    a1 := calsPerElf[len(calsPerElf) - 1]

    fmt.Printf("Answer one: %d\n", a1)

    a2 := calsPerElf[len(calsPerElf) - 1] + calsPerElf[len(calsPerElf) - 2] + calsPerElf[len(calsPerElf) - 3]

    fmt.Printf("Answer two: %d\n", a2)
}


