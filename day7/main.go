package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabs, min, max := read()
	for _, testCase := range []struct {
		desc string
		fn   func(int, []int) int
	}{
		{"diff", diff},
		{"seqsum", sequenceSum},
	} {
		minVal, minX := findMin(testCase.desc, testCase.fn, crabs, min, max)
		fmt.Printf(":: Min f_%s(%04d) => %d\n", testCase.desc, minX, minVal)
	}
}

func findMin(
	desc string,
	fn func(int, []int) int,
	seq []int,
	lb int,
	ub int,
) (minVal, minX int) {
	minVal = math.MaxInt32
	if len(seq) < 20 {
		fmt.Println()
	}
	for x := lb; x <= ub; x++ {
		y := fn(x, seq)
		if y < minVal {
			minVal = y
			minX = x
		}
		if len(seq) < 20 {
			fmt.Printf("f_%s(%04d) => %d\n", desc, x, y)
		}
	}
	if len(seq) < 20 {
		fmt.Println()
	}
	return
}

func diff(x int, crabs []int) int {
	sum := 0
	for _, crab := range crabs {
		f := x - crab
		if f < 0 {
			sum -= f
		} else {
			sum += f
		}
	}
	return sum
}

func sequenceSum(x int, crabs []int) int {
	sum := 0
	for _, crab := range crabs {
		n := x - crab
		if n < 0 {
			n = -n
		}
		sum += n * (n + 1) / 2
	}
	return sum
}

// boring input read
func read() (crabs []int, min, max int) {
	min = math.MaxInt32
	max = math.MinInt32

	reader, closer := selectInput()
	defer closer()
	s := bufio.NewScanner(reader)

	if s.Scan() {
		xs := strings.Split(s.Text(), ",")
		crabs = make([]int, len(xs), len(xs))
		for i, x := range xs {
			n, err := strconv.Atoi(x)
			if err != nil {
				log.Fatalf("Can't parse crab[%d] = %s as a number: %v", i, x, err)
			}
			crabs[i] = n
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner errors: %v\n", err)
	}

	return
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
