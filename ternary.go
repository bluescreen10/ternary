/*
Package ternary implements a ternary operator

# Usage

A ternary operator is a short version of an if-else statement that can make
the code more readable for some use cases such is variable initialization:

	import "github.com/bluescreen10/ternary"
	port := ternary.T( input > 0, input, 8080)

*/

package ternary

// T creates a ternary expression. The first argument is a bool called cond,
// the second is the value that gets returned if cond is true, and the third
// argument is the value that gets returned if cond is false.
func T[T any](cond bool, ifTrue, ifFalse T) T {
	switch cond {
	case true:
		return ifTrue
	default:
		return ifFalse
	}
}
