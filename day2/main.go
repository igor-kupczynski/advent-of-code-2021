package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var horizontal, depth int
	s := &inputReader{bufio.NewScanner(os.Stdin)}
	for i := 0; ; i++ {
		dir, n, ok := s.MustNext()
		if !ok {
			break
		}
		switch dir {
		case "forward":
			horizontal += n
		case "down":
			depth += n
		case "up":
			depth -= n
		default:
			log.Fatalf("Illegal direction: %s\n", dir)
		}
	}
	fmt.Println(horizontal * depth)
	if err := s.Err(); err != nil {
		panic(err)
	}
}

type inputReader struct {
	scanner *bufio.Scanner
}

func (r *inputReader) MustNext() (string, int, bool) {
	if ok := r.scanner.Scan(); ok {
		var dir string
		var n int
		_, err := fmt.Sscanf(r.scanner.Text(), "%s %d", &dir, &n)
		if err != nil {
			panic(err)
		}
		return dir, n, true
	}
	return "", 0, false
}

func (r *inputReader) Err() error {
	return r.scanner.Err()
}
