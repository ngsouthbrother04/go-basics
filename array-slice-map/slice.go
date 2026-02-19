package main

import (
	"fmt"
	"reflect"
	"slices"
)

func printSection(title string) {
	fmt.Printf("\n========== %s ==========\n", title)
}

func printCase(label string, result any) {
	fmt.Printf("%-25s : %v\n", label, result)
}

func lyThuyetSlice() {
	//Khai báo slice
	var s []int
	fmt.Println("un init:", s, s == nil, len(s) == 0)

	arr := [5]int{10, 20, 30, 40, 50}
	slice := []int{100, 200, 300}

	isASlice := reflect.TypeOf(slice).Kind() == reflect.Slice
	isAnArray := reflect.TypeOf(arr).Kind() == reflect.Array
	fmt.Println("Array: ", reflect.TypeOf(arr).Kind())
	fmt.Println("Slice: ", reflect.TypeOf(slice).Kind())
	fmt.Println("isASlice:", isASlice)
	fmt.Println("isAnArray:", isAnArray)
}

func khoiTaoSliceTuArray() {
	//Cú pháp: slice := arr[startIndex:endIndex], trong đó endIndex không được bao gồm phần tử tại vị trí endIndex
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Type: ", reflect.TypeOf(slice).Kind())
}

func khoiTaoSliceBangMake() {
	// Cú pháp: make([]Type, length, capacity)
	/*
		- Type: kiểu dữ liệu của slice
		- length: độ dài ban đầu của slice
		- capacity: sức chứa tối đa của slice
	*/

	slice := make([]int, 3, 6)
	sl := []string{"A", "B", "C"}

	fmt.Println("Slice of int:", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	fmt.Println()

	fmt.Println("Slice of string:", sl)
	fmt.Println("Length:", len(sl))
	fmt.Println("Capacity:", cap(sl))

	fmt.Println()

	slice = append(slice, 1, 2)
	slice = append(slice, 3)
	fmt.Println("After append:", slice)
	fmt.Println("New Length:", len(slice))
	fmt.Println("New Capacity:", cap(slice))

	fmt.Println()

	// Thêm phần tử vượt quá capacity ban đầu -> tự động tăng gấp đôi capacity
	slice2 := []int{100, 200, 300}
	slice3 := []int{400, 500}
	slice = append(slice, slice2...) // Không thể append cùng lúc nhiều slice vào 1 slice
	slice = append(slice, slice3...)
	fmt.Println("After exceeding capacity:", slice)
	fmt.Println("Final Length:", len(slice))
	fmt.Println("Final Capacity:", cap(slice))

	fmt.Println()

	c := make([]int, len(slice))
	copy(c, slice)
	fmt.Println("After copy:", c)

}

func subSlice() {
	// Tạo subslice từ slice ban đầu
	slice := []int{10, 20, 30, 40, 50}
	subSlice := slice[1:4]
	subSlice2 := slice[:3] // Tương đương slice[0:3]
	subSlice3 := slice[2:] // Tương đương slice[2:len(slice)]
	fmt.Println("Original Slice:", slice)
	fmt.Println("Subslice:", subSlice)
	fmt.Println("Is subslice is a slice:", reflect.TypeOf(subSlice).Kind() == reflect.Slice)
	fmt.Println("Subslice2:", subSlice2)
	fmt.Println("Subslice3:", subSlice3)

	fmt.Println()

	//So sánh length và capacity giữa slice ban đầu và subslice
	fmt.Println("Original Slice Length:", len(slice), "Capacity:", cap(slice))
	fmt.Println("Subslice Length:", len(subSlice), "Capacity:", cap(subSlice))
	fmt.Println("Subslice2 Length:", len(subSlice2), "Capacity:", cap(subSlice2))
	fmt.Println("Subslice3 Length:", len(subSlice3), "Capacity:", cap(subSlice3))
}

func sliceUtilities() {
	/** ================== CLONE & COMPARE ================== **/
	printSection("CLONE & COMPARE")
	printCase("Clone", slices.Clone([]int{1, 2, 3}))
	printCase("Equal [1 2 3] vs [1 2]", slices.Equal([]int{1, 2, 3}, []int{1, 2}))

	/** ================== SEARCH ================== **/
	printSection("SEARCH")
	printCase("Index of 2 in [1 2 3 4 2]", slices.Index([]int{1, 2, 3, 4, 2}, 2))
	printCase("Contains 5 in [1 2 3]", slices.Contains([]int{1, 2, 3}, 5))

	/** ================== MODIFY ================== **/
	printSection("MODIFY")
	printCase("Insert 2 at pos 2 in [1 3]", slices.Insert([]int{1, 3}, 2, 2))
	printCase("Delete index [1,3) in [1 2 3 4]", slices.Delete([]int{1, 2, 3, 4}, 1, 3))

	s1 := []int{1, 2, 3}
	slices.Reverse(s1)
	printCase("Reverse [1 2 3]", s1)

	/** ================== ORDERING ================== **/
	printSection("ORDERING")

	s2 := []int{3, 1, 2}
	slices.Sort(s2)
	printCase("Sort asc [3 1 2]", s2)

	s3 := []int{3, 1, 2}
	slices.SortFunc(s3, func(a, b int) int {
		return a - b
	})
	printCase("SortFunc asc [3 1 2]", s3)

	/** ================== AGGREGATE ================== **/
	printSection("AGGREGATE")
	printCase("Max [1 3 2]", slices.Max([]int{1, 3, 3, 2}))
	printCase("Min [1 3 2]", slices.Min([]int{1, 3, 2, 1}))
}
