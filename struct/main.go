package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GiangVien struct {
	Name    string `json:"name"`
	Gender  int    `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (gv *GiangVien) printInfo() {
	fmt.Printf("Name field: %s\n", gv.Name)
	fmt.Printf("Gender field: %d\n", gv.Gender)
	fmt.Printf("Age field: %d\n", gv.Age)
	fmt.Printf("Address field: %s\n", gv.Address)
}

func (gv *GiangVien) clearInfo() {
	gv.Name = ""
	gv.Gender = 0
	gv.Age = 0
	gv.Address = ""
}

func lyThuyetStruct() {

	gv := GiangVien{
		Name:    "NNA",
		Gender:  1,
		Age:     22,
		Address: "HCMC",
	}

	//Not recommended way
	// gvWithoutKey := GiangVien{
	// 	"NNA - 2",
	// 	1,
	// 	22,
	// 	"HCMC - 2",
	// }

	gv.printInfo()
	gv.clearInfo()
	fmt.Println()
	gv.printInfo()
	gvInJson, err := json.Marshal(gv)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(gvInJson))
	}
}

func main() {
	// lyThuyetStruct()
	var err error
	var chieuDai, chieuRong float64

	chieuDai, err = readInput("Nhap chieu dai: ")
	if err != nil {
		fmt.Println("Loi:", err)
		return
	}

	chieuRong, err = readInput("Nhap chieu rong: ")
	if err != nil {
		fmt.Println("Loi:", err)
		return
	}

	cn := ChuNhat{
		ChieuDai:  chieuDai,
		ChieuRong: chieuRong,
	}

	fmt.Printf("Dien tich hinh chu nhat: %.2f\n", cn.tinhDienTich())
	fmt.Printf("Chu vi hinh chu nhat: %.2f\n", cn.tinhChuVi())
}
