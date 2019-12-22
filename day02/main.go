package main

import (
	"fmt"
	"gnodivad/aoc/intcode"
	"gnodivad/aoc/utils"
)

func main() {
	inputs := utils.ReadNumbers("day02/input.txt")
	output := intcode.Run(inputs)
	fmt.Println(output)
}
