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
	for _, permutations := range allPermutations([]int{5, 6, 7, 8, 9}) {
		signal := 0

		var memories [][]string
		for i := 0; i < 5; i++ {
			memory := make([]string, len(instructions))
			copy(memory, instructions)
			memories = append(memories, memory)
		}
		programCounters := make([]int, 5)

		for feedbackLoop := 0; programCounters[4] != -1; feedbackLoop++ {
			for amplifier, permutation := range permutations {
				if feedbackLoop == 0 {
					signal, programCounters[amplifier] = calculate(memories[amplifier], []int{permutation, signal}, programCounters[amplifier])
				} else {
					signal, programCounters[amplifier] = calculate(memories[amplifier], []int{signal}, programCounters[amplifier])
				}
			}
		}
		bestSignal = max(bestSignal, signal)
	}
	fmt.Println(bestSignal)
}

func calculate(instructions []string, input []int, programCounter int) (int, int) {
	output := 0

	for programCounter > -1 {
		instruction := LeftPad2Len(instructions[programCounter], "0", 4)
		opcode, _ := strconv.Atoi(string(instruction[2:len(instruction)]))
		p1mode, _ := strconv.Atoi(string(instruction[1:2]))
		p2mode, _ := strconv.Atoi(string(instruction[0:1]))

		if opcode == 99 {
			programCounter = -1
		}

		var parameter1 int
		var parameter2 int
		if opcode == 1 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[programCounter+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[programCounter+2])
			}
			result := parameter1 + parameter2

			i, _ := strconv.Atoi(instructions[programCounter+3])
			instructions[i] = strconv.Itoa(result)
			programCounter += 4
		}

		if opcode == 2 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[programCounter+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[programCounter+2])
			}

			result := parameter1 * parameter2

			i, _ := strconv.Atoi(instructions[programCounter+3])
			instructions[i] = strconv.Itoa(result)
			programCounter += 4
		}

		if opcode == 3 {
			if len(input) == 0 {
				return output, programCounter
			}

			i, _ := strconv.Atoi(instructions[programCounter+1])

			instructions[i] = strconv.Itoa(input[0])
			input = input[1:]
			programCounter += 2
		}

		if opcode == 4 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[programCounter+1])
			}

			output = parameter1
			programCounter += 2
		}

		if opcode == 5 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[programCounter+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[programCounter+2])
			}

			if parameter1 != 0 {
				programCounter = parameter2
			} else {
				programCounter += 3
			}
		}

		if opcode == 6 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[programCounter+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[programCounter+2])
			}

			if parameter1 == 0 {
				programCounter = parameter2
			} else {
				programCounter += 3
			}
		}

		if opcode == 7 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[programCounter+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[programCounter+2])
			}

			parameter3, _ := strconv.Atoi(instructions[programCounter+3])

			if parameter1 < parameter2 {
				instructions[parameter3] = strconv.Itoa(1)
			} else {
				instructions[parameter3] = strconv.Itoa(0)
			}

			programCounter += 4
		}

		if opcode == 8 {
			if p1mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+1])
				parameter1, _ = strconv.Atoi(instructions[i])
			} else {
				parameter1, _ = strconv.Atoi(instructions[programCounter+1])
			}

			if p2mode == 0 {
				i, _ := strconv.Atoi(instructions[programCounter+2])
				parameter2, _ = strconv.Atoi(instructions[i])
			} else {
				parameter2, _ = strconv.Atoi(instructions[programCounter+2])
			}

			parameter3, _ := strconv.Atoi(instructions[programCounter+3])

			if parameter1 == parameter2 {
				instructions[parameter3] = strconv.Itoa(1)
			} else {
				instructions[parameter3] = strconv.Itoa(0)
			}

			programCounter += 4
		}
	}
	
	return output, programCounter
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
