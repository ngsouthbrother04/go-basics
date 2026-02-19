
# Ghi nhớ về `struct` trong Go

Đây là các kiến thức trọng tâm, ngắn gọn và thực dụng về `struct` trong Go.

## 1. Định nghĩa và literal

- Khai báo cơ bản:

```go
type Person struct {
	Name string
	Age  int
}

// Khởi tạo (composite literal)
p := Person{Name: "An", Age: 30}
```

- Không đề xuất dùng theo thứ tự không có key trừ khi rất ngắn: `Person{"An", 30}` — dễ bị nhầm khi có nhiều field.

## 2. Zero value & visibility

- Zero-value cho struct là mỗi field được khởi tạo về zero value tương ứng ("", 0, false, nil...).
- Field bắt đầu bằng chữ hoa => exported (có thể encode/json, truy cập từ package khác). Chữ thường => unexported.

## 3. Method receivers: value vs pointer

- Receiver giá trị (value): `func (p Person) String() string` — nhận một bản sao. Thích hợp khi method không sửa đổi struct.
- Receiver con trỏ (pointer): `func (p *Person) SetName(n string)` — sửa trực tiếp object, tránh copy cho struct lớn.
- Quy tắc: nếu một method cần pointer, nên dùng pointer cho tất cả method để tránh receiver set không nhất quán.

## 4. Copy behavior

- Gán một struct cho biến khác sẽ copy toàn bộ value (shallow copy). Nếu struct chứa slices, maps, con trỏ thì các reference bên trong vẫn trỏ cùng underlying data.

## 5. Embedded fields (composition)

```go
type Address struct { City string }
type Employee struct {
	Person
	Address
}

e := Employee{Person: Person{Name: "B"}, Address: Address{City: "HCM"}}
// Truy cập promoted fields: e.Name, e.City
```

- Embedding cho phép "promote" field/method — cách đơn giản thay thế kế thừa.

## 6. Tags và JSON

- Field tags dùng cho encoding/decoding, validation, ORM, v.v:

```go
type GiangVien struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}
```

- `omitempty` giúp loại bỏ field có zero value khi marshal.
- Khi marshal/unmarshal JSON, chỉ những field exported mới được xử lý.

## 7. Anonymous structs

- Dùng khi cần struct tạm thời, ví dụ trả về JSON nhanh:

```go
data := struct{ Msg string `json:"msg"` }{Msg: "Hello"}
```

## 8. So sánh struct

- Struct có thể so sánh bằng `==` nếu tất cả field đều comparable. Nếu chứa slice/map/function thì không comparable.

## 9. Reflection và tags

- Dùng `reflect` để đọc tag runtime:

```go
t := reflect.TypeOf(GiangVien{})
f, _ := t.FieldByName("Name")
fmt.Println(f.Tag.Get("json"))
```

## 10. Best practices

- Prefer explicit field names in literals.
- Sử dụng pointer receivers khi method sửa đổi đối tượng hoặc tránh copy lớn.
- Giữ struct nhỏ và chuyên biệt; dùng embedding/ composition thay vì kế thừa sâu.
- Dùng tags rõ ràng cho JSON/DB và đặt tên exported field theo convention.
- Khi cần constructor-like function, dùng pattern:

```go
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}
```

## 11. Ví dụ tổng hợp

```go
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (u *User) Anonymize() {
	u.Email = ""
}
```

-- Kết luận: hiểu rõ copy semantics, receiver types, exported vs unexported, và tags là then chốt khi làm việc với `struct` trong Go.

