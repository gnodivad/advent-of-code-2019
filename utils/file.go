package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadNumbers return an array consists all the numbers that read from filepath.
func ReadNumbers(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int
	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), ",") {
			number, _ := strconv.Atoi(s)
			numbers = append(numbers, number)
		}
	}

	return numbers
}
