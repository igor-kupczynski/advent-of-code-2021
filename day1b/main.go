package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	increases := 0
	buf := make([]int, 0)
	s := &intReader{bufio.NewScanner(os.Stdin)}
	for i := 0; ; i++ {
		n, ok := s.MustNext()
		if !ok {
			break
		}
		buf = append(buf, n)
		if i >= 3 {
			prev := buf[i-1] + buf[i-2] + buf[i-3]
			curr := buf[i] + buf[i-1] + buf[i-2]
			if curr > prev {
				increases++
			}
		}
	}
	fmt.Println(increases)
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

type intReader struct {
	scanner *bufio.Scanner
}

func (r *intReader) MustNext() (int, bool) {
	if ok := r.scanner.Scan(); ok {
		var item int
		_, err := fmt.Sscanf(r.scanner.Text(), "%d", &item)
		if err != nil {
			panic(err)
		}
		return item, true
	}
	return 0, false
}

func (r *intReader) Err() error {
	return r.scanner.Err()
}
