package main

import (
	"errors"
	"fmt"
	"strconv"
)

var operators = map[string]Arithmetic{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
	"%": modulo,
	"^": pow,
}

func perform(stack *Stack, fn Arithmetic) (float64, error) {
	val1, err := stack.Shift()

	if err != nil {
		return 0, errors.New("invalid RPN expression")
	}

	val2, err := stack.Shift()

	if err != nil {
		return 0, errors.New("invalid RPN expression")
	}

	return fn(val1, val2), nil
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
			result, err := perform(&stack, arithmetic)
			if err != nil {
				return 0, err
			}
			stack.Push(result)
			continue
		}

		if item == "." {
			subExpression += item
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
	result, err := RPN("5.3 4 * 5 + 10 -")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
