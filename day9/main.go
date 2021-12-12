package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func neighbourhood(hMap [][]int, y, x int) []int {
	neighbourhood := make([]int, 0)

	// top
	if y > 0 {
		neighbourhood = append(neighbourhood, hMap[y-1][x])
	}

	// right
	if x < len(hMap[y])-1 {
		neighbourhood = append(neighbourhood, hMap[y][x+1])
	}

	// bottom
	if y < len(hMap)-1 {
		neighbourhood = append(neighbourhood, hMap[y+1][x])
	}

	// left
	if x > 0 {
		neighbourhood = append(neighbourhood, hMap[y][x-1])
	}

	return neighbourhood
}

func lowest(xs []int, x int) bool {
	for _, item := range xs {
		if item <= x {
			return false
		}
	}
	return true
}

func main() {
	hMap := read()
	if len(hMap) < 20 {
		fmt.Printf("%v\n", hMap)
	}

	var part1 int
	for y := 0; y < len(hMap); y++ {
		for x := 0; x < len(hMap[y]); x++ {
			height := hMap[y][x]
			neighbours := neighbourhood(hMap, y, x)
			if lowest(neighbours, height) {
				if len(hMap) < 20 {
					fmt.Printf("(%d, %d) = %d, neighbourhood = %v\n", y, x, height, neighbours)
				}
				riskLevel := 1 + height
				part1 += riskLevel
			}
		}
	}

	fmt.Printf("part 1: %d\n", part1)
}

// boring input read
func read() [][]int {
	lines := make([][]int, 0)

	reader, closer := selectInput()
	defer closer()
	s := bufio.NewScanner(reader)

	i := 0
	for s.Scan() {
		lines = append(lines, nil)
		for _, c := range s.Text() {
			lines[i] = append(lines[i], int(c-'0'))
		}
		i++
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
