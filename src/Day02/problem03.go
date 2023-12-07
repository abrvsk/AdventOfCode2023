package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Problem03() int {
	buffer, err := os.Open("src/Day02/problem03.csv")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanLines)

	res := 0
	maxReds, maxGreens, maxBlues := 12, 13, 14
	for scanner.Scan() {
		res += calculateAmountOfCubes(scanner.Text(), maxReds, maxGreens, maxBlues)
	}

	return res
}

func calculateAmountOfCubes(str string, maxReds, maxGreens, maxBlues int) int {
	split := strings.Split(str, ": ")
	game := strings.Split(split[0], " ")
	gameNumber, err := strconv.Atoi(game[1])
	if err != nil {
		fmt.Println(err)
	}
	draws := strings.Split(split[1], "; ")

	for _, draw := range draws {
		cubes := strings.Split(draw, ", ")
		for _, cubeDraw := range cubes {
			if strings.Contains(cubeDraw, "red") {
				reds := strings.Split(cubeDraw, " ")
				redAmount, err := strconv.Atoi(reds[0])
				if err == nil && redAmount > maxReds {
					gameNumber = 0
				}
			}
			if strings.Contains(cubeDraw, "green") {
				greens := strings.Split(cubeDraw, " ")
				greenAmount, err := strconv.Atoi(greens[0])
				if err == nil && greenAmount > maxGreens {
					gameNumber = 0
				}
			}
			if strings.Contains(cubeDraw, "blue") {
				blues := strings.Split(cubeDraw, " ")
				blueAmount, err := strconv.Atoi(blues[0])
				if err == nil && blueAmount > maxBlues {
					gameNumber = 0
				}
			}
		}
	}
	return gameNumber
}
