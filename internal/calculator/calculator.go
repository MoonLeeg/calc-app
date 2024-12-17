package calculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func Calc(expression string) (float64, error) {
	re := regexp.MustCompile(`(\d+\.?\d*|\+|\-|\*|\/)`)
	tokens := re.FindAllString(expression, -1)

	stack := make([]float64, 0)
	operators := make([]string, 0)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(operators) > 0 && getPrecedence(token) <= getPrecedence(operators[len(operators)-1]) {
				if len(stack) < 2 {
					return 0, errors.New("invalid expression")
				}

				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				op := operators[len(operators)-1]
				operators = operators[:len(operators)-1]

				switch op {
				case "+":
					stack = append(stack, a+b)
				case "-":
					stack = append(stack, a-b)
				case "*":
					stack = append(stack, a*b)
				case "/":
					if b == 0 {
						return 0, errors.New("division by zero")
					}
					stack = append(stack, a/b)
				}
			}
			operators = append(operators, token)
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid character: %s", token)
			}
			stack = append(stack, num)
		}
	}

	for len(operators) > 0 {
		if len(stack) < 2 {
			return 0, errors.New("invalid expression")
		}

		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		op := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		switch op {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			if b == 0 {
				return 0, errors.New("division by zero")
			}
			stack = append(stack, a/b)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

func getPrecedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}
