package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Box[T any] struct {
	Content T
}

type Number interface {
	constraints.Integer | constraints.Float
}

func sum[T Number](a, b T) T {
	return a + b
}

func printValue[T any](value T) {
	fmt.Println(value)
}

func compare[T comparable](a, b T) bool {
	return a == b
}

func findMax[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func longestString[T string](a, b T) T {
	if len(a) > len(b) {
		return a
	}
	return b
}

func main() {
	// fmt.Println(compare(3, 2))
	// fmt.Println(compare("hello", "hello"))

	// fmt.Println(findMax(3.3, 2))

	// fmt.Println(longestString("short", "longerString"))

	// intBox := Box[int]{
	// 	Content: 42,
	// }

	// stringBox := Box[string]{
	// 	Content: "Generics in Go",
	// }

	// printValue(intBox.Content)
	// printValue(stringBox.Content)

	printValue(sum(5.5, 23))
}
