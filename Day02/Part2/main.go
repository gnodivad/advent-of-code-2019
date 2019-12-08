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

		for i := 0; i < 100; i++ {
			for j:= 0; j < 100; j++ {
				if result := calculatePosition0(numbers, i, j); result == 19690720 {
					fmt.Println(100 * i + j)
				}
			}
		}
	}
}

func calculatePosition0(numbers []int, noun int, verb int) int {
	mem := make([] int, len(numbers))
	copy(mem, numbers)

	index := 0
	mem[1] = noun
	mem[2] = verb
	for true {
		if  mem[index] == 99 || index == len(mem)-1 {
			break
		}

		if mem[index] == 1 {
			mem[mem[index+3]] = mem[mem[index+1]] + mem[mem[index+2]]
		}

		if mem[index] == 2 {
			mem[mem[index+3]] = mem[mem[index+1]] * mem[mem[index+2]]			
		}

		index += 4
	}

	return mem[0]
}
