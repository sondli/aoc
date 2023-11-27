package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instructuion struct {
	move int
	from int
	to   int
}

func main() {
	c, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(c), "\n")
	instructions := []*Instructuion{}
	for _, row := range rows {
		if len(row) == 0 || string(row[0]) != "m" {
			continue
		}

		inst := strings.Split(row, " ")
		move, _ := strconv.Atoi(inst[1])
		from, _ := strconv.Atoi(inst[3])
		to, _ := strconv.Atoi(inst[5])
		instructions = append(instructions, &Instructuion{
			move: move,
			from: from,
			to:   to,
		})
	}

	stacks := getStartingStacks()

	for _, inst := range instructions {
		for i := 0; i < inst.move; i++ {
			stacks[inst.to] = append(stacks[inst.to], stacks[inst.from][len(stacks[inst.from])-1])
			stacks[inst.from] = stacks[inst.from][:len(stacks[inst.from])-1]
		}
	}

	message := ""
	for i, stack := range stacks {
		if i == 0 || len(stack) == 0 {
			continue
		}

		message += string(stack[len(stack)-1])
	}

	fmt.Println("Answer one: " + message)

	stacksTwo := getStartingStacks()

	for _, inst := range instructions {
		for i := inst.move; i > 0; i-- {
			indexToMove := len(stacksTwo[inst.from]) - i
			stacksTwo[inst.to] = append(stacksTwo[inst.to], stacksTwo[inst.from][indexToMove])
		}
        rangeToRemove := len(stacksTwo) - inst.move
        stacksTwo[inst.from] = stacksTwo[inst.from][0:rangeToRemove -1]
	}

	messageTwo := ""
	for i, stack := range stacksTwo {
		if i == 0 || len(stack) == 0 {
			continue
		}

		messageTwo += string(stack[len(stack)-1])
	}

	fmt.Println("Answer two: " + messageTwo)
	fmt.Println(stacksTwo)
}

func getStartingStacks() [][]string {
	return [][]string{
		{},
		{"G", "F", "V", "H", "P", "S"},
		{"G", "J", "F", "B", "V", "D", "Z", "M"},
		{"G", "M", "L", "J", "N"},
		{"N", "G", "Z", "V", "D", "W", "P"},
		{"V", "R", "C", "B"},
		{"V", "R", "S", "M", "P", "W", "L", "Z"},
		{"T", "H", "P"},
		{"Q", "R", "S", "N", "C", "H", "Z", "V"},
		{"F", "L", "G", "P", "V", "Q", "J"},
	}
}
