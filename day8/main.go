package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	lines := read()
	total := 0
	for _, line := range lines {
		if len(lines) < 20 {
			fmt.Printf("%v\n", line)
		}
		for _, digit := range line {
			if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
				total++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", total)
}

// boring input read
func read() [][]string {
	lines := make([][]string, 0)

	reader, closer := selectInput()
	defer closer()
	s := bufio.NewScanner(reader)

	for s.Scan() {
		xs := strings.SplitN(s.Text(), " | ", 2)
		signals := strings.Split(xs[0], " ")
		_ = signals // ignored in part 1
		digits := strings.Split(xs[1], " ")
		lines = append(lines, digits)
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner errors: %v\n", err)
	}

	return lines
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
