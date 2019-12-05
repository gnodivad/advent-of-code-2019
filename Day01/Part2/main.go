package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("Day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		number = (number / 3) - 2

		for number > 0 {
			sum += number
			number = (number / 3) - 2
		}
	}

	fmt.Println(sum)
}
