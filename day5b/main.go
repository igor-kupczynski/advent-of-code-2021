package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	lines, maxX, maxY := read(os.Stdin)
	diagram := NewDiagram(maxX, maxY)
	for _, line := range lines {
		diagram.draw(line)
	}

	if maxX < 20 && maxY < 20 {
		printDiagram(diagram)
	}
	fmt.Printf("\nAnswer: %d\n", diagram.count(func(n int) bool {
		return n >= 2
	}))
}

func printDiagram(diagram *diagram) {
	for i := 0; i < len(diagram.board); i++ {
		for j := 0; j < len(diagram.board[0]); j++ {
			pixel := diagram.board[j][i]
			if pixel == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", pixel)
			}
		}
		fmt.Printf("\n")
	}
}

// data model
type line struct {
	x1, y1 int
	x2, y2 int
}

func (l line) vector() (dx int, dy int, length int) {
	if l.x2 > l.x1 {
		dx = +1
		length = l.x2 - l.x1 + 1
	} else if l.x2 == l.x1 {
		dx = 0
	} else {
		dx = -1
		length = l.x1 - l.x2 + 1
	}
	if l.y2 > l.y1 {
		dy = +1
		length = l.y2 - l.y1 + 1
	} else if l.y2 == l.y1 {
		dy = 0
	} else {
		dy = -1
		length = l.y1 - l.y2 + 1
	}
	return
}

type diagram struct {
	board [][]int
}

func NewDiagram(maxX, maxY int) *diagram {
	board := make([][]int, maxX+1, maxX+1)
	for i := 0; i < maxX+1; i++ {
		board[i] = make([]int, maxY+1, maxY+1)
	}
	return &diagram{board: board}
}

func (d *diagram) draw(l line) {
	dx, dy, length := l.vector()
	x := l.x1
	y := l.y1
	for n := 0; n < length; n++ {
		d.board[x][y]++
		x += dx
		y += dy
	}
}

func (d *diagram) count(filter func(int) bool) int {
	count := 0
	for i := 0; i < len(d.board); i++ {
		for j := 0; j < len(d.board[0]); j++ {
			if filter(d.board[i][j]) {
				count++
			}
		}
	}
	return count
}

// boring input read
func read(r io.Reader) (lines []line, maxX int, maxY int) {
	lines = make([]line, 0)

	s := bufio.NewScanner(r)

	for s.Scan() {
		var x1, x2, y1, y2 int
		if _, err := fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2); err != nil {
			log.Fatalf("Can't parse line of input: %v", err)
		}
		lines = append(lines, line{x1: x1, y1: y1, x2: x2, y2: y2})
		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = y1
		}
		if y2 > maxY {
			maxY = y2
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner errors: %v\n", err)
	}

	return
}
