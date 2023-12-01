package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("values.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	twoDigitNumbers := make([]int, 0)
	firstNumberIndex := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		twoDigitNumber := ""

		for i, char := range str {

			if isDigit(char) {
				if len(twoDigitNumber) == 0 {
					twoDigitNumber = string(char)
					firstNumberIndex = i
				}
				println("Primeiro: " + twoDigitNumber)
				break
			}
		}

		for i := len(str) - 1; i >= 0; i-- {

			if i <= firstNumberIndex {
				twoDigitNumber = twoDigitNumber + string(str[firstNumberIndex])
				println("Segundo: " + twoDigitNumber)

				digit, err := strconv.Atoi(twoDigitNumber)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("twoDigit int: %d \n", digit)
				twoDigitNumbers = append(twoDigitNumbers, digit)

				break
			}

			if isDigit(rune(str[i])) {
				twoDigitNumber = twoDigitNumber + string(str[i])
				println("Segundo: " + twoDigitNumber)

				digit, err := strconv.Atoi(twoDigitNumber)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("twoDigit int: %d \n", digit)
				twoDigitNumbers = append(twoDigitNumbers, digit)

				break
			}

		}

	}

	for _, digit := range twoDigitNumbers {
		sum = sum + digit
	}

	println(sum)

	readFile.Close()
}

// Helper function to check if a character is a digit
func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
