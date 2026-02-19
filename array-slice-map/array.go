package main

import "fmt"

func lyThuyetArray() {
	//Khởi tạo chỉ định kích thước
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println(nums)

	//Khởi tạo không chỉ định kích thước, tuy nhiên không thể tăng hoặc giảm kích thước mảng sau khi khai báo
	var numbers = [...]int{4, 5, 6, 7, 8}
	fmt.Println(numbers)
}

func mang2Chieu() {
	//Khai báo mảng 2 chiều
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
}

func forRangeArray() {
	nums := [6]int{10, 20, 30, 40, 50, 60}

	//Duyệt mảng sử dụng for range (2 biến trả về: index và value)
	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	fmt.Println()

	//Bỏ qua giá trị index hoặc value bằng cách sử dụng dấu gạch dưới _
	fmt.Println("Bỏ qua index:")
	for _, v := range nums {
		fmt.Printf("Value: %d\n", v)
	}
	fmt.Println()

	//Có thể dùng để duyệt qua 1 map
	fmt.Println("\nDuyệt map:")
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println()

	//Chỉ lấy index hoặc key
	fmt.Println("Chỉ lấy index:")
	for i := range nums {
		fmt.Printf("Index: %d\n", i)
	}
	fmt.Println("\nChỉ lấy key:")
	for k := range m {
		fmt.Printf("Key: %s\n", k)
	}
}

func arrayAndStruct() {
	type Employee struct {
		ID   int
		Name string
		Age  int
	}

	employees := [...]Employee{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
		{ID: 4, Name: "Diana", Age: 28},
	}

	for _, e := range employees {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", e.ID, e.Name, e.Age)
	}
}
