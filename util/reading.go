package util

import (
	"bufio"
	"os"
	"strconv"
)

func NumbersFromFile(path string) []int {
	fh, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	numbers := make([]int, 0)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}
