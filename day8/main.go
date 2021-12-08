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

type signal uint8

const (
	digit0 signal = signal(segmentA | segmentB | segmentC | segmentE | segmentF | segmentG)
	digit1 signal = signal(segmentC | segmentF)
	digit2 signal = signal(segmentA | segmentC | segmentD | segmentE | segmentG)
	digit3 signal = signal(segmentA | segmentC | segmentD | segmentF | segmentG)
	digit4 signal = signal(segmentB | segmentC | segmentD | segmentF)
	digit5 signal = signal(segmentA | segmentB | segmentD | segmentF | segmentG)
	digit6 signal = signal(segmentA | segmentB | segmentD | segmentE | segmentF | segmentG)
	digit7 signal = signal(segmentA | segmentC | segmentF)
	digit8 signal = signal(segmentA | segmentB | segmentC | segmentD | segmentE | segmentF | segmentG)
	digit9 signal = signal(segmentA | segmentB | segmentC | segmentD | segmentF | segmentG)
)

func (s signal) length() int {
	return bits.OnesCount8(uint8(s))
}

func (s signal) hasSegment(seg segment) bool {
	if uint8(s)&uint8(seg) > 0 {
		return true
	}
	return false
}

func (s signal) format() string {
	buf := make([]byte, s.length())
	for _, segment := range segments {
		if s.hasSegment(segment) {
			buf = append(buf, segment.format())
		}
	}
	return string(buf)
}

func asSignal(in string) signal {
	var s uint8
	for _, c := range in {
		s |= uint8(asSegment(c))
	}
	return signal(s)
}

func main() {
	lines := read()

	// Part 1
	total := 0
	for _, line := range lines {
		if len(lines) < 20 {
			fmt.Printf("%s\n", line.format())
		}
		for _, digit := range line.digits {
			if digit.length() == digit1.length() ||
				digit.length() == digit4.length() ||
				digit.length() == digit7.length() ||
				digit.length() == digit8.length() {
				total++
			}
		}
	}
	if len(lines) < 20 {
		fmt.Println()
	}

	fmt.Printf("Part 1: %d\n", total)
}

// boring input read
type line struct {
	signals []signal
	digits  []signal
}

func (l *line) format() string {
	signals := make([]string, 0)
	for _, s := range l.signals {
		signals = append(signals, s.format())
	}

	digits := make([]string, 0)
	for _, s := range l.digits {
		digits = append(digits, s.format())
	}
	return fmt.Sprintf("%s | %s", strings.Join(signals, " "), strings.Join(digits, " "))
}

func read() []line {
	lines := make([]line, 0)

	reader, closer := selectInput()
	defer closer()
	s := bufio.NewScanner(reader)

	for s.Scan() {
		xs := strings.SplitN(s.Text(), " | ", 2)
		signals := make([]signal, 0)
		for _, x := range strings.Split(xs[0], " ") {
			signals = append(signals, asSignal(x))
		}

		digits := make([]signal, 0)
		for _, x := range strings.Split(xs[1], " ") {
			digits = append(digits, asSignal(x))
		}

		lines = append(lines, line{signals, digits})
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
