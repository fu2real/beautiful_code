package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	input = strings.ReplaceAll(input, " ", "")
	operArr := []int{}
	symArr := []byte{}

	for i, v := range input {
		x, err := strconv.Atoi(string(v))
		if err == nil {
			if i-1 >= 0 {
				symArr = append(symArr, input[i-1])
			} else {
				symArr = append(symArr, 0)
			}
			operArr = append(operArr, x)
		}
	}

	if len(operArr) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	sym1 := string(symArr[0])
	sym2 := string(symArr[1])
	oper1 := operArr[0]
	oper2 := operArr[1]
	if sym1 == "-" {
		oper1 = -oper1
	}
	if sym2 == "-" {
		oper2 = -oper2
	}

	answer := oper1 + oper2
	return strconv.Itoa(answer), nil
}
