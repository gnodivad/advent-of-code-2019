package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("Day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	minDistance := math.MaxInt64
	mapData := make(map[Point]bool)
	idx, originX, originY := 0, 0, 0

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")

		currentX := originX
		currentY := originY

		for _, s := range input {
			action := string(s[0])
			step, _ := strconv.Atoi(s[1:])

			dirX, dirY := 0, 0
			if action == "R" {
				dirX = 1
				dirY = 0
			} else if action == "L" {
				dirX = -1
				dirY = 0
			} else if action == "U" {
				dirX = 0
				dirY = 1
			} else if action == "D" {
				dirX = 0
				dirY = -1
			}

			for i := 0; i < step; i++ {
				currentX += dirX
				currentY += dirY

				if idx == 0 {
					mapData[Point{currentX, currentY}] = true
				}

				if idx == 1 && mapData[Point{currentX, currentY}] == true {
					if distance := calculateManhanttanDistance(originX, originY, currentX, currentY); minDistance > distance {
						minDistance = distance
					}
				}
			}
		}

		idx++
	}

	fmt.Println(minDistance)
}

func calculateManhanttanDistance(originX, originY, currentX, currentY int) int {
	return Abs(originX-currentX) + Abs(originY-currentY)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
