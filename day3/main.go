package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := FromStdio()
	var sum []int
	var termsNo int
	for ; ; termsNo++ {
		line, ok := in.MustNext()
		if !ok {
			break
		}
		if sum == nil {
			sum = make([]int, len(line), len(line))
		}
		add(sum, line)
	}
	most, least := getMostLeastSig(sum, termsNo)
	fmt.Println(sum, termsNo)
	fmt.Printf(" most: %012b => %d\n", most, most)
	fmt.Printf("least: %012b => %d\n", least, least)
	fmt.Printf("%d\n", most*least)
	if err := in.Err(); err != nil {
		panic(err)
	}
}

func add(acc []int, term []int) {
	if len(acc) != len(term) {
		log.Fatalf("Different lengths of sum %v and term %v\n", acc, term)
	}
	for i := 0; i < len(term); i++ {
		acc[i] += term[i]
	}
}

func getMostLeastSig(acc []int, termsNo int) (most, least uint32) {
	for i := 0; i < len(acc); i++ {
		ones := acc[i]
		zeros := termsNo - acc[i]
		pos := len(acc) - i - 1
		if ones > zeros {
			most |= 1 << pos
		} else {
			least |= 1 << pos
		}
	}
	return
}

// boring input reader

type inputReader struct {
	scanner *bufio.Scanner
}

func FromStdio() *inputReader {
	return &inputReader{bufio.NewScanner(os.Stdin)}
}

func (r *inputReader) MustNext() ([]int, bool) {
	if ok := r.scanner.Scan(); ok {
		var input string
		_, err := fmt.Sscanf(r.scanner.Text(), "%s", &input)
		if err != nil {
			panic(err)
		}

		result := make([]int, len(input), len(input))
		for i, ch := range input {
			result[i] = int(ch - '0')
		}
		return result, true
	}
	return nil, false
}

func (r *inputReader) Err() error {
	return r.scanner.Err()
}
