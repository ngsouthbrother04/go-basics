**Ghi chú về Module, Package và chương trình trong Go**

**Khái niệm chính**

- **Module**: Một module được khai báo bằng file `go.mod`. Module là một tập hợp các package và xác định gốc của đường dẫn import (module path). Một repository có thể chứa một hoặc nhiều module.

- **Package**: Mỗi package tương đương với một thư mục chứa các file `.go` (ngoại trừ trường hợp đặc biệt với test). Tất cả file `.go` (không bao gồm `_test.go`) trong cùng một thư mục phải khai báo cùng một tên package.

- **File Go**: Một package có thể gồm nhiều file `.go`. Khi biên dịch, tất cả file trong cùng package sẽ được biên dịch chung thành một đơn vị.

**Package `main` và hàm `main`**

- **Package `main`**: Nếu một package có tên là `main`, khi build nó sẽ sinh ra một chương trình thực thi (executable/binary).

- **Hàm `main`**: Để tạo chương trình có thể chạy được, package `main` phải định nghĩa một hàm `func main()`. Thông thường chỉ nên có một định nghĩa `main()` trong package `main` (nhiều định nghĩa sẽ gây lỗi khi build).

- **Lưu ý**: Các package khác (không phải `main`) là thư viện và không sinh ra executable.

**Quy tắc và thực hành thường gặp**

- **Một package — một thư mục**: Quy tắc phổ biến là một thư mục = một package. File test (`_test.go`) có thể khai báo package khác (ví dụ `pkg_test`) để viết test độc lập.

- **Export**: Identifier (hàm, biến, type, const) bắt đầu bằng chữ in hoa là được export (công khai), bắt đầu bằng chữ thường là private trong package.

- **`go.mod`**: File `go.mod` nằm ở gốc module, chứa tên module và phiên bản Go dùng. Ví dụ:

```go
module example.com/mymodule

go 1.20
```

**Ví dụ đơn giản**
- File `main.go` trong package `main`:

```go
package main

import "fmt"
````markdown
**Ghi chú về Module, Package và chương trình trong Go**

**Khái niệm chính**

- **Module**: Một module được khai báo bằng file `go.mod`. Module là một tập hợp các package và xác định gốc của đường dẫn import (module path). Một repository có thể chứa một hoặc nhiều module.

- **Package**: Mỗi package tương đương với một thư mục chứa các file `.go` (ngoại trừ trường hợp đặc biệt với test). Tất cả file `.go` (không bao gồm `_test.go`) trong cùng một thư mục phải khai báo cùng một tên package.

- **File Go**: Một package có thể gồm nhiều file `.go`. Khi biên dịch, tất cả file trong cùng package sẽ được biên dịch chung thành một đơn vị.

**Package `main` và hàm `main`**

- **Package `main`**: Nếu một package có tên là `main`, khi build nó sẽ sinh ra một chương trình thực thi (executable/binary).

- **Hàm `main`**: Để tạo chương trình có thể chạy được, package `main` phải định nghĩa một hàm `func main()`. Thông thường chỉ nên có một định nghĩa `main()` trong package `main` (nhiều định nghĩa sẽ gây lỗi khi build).

- **Lưu ý**: Các package khác (không phải `main`) là thư viện và không sinh ra executable.

**Quy tắc và thực hành thường gặp**

- **Một package — một thư mục**: Quy tắc phổ biến là một thư mục = một package. File test (`_test.go`) có thể khai báo package khác (ví dụ `pkg_test`) để viết test độc lập.

- **Export**: Identifier (hàm, biến, type, const) bắt đầu bằng chữ in hoa là được export (công khai), bắt đầu bằng chữ thường là private trong package.

- **`go.mod`**: File `go.mod` nằm ở gốc module, chứa tên module và phiên bản Go dùng. Ví dụ:

```go
module example.com/mymodule

go 1.20
```

**Ví dụ đơn giản**
- File `main.go` trong package `main`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
}
```

**Tóm tắt ngắn**
- Module chứa nhiều package; package tương ứng một thư mục.
- Package có thể có nhiều file `.go` — tất cả được biên dịch chung.
- Package `main` + một `func main()` tạo thành chương trình thực thi.

Nếu muốn, tôi có thể thêm ví dụ về nhiều package trong cùng một module hoặc giải thích `go mod`/`go build`/`go run` chi tiết hơn.

Tốt. Ta **bắt đầu PHASE 1 – GO CORE & MINDSET** ngay bây giờ, theo đúng giáo trình và style big tech.
Mục tiêu của buổi này: **đặt nền tư duy Go + viết được chương trình Go đầu tiên, hiểu rõ cấu trúc và zero-value**.

---

## 4. Biến & Zero Value (KHÁC Java / JS)

### Trích từ Chapter 2–3: *Types* & *Variables* 

**Ghi chú về Module, Package và chương trình trong Go**

---

1. Khái niệm cơ bản
- Module: khai báo bằng `go.mod`. Module gom nhiều package và xác định module path (gốc import). Một repo có thể chứa nhiều module.
- Package: tương ứng một thư mục chứa các file `.go`. Tất cả file `.go` (không bao gồm `_test.go`) trong cùng thư mục phải khai báo cùng tên package.
- File Go: Một package có thể có nhiều file `.go`; khi build chúng được biên dịch chung.

2. Package `main` và chương trình thực thi
- Package `main` → khi build sẽ tạo executable.
- Để chương trình chạy được phải có `func main()` trong package `main`.
- Các package khác (không phải `main`) là thư viện (library), không tạo executable.

---

4. Biến & Zero Value (khác Java/JS)

Trích từ Chapter 2–3: Types & Variables

Khai báo biến (mẫu):

```go
var a int
var b string
var c bool
```

Nếu không gán giá trị ban đầu, biến vẫn có giá trị mặc định (zero value).

Bảng zero value (cần nhớ):

| Type    | Zero value |
| ------- | ---------- |
| int     | 0          |
| float   | 0.0        |
| bool    | false      |
| string  | ""        |
| pointer | nil        |
| slice   | nil        |
| map     | nil        |
| struct  | mọi field bằng zero value |

Lưu ý: Go không có `null` như Java/JS; có `nil` cho các kiểu tham chiếu.

Short declaration (chỉ dùng trong function):

```go
x := 10
s := "go"
```

⚠️ Không dùng `:=` ở package scope (ngoại trừ trong body của function).

---

Hằng số (Constants)

- Khai báo hằng số dùng từ khóa `const`.

```go
const Pi = 3.14159
const Greeting string = "hello"
```

- Hằng số trong Go phải là giá trị biên dịch thời gian (compile-time constant). Không thể gán kết quả của hàm runtime.
- Có hằng số kiểu (typed) và không kiểu (untyped). Untyped constants linh hoạt khi kết hợp với các biểu thức.

Grouped constants và `iota` (thường dùng để tạo enum):

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)

const (
    KB = 1 << (10 * iota) // 1<<0, 1<<10, 1<<20 ...
    MB
    GB
)
```

- Hằng số hữu ích để đặt tên giá trị cố định (config, flags, enum) và giúp đọc code rõ ràng hơn.

---

5. So sánh nhanh với Java / JS (để tránh lỗi tư duy)

| Java / JS       | Go                  |
| --------------- | ------------------- |
| null everywhere | zero-value          |
| exception       | error return        |
| class           | struct              |
| inheritance     | composition         |
| async/await     | goroutine + channel |

---

### Shadowing (Ghi đè biến)

Shadowing xảy ra khi một biến khai báo trong scope nội bộ (ví dụ trong `if`, `for` hoặc một block) có cùng tên với biến ở scope ngoài; biến nội bộ sẽ "che khuất" (shadow) biến ngoài trong phạm vi của nó.

Ví dụ:

```go
package main

import "fmt"

func main() {
    x := 1
    if true {
        x := 2 // shadow biến x bên ngoài
        fmt.Println("inner x:", x) // 2
    }
    fmt.Println("outer x:", x) // 1
}
```

Một bẫy phổ biến là dùng `:=` vô tình tạo biến mới thay vì gán:

```go
func main() {
    n := 0
    if true {
        n := n + 1 // tạo biến mới shadow, dùng giá trị n bên ngoài để tính
        fmt.Println(n) // 1 (biến mới)
    }
    fmt.Println(n) // 0 (biến ngoài không thay đổi)
}
```

Hướng dẫn:
- Dùng phép gán `=` khi muốn cập nhật biến ngoài.
- Đổi tên biến nếu cần để tránh nhầm lẫn.
- Dùng `golangci-lint` hoặc `go vet` (rule `shadow`) để phát hiện shadowing.

---
