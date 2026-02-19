package main

import "fmt"

func skipRandomNums(num int) bool {
	if num == 6 || num == 48 || num == 75 || num == 89 {
		return true
	}
	return false
}

func A1() {
	for i := 1; i <= 100; i++ {
		if skipRandomNums(i) {
			continue
		}
		if i != 100 {
			fmt.Printf("%#v, ", i)
		} else {
			fmt.Printf("%#v\n", i)
		}
	}
}

func A2() {
	cnt := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}

		if cnt == 3 {
			fmt.Println()
			cnt = 0
		}

		if cnt > 0 {
			fmt.Printf(", ")
		}

		fmt.Printf("%#v", i)
		cnt++
	}
}
