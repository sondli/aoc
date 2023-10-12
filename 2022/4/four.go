package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	r, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(r)
	ch := make(chan bool, 1000)
	ch2 := make(chan bool, 1000)

	for s.Scan() {
		wg.Add(1)
		go processLine(s.Text(), ch, ch2)
	}

	wg.Wait()
	close(ch)
	close(ch2)

	answerOne := 0
	answerTwo := 0

	for isCompleteOverlap := range ch {
		if isCompleteOverlap {
			answerOne++
		}
	}

	for isAnyOverlap := range ch2 {
		if isAnyOverlap {
			answerTwo++
		}
	}

	fmt.Printf("Answer one: %d\n", answerOne)
	fmt.Printf("Answer two: %d\n", answerTwo)

}

func processLine(line string, ch chan bool, ch2 chan bool) {
	defer wg.Done()
	e1, e2, _ := strings.Cut(line, ",")

	e1Start, e1End, _ := strings.Cut(e1, "-")
	e2Start, e2End, _ := strings.Cut(e2, "-")

	e1StartInt, _ := strconv.Atoi(e1Start)
	e1EndInt, _ := strconv.Atoi(e1End)
	e2StartInt, _ := strconv.Atoi(e2Start)
	e2EndInt, _ := strconv.Atoi(e2End)

	e1Range := makeRange(e1StartInt, e1EndInt)
	e2Range := makeRange(e2StartInt, e2EndInt)

	ch <- slices.Contains(e1Range, e2StartInt) && slices.Contains(e1Range, e2EndInt) ||
		slices.Contains(e2Range, e1StartInt) && slices.Contains(e2Range, e1EndInt)

	ch2 <- slices.Contains(e1Range, e2StartInt) || slices.Contains(e1Range, e2EndInt) ||
		slices.Contains(e2Range, e1StartInt) || slices.Contains(e2Range, e1EndInt)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
