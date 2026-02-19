**Ghi chú: Hàm (Functions) trong Go — tóm tắt từ `function/main.go`**

1) Định nghĩa hàm cơ bản
- Cú pháp: `func name(params) returnType { ... }`.
- Ví dụ:

```go
func congHaiSoNguyen(a, b int) int {
	return a + b
}
```

2) Nhiều giá trị trả về
- Go cho phép return nhiều giá trị cùng lúc (thường dùng cho `value, err`):

```go
func hamTraVeNhieuGiaTri(a, b, c int) (int, int, int) {
	return a+b+c, a*b*c, a-b-c
}
```

3) Named return (tên biến trả về)
- Bạn có thể đặt tên cho giá trị trả về và gán trực tiếp trong thân hàm, sau đó `return` không cần liệt kê biến:

```go
func phepTru2So(a, b float64) (ketQua float64) {
	ketQua = a - b

  //Có thể sử dụng 1 trong 2 cách
	return // trả về ketQua
  //return ketQua
}
```

4) Xử lý input bất thường (ví dụ chia cho 0)
- Kiểm tra và xử lý sớm trong hàm, ví dụ `if num2 == 0 { num2 = 1 }` trong `phepToan` để tránh division-by-zero.


5) Gợi ý thêm
- Tránh side-effects bất ngờ; prefer return values over mutating globals.
- Dùng multiple returns để trả `value, err` và xử lý lỗi rõ ràng.
- Khi cần hàm flexible, xem `variadic` functions: `func sum(nums ...int) int {}`.

---
