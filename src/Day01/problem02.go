package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Problem02() int {
	buffer, err := os.Open("src/Day01/problem01.csv")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanLines)


	res := 0
	for scanner.Scan() {
		current := getFirstAndLastDigit(scanner.Text())
		res += current
	}

	return res
}

func getFirstAndLastDigit(str string) int {
	numNames := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	res := []string{}
	for _, char := range strings.Split(str, "") {
		if _, err := strconv.Atoi(char); err == nil {
			res = append(res, char)
		}

		for i, num := range numNames {
			if strings.HasPrefix(str, num) {
				res = append(res, strconv.Itoa(i+1))
			}
		}

		str = str[1:]
	}

	first := res[:1][0]
	last := res[len(res)-1:][0]
	finalInt, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Println(err)
	}
	return finalInt
}
