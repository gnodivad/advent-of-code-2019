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
	file, err := os.Open("Day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), ",")

		instructions := make([]string, len(inputs))
		for i, s := range inputs {
			instructions[i] = s
		}

		output := calculatePosition0(instructions)

		fmt.Println(output)
	}
}

func calculatePosition0(instructions []string) int {
	index := 0
	output := 0

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

			instructions[i] = strconv.Itoa(1)
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
	}

	return output
}

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s

	return retStr[(len(retStr) - overallLen):]
}
