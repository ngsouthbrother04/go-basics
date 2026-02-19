# Array, Slice và Map trong Go

---

# I. Array trong Go

## 1. Khái niệm cơ bản

- **Array** là tập hợp các phần tử **cùng kiểu dữ liệu**.
- **Kích thước cố định** tại thời điểm compile.
- Kích thước là **một phần của type** → `[3]int` và `[5]int` là hai type hoàn toàn khác nhau.

### Đặc điểm quan trọng

| Thuộc tính       | Mô tả                                                                 |
|------------------|-----------------------------------------------------------------------|
| Fixed size       | Không thể thay đổi kích thước sau khi khai báo                        |
| Part of type     | `[3]int` ≠ `[4]int`                                                   |
| Value type       | Khi truyền vào function → copy toàn bộ array                          |
| Ít sử dụng thực tế| Thường thay thế bằng slice                                            |

## 2. Khai báo array

```go
// Cách 1: Khai báo đầy đủ
var a [3]int = [3]int{1, 2, 3}

// Cách 2: Compiler suy luận kích thước
b := [...]int{1, 2, 3, 4}

// Cách 3: Short declaration
c := [4]int{10, 20, 30, 40}
```

## 3. Array là value type

```go
func modify(arr [3]int) {
    arr[0] = 999
}

nums := [3]int{1, 2, 3}
modify(nums)
fmt.Println(nums) // [1 2 3] → không thay đổi
```

**Muốn thay đổi gốc → dùng pointer**:

```go
func modify(arr *[3]int) {
    arr[0] = 999
}

nums := [3]int{1, 2, 3}
modify(&nums)
fmt.Println(nums) // [999 2 3]
```

## 4. Array đa chiều

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

## 5. Khi nào nên dùng array?

- Kích thước cố định và biết trước tại compile-time.
- Ma trận toán học, màu RGB `[3]uint8`, buffer cố định.
- Cần hiệu suất cao với dữ liệu nhỏ (không copy khi truyền pointer).

---

# II. Slice trong Go

## 1. Bản chất của slice

> **Slice là một "view" (cửa sổ) trỏ vào một underlying array.**

Slice **không lưu dữ liệu trực tiếp** — dữ liệu thực sự nằm trong underlying array.

```
Underlying array: [10, 20, 30, 40, 50, ...]
Slice header:     pointer → vị trí bắt đầu + len + cap
```

## 2. Cấu trúc nội bộ của slice (Slice Header)

Trong runtime của Go, một slice được biểu diễn bởi struct sau (trong package `reflect` là `SliceHeader`, trong runtime là tương tự):

```go
type SliceHeader struct {
    Data uintptr // pointer tới phần tử đầu tiên của view
    Len  int
    Cap  int
}
```

- **`Data` (pointer)**: Trỏ đến phần tử đầu tiên mà slice đang "nhìn thấy" trong underlying array. Đây là kiểu `unsafe.Pointer` trong runtime, nhưng có thể hiểu là `*T` (con trỏ tới kiểu phần tử).
- **`Len`**: Số phần tử hiện tại slice có thể truy cập (`len(s)`).
- **`Cap`**: Dung lượng tối đa từ vị trí bắt đầu của slice đến cuối underlying array (`cap(s)`).

**Slice là value type** — khi copy slice, chỉ copy 3 trường này (pointer + len + cap), không copy dữ liệu.

```go
s1 := []int{1, 2, 3}
s2 := s1        // copy header → s2 trỏ cùng underlying array
s2[0] = 999
fmt.Println(s1) // [999 2 3] → thay đổi chung
```

## 3. Khai báo slice

```go
// Nil slice
var s []int                 // s == nil, len=0, cap=0

// Slice literal (tạo underlying array ngầm)
s := []int{1, 2, 3}

// make
s := make([]int, 3)      // len=3, cap=3
s := make([]int, 3, 10)  // len=3, cap=10 (pre-allocate)
```

## 4. Length vs Capacity

```go
s := make([]int, 3, 10)
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 10
```

- `len`: số phần tử hiện có.
- `cap`: số phần tử tối đa có thể thêm mà không cần realloc underlying array.

## 5. Slicing và shared underlying array

```go
original := []int{10, 20, 30, 40, 50}
sub := original[1:4]    // len=3, cap=4, pointer trỏ đến original[1]

sub[0] = 999
fmt.Println(original)   // [10 999 30 40 50]
```

**Cảnh báo**: Các slice con chia sẻ cùng underlying array → thay đổi một slice ảnh hưởng các slice khác.

## 6. Append và realloc

```go
s := []int{1, 2, 3}
s = append(s, 4)                 // phải gán lại
s = append(s, 5, 6, 7)
s = append(s, anotherSlice...)
```

**Khi nào xảy ra realloc?**

| Trường hợp          | Hậu quả                                      |
|---------------------|----------------------------------------------|
| Còn capacity        | Dùng chung underlying array                  |
| Hết capacity        | Tạo underlying array mới (thường gấp đôi), copy dữ liệu, **pointer thay đổi** |

Sau realloc, các slice cũ vẫn trỏ vào array cũ → có thể gây bug logic.

## 7. Copy slice độc lập

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)          // dst hoàn toàn độc lập
// Hoặc Go 1.21+: clone := slices.Clone(src)
```

## 8. Best practices với slice

- **Pre-allocate** khi biết kích thước gần đúng: `make([]T, 0, expectedSize)`.
- Tránh bug shared memory: nếu cần slice độc lập → dùng `copy` hoặc `slices.Clone`.
- Cẩn thận khi giữ reference đến slice con sau khi append có thể realloc.

---

# III. So sánh Array vs Slice

| Tiêu chí          | Array              | Slice                          |
|-------------------|--------------------|--------------------------------|
| Kích thước        | Cố định            | Động                           |
| Type              | `[n]T`             | `[]T`                          |
| Bộ nhớ            | Inline (trong stack)| Pointer tới underlying array   |
| Truyền hàm        | Copy toàn bộ       | Copy header (ptr+len+cap)      |
| Append            | Không              | Có                             |
| Sử dụng thực tế   | Hiếm               | Rất phổ biến                   |

---

# IV. Package `slices` (Go 1.21+)

```go
import "slices"
```

Một số hàm hữu ích:

- `slices.Equal`, `slices.Compare`, `slices.Contains`, `slices.Index`
- `slices.Insert`, `slices.Delete`, `slices.Replace`
- `slices.Clone`, `slices.Grow`, `slices.Clip`
- `slices.Sort`, `slices.SortFunc`, `slices.Reverse`
- `slices.Max`, `slices.Min`, `slices.Concat`

---

# V. Map trong Go

## 1. Khái niệm cơ bản

- **Map** là hash table lưu cặp **key-value**.
- **Reference type** → truyền vào hàm không copy dữ liệu.
- **Unordered** → thứ tự iterate không cố định.
- Key phải là **comparable type**.

## 2. Khai báo

```go
// Tốt nhất: dùng make
m := make(map[string]int)

// Literal
m := map[string]int{"Alice": 25, "Bob": 30}

// Nil map (chỉ đọc được)
var m map[string]int // m == nil → không thể ghi
```

**Lỗi phổ biến**: Ghi vào nil map → panic.

## 3. Thao tác cơ bản

```go
m["Alice"] = 26                  // thêm/cập nhật
value := m["Alice"]              // đọc (zero value nếu không tồn tại)

value, ok := m["Bob"]            // comma-ok idiom kiểm tra tồn tại
delete(m, "Alice")               // xóa (an toàn với key không tồn tại)
len(m)                           // số phần tử
```

## 4. Duyệt map

```go
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
```

Thứ tự **không đảm bảo**.

## 5. Map là reference type

```go
func modify(m map[string]int) {
    m["Alice"] = 999
}

m := map[string]int{"Alice": 25}
modify(m)
fmt.Println(m["Alice"]) // 999
```

## 6. Best practices

- Luôn `make` trước khi dùng.
- Dùng comma-ok để phân biệt zero value và key không tồn tại.
- Pre-allocate capacity nếu biết trước: `make(map[K]V, hint)`.
- Không dựa vào thứ tự iterate → sort keys nếu cần.

## 7. Package `maps` (Go 1.21+)

```go
import "maps"
```

- `maps.Clone`, `maps.Copy`, `maps.Equal`, `maps.DeleteFunc`

---

# VI. Tổng kết cốt lõi

### Array

- Fixed size, value type, ít dùng.

### Slice

- View vào underlying array.
- Header: **pointer + len + cap**.
- Shared memory → dễ bug nếu không cẩn thận.
- Append có thể thay đổi pointer (realloc).

### Map

- Hash table, reference type, unordered.
- Key comparable, nil map không ghi được.
- Dùng comma-ok để kiểm tra tồn tại.

**Mental model quan trọng nhất với slice**:

> Slice ≠ dữ liệu  
> Slice = Pointer + Metadata (len, cap)  
> Dữ liệu thực nằm trong underlying array
