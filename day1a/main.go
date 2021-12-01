package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	var increases int = 0
	var previous int = math.MaxInt32
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var current int
		_, err := fmt.Sscanf(s.Text(), "%d", &current)
		if err != nil {
			log.Fatal(err)
		}
		if current > previous {
			increases++
		}
		previous = current
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(increases)
}
