package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for s.Scan() {
		wg.Add(1)
		go processLine(s.Text(), ch)
	}

	wg.Wait()
    close(ch)

    answer := 0

    for isOverlap := range ch {
        if isOverlap {
            answer++
        }
    }

    fmt.Printf("Answer one: %d\n", answer)

}

func processLine(line string, ch chan bool) {
	defer wg.Done()
    e1, e2, _ := strings.Cut(line, ",")

    e1Start, e1End, _ := strings.Cut(e1, "-") 
    e2Start, e2End, _ := strings.Cut(e2, "-") 

    e1StartInt, _ := strconv.Atoi(e1Start)
    e1EndInt, _ := strconv.Atoi(e1End)
    e2StartInt, _ := strconv.Atoi(e2Start)
    e2EndInt, _ := strconv.Atoi(e2End)

    e1Range := e1EndInt - e1StartInt
    e2Range := e2EndInt - e2StartInt

    if e1Range < e2Range {
        ch <- e1StartInt >= e2StartInt && e1EndInt <= e2EndInt
    } else if e1Range > e2Range {
        ch <- e2StartInt >= e1StartInt && e2EndInt <= e1EndInt 
    } else {
        ch <- e1StartInt == e2StartInt && e1EndInt == e2EndInt 
    } 
}
