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
                gameNumber = checkColorDraw(cubeDraw, "red", maxReds, gameNumber)
			}
			if strings.Contains(cubeDraw, "green") {
                gameNumber = checkColorDraw(cubeDraw, "green", maxGreens, gameNumber)
			}
			if strings.Contains(cubeDraw, "blue") {
                gameNumber = checkColorDraw(cubeDraw, "blue", maxBlues, gameNumber)
			}
		}
	}
	return gameNumber
}

func checkColorDraw(cubeDraw, color string, maxAmount, gameNumber int) int {
	cubes := strings.Split(cubeDraw, " ")
	amount, err := strconv.Atoi(cubes[0])
	if err == nil && amount > maxAmount {
		return 0
	}
	return gameNumber
}
