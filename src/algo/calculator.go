package Algo

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func Evaluate(s string) float64{
	expression, err := govaluate.NewEvaluableExpression(s)
    if err != nil {
        fmt.Println("Error parsing expression:", err)
        return 0
    }

    result, err := expression.Evaluate(nil)
    if err != nil {
        fmt.Println("Error evaluating expression:", err)
        return 0
    }

	return result.(float64)
}