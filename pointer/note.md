# Ghi chú Cốt lõi: Pointer trong Golang

## Khái niệm cơ bản

- Biến thông thường: lưu trữ giá trị (ví dụ: số `10`, chuỗi `"Hello"`).
- Pointer (con trỏ): lưu trữ địa chỉ bộ nhớ của một biến khác ("số nhà").
- Mục đích: tiết kiệm bộ nhớ (không cần copy dữ liệu lớn) và cho phép hàm thay đổi giá trị của biến gốc.

## Hai ký tự cần nhớ

| Ký hiệu | Tên gọi | Ý nghĩa | Ví dụ |
| ------: | ------- | ------- | ----- |
| `&` | Toán tử lấy địa chỉ | "Hỏi số nhà" của một biến | `p = &a` |
| `*` | Toán tử giải tham chiếu | "Mở cửa nhà" để xem/sửa giá trị bên trong | `fmt.Print(*p)` hoặc `*p = 20` |

## Trạng thái `nil` (Nguy hiểm)

- Một con trỏ chưa trỏ đi đâu sẽ có giá trị là `nil`.
- Lưu ý: Thao tác `*p` trên một con trỏ `nil` sẽ gây sập chương trình (runtime panic).
- Kiểm tra an toàn:

```go
if p != nil {
	// dùng p an toàn
}
```

## Khi nào nên dùng Pointer?

- ✅ Khi cần sửa đổi dữ liệu gốc bên trong một hàm.
- ✅ Khi làm việc với struct lớn để tránh việc copy tốn tài nguyên.
- ❌ Không cần dùng con trỏ để quản lý các nhóm kiểu dữ liệu sau: **Reference Types** (`Slice`, `Map`, `Channel`, `Interface`, `Function`), **Basic value types** (`int`, `uint`, `float32`, `float64`, `bool`, `rune`, `byte`) và **String** (vì `string` trong Golang là _*immutable*_, nên việc truyền string vào 1 hàm là rất nhanh và an toàn)
