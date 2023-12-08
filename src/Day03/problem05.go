package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Problem05() int {
	buffer, err := os.Open("src/Day03/problem05.csv")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanLines)

	prevLine := ""
	scanner.Scan()
	current := scanner.Text()
	res := 0
	for scanner.Scan() {
		nextLine := scanner.Text()
		for i := 0; i < len(current); i++ {
			res += checkCharAndGetNumber(prevLine, current, nextLine, i)
		}
		prevLine = current
		current = nextLine
	}

	return res
}

func checkCharAndGetNumber(prev, curr, next string, index int) int {
	char := string(curr[index])
	nums := make([]int, 0)
	if regexp.MustCompile(`[^\d^.^\n]`).MatchString(char) {
		for i := index; i <= index+1; i++ {
			if i < 0 || i >= len(curr) {
				continue
			}
			nums = append(nums, getFullNumber(prev, i), getFullNumber(curr, i), getFullNumber(next, i))
		}
	}

	res := 0
	for i, num := range nums {
		newNums := nums[i+1:]
		if !slices.Contains(newNums, num) {
			res += num
		}
	}
	return res
}

func getFullNumber(str string, index int) int {
	before := strings.Split(str[:index], "")
	after := strings.Split(str[index:], "")

	tempNum := ""
	re := regexp.MustCompile(`\d`)
	for i := len(before) - 1; i >= 0; i-- {
		char := before[i]
		if !re.MatchString(char) {
			break
		}
		tempNum = char + tempNum
	}

	for i := 0; i < len(after); i++ {
		char := after[i]
		if !re.MatchString(char) {
			break
		}
		tempNum += char
	}
	num, err := strconv.Atoi(tempNum)
	if err != nil {
		return 0
	}
	return num
}
