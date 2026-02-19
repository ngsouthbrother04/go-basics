package main

import "fmt"

func basicOperators() {
	diem := 5

	tong := diem + 2
	hieu := diem - 2
	tich := diem * 2
	thuong := float32(diem) / 2
	modulo := diem % 2

	fmt.Println("Điểm ban đầu:", diem)
	fmt.Println("Tổng:", tong)
	fmt.Println("Hiệu:", hieu)
	fmt.Println("Tích:", tich)
	fmt.Printf("Thương: %.2f \n", thuong)
	fmt.Println("Modulo:", modulo)

	diem++
	plusPlus := diem
	fmt.Printf("Điểm sau khi cộng 1: %v \n", plusPlus)

	diem--
	minusMinus := diem
	fmt.Printf("Điểm sau khi trừ 1: %v \n", minusMinus)
}

func comparingOperators() {
	s1 := 10
	s2 := 20

	phepSoSanhBang := s1 == s2
	phepSoSanhLonHon := s1 > s2
	phepSoSanhNhoHon := s1 < s2
	phepSoSanhLonHonHoacBang := s1 >= s2
	phepSoSanhNhoHonHoacBang := s1 <= s2
	phepSoSanhKhac := s1 != s2

	fmt.Printf("Kết quả phép so sánh s1 == s2: %#v \n", phepSoSanhBang)
	fmt.Printf("Kết quả phép so sánh s1 > s2: %#v \n", phepSoSanhLonHon)
	fmt.Printf("Kết quả phép so sánh s1 < s2: %#v \n", phepSoSanhNhoHon)
	fmt.Printf("Kết quả phép so sánh s1 >= s2: %#v \n", phepSoSanhLonHonHoacBang)
	fmt.Printf("Kết quả phép so sánh s1 <= s2: %#v \n", phepSoSanhNhoHonHoacBang)
	fmt.Printf("Kết quả phép so sánh s1 != s2: %#v \n", phepSoSanhKhac)
}

func logicalOperators() {
	a := false
	b := true
	c := false

	andOperator := a && b
	orOperator := a || b || c
	notOperatorA := !a
	notOperatorB := !b

	fmt.Printf("Kết quả phép toán AND (a && b): %#v \n", andOperator)
	fmt.Printf("Kết quả phép toán OR (a || b): %#v \n", orOperator)
	fmt.Printf("Kết quả phép toán NOT (!a): %#v \n", notOperatorA)
	fmt.Printf("Kết quả phép toán NOT (!b): %#v \n", notOperatorB)
}

func main() {
	// basicOperators()
	// comparingOperators()
	logicalOperators()
}
