package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	pixels := readFromFile("day08/input.txt")[0]

	wide, tall := 25, 6
	sizeOfPixel := wide * tall
	totalOfLayers := len(pixels) / (sizeOfPixel)

	{
		fmt.Println("--- Part One ---")
		pixelCount := make(map[string]int)
		totalOfZeroDigit := math.MaxInt64
		for i := 0; i < totalOfLayers; i++ {
			layer := string(pixels[i*sizeOfPixel : i*sizeOfPixel+sizeOfPixel])

			currentPixelCount := make(map[string]int)
			for _, pixel := range layer {
				currentPixelCount[string(pixel)] = currentPixelCount[string(pixel)] + 1
			}

			if currentPixelCount["0"] < totalOfZeroDigit {
				totalOfZeroDigit = currentPixelCount["0"]
				pixelCount = currentPixelCount
			}
		}
		fmt.Println(pixelCount["1"] * pixelCount["2"])
	}

	{
		fmt.Println("--- Part Two ---")
		positionDigitGroup := make(map[int][]string)
		for i := 0; i < totalOfLayers; i++ {
			layer := string(pixels[i*sizeOfPixel : i*sizeOfPixel+sizeOfPixel])

			for position, pixel := range layer {
				positionDigitGroup[position] = append(positionDigitGroup[position], string(pixel))
			}
		}

		message := ""
		for i := 0; i < sizeOfPixel; i++ {
			for _, digit := range positionDigitGroup[i] {
				if digit != "2" {
					message += digit
					break
				}
			}
		}

		for h := 0; h < tall; h++ {
			for w := 0; w < wide; w++ {
				printColor(string(message[(h*wide)+w]))
			}
			fmt.Println()
		}
	}
}

func printColor(pixel string) {
	switch pixel {
	case "0":
		fmt.Print("█")
	case "1":
		fmt.Print(" ")
	default:
		fmt.Print("?")
	}
}

func readFromFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}
