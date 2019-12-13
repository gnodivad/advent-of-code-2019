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

	path := make(map[string]int)
	object, distance := orbits["YOU"], 0

	for true {
		path[object] = distance

		if _, ok := orbits[object]; !ok {
			break
		}

		object = orbits[object]
		distance++
	}

	santa, sum := orbits["SAN"], 0
	for true {
		if _, ok := path[santa]; ok {
			sum += path[santa]
			break
		}

		sum++
		santa = orbits[santa]
	}
	fmt.Println(sum)
}
