package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func day01First(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	firstNumber := 0
	secondNumber := 0
	runningCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		for _, character := range line {
			if number := int(character); 47 < number && number < 58 {
				firstNumber = number - 48
				break
			}
		}

		for i := len(line); i > 0; i-- {
			if number := int(line[i-1]); 47 < number && number < 58 {
				secondNumber = number - 48
				break
			}
		}
		runningCount += 10*firstNumber + secondNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return runningCount
}

func stringToDigit(line string, index int) (int, error) {
	switch line[index] {
	case 'o':
		if len(line) >= index+3 && line[index:index+3] == "one" {
			return 1, nil
		}
	case 't':
		if len(line) >= index+3 && line[index:index+3] == "two" {
			return 2, nil
		} else if len(line) >= index+5 && line[index:index+5] == "three" {
			return 3, nil
		}
	case 'f':
		if len(line) >= index+4 && line[index:index+4] == "four" {
			return 4, nil
		} else if len(line) >= index+4 && line[index:index+4] == "five" {
			return 5, nil
		}
	case 's':
		if len(line) >= index+3 && line[index:index+3] == "six" {
			return 6, nil
		} else if len(line) >= index+5 && line[index:index+5] == "seven" {
			return 7, nil
		}
	case 'e':
		if len(line) >= index+5 && line[index:index+5] == "eight" {
			return 8, nil
		}
	case 'n':
		if len(line) >= index+4 && line[index:index+4] == "nine" {
			return 9, nil
		}
	default:
		return 0, errors.New("No digit")
	}
	return 0, errors.New("No digit at the end")
}

func day01Second(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	firstNumber := 0
	secondNumber := 0
	runningCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		for index, character := range line {
			if number := int(character); 47 < number && number < 58 {
				firstNumber = number - 48
				break
			} else if number, err := stringToDigit(line, index); err == nil {
				firstNumber = number
				break
			}
		}

		for i := len(line); i > 0; i-- {
			if number := int(line[i-1]); 47 < number && number < 58 {
				secondNumber = number - 48
				break
			} else if number, err := stringToDigit(line, i-1); err == nil {
				secondNumber = number
				break
			}
		}
		runningCount += 10*firstNumber + secondNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return runningCount
}

func main() {
	ret := day01First("input1.txt")
	fmt.Printf("The Answer for 1 is: '%d'\n", ret)
	ret = day01Second("input1.txt")
	fmt.Printf("The Answer for 2 is: '%d'\n", ret)
}
