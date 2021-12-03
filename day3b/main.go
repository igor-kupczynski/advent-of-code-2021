package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readAll()
	sum := make([]int, len(input[0]), len(input[0]))
	for _, term := range input {
		add(sum, term)
	}
	oxygenBits := filterBitByBit(input, getMostCommon)
	oxygen := bitSliceToNumber(oxygenBits)
	co2Bits := filterBitByBit(input, getLeastCommon)
	co2 := bitSliceToNumber(co2Bits)
	fmt.Printf("oxy: %v, %012b, %d\n", oxygenBits, oxygen, oxygen)
	fmt.Printf("co2: %v, %012b, %d\n", co2Bits, co2, co2)
	fmt.Printf("%d\n", oxygen*co2)
}

func add(acc []int, term []int) {
	if len(acc) != len(term) {
		log.Fatalf("Different lengths of sum %v and term %v\n", acc, term)
	}
	for i := 0; i < len(term); i++ {
		acc[i] += term[i]
	}
}

func filterBitByBit(input [][]int, selection func(input [][]int, pos int) int) []int {
	var result [][]int
	for i := 0; i < len(input); i++ {
		prevResult := result
		if result == nil {
			prevResult = input
		}
		toKeep := selection(prevResult, i)
		result = keepWithBitAtPos(prevResult, toKeep, i)
		if len(result) == 1 {
			return result[0]
		}
	}
	log.Fatalf("Can't filter to a single number, got: %v\n", result)
	return nil
}

func keepWithBitAtPos(input [][]int, toKeep int, pos int) [][]int {
	result := make([][]int, 0)
	for _, item := range input {
		if item[pos] == toKeep {
			result = append(result, item)
		}
	}
	return result
}

func getMostCommon(input [][]int, pos int) int {
	var zeros, ones int
	for _, item := range input {
		if item[pos] == 1 {
			ones++
		} else {
			zeros++
		}
	}
	if ones >= zeros {
		return 1
	}
	return 0
}

func getLeastCommon(input [][]int, pos int) int {
	var zeros, ones int
	for _, item := range input {
		if item[pos] == 1 {
			ones++
		} else {
			zeros++
		}
	}
	if zeros <= ones {
		return 0
	}
	return 1
}

func bitSliceToNumber(acc []int) uint32 {
	var result uint32
	for i := 0; i < len(acc); i++ {
		pos := len(acc) - i - 1
		if acc[i] == 1 {
			result |= 1 << pos
		}
	}
	return result
}

// boring input reader

func readAll() [][]int {
	result := make([][]int, 0)
	in := FromStdio()
	for {
		line, ok := in.MustNext()
		if !ok {
			break
		}
		result = append(result, line)
	}
	if err := in.Err(); err != nil {
		panic(err)
	}
	return result
}

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
