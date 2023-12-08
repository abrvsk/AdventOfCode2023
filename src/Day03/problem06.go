package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func Problem06() int {
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
			res += checkAstericsAndGetNumber(prevLine, current, nextLine, i)
		}
		prevLine = current
		current = nextLine
	}

	return res
}

func checkAstericsAndGetNumber(prev, curr, next string, index int) int {
	char := string(curr[index])
	nums := make([]int, 0)
	if regexp.MustCompile(`\*`).MatchString(char) {
		for i := index; i <= index+1; i++ {
			if i < 0 || i >= len(curr) {
				continue
			}
			nums = append(nums, getFullNumber(prev, i), getFullNumber(curr, i), getFullNumber(next, i))
		}
	}

	uniqueNums := make([]int, 0)
	for _, num := range nums {
		if num > 0 && !slices.Contains(uniqueNums, num) {
			uniqueNums = append(uniqueNums, num)
		}
	}

	if len(uniqueNums) > 1 {
        res := 1
		for _, v := range uniqueNums {
			res *= v
		}
        return res
	}
	return 0
}

