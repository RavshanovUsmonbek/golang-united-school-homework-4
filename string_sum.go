package string_sum

import (
	"errors"
	"fmt"
	"regexp"
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

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return "", fmt.Errorf("errorEmptyInput: %w", errorEmptyInput)
	}
	// cleaning input string from trailing spaces
	input = strings.TrimSpace(input)
	operands := findOperandsWithSplit(input)
	if checkForManyOperands(input, operands) {
		return "", fmt.Errorf("errorNotTwoOperands: %w", errorNotTwoOperands)
	}
	err = checkForNanOperands(input, operands)
	if err != nil {
		return "", err
	}

	output = calculate(input)
	return output, nil
}

func findOperands(input string) []string {
	re := regexp.MustCompile(`[+,-]?[0-9]+`)
	return re.FindAllString(input, -1)
}

func findOperandsWithSplit(input string) []string {
	pattern := regexp.MustCompile(`[+,-]{1}`)
	operands := pattern.Split(input, -1)
	return operands
}

func checkForManyOperands(input string, operands []string) bool {
	if len(getEffectiveOperands(input, operands)) == 2 {
		return false
	}
	return true
}

func getEffectiveOperands(input string, operands []string) []string {
	if strings.HasPrefix(input, "+") || strings.HasPrefix(input, "-") {
		return operands[1:]
	}
	return operands
}

func checkForNanOperands(input string, operands []string) error {
	operands = getEffectiveOperands(input, operands)
	for _, op := range operands {
		_, err := strconv.Atoi(op)
		if err != nil {
			return fmt.Errorf("NaNOperandError: %w", err)
		}
	}
	return nil
}

func calculate(input string) string {
	operands := findOperands(input)
	var sum int
	for _, op := range operands {
		n, _ := strconv.Atoi(op)
		sum += n
	}
	return strconv.Itoa(sum)
}
