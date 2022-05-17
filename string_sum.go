package string_sum

// package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func sliceCalcSum(temp []string, firstElementIsNegative bool) (int, error) {
	tempInt := 0
	var sum int = 0
	var err error = nil
	for index, num := range temp {
		tempInt, err = strconv.Atoi(num)
		if err == nil {
			if index == 0 && firstElementIsNegative {
				sum -= tempInt
			} else {
				sum += tempInt
			}
		} else {
			return 0, fmt.Errorf("provided not a number: %w", err)
		}
	}
	return sum, nil
}

func StringSum(input string) (string, error) {
	var err error = nil
	var sum int = 0
	if input == "" {
		return "", fmt.Errorf("empty input %w", errorEmptyInput)
	}
	output := strings.ReplaceAll(input, " ", "")
	output = strings.ReplaceAll(output, "+", " +")
	output = strings.ReplaceAll(output, "-", " -")

	firstElementIsNegative := false

	if strings.HasPrefix(output, " -") {
		firstElementIsNegative = true
		output = strings.TrimLeft(output, " -")
	}
	if strings.HasPrefix(output, " +") {
		output = strings.TrimLeft(output, " +")
	}

	temp := strings.Split(output, " ")

	if len(temp) == 2 {
		sum, err = sliceCalcSum(temp, firstElementIsNegative)
		if err == nil {
			return strconv.FormatInt(int64(sum), 10), err
		} else {
			return "", err
		}
		// }
	} else {
		return "", fmt.Errorf("something went wrong: %w", errorNotTwoOperands)
	}
}

// func main() {
// 	//	m := "+1-2 -3 -5"
// 	m := "-a-2"
// 	var a string
// 	fmt.Println(m)
// 	a, e := StringSum(m)
// 	fmt.Println("string", a)
// 	fmt.Println("error: ", e)
// }
