package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("Day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	orbits := make(map[string]string)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ")")

		orbits[input[1]] = input[0]
	}

	sum := 0
	for k := range orbits {
		from := k
		for true {
			if _, ok := orbits[from]; !ok {
				break
			}

			sum++

			from = orbits[from]
		}
	}

	fmt.Println(sum)
}
