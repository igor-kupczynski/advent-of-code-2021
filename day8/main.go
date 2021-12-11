package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"strings"
)

type segment uint8

const (
	segmentA segment = 1 << 0
	segmentB segment = 1 << 1
	segmentC segment = 1 << 2
	segmentD segment = 1 << 3
	segmentE segment = 1 << 4
	segmentF segment = 1 << 5
	segmentG segment = 1 << 6
)

var segments = []segment{segmentA, segmentB, segmentC, segmentD, segmentE, segmentF, segmentG}

func (s segment) format() byte {
	switch s {
	case segmentA:
		return 'a'
	case segmentB:
		return 'b'
	case segmentC:
		return 'c'
	case segmentD:
		return 'd'
	case segmentE:
		return 'e'
	case segmentF:
		return 'f'
	case segmentG:
		return 'g'
	default:
		log.Fatalf("Can't match %b to a segment", s)
		return ' '
	}
}

func asSegment(in rune) segment {
	switch in {
	case 'a':
		return segmentA
	case 'b':
		return segmentB
	case 'c':
		return segmentC
	case 'd':
		return segmentD
	case 'e':
		return segmentE
	case 'f':
		return segmentF
	case 'g':
		return segmentG
	default:
		log.Fatalf("Can't match %c to a segment", in)
		return 0
	}
}

type segmentSet uint8

const (
	digit0 = segmentSet(segmentA | segmentB | segmentC | segmentE | segmentF | segmentG)
	digit1 = segmentSet(segmentC | segmentF)
	digit2 = segmentSet(segmentA | segmentC | segmentD | segmentE | segmentG)
	digit3 = segmentSet(segmentA | segmentC | segmentD | segmentF | segmentG)
	digit4 = segmentSet(segmentB | segmentC | segmentD | segmentF)
	digit5 = segmentSet(segmentA | segmentB | segmentD | segmentF | segmentG)
	digit6 = segmentSet(segmentA | segmentB | segmentD | segmentE | segmentF | segmentG)
	digit7 = segmentSet(segmentA | segmentC | segmentF)
	digit8 = segmentSet(segmentA | segmentB | segmentC | segmentD | segmentE | segmentF | segmentG)
	digit9 = segmentSet(segmentA | segmentB | segmentC | segmentD | segmentF | segmentG)
)

func (s segmentSet) value() int {
	switch s {
	case digit0:
		return 0
	case digit1:
		return 1
	case digit2:
		return 2
	case digit3:
		return 3
	case digit4:
		return 4
	case digit5:
		return 5
	case digit6:
		return 6
	case digit7:
		return 7
	case digit8:
		return 8
	case digit9:
		return 9
	default:
		return -1
	}
}

func (s segmentSet) length() int {
	return bits.OnesCount8(uint8(s))
}

func (s segmentSet) has(seg segment) bool {
	if uint8(s)&uint8(seg) > 0 {
		return true
	}
	return false
}

func (s segmentSet) countSegments() (segment, int) {
	var count int
	var got segment
	for _, seg := range segments {
		if s.has(seg) {
			count++
			got = seg
		}
	}

	return got, count
}

func (s segmentSet) format() string {
	buf := make([]byte, 0, s.length())
	for _, segment := range segments {
		if s.has(segment) {
			buf = append(buf, segment.format())
		}
	}
	return string(buf)
}

func asSegmentSet(in string) segmentSet {
	var s uint8
	for _, c := range in {
		s |= uint8(asSegment(c))
	}
	return segmentSet(s)
}

func diff(a segmentSet, b segmentSet, xs ...segmentSet) segmentSet {
	result := a
	items := append([]segmentSet{b}, xs...)
	for i := 0; i < len(items); i++ {
		result = segmentSet(uint8(result) ^ uint8(items[i]))
	}
	return result
}

type scrambledSegmentsSolver struct {
	inputWires  []segmentSet
	inputDigits []segmentSet

	matchedDigits   map[segmentSet]segmentSet
	matchedSegments map[segment]segment
}

func NewSolver(p puzzle) *scrambledSegmentsSolver {
	inputWires := make([]segmentSet, 0, 10)
	inputDigits := make([]segmentSet, 0, 4)
	for _, wire := range p.wires {
		inputWires = append(inputWires, wire)
	}
	for _, digit := range p.digits {
		inputDigits = append(inputDigits, digit)
	}
	return &scrambledSegmentsSolver{
		inputWires:      inputWires,
		inputDigits:     inputDigits,
		matchedDigits:   map[segmentSet]segmentSet{},
		matchedSegments: map[segment]segment{},
	}
}

func (s *scrambledSegmentsSolver) Solve() map[segmentSet]segmentSet {
	// Find easy digits
	s.findEasyDigits()

	// Deduce other digits
	s.deduceOtherDigits()

	// Invert the matched digits (so that we can transform the other way)
	solution := map[segmentSet]segmentSet{}
	for k, v := range s.matchedDigits {
		solution[v] = k
	}
	return solution
}

func (s *scrambledSegmentsSolver) findEasyDigits() {
	unmatchedDigits := make([]segmentSet, 0, 6)
	for _, wire := range s.inputWires {
		switch wire.length() {
		case digit1.length():
			s.matchedDigits[digit1] = wire
		case digit4.length():
			s.matchedDigits[digit4] = wire
		case digit7.length():
			s.matchedDigits[digit7] = wire
		case digit8.length():
			s.matchedDigits[digit8] = wire
		default:
			unmatchedDigits = append(unmatchedDigits, wire)
		}
	}
	s.inputWires = unmatchedDigits
}

func (s *scrambledSegmentsSolver) deduceOtherDigits() {
	// Find segment A  (difference between digits 7 and 1)
	// {A} = [7] - [1]
	s.matchedSegments[segmentA] = mustOne(
		"{A}",
		diff(
			s.matchedDigits[digit7],
			s.matchedDigits[digit1],
		),
	)

	// Find digit 3  (like digit 7, but with two extra segments D, G)
	// [3] - [7] = {D, G}
	for i := 0; i < len(s.inputWires); i++ {
		if _, count := diff(
			s.inputWires[i],
			s.matchedDigits[digit7],
		).countSegments(); count == 2 {
			s.matchedDigits[digit3] = s.inputWires[i]
			s.inputWires = append(s.inputWires[:i], s.inputWires[i+1:]...)
			break
		}
	}

	// Find digit 9 and segment B
	// [9] - [3] = {B}
	for i := 0; i < len(s.inputWires); i++ {
		if got, count := diff(
			s.inputWires[i],
			s.matchedDigits[digit3],
		).countSegments(); count == 1 {
			s.matchedDigits[digit9] = s.inputWires[i]
			s.inputWires = append(s.inputWires[:i], s.inputWires[i+1:]...)
			s.matchedSegments[segmentB] = got
			break
		}
	}

	// Find segment E
	// {E} = [8] - [9]
	s.matchedSegments[segmentE] = mustOne(
		"{E}",
		diff(
			s.matchedDigits[digit8],
			s.matchedDigits[digit9],
		),
	)

	// Find segment D
	// {D} = [4] - [1] - B
	s.matchedSegments[segmentD] = mustOne(
		"{D}",
		diff(
			s.matchedDigits[digit4],
			s.matchedDigits[digit1],
			segmentSet(s.matchedSegments[segmentB]),
		),
	)

	// Find segment G
	// {G} = [8] - [1] - {B} - {D} - {E}
	s.matchedSegments[segmentG] = mustOne(
		"{G}",
		diff(
			s.matchedDigits[digit8],
			s.matchedDigits[digit1],
			segmentSet(s.matchedSegments[segmentA]),
			segmentSet(s.matchedSegments[segmentB]),
			segmentSet(s.matchedSegments[segmentD]),
			segmentSet(s.matchedSegments[segmentE]),
		),
	)

	// Find digit 6 and segment F
	// {F} = [6] - {A} - {B} - {D} - {E} - {G}
	for i := 0; i < len(s.inputWires); i++ {
		if got, count := diff(
			s.inputWires[i],
			segmentSet(s.matchedSegments[segmentA]),
			segmentSet(s.matchedSegments[segmentB]),
			segmentSet(s.matchedSegments[segmentD]),
			segmentSet(s.matchedSegments[segmentE]),
			segmentSet(s.matchedSegments[segmentG]),
		).countSegments(); count == 1 {
			s.matchedDigits[digit6] = s.inputWires[i]
			s.inputWires = append(s.inputWires[:i], s.inputWires[i+1:]...)
			s.matchedSegments[segmentF] = got
			break
		}
	}

	// Find segment C
	// {C} = [1] - {F}
	s.matchedSegments[segmentC] = mustOne(
		"{C}",
		diff(
			s.matchedDigits[digit1],
			segmentSet(s.matchedSegments[segmentF]),
		),
	)

	// Find digit 5
	// {5} = {6} - {E}
	for i := 0; i < len(s.inputWires); i++ {
		if diff(
			s.inputWires[i],
			segmentSet(s.matchedSegments[segmentE]),
		) == s.matchedDigits[digit6] {
			s.matchedDigits[digit5] = s.inputWires[i]
			s.inputWires = append(s.inputWires[:i], s.inputWires[i+1:]...)
			break
		}
	}

	// Find digit 0
	// [0] = [8] - {D}
	for i := 0; i < len(s.inputWires); i++ {
		if diff(
			s.matchedDigits[digit8],
			segmentSet(s.matchedSegments[segmentD]),
		) == s.inputWires[i] {
			s.matchedDigits[digit0] = s.inputWires[i]
			s.inputWires = append(s.inputWires[:i], s.inputWires[i+1:]...)
			break
		}
	}

	// Find digit 2
	// The last remaining digit
	s.matchedDigits[digit2] = s.inputWires[0]
	s.inputWires = s.inputWires[1:]

	// Assert that we've processed the input and have the solution
	if len(s.inputWires) > 0 {
		log.Fatalf("Left with input to process: %v", s.inputWires)
	}
	if len(s.matchedDigits) != 10 {
		log.Fatalf("Want all 10 digits, got: %v", s.matchedDigits)
	}
	if len(s.matchedSegments) != 7 {
		log.Fatalf("Want all 7 digits, got: %v", s.matchedSegments)
	}
}

func mustOne(desc string, s segmentSet) segment {
	if got, count := s.countSegments(); count == 1 {
		return got
	}
	panic(fmt.Errorf("the segment set %s: %0b contains more than a single segment", desc, s))
	return 0
}

func main() {
	lines := read()

	countPart1 := 0
	sumPart2 := 0
	for _, line := range lines {
		if len(lines) < 20 {
			fmt.Printf("%s\n", line.format())
		}

		solver := NewSolver(line)
		solution := solver.Solve()

		for i, digit := range line.digits {
			switch solution[digit] {
			case digit1, digit4, digit7, digit8:
				countPart1++
			}
			digit := solution[digit].value()
			for j := 0; j < 3-i; j++ {
				digit *= 10
			}
			sumPart2 += digit
		}
	}
	if len(lines) < 20 {
		fmt.Println()
	}

	fmt.Printf("Part 1: %d\n", countPart1)
	fmt.Printf("Part 2: %d\n", sumPart2)
	for _, line := range lines {
		_ = line
	}
}

// boring input read
type puzzle struct {
	wires  []segmentSet
	digits []segmentSet
}

func (l *puzzle) format() string {
	signals := make([]string, 0)
	for _, s := range l.wires {
		signals = append(signals, s.format())
	}

	digits := make([]string, 0)
	for _, s := range l.digits {
		digits = append(digits, s.format())
	}
	return fmt.Sprintf("%s | %s", strings.Join(signals, " "), strings.Join(digits, " "))
}

func read() []puzzle {
	lines := make([]puzzle, 0)

	reader, closer := selectInput()
	defer closer()
	s := bufio.NewScanner(reader)

	for s.Scan() {
		xs := strings.SplitN(s.Text(), " | ", 2)
		signals := make([]segmentSet, 0)
		for _, x := range strings.Split(xs[0], " ") {
			signals = append(signals, asSegmentSet(x))
		}

		digits := make([]segmentSet, 0)
		for _, x := range strings.Split(xs[1], " ") {
			digits = append(digits, asSegmentSet(x))
		}

		lines = append(lines, puzzle{signals, digits})
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
