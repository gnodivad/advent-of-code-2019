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

	mapData := make(map[string]int)

	adjacentDigits := false

	for pos, char := range s {
		if pos > 0 && string(char) < string(s[pos-1]) {
			return false
		}

		mapData[string(char)] = mapData[string(char)] + 1
	}

	for _, v := range mapData {
		if v%2 == 0 && v/2 < 2 {
			adjacentDigits = true
		}
	}

	return adjacentDigits
}
