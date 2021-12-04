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
	numbers, boards := read(os.Stdin)

	var boardsLeft = len(boards)
	var lastWinBoard *board
	var lastWinNum int
bingo:
	for _, n := range numbers {
		for _, b := range boards {
			if b.won {
				continue
			}
			i, j, ok := b.mark(n)
			if !ok {
				continue
			}
			if b.markedInRows[i] == 5 || b.markedInColumns[j] == 5 {
				// winning board
				b.won = true
				lastWinBoard = b
				lastWinNum = n

				boardsLeft--
				if boardsLeft == 0 {
					break bingo
				}
			}
		}
	}

	if lastWinBoard == nil {
		log.Fatalf("No winning board")
	}

	fmt.Printf("%d\n", lastWinBoard.sumUnmarked()*lastWinNum)
}

// data model
type board struct {
	numbers         [5][5]int
	marked          [5][5]bool
	markedInRows    [5]int
	markedInColumns [5]int
	won             bool
}

func (b *board) mark(number int) (int, int, bool) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] == number {
				b.marked[i][j] = true
				b.markedInRows[i]++
				b.markedInColumns[j]++
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func (b *board) sumUnmarked() int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum
}

// boring input read
func read(r io.Reader) ([]int, []*board) {
	numbers := make([]int, 0)
	boards := make([]*board, 0)

	s := bufio.NewScanner(r)

	// Let's get the numbers first
	if !s.Scan() {
		log.Fatalf("No input\n")
	}
	for _, x := range strings.Split(s.Text(), ",") {
		n, err := strconv.Atoi(x)
		if err != nil {
			log.Fatalf("Can't convert item [%s] to int\n", x)
		}
		numbers = append(numbers, n)
	}

	// Now the boards
	for {
		// new line
		if !s.Scan() {
			break
		}
		var numbers [5][5]int
		// 5 lines of 5 numbers
		for i := 0; i < 5; i++ {
			if !s.Scan() {
				log.Fatalf("Expected another line of the board")
			}

			j := 0
			for _, x := range strings.Split(s.Text(), " ") {
				if x == "" {
					// extra space in the input
					continue
				}
				n, err := strconv.Atoi(x)
				if err != nil {
					log.Fatalf("Can't convert item [%s] to int\n", x)
				}
				numbers[i][j] = n
				j++
			}
		}

		boards = append(boards, &board{numbers: numbers})
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner errors: %v\n", err)
	}

	return numbers, boards
}
