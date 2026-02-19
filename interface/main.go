package main

import (
	"fmt"

	"nnama.com/interface/mouse"
	"nnama.com/interface/services"
)

// Minh hoạ empty interface (interface{} = any)
func PrintAny(a interface{}) {
	//Type assertion
	// str, ok := a.(string)

	// if ok {
	// 	fmt.Printf("String: %s\n", str)
	// } else {
	// 	panic("Please send a string")
	// }

	// Type switch
	switch a.(type) {
	case string:
		fmt.Printf("String: %s\n", a)
	case int:
		fmt.Printf("Int: %d\n", a)
	default:
		panic("Please send a string or an int")
	}
}

func main() {
	m, err := mouse.New("Mickey")
	if err != nil {
		panic(err)
	}

	fmt.Println(services.Daily(m))
	PrintAny("NNA")
	PrintAny(3)
}
