package Algo

import (
    "strconv"
	"strings"
    "unicode"
    "math"
)

func applyOperation(a, b float64, op rune) float64 {
    switch op {
    case '+':
        return a + b
    case '-':
        return a - b
    case '*':
        return a * b
    case '/':
        return a / b
    case '^':
        return math.Pow(a, b)
    }
    return 0
}

func Calculate(expression string) (float64, error) {
	var numStack []float64
	var opStack []rune

	precedence := map[rune]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'^': 3,
	}

	for len(expression) > 0 {
		token := expression[0]
		expression = expression[1:]

		if token == '(' {
			opStack = append(opStack, rune(token))
		} else if token == ')' {
			for opStack[len(opStack)-1] != '(' {
				num2 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				num1 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				numStack = append(numStack, applyOperation(num1, num2, op))
			}
			opStack = opStack[:len(opStack)-1] // Remove the '(' from the stack
		} else if strings.ContainsRune("+-*/^", rune(token)) {
			for len(opStack) > 0 && precedence[opStack[len(opStack)-1]] >= precedence[rune(token)] {
				num2 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				num1 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				numStack = append(numStack, applyOperation(num1, num2, op))
			}
			opStack = append(opStack, rune(token))
		} else if unicode.IsDigit(rune(token)) {
			value, _ := strconv.Atoi(string(token))
			numStack = append(numStack, float64(value))
		}
	}

	for len(opStack) > 0 {
		num2 := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		num1 := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		numStack = append(numStack, applyOperation(num1, num2, op))
	}

	return numStack[0], nil
}