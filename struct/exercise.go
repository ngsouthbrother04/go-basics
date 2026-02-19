package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ChuNhat struct {
	ChieuDai  float64
	ChieuRong float64
}

func (cn *ChuNhat) tinhDienTich() float64 {
	return cn.ChieuDai * cn.ChieuRong
}

func (cn *ChuNhat) tinhChuVi() float64 {
	return 2 * (cn.ChieuDai + cn.ChieuRong)
}

func readInput(msg string) (float64, error) {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input) //Remove whitespace and newline characters

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}
