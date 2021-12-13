package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type token byte

const (
	empty       token = 0
	roundOpen   token = '('
	roundClose  token = ')'
	squareOpen  token = '['
	squareClose token = ']'
	curlyOpen   token = '{'
	curlyClose  token = '}'
	angleOpen   token = '<'
	angleClose  token = '>'
)

var tokens = map[token]struct{}{
	roundOpen: {}, roundClose: {},
	squareOpen: {}, squareClose: {},
	curlyOpen: {}, curlyClose: {},
	angleOpen: {}, angleClose: {},
}

func isOpeningBracket(t token) bool {
	switch t {
	case roundOpen:
		return true
	case squareOpen:
		return true
	case curlyOpen:
		return true
	case angleOpen:
		return true
	default:
		return false
	}
}

func isClosingBracket(t token) bool {
	switch t {
	case roundClose:
		return true
	case squareClose:
		return true
	case curlyClose:
		return true
	case angleClose:
		return true
	default:
		return false
	}
}

func pair(t token) token {
	switch t {
	case roundOpen:
		return roundClose
	case squareOpen:
		return squareClose
	case curlyOpen:
		return curlyClose
	case angleOpen:
		return angleClose
	case roundClose:
		return roundOpen
	case squareClose:
		return squareOpen
	case curlyClose:
		return curlyOpen
	case angleClose:
		return angleOpen
	default:
		return t
	}
}

//func matchesScope(from, to token) bool {
//	if (from == roundOpen && to == roundClose) ||
//		(from == squareOpen && to == squareClose) ||
//		(from == curlyOpen && to == curlyClose) ||
//		(from == angleOpen && to == angleClose) {
//		return true
//	}
//	return false
//}

type scanner struct {
	buf []byte
}

func NewScanner(input string) *scanner {
	return &scanner{buf: []byte(input)}
}

func (s *scanner) Next() (token, bool) {
	for len(s.buf) > 0 {
		var x byte
		x, s.buf = s.buf[0], s.buf[1:]
		candidate := token(x)
		if _, ok := tokens[candidate]; !ok {
			// not a token, ignore
			continue
		}
		return candidate, true
	}
	return 0, false
}

func main() {
	lines := read()
	autocompletePoints := make([]int, 0)
	var part1, part2 int

	for _, line := range lines {
		scopes := make([]token, 0)
		scopeError := empty

		scanner := NewScanner(line)
	nextToken:
		for {
			t, ok := scanner.Next()
			if !ok {
				break nextToken
			}
			if isOpeningBracket(t) {
				// Add new scope to stack
				scopes = append(scopes, t)
				continue nextToken
			}
			if isClosingBracket(t) {
				// Remove the scope from scopes if valid
				top := scopes[len(scopes)-1]
				if pair(top) == t {
					// Valid, closes the scope
					scopes = scopes[:len(scopes)-1]
					continue nextToken
				}
				scopeError = t
				break nextToken
			}
			// What are we even doing here, panic
			panic(fmt.Errorf("illegal token %c", t))
		}

		if scopeError == empty && len(scopes) > 0 {
			// Incomplete line
			if len(lines) < 20 {
				fmt.Printf("- `%s` - Incomplete, open scopes: %v, ", line, scopes)
			}

			points := 0
			// Complete
			for i := len(scopes) - 1; i >= 0; i-- {
				switch pair(scopes[i]) {
				case roundClose:
					points = points*5 + 1
				case squareClose:
					points = points*5 + 2
				case curlyClose:
					points = points*5 + 3
				case angleClose:
					points = points*5 + 4
				}
			}
			autocompletePoints = append(autocompletePoints, points)
			if len(lines) < 20 {
				fmt.Printf("points: %d\n", points)
			}
		}
		if scopeError != empty {
			// Scope error
			if len(lines) < 20 {
				fmt.Printf("- `%s` - Expected: %c, but found %c instead.\n", line, pair(scopes[len(scopes)-1]), scopeError)
			}

			points := 0
			switch scopeError {
			case roundClose:
				points = 3
			case squareClose:
				points = 57
			case curlyClose:
				points = 1197
			case angleClose:
				points = 25137
			}
			part1 += points
		}
	}

	sort.Ints(autocompletePoints)
	part2 = autocompletePoints[len(autocompletePoints)/2]

	fmt.Printf("part 1: %d\n", part1)
	fmt.Printf("part 2: %d\n", part2)
}

// boring input read
func read() []string {
	lines := make([]string, 0)

	reader, closer := selectInput()
	defer closer()
	s := bufio.NewScanner(reader)

	for s.Scan() {
		lines = append(lines, s.Text())
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
