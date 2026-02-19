package main

import "fmt"

func lyThuyetMap() {
	//? Cách 1: Khởi tạo map bằng cách sử dụng literal
	drink := map[string]string{
		"coke":  "Coca Cola",
		"pepsi": "Pepsi",
	}
	fmt.Println(drink)
	fmt.Println(drink["coke"])

	fmt.Println()

	//? Cách 2: Khởi tạo map bằng cách sử dụng make
	student := make(map[int]string)
	student[1] = "Alice"
	student[2] = "Bob"
	fmt.Println(student)

	fmt.Println()

	//? Cách 3: Khởi tạo map rỗng và sau đó phải cấp phát bộ nhớ bằng make trước khi thêm phẩn tử vào map
	var employee map[string]int
	employee = make(map[string]int)
	employee["Alice"] = 30
	employee["Bob"] = 25
	fmt.Println(employee)

	fmt.Println()

	//? Kiểm tra xem một key có tồn tại trong map hay không
	v, existed := employee["Alice"]
	if existed {
		fmt.Printf("Alice exists in employee map with value: %d\n", v)
	} else {
		fmt.Println("Alice does not exist in employee map")
	}

	fmt.Println()

	//? for range trong map
	for key, value := range drink {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	fmt.Println()

}

func mapKetHopStruct() {
	type Employee struct {
		Name string
		Age  int
		Role string
	}

	m := map[string]Employee{
		"e1": {"Alice", 30, "Developer"},
		"e2": {"Bob", 25, "Designer"},
	}

	for _, v := range m {
		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("Age: %d\n", v.Age)
		fmt.Printf("Role: %s\n", v.Role)
		fmt.Println()
	}
}

func mapKetHopSlice() {
	studentsSubjects := map[string][]string{
		"Tèo": {"Toán", "Lý", "Hóa"},
		"Tí":  {"Văn", "Sử", "Địa"},
		"Tủn": {"Anh", "Pháp", "Đức"},
	}

	for k, v := range studentsSubjects {
		fmt.Printf("Học sinh: %s\n", k)
		for _, s := range v {
			fmt.Printf("- Môn học: %s\n", s)
		}
		fmt.Println()
	}
}
