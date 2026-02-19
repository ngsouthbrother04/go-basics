package services

import "fmt"

// Animal kế thừa (embed) cả Speaker và Eater
// Type thỏa mãn Animal phải có tất cả methods của Speaker + methods của Eater
type Animal interface {
	Speaker
	Eater
	Sleep() bool
}

func Daily(a Animal) string {
	return a.Speak() + " -> after speak -> " + a.Eat() + " -> after eat -> " + fmt.Sprintf("%v", a.Sleep())
}
