package main

import (
	"fmt"
)

func ifElse() {
	fmt.Println("=====================")
	diem := 2

	if diem > 8 {
		fmt.Println("Gioi")
	} else if diem == 5 {
		fmt.Println("Trung Binh")
	} else {
		fmt.Println("Khong phai Gioi")
	}

	//If with short statement. The scope of x is limited to the if block.
	if x := 10; x > 9 {
		fmt.Println("X lon hon 9")
	}
	fmt.Println("=====================")
}

func switchCase() {
	fmt.Println("=====================")
	diem := 8
	authenticated := false

	switch diem {
	case 9, 10:
		fmt.Println("Hoc sinh gioi")
	case 6, 7, 8:
		fmt.Println("Hoc sinh kha")
	case 5:
		fmt.Println("Hoc sinh trung binh")
	default:
		fmt.Println("Hoc sinh yeu")
	}

	switch {
	case authenticated:
		fmt.Println("Welcome back!")
	case !authenticated:
		fmt.Println("Please log in!")
	default:
		fmt.Println("Unknown authentication status.")
	}
	fmt.Println("=====================")
}

func forLoop() {
	fmt.Println("=====================")

	//Basic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Lan lap:", i)
	}

	//With break and continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue //skip even numbers
		}
		if i > 7 {
			break //stop the loop if i > 7
		}

		fmt.Println("So le nho hon hoac bang 7:", i)
	}

	fmt.Println("=====================")
}

func main() {
	// ifElse()
	// switchCase()
	// forLoop()
	A1()
	A2()
}
