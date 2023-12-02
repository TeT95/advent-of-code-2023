package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("values.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		convertSpelledNumbers(str, &sum)
	}

	println(sum)

	readFile.Close()
}

// Helper function to check if a character is a digit
func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func convertSpelledNumbers(str string, sum *int) {
	spelledNumbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	minPos := len(str)
	maxPos := 0

	var firstDigitAtLeft string
	var lastDigitAtRight string

	for k, v := range spelledNumbers {
		idx := strings.Index(str, k)

		lastIdx := strings.LastIndex(str, k)

		if idx != -1 && minPos > idx {
			minPos = idx
			firstDigitAtLeft = v
		}

		if lastIdx != -1 && maxPos < lastIdx {
			maxPos = lastIdx
			lastDigitAtRight = v
		}

	}

	beforeFistPos := str[:minPos]

	for _, char := range beforeFistPos {
		if isDigit(rune(char)) {
			firstDigitAtLeft = string(char)
			break
		}
	}

	afterLastPos := str[maxPos:]

	for _, char := range afterLastPos {
		if isDigit(rune(char)) {
			lastDigitAtRight = string(char)
		}
	}

	twoDigit := firstDigitAtLeft + lastDigitAtRight

	digit, err := strconv.Atoi(twoDigit)

	if err != nil {
		println(err)
	}

	*sum += digit
}
