package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	hMap := read()
	if len(hMap) < 20 {
		fmt.Printf("%v", hMap)
	}
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
