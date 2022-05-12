package string_sum

// package main

import (
	"errors"
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
	output := strings.ReplaceAll(input, " ", "")
	output = strings.ReplaceAll(input, "+", " +")
	output = strings.ReplaceAll(input, "-", " -")

	firstElementIsNegative := false

	if strings.HasPrefix(output, " -") {
		firstElementIsNegative = true
		output = strings.TrimLeft(output, " -")
	}
	if strings.HasPrefix(output, " +") {
		output = strings.TrimLeft(output, " +")
	}

	temp := strings.Split(output, " ")
	var sum int64 = 0
	if len(temp) > 1 {
		for index, num := range temp {
			temp_int, _ := strconv.Atoi(num)
			if index == 0 && firstElementIsNegative {
				sum -= int64(temp_int)
			} else {
				sum += int64(temp_int)
			}
		}
	}

	// fmt.Println("temp sum: ", sum)
	// fmt.Println(temp)
	return strconv.FormatInt(sum, 10), nil
}

// func main() {
// 	m := "+1-2 -3 -5"
// 	var a string
// 	fmt.Println(m)
// 	a, _ = StringSum(m)
// 	fmt.Println(a)
// }
