
**Control Flow — if / switch / for (tóm tắt từ code)**

1) If / else

- Cú pháp cơ bản:

```go
if cond {
	// ...
} else if other {
	// ...
} else {
	// ...
}
```

- Short statement trong `if` (scope của biến giới hạn trong block):

```go
if x := 10; x > 9 {
	fmt.Println("x > 9")
}
// x không tồn tại ở đây
```

2) Switch -> Switch dùng thay cho việc dùng if để so sánh bằng

- `switch` cho nhiều `case` trên cùng một giá trị:

```go
switch score {
case 9, 10:
	// ...
case 6,7,8:
	// ...
default:
	// ...
}
```

- `switch` không có biểu thức dùng để xét các điều kiện boolean (tương tự chuỗi `if`):

```go
switch {
case authenticated:
	// ...
case !authenticated:
	// ...
}
```

3) For (vòng lặp)

- Ba dạng chính:
  - Kiểu `for` truyền thống: `for i := 0; i < n; i++ {}`
  - Kiểu `while`-style: `for i < n {}`
  - `for range` để duyệt slice, map, string.
- `break` dừng vòng lặp; `continue` bỏ qua phần còn lại của lần lặp hiện tại.

Ví dụ mẫu (từ `forLoop()`):

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}

for i := 1; i <= 10; i++ {
	if i%2 == 0 { continue }
	if i > 7 { break }
	fmt.Println(i)
}
```

4) Những lưu ý thực tế

- `++` và `--` là statements — không trả về giá trị; không dùng được trong biểu thức như `x := a++`.
- Phạm vi biến (scope) của short `if` chỉ trong block `if`.
- `switch` cho phép liệt kê nhiều giá trị trong một `case` (ví dụ `case 6,7,8:`).

---
