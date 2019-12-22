package main

import (
	"fmt"
	"gnodivad/aoc/intcode"
	"gnodivad/aoc/utils"
)

func main() {
	inputs := utils.ReadNumbers("day02/input.txt")
	{
		fmt.Println("--- Part One ---")
		mem := make([]int, len(inputs))
		copy(mem, inputs)
		output := intcode.Run(mem)
		fmt.Println(output)
	}
	{
		fmt.Println("--- Part Two ---")
		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				mem := make([]int, len(inputs))
				copy(mem, inputs)
				mem[1] = noun
				mem[2] = verb
				if output := intcode.Run(mem); output == 19690720 {
					fmt.Println(100*noun + verb)
				}
			}
		}
	}
}
