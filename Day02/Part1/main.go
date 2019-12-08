package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("Day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), ",")

		numbers := make([]int, len(inputs))
		for i, s := range inputs {
			number, _ := strconv.Atoi(s)
			numbers[i] = number
		}

		calculatePosition0(numbers)
	}
}

func calculatePosition0(numbers []int) {
	index := 0
	for true {
		if numbers[index] == 99 || index == len(numbers)-1 {
			break
		}

		if numbers[index] == 1 {
			result := numbers[numbers[index+1]] + numbers[numbers[index+2]]
			numbers[numbers[index+3]] = result
			index += 4
		}

		if numbers[index] == 2 {
			result := numbers[numbers[index+1]] * numbers[numbers[index+2]]
			numbers[numbers[index+3]] = result
			index += 4
		}
	}

	fmt.Println(numbers[0])
}
