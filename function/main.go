package main

import "fmt"

func phepToan(num1, num2 int) {
	//Trường hợp truyền num2 = 0
	if num2 == 0 {
		num2 = 1
	}

	sum := num1 + num2
	diff := num1 - num2
	prod := num1 * num2
	quot := float32(num1) / float32(num2)
	mod := num1 % num2

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot)
	fmt.Println("Modulus:", mod)
}

func congHaiSoNguyen(a, b int) int {
	return a + b
}

func phepTru2So(a, b float64) (ketQua float64) {
	ketQua = a - b

	return ketQua
}

func hamTraVeNhieuGiaTri(a, b, c int) (int, int, int) {

	tong := a + b + c
	tich := a * b * c
	hieu := a - b - c

	//Có 2 cách để return nhiều giá trị
	return tong, tich, hieu
	// return a + b + c, a * b * c, a - b - c
}

func main() {
	// phepToan(10, 3)
	// phepToan(22, 0)

	// res := congHaiSoNguyen(3, 6)
	// fmt.Println("Result:", res)

	// ans := phepTru2So(3, 1.5)
	// fmt.Println("Answer:", ans)

	tong, tich, hieu := hamTraVeNhieuGiaTri(5, 2, 1)
	fmt.Println("Tong:", tong)
	fmt.Println("Tich:", tich)
	fmt.Println("Hieu:", hieu)
}
