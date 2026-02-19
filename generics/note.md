# Generics trong Golang

## Generics là gì

Generics cho phép bạn viết hàm hoặc kiểu dữ liệu tổng quát, hoạt động
với nhiều kiểu dữ liệu khác nhau mà vẫn đảm bảo type-safety tại compile
time.

``` go
func Print[T any](value T) {
    fmt.Println(value)
}
```

## Comparable

### Khi nào sử dụng được Comparable

- Chỉ có thể sử dụng được với các kiểu dữ liệu có toán tử == và !=
- Đó là: string, int, float, bool, struct (KHÔNG chứa slice, map và
    function)
- Các kiểu slice, map và function không thể dùng Comparable vì không
    hỗ trợ toán tử so sánh ==.

``` go
func compare[T comparable](a, b T) bool {
 return a == b
}
```

## cmp.Ordered

- Dùng để so sánh lớn bé

------------------------------------------------------------------------

# So sánh giữa cmp và constraints

## constraints

- Thuộc gói experimental (x/exp)
- Dùng từ Go 1.18
- Cung cấp các type constraint như:
  - constraints.Ordered
  - constraints.Integer
  - constraints.Float

Ví dụ:

``` go
import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

Nhược điểm: - Không thuộc standard library - Không đảm bảo ổn định API
lâu dài - Không nên dùng trong production mới

------------------------------------------------------------------------

## cmp (Go 1.21+)

- Thuộc standard library
- Ổn định và được khuyến nghị sử dụng
- Cung cấp:
  - cmp.Ordered
  - cmp.Compare()

Ví dụ:

``` go
import "cmp"

func Max[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

Hoặc dùng Compare:

``` go
cmp.Compare(a, b)
```

Giá trị trả về: - -1 nếu a \< b - 0 nếu a == b - 1 nếu a \> b

------------------------------------------------------------------------

## So sánh chi tiết

  Tiêu chí           constraints         cmp
  ------------------ ------------------- ------------------
  Thuộc về           x/exp               Standard library
  Stable API         Không đảm bảo       Có
  Go version         1.18+               1.21+
  Ordered            Có                  Có
  Compare function   Không               Có
  Production ready   Không khuyến nghị   Khuyến nghị

------------------------------------------------------------------------

## Khi nào dùng cái nào?

- Go \>= 1.21 → Dùng `cmp`
- Code legacy (1.18--1.20) → Có thể dùng `constraints`
- Code production mới → Không dùng `x/exp`

------------------------------------------------------------------------

## Tư duy quan trọng

- `Ordered` dùng khi cần so sánh \< \> \<= \>=
- `comparable` dùng khi cần == hoặc làm key map
- Generics nên dùng cho data structure và utility layer
- Không lạm dụng generics trong business logic

## Type constraints

### Khái niệm

Type constraints xác định những kiểu dữ liệu nào có thể được dùng với generic. 
Đó là các "yêu cầu" mà một kiểu phải thỏa mãn để được phép sử dụng.

### Các loại constraitns cơ bản

#### 1. `any` - Không giới hạn

```go
func Print[T any](value T) {
    fmt.Println(value)
}

// Dùng được với bất kỳ kiểu nào
Print(42)           // int
Print("hello")      // string
Print([]int{1,2,3}) // slice
```

#### 2. `comparable` - Có thể so sánh bằng ==

```go
func Contains[T comparable](slice []T, value T) bool {
    for _, v := range slice {
        if v == value {
            return true
        }
    }
    return false
}

// Dùng được
Contains([]int{1, 2, 3}, 2)           // ✓ int comparable
Contains([]string{"a", "b"}, "a")     // ✓ string comparable

// Không dùng được
Contains([][]int{{1}, {2}}, []int{1}) // ✗ slice không comparable
```

#### 3. `cmp.Ordered` - Có thể so sánh lớn bé

```go
import "cmp"

func Max[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Dùng được
Max(10, 20)        // ✓ int
Max(3.5, 2.1)      // ✓ float64
Max("apple", "zoo") // ✓ string

// Không dùng được
Max([]int{1}, []int{2}) // ✗ slice không ordered
```

### Custom Type Constraints

#### Interface-based constraints

```go
// Định nghĩa constraint
type Number interface {
    int | int64 | float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

Sum([]int{1, 2, 3})        // ✓ 
Sum([]float64{1.5, 2.5})   // ✓
Sum([]string{"a", "b"})    // ✗ string không thuộc Number
```

#### Kết hợp nhiều constraints

```go
// Constraint phải là int, int64, hoặc có thể so sánh
type ComparableInt interface {
    int | int32 | int64
    comparable // phải thỏa mãn cả comparable
}

func FindMin[T ComparableInt](nums []T) T {
    min := nums[0]
    for _, n := range nums {
        if n < min {
            min = n
        }
    }
    return min
}
```

#### Kết hợp interface + type set (Go 1.18+)

```go
type Reader interface {
    Read([]byte) (int, error)
}

type ReadSeeker interface {
    Reader
    Seek(int64, int) (int64, error)
}

// Generic với interface constraint
func CopyData[T ReadSeeker](src T) error {
    data := make([]byte, 1024)
    _, err := src.Read(data)
    return err
}
```

### Ví dụ thực tế

#### Sorting generic

```go
func BubbleSort[T cmp.Ordered](slice []T) {
    for i := 0; i < len(slice); i++ {
        for j := 0; j < len(slice)-i-1; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
}

BubbleSort([]int{3, 1, 4, 1, 5})
BubbleSort([]string{"c", "a", "b"})
```

#### Generic data structure (Stack)

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(value T) {
    s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, fmt.Errorf("stack empty")
    }
    value := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return value, nil
}

// Dùng
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
val, _ := intStack.Pop() // 2
```

### Bảng cheat sheet Constraints

| Constraint | Dùng khi | Ví dụ kiểu |
|-----------|---------|-----------|
| `any` | Không cần giới hạn | Bất kỳ kiểu nào |
| `comparable` | Cần dùng ==, != hoặc làm map key | int, string, struct |
| `cmp.Ordered` | Cần <, >, <=, >= | int, float64, string |
| `int \| float64` | Giới hạn số kiểu cụ thể | Chỉ int hoặc float64 |
| Interface | Cần implement method | Reader, Writer |

### Best practices

- ✓ Dùng constraint **càng general càng tốt** (any > comparable > Ordered)
- ✓ Tạo **custom constraint** khi cần reuse logic
- ✓ Dùng generics cho **utility functions** và **data structures**
- ✗ Không dùng generics trong **complex business logic**
- ✗ Không dùng `any` khi có constraint cụ thể hơn
