package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type point struct {
	y, x int
}

type heightMap [][]int

func (m heightMap) LenY() int {
	return len(m)
}

func (m heightMap) LenX() int {
	return len(m[0])
}

func (m heightMap) HeightAt(p point) int {
	return m[p.y][p.x]
}

func (m heightMap) RiskLevelAt(p point) int {
	return m.HeightAt(p) + 1
}

func (m heightMap) LowPoints() []point {
	points := make([]point, 0)
	for y := 0; y < m.LenY(); y++ {
		for x := 0; x < m.LenX(); x++ {
			p := point{y: y, x: x}
			height := m.HeightAt(p)
			neighbours := neighbourhoodHeights(m, p)
			if lowest(neighbours, height) {
				points = append(points, p)
			}
		}
	}
	return points
}

func neighbourhoodHeights(m heightMap, p point) []int {
	heights := make([]int, 0)
	for _, np := range neighbourhood(m, p) {
		heights = append(heights, m.HeightAt(np))
	}
	return heights
}

func neighbourhood(m heightMap, p point) []point {
	neighbourhood := make([]point, 0)

	// top
	if p.y > 0 {
		neighbourhood = append(neighbourhood, point{p.y - 1, p.x})
	}

	// right
	if p.x < m.LenX()-1 {
		neighbourhood = append(neighbourhood, point{p.y, p.x + 1})
	}

	// bottom
	if p.y < m.LenY()-1 {
		neighbourhood = append(neighbourhood, point{p.y + 1, p.x})
	}

	// left
	if p.x > 0 {
		neighbourhood = append(neighbourhood, point{p.y, p.x - 1})
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
	m := read()
	if m.LenY() < 20 {
		fmt.Printf("%v\n", m)
	}

	var part1 int

	lowPoints := m.LowPoints()
	for _, lp := range lowPoints {
		if m.LenY() < 20 {
			fmt.Printf("%v = %d\n", lp, m.HeightAt(lp))
		}
		part1 += m.RiskLevelAt(lp)
	}

	fmt.Printf("part 1: %d\n", part1)
}

// boring input read
func read() heightMap {
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
