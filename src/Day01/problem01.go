package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Problem01() int {
	buffer, err := os.Open("src/Day01/problem01.csv")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanLines)

	res := 0
	for scanner.Scan() {
		res += getFirstAndLastDigits(scanner.Text())
	}

	return res
}

func getFirstAndLastDigits(str string) int {
	intStr := regexp.MustCompile(`\d*`).FindAllString(str, -1)
	joined := strings.Join(intStr, "")
	first := joined[:1]
	last := joined[len(joined)-1:]

	if len(first)+len(last) == 0 {
		return 0
	}

	res, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
