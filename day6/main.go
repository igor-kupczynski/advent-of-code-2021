package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	school := read()
	var (
		acc      = make(map[cacheKey]uint64)
		sumPart1 uint64
		sumPart2 uint64
	)
	for _, f := range school {
		sumPart1 += grow(acc, f, 80)
		sumPart2 += grow(acc, f, 256)
	}
	fmt.Printf("Answer part 1: %d\n", sumPart1)
	fmt.Printf("Answer part 2: %d\n", sumPart2)
}

const timerReset = 6
const timerNew = 8

type cacheKey struct {
	timer int
	days  int
}

func grow(acc map[cacheKey]uint64, timer int, days int) (result uint64) {
	if n, ok := acc[cacheKey{timer, days}]; ok {
		return n
	}
	if timer >= days { // no time for any more spawns
		result = 1
	} else if timer > 0 { // yet another day
		result = grow(acc, timer-1, days-1)
	} else { // neat, new fish
		result = grow(acc, timerReset, days-1) + grow(acc, timerNew, days-1)
	}
	acc[cacheKey{timer, days}] = result
	return
}

// boring input read
func read() []int {
	var school []int

	reader, closer := selectInput()
	defer closer()
	s := bufio.NewScanner(reader)

	if s.Scan() {
		xs := strings.Split(s.Text(), ",")
		school = make([]int, len(xs), len(xs))
		for i, x := range xs {
			f, err := strconv.Atoi(x)
			if err != nil {
				log.Fatalf("Can't parse fish[%d] = %s as a number: %v", i, x, err)
			}
			school[i] = f
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner errors: %v\n", err)
	}

	return school
}

func selectInput() (reader io.Reader, closer func()) {
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Can't open %s: %v\n", os.Args[1], err)
		}
		return f, func() {
			_ = f.Close()
		}
	}
	return os.Stdin, func() {
		// do nothing
	}
}
