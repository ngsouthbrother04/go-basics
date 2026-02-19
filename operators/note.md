
**Ghi chú: Toán tử trong Go (dựa trên `operators/app.go`)**

1) Toán tử số học (Arithmetic)

- `+` (cộng), `-` (trừ), `*` (nhân), `/` (chia), `%` (modulo).
- Lưu ý về phép chia: chia giữa hai số nguyên cho kết quả là số nguyên (cắt phần thập phân). Nếu cần kết quả thực (float), phải ép kiểu trước khi chia, ví dụ: `float32(diem) / 2`.
- Ví dụ (tương tự `basicOperators()`):

```go
diem := 5
tong := diem + 2
thuong := float32(diem) / 2 // ép kiểu để có float
modulo := diem % 2
```

1) Toán tử tăng/giảm (`++` / `--`)

- Trong Go, `diem++` và `diem--` là statements — chúng không phải expressions và **không trả về giá trị**.
- Do đó **không thể** viết `x := diem++` hay dùng `diem++` trong một biểu thức. Phải tách ra thành hai bước nếu cần lấy giá trị trước/sau khi tăng:

```go
x := diem
diem++

// hoặc
diem++
x = diem
```

1) Toán tử so sánh (Comparison)

- Các toán tử: `==`, `!=`, `>`, `<`, `>=`, `<=`.
- Kết quả luôn là `bool`. Ví dụ từ `comparingOperators()`:

```go
s1 := 10
s2 := 20
fmt.Println(s1 == s2) // false
fmt.Println(s1 < s2)  // true
```

1) Toán tử logic (Logical)

- `&&` (AND), `||` (OR), `!` (NOT).
- Toán tử này áp dụng cho boolean. Ví dụ từ `logicalOperators()`:

```go
a := false
b := true
fmt.Println(a && b) // false
fmt.Println(a || b) // true
fmt.Println(!a)     // true
```

1) Kiểu dữ liệu và chuyển đổi (Type conversion)

- Go không tự động convert giữa các numeric types (ví dụ `int` ↔ `float32`) khi thực hiện phép toán; phải ép kiểu rõ ràng.
- Ví dụ: `float32(diem) / 2` để có phép chia chính xác với phần thập phân.

1) Gợi ý thực hành

- Khi in kết quả dùng `fmt.Printf` với verb phù hợp: `%.2f` cho float, `%d` cho int, `%v`/`%#v` cho debug.
- Tránh dùng `++`/`--` trong biểu thức; tách rõ bước tăng/giảm.
- Kiểm tra kiểu khi làm phép chia hoặc khi phối hợp nhiều kiểu numeric.

---

Ví dụ ứng dụng thực tế: Quyết định quyền truy cập trong Backend

Tóm tắt ý tưởng:

- `canRead` được tính bằng `Authenticated && !IsBanned`.
- `canWrite` cần `Authenticated && AccountActive && (Role=="admin" || Role=="editor") && !IsBanned`.
- `showBeta` = `globalFeatureFlag && (Role=="admin" || IsBetaTester)`.
- `cacheResponse` có thể là `!Authenticated || (Role=="guest" && !AccountActive)`, nhưng phải loại trừ `IsBanned`.

- Code minh hoạ:

```go
package main

import "fmt"

// Ví dụ thực tế (mini): quyết định quyền truy cập / hiển thị feature
// dựa trên các toán tử logic trong một backend request handler.

type User struct {
 ID            int
 Role          string // e.g., "admin", "editor", "user", "guest"
 Authenticated bool
 AccountActive bool
 IsBanned      bool
 IsBetaTester  bool
}

func evaluate(user User, globalFeatureBeta bool) {
 // Quyết định quyền đọc: phải authenticated và không bị banned
 canRead := user.Authenticated && !user.IsBanned

 // Quyết định quyền ghi: authenticated, account active, và role phù hợp
 canWrite := user.Authenticated && user.AccountActive && (user.Role == "admin" || user.Role == "editor") && !user.IsBanned

 // Hiển thị feature beta: bật global flag và người dùng là admin hoặc beta-tester
 showBeta := globalFeatureBeta && (user.Role == "admin" || user.IsBetaTester)

 // Quyết định cache: cache cho guest hoặc không-authenticated; không cache nếu banned
 cacheResponse := !user.Authenticated || (user.Role == "guest" && !user.AccountActive)
 cacheResponse = cacheResponse && !user.IsBanned

 // Gán giá trị quyền dạng string để dùng trong logging/response
 permission := "none"
 if canWrite {
  permission = "write"
 } else if canRead {
  permission = "read"
 }

 fmt.Printf("User %d (role=%s, auth=%v, active=%v, banned=%v, beta=%v) -> permission=%s, betaFeature=%v, cache=%v\n",
  user.ID, user.Role, user.Authenticated, user.AccountActive, user.IsBanned, user.IsBetaTester,
  permission, showBeta, cacheResponse)
}

func main() {
 users := []User{
  {ID: 1, Role: "admin", Authenticated: true, AccountActive: true, IsBanned: false, IsBetaTester: false},
  {ID: 2, Role: "editor", Authenticated: true, AccountActive: true, IsBanned: false, IsBetaTester: false},
  {ID: 3, Role: "user", Authenticated: true, AccountActive: true, IsBanned: false, IsBetaTester: true},
  {ID: 4, Role: "guest", Authenticated: false, AccountActive: false, IsBanned: false, IsBetaTester: false},
  {ID: 5, Role: "user", Authenticated: true, AccountActive: false, IsBanned: false, IsBetaTester: false},
  {ID: 6, Role: "user", Authenticated: true, AccountActive: true, IsBanned: true, IsBetaTester: true},
 }

 globalFeatureBeta := true

 fmt.Println("Backend decision demo — logical operator based assignments:")
 for _, u := range users {
  evaluate(u, globalFeatureBeta)
 }
}
```
