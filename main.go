// This package is part of Jon Bodner's *Learning Go* chapter 10 exercise
// The purpose is to expose an Add method that adds together two integers.
package learningGoEx10

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// a: integer
// b: integer
// returns the sum of a and b
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a,b T) T{
	return a + b
}