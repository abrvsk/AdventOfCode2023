package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Problem04() int {
	buffer, err := os.Open("src/Day02/problem03.csv")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanLines)

	res := 0
	for scanner.Scan() {
		res += calculatePowerOfCubes(scanner.Text())
	}

	return res
}

func calculatePowerOfCubes(str string) int {
	split := strings.Split(str, ": ")
	draws := strings.Split(split[1], "; ")

	minRed, minGreen, minBlue := 1, 1, 1
	for _, draw := range draws {
		cubes := strings.Split(draw, ", ")
		for _, cubeDraw := range cubes {
            minRed = checkColorDraw(cubeDraw, "red", minRed)
            minGreen = checkColorDraw(cubeDraw, "green", minGreen)
            minBlue = checkColorDraw(cubeDraw, "blue", minBlue)
		}
	}
	return minRed * minGreen * minBlue
}

func checkColorDraw(cubeDraw, color string, minAmount int) int {
	cubes := strings.Split(cubeDraw, " ")
	amount, err := strconv.Atoi(cubes[0])
	if err == nil && amount > minAmount {
		return amount
	}
    return minAmount
}
