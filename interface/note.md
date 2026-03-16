# Interface trong Go – Ghi chú ôn tập ngắn gọn và có hệ thống

## Tư duy cốt lõi

- **Struct** định nghĩa **cái gì** (dữ liệu + trạng thái).
- **Interface** định nghĩa **cái gì làm được** (hành vi).

Go dùng interface để **ràng buộc hành vi**, không phải để mô tả toàn bộ object.

---

## 1. Interface là gì?

Interface là **tập hợp các khai báo phương thức** (method signatures), không chứa implementation.

```go
type Speaker interface {
    Speak() string
    GetName() string
}
```

Một type **thỏa mãn** interface nếu nó triển khai **đúng và đủ** các phương thức đó.  
Không cần từ khóa `implements` — Go dùng **triển khai ngầm định** (implicit implementation).

---

## 2. Triển khai ngầm định (Duck Typing)

```go
type Mouse struct {
    name string
}

func (m *Mouse) Speak() string   { return "chít chít" }
func (m *Mouse) GetName() string { return m.name }
```

→ `*Mouse` tự động thỏa mãn `Speaker` mà không cần khai báo gì thêm.

Kiểm tra tính thỏa mãn diễn ra **tại điểm sử dụng**, không phải tại điểm định nghĩa struct.

---

## 3. Kết hợp interface (Interface Embedding)

Go không có kế thừa, nhưng interface có thể nhúng các interface khác.

```go
type Eater interface {
    Eat() string
}

type Animal interface {
    Speaker     // nhúng Speaker
    Eater       // nhúng Eater
    Sleep() bool
}
```

Để thỏa mãn `Animal`, một type phải triển khai **tất cả** các phương thức của `Speaker`, `Eater` và `Sleep()`.

---

## 4. Bảng tóm tắt khả năng thỏa mãn interface

| Struct | Các phương thức đã triển khai                  | Interface thỏa mãn              |
|--------|------------------------------------------------|---------------------------------|
| Mouse  | Speak, GetName, Eat, Sleep                     | Speaker, Eater, Animal          |
| Dog    | Speak, GetName, Eat                            | Speaker, Eater                  |
| Cat    | Speak, GetName                                 | Speaker                         |

**Quy tắc**: Thiếu dù chỉ một phương thức → không thỏa mãn interface (lỗi biên dịch).

---

## 5. Đa hình (Polymorphism) trong Go

Đa hình đạt được khi hàm nhận tham số là interface:

```go
func DailyRoutine(a Animal) string {
    return a.GetName() + ": " + a.Speak() + " → " + a.Eat()
}
```

- `*Mouse` → hợp lệ (có đủ phương thức).
- `*Dog`   → lỗi biên dịch (thiếu `Sleep`).

→ Đa hình trong Go là **an toàn ở thời điểm biên dịch**.

---

## 6. Mẫu thiết kế: Constructor trả về interface

```go
func NewAnimal(name string) (Animal, error) {
    return &Mouse{name: name}, nil
}
```

### Lý do trả về interface thay vì concrete type

- Ẩn chi tiết triển khai (implementation hiding).
- Giới hạn hành vi mà caller có thể sử dụng.
- Dễ thay đổi implementation sau này mà không phá vỡ code bên ngoài.
- Tăng tính đóng gói (encapsulation).

Caller chỉ biết đối tượng là một `Animal`, không cần biết nó là `Mouse`, `Dog`, hay gì khác.

---

## 7. Quy tắc vàng: Accept Interfaces, Return Structs

- **Nhận vào** interface → hàm linh hoạt, dễ mở rộng.
- **Trả về** struct → caller có đầy đủ thông tin và phương thức.

**Ngoại lệ có chủ đích**: Trả về interface khi muốn **ẩn implementation** hoặc **kiểm soát API surface** (như constructor ở phần 6).

---

## 8. Type Assertion – Ép kiểu từ interface

```go
var a Animal = &Mouse{name: "Jerry"}

m, ok := a.(*Mouse)  // assertion về concrete type
if ok {
    fmt.Println(m.name)  // truy cập field riêng (nếu có)
}
```

Dùng type assertion khi thực sự cần hành vi đặc thù của concrete type.  
Tránh lạm dụng để giữ tính đa hình.

---

## 9. Empty Interface (`interface{}` hoặc `any`)

### 9.1 Đặc điểm

```go
type Any interface{}  // hoặc dùng alias built-in: any (từ Go 1.18)
```

Không yêu cầu phương thức nào → **mọi type trong Go đều thỏa mãn**.

```go
var x any
x = 42
x = "hello"
x = &Mouse{}
```

### 9.2 Ứng dụng phổ biến

- `fmt.Println(...any)`
- `json.Unmarshal` (dữ liệu động)
- `map[string]any` (cấu trúc JSON không cố định)
- Các thư viện/framework cũ (trước generics)

Thường xuất hiện ở **ranh giới hệ thống** (I/O, serialization, reflection).

### 9.3 Nhược điểm

- Mất **type safety**.
- Phải dùng type assertion hoặc type switch → dễ lỗi runtime.

```go
func BadAdd(a, b any) any {
    return a.(int) + b.(int)  // panic nếu sai kiểu
}
```

### 9.4 Type switch

```go
switch v := value.(type) {
case int:
    fmt.Println(v + 1)
case string:
    fmt.Println(len(v))
default:
    fmt.Println("kiểu không hỗ trợ")
}
```

### 9.5 So sánh

| Tiêu chí               | `any` / `interface{}`        | Interface có phương thức     |
|-----------------------|------------------------------|------------------------------|
| Type safety           | Không                        | Có                           |
| Kiểm tra biên dịch    | Không                        | Có                           |
| Thể hiện ý nghĩa domain | Không                        | Có                           |
| Dùng trong business logic | Nên tránh                  | Nên dùng                     |

### 9.6 Giải pháp hiện đại (Go ≥ 1.18)

Ưu tiên **generics** thay vì `any`:

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b { return a }
    return b
}
```

→ Type-safe, không cần assertion.

---

## 10. Best Practices

1. Giữ interface **nhỏ** (1–3 phương thức) → dễ tái sử dụng.
2. Định nghĩa interface **tại nơi sử dụng** (consumer side), không phải nơi triển khai.
3. Tránh export interface quá sớm — dễ gây breaking change.
4. Chỉ trả về interface khi có lý do thiết kế rõ ràng (ẩn implementation).
5. Hạn chế tối đa `any` trong logic nghiệp vụ.
6. Dùng generics thay cho `any` khi có thể.

---

## 11. Các câu nhớ nhanh cho phỏng vấn / ôn thi

- Interface định nghĩa **hành vi**, không phải cấu trúc.
- Triển khai là **ngầm định** — “nếu có đủ method thì nó là interface đó”.
- Thiếu một method → lỗi biên dịch ngay.
- Trả về interface từ constructor → ẩn implementation, kiểm soát API.
- `any` chấp nhận mọi thứ, nhưng đánh đổi type safety.
- Package là đơn vị đóng gói thực sự trong Go.

> **"Empty interface accepts anything; meaningful interfaces accept only what they need."**
