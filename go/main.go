package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Arithmetic func(a float64, b float64) float64

func add(a, b float64) float64 {
	return a + b
}

var operators = map[string]Arithmetic{
	"+": add,
}

func RPN(exp string) (float64, error) {
	stack := Stack{}
	subExpression := ""

	for i := range exp {
		item := exp[i : i+1]

		if item == " " {
			if subExpression == "" {
				continue
			}
			num, err := strconv.ParseFloat(subExpression, 64)

			if err != nil {
				fmt.Println(err)
				break
			}

			stack.Push(num)
			subExpression = ""
			continue
		}

		arithmetic, ok := operators[item]

		if ok {
			val1, err := stack.Shift()

			if err != nil {
				return 0, errors.New("invalid RPN expression")
			}

			val2, err := stack.Shift()

			if err != nil {
				return 0, errors.New("invalid RPN expression")
			}

			result := arithmetic(val1, val2)
			stack.Push(result)
			continue
		}

		_, err := strconv.ParseFloat(item, 64)

		if err != nil {
			return 0, errors.New("invalid RPN expression: Invalid character")
		}

		subExpression += item
	}

	result, err := stack.Shift()

	if err != nil {
		return 0, err
	}

	return result, nil

}

func main() {

	result, err := RPN("5 4 + 13 + 104 + *")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
