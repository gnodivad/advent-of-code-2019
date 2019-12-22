package intcode

const (
	add      = 1
	multiply = 2
	halt     = 99
)

// Run perform calculation based on opcode provided in inputs
func Run(inputs []int) int {
	pointer := 0

	for {
		opcode := inputs[pointer]

		switch opcode {
		case halt:
			return inputs[0]
		case add:
			result := inputs[inputs[pointer+1]] + inputs[inputs[pointer+2]]
			inputs[inputs[pointer+3]] = result
			pointer += movePointer(opcode)
		case multiply:
			result := inputs[inputs[pointer+1]] * inputs[inputs[pointer+2]]
			inputs[inputs[pointer+3]] = result
			pointer += movePointer(opcode)
		}
	}
}

func movePointer(opcode int) int {
	switch opcode {
	case add:
		return 4
	case multiply:
		return 4
	}

	return 0
}
