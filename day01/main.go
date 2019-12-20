package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	masses := readNumbersFromFile("day01/input.txt")

	fuelRequired, exactFuelRequired := 0, 0
	for _, mass := range masses {
		mass = CalculateFuelRequired(mass)
		fuelRequired += mass

		for mass > 0 {
			exactFuelRequired += mass
			mass = CalculateFuelRequired(mass)
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(fuelRequired)

	fmt.Println("--- Part Two ---")
	fmt.Println(exactFuelRequired)
}

// CalculateFuelRequired return total of the fuel required per unit of mass
func CalculateFuelRequired(mass int) int {
	return (mass / 3) - 2
}

func readNumbersFromFile(filepath string) (numbers []int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers = append(numbers, stringToInt(scanner.Text()))
	}

	return
}

func stringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
