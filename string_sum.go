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

func StringSum(input string) (string, error) {
	var err error = nil
	var out string = ""
	if input != "" {
		output := strings.ReplaceAll(input, " ", "")
		output = strings.ReplaceAll(output, "+", " +")
		output = strings.ReplaceAll(output, "-", " -")

		firstElementIsNegative := false

		if strings.HasPrefix(output, " -") {
			firstElementIsNegative = true
			output = strings.TrimLeft(output, " -")
		} else if strings.HasPrefix(output, " +") {
			output = strings.TrimLeft(output, " +")
		}

		temp := strings.Split(output, " ")
		var sum int64 = 0

		if len(temp) == 2 {
			tempInt := 0
			for index, num := range temp {
				tempInt, err = strconv.Atoi(num)
				if err == nil {
					if index == 0 && firstElementIsNegative {
						sum -= int64(tempInt)
					} else {
						sum += int64(tempInt)
					}
				} else {
					return "", fmt.Errorf("provided not a number: %w", err)
				}
			}
			out = strconv.FormatInt(sum, 10)
		} else {
			err = fmt.Errorf("something went wrong: %w", errorNotTwoOperands)
		}
	} else {
		err = fmt.Errorf("empty input %w", errorEmptyInput)
	}

	// fmt.Println("temp sum: ", sum)
	// fmt.Println(temp)
	return out, err
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
