package main

import "fmt"

func basicAboutFmt() {
	//==============Printf & Println & Print=================
	// var name string = "NA"
	// var age int = 22

	// fmt.Printf("Hello %s, and your age is %d. \n", name, age)

	// fmt.Print(name)
	// fmt.Print(age)
	// fmt.Println()

	// fmt.Println(name)
	// fmt.Println(age)

	//==============Scanf & Scanln & Scan=================
	/*
		Scan:
			- Đọc các giá trị phân tách bởi whitespace (space, tab, newline (Enter)).
			- Dừng khi đã đọc đủ tham số (số biến truyền vào) hoặc gặp EOF/error.
			- Newline được xem là whitespace; không yêu cầu kết thúc bằng newline.

		Scanln:
			- Tương tự Scan nhưng dừng tại newline (Enter).
			- Sau tham số cuối phải có newline hoặc EOF; nếu còn token trên cùng dòng, Scanln trả về lỗi.

		Scanf:
			- Đọc theo chuỗi định dạng (format string) giống fmt.Printf nhưng theo chiều ngược lại.
			- Nếu input không khớp định dạng, Scanf trả về lỗi và số giá trị đã đọc.
			- Dùng khi cần kiểm soát định dạng nhập (ví dụ "%d %s").
	*/

	// var (
	// 	name string
	// 	age  int
	// )
	// fmt.Print("Enter your name and age: ")

	// fmt.Scan(&name, &age)
	// fmt.Printf("Your name is %s, and your age is %d. \n", name, age)

	// fmt.Scanln(&name, &age)
	// fmt.Printf("Your name is %s, and your age is %d. \n", name, age)

	// fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Print("Enter your age: ")
	// fmt.Scanf("%d", &age)
	// fmt.Printf("Your name is %s, and your age is %d. \n", name, age)

	//==============Sprintf & Sprintln & Sprint=================
	/*
		Sprint:
			- Tương tự fmt.Print nhưng trả về chuỗi thay vì in ra màn hình.
			- Tham số được chuyển đổi thành chuỗi và nối lại với nhau.

		Sprintln:
			- Tương tự Sprint
			- Tham số được chuyển đổi thành chuỗi, nối lại với nhau và thêm newline ở cuối.

		Sprintf:
			- Tương tự Sprint
			- Sử dụng chuỗi định dạng để định dạng các tham số.
	*/

	// msg := fmt.Sprint("My name is ", "NA")
	// fmt.Print(msg)
	// fmt.Print("---------------- \n")

	// msg := fmt.Sprintln("Hello ", "NA")
	// fmt.Print(msg)
	// fmt.Print("--------------- \n")

	// city := "HCM"
	// country := "Vietnam"
	// time := 4
	// msg := fmt.Sprintf("Live in %s, %s for %d years. \n", city, country, time)
	// fmt.Print(msg)
	// fmt.Print("------------- \n")
}

func formattingVerbs() {
	ten := "NA"
	tuoi := 22
	chieuCao := 1.6632213
	daTotNghiep := false
	phanTram := 10

	//%T - type: kiểu dữ liệu
	fmt.Printf("ten: %T \n", ten)
	fmt.Printf("tuoi: %T \n", tuoi)
	fmt.Printf("chieuCao: %T \n", chieuCao)
	fmt.Printf("daTotNghiep: %T \n", daTotNghiep)
	fmt.Printf("phanTram: %T \n", phanTram)
	fmt.Print("------------- \n")

	//%v - value: giá trị của biến
	fmt.Printf("ten: %v \n", ten)
	fmt.Printf("tuoi: %v \n", tuoi)
	fmt.Printf("chieuCao: %v \n", chieuCao)
	fmt.Printf("daTotNghiep: %v \n", daTotNghiep)
	fmt.Printf("phanTram: %v \n", phanTram)
	fmt.Print("------------- \n")

	//%#v - value: giá trị của biến theo định dạng Go-syntax
	fmt.Printf("ten: %#v \n", ten)
	fmt.Printf("tuoi: %#v \n", tuoi)
	fmt.Printf("chieuCao: %#v \n", chieuCao)
	fmt.Printf("daTotNghiep: %#v \n", daTotNghiep)
	fmt.Printf("phanTram: %#v \n", phanTram)
	fmt.Print("------------- \n")

	//%d - decimal: định dạng số thập phân (chỉ áp dụng cho kiểu số nguyên)
	fmt.Printf("tuoi: %d \n", tuoi)
	fmt.Printf("phanTram: %d \n", phanTram)
	fmt.Print("------------- \n")

	//%.f - float: định dạng số dấu chấm động
	fmt.Printf("chieuCao: %.f \n", chieuCao)
	fmt.Printf("chieuCao: %.3f \n", chieuCao)
	fmt.Print("------------- \n")

	//%t - true/false: định dạng boolean
	fmt.Printf("daTotNghiep: %t \n", daTotNghiep)
	fmt.Print("------------- \n")

	//%% - ký tự phần trăm
	fmt.Printf("phanTram: %d%% \n", phanTram)
	fmt.Print("------------- \n")

	//%c - character (rune): định dạng ký tự từ mã Unicode (chỉ áp dụng cho kiểu số nguyên)
	fmt.Printf("Mã Unicode của ký tự 'A' là: %c \n", 65)
	fmt.Printf("Mã Unicode của ký tự 'a' là: %c \n", 97)
	fmt.Print("------------- \n")

	//%p - pointer: định dạng địa chỉ bộ nhớ
	fmt.Printf("Địa chỉ bộ nhớ của biến ten: %p \n", &ten)
	fmt.Printf("Địa chỉ bộ nhớ của biến tuoi: %p \n", &tuoi)
	fmt.Print("------------- \n")
}

func main() {
	// basicAboutFmt()
	formattingVerbs()
}
