package main

import (
	"fmt"
	"strconv"
)

func main() {
	total := 0

	for i := 273025; i <= 767253; i++ {
		if IsValidNumber(i) {
			total++
		}
	}

	fmt.Println(total)
}

func IsValidNumber(number int) bool {
	s := strconv.Itoa(number)
	adjacentDigits := false

	for pos, char := range s {
		if pos > 0 && string(char) < string(s[pos-1]) {
			return false
		}

		if pos > 0 && string(char) == string(s[pos-1]) {
			adjacentDigits = true
		}
	}

	return adjacentDigits
}
