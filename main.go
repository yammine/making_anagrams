package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	first := scanner.Text()
	scanner.Scan()
	second := scanner.Text()

	var frequency [26]int
	var count int

	for i := 0; i < len(first); i++ {
		letter := first[i]
		frequency[letter-'a'] += 1
	}

	for j := 0; j < len(second); j++ {
		letter := second[j]
		frequency[letter-'a'] -= 1
	}

	for _, val := range frequency {
		count += int(math.Abs(float64(val)))
	}

	fmt.Println(count)
}
