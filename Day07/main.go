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
	file, err := os.Open("Day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := make([]string, 0)
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), ",")

		instructions = make([]string, len(inputs))
		for i, s := range inputs {
			instructions[i] = s
		}
	}

	bestSignal := -1
	for _, permutations := range allPermutations([]int{0, 1, 2, 3, 4}) {
		input := 0
		for _, permutation := range permutations {
			memory := make([]string, len(instructions))
			copy(memory, instructions)
			output := calculate(memory, []int{permutation, input})
			input = output
			bestSignal = max(bestSignal, output)
		}
	}
	fmt.Println(bestSignal)
}

func calculate(instructions []string, input []int) (output int) {
	index := 0
	output = 0

	for true {
		instruction := LeftPad2Len(instructions[index], "0", 4)
		opcode, _ := strconv.Atoi(string(instruction[2:len(instruction)]))
		p1mode, _ := strconv.Atoi(string(instruction[1:2]))
		p2mode, _ := strconv.Atoi(string(instruction[0:1]))

		if opcode == 99 || index == len(instructions)-1 {
			break
		}

		var parameter1 int
		var parameter2 int
		if opcode == 1 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[index+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[index+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[index+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[index+2])
			}
			result := parameter1 + parameter2

			i, _ := strconv.Atoi(instructions[index+3])
			instructions[i] = strconv.Itoa(result)
			index += 4
		}

		if opcode == 2 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[index+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[index+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[index+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[index+2])
			}

			result := parameter1 * parameter2

			i, _ := strconv.Atoi(instructions[index+3])
			instructions[i] = strconv.Itoa(result)
			index += 4
		}

		if opcode == 3 {
			i, _ := strconv.Atoi(instructions[index+1])

			instructions[i] = strconv.Itoa(input[0])
			input = input[1:]
			index += 2
		}

		if opcode == 4 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[index+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[index+1])
			}

			output = parameter1
			index += 2
		}

		if opcode == 5 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[index+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[index+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[index+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[index+2])
			}

			if parameter1 != 0 {
				index = parameter2
			} else {
				index += 3
			}
		}

		if opcode == 6 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[index+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[index+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[index+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[index+2])
			}

			if parameter1 == 0 {
				index = parameter2
			} else {
				index += 3
			}
		}

		if opcode == 7 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[index+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[index+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[index+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[index+2])
			}

			parameter3, _ := strconv.Atoi(instructions[index+3])

			if parameter1 < parameter2 {
				instructions[parameter3] = strconv.Itoa(1)
			} else {
				instructions[parameter3] = strconv.Itoa(0)
			}

			index += 4
		}

		if opcode == 8 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[index+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[index+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[index+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[index+2])
			}

			parameter3, _ := strconv.Atoi(instructions[index+3])

			if parameter1 == parameter2 {
				instructions[parameter3] = strconv.Itoa(1)
			} else {
				instructions[parameter3] = strconv.Itoa(0)
			}

			index += 4
		}
	}

	return
}

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s

	return retStr[(len(retStr) - overallLen):]
}

func allPermutations(values []int) (result [][]int) {
	if len(values) == 1 {
		result = append(result, values)
		return
	}
	for i, current := range values {
		others := make([]int, 0, len(values)-1)
		others = append(others, values[:i]...)
		others = append(others, values[i+1:]...)
		for _, route := range allPermutations(others) {
			result = append(result, append(route, current))
		}
	}
	return
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
