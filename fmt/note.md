
**Cheatsheet: Formatting verbs (quick reference)**

| Verb | Loại | Mô tả (ví dụ) |
|------|------|----------------|
| `%v` | general | Giá trị theo định dạng mặc định
| `%+v` | general | Như `%v` nhưng show tên trường struct
| `%#v` | general | Biểu diễn theo cú pháp Go (hữu ích để debug)
| `%T` | general | Kiểu của giá trị
| `%t` | boolean | `true` / `false`
| `%d` | int | Decimal (base 10)
| `%c` | int | Unicode code point -> char
| `%.f` | float | Fixed-point (use precision `.2f`)
| `%s` | string/bytes | Raw string
| `%p` | pointer | Pointer address (hex)

---

**Ghi chú về package fmt và Formatting Verbs trong Go**

1) Các hàm in/định dạng chính
- `fmt.Print(...)`: in các giá trị, không thêm khoảng trắng hoặc newline ngoài những gì các giá trị cung cấp.
- `fmt.Println(...)`: giống `Print` nhưng thêm khoảng trắng giữa các tham số và newline ở cuối.
- `fmt.Printf(format, ...)`: in theo chuỗi định dạng (format string) với các verbs (ví dụ `%d`, `%s`).
- `fmt.Sprint/Sprintln/Sprintf`: giống `Print/Println/Printf` nhưng trả về chuỗi thay vì in ra stdout.
- `fmt.Fprint/Fprintln/Fprintf`: tương tự nhưng in vào `io.Writer` (ví dụ `os.Stdout`, file, buffer).

2) Các hàm đọc input
- `fmt.Scan(...)`: đọc token phân tách bởi whitespace vào các biến (phải truyền địa chỉ `&v`).
- `fmt.Scanln(...)`: như `Scan` nhưng dừng tại newline; nếu còn token trên cùng dòng có thể trả về lỗi.
- `fmt.Scanf(format, ...)`: đọc theo chuỗi định dạng; trả về số giá trị đã đọc và lỗi nếu input không phù hợp.

3) Nguyên tắc khi dùng Scan/Scanf
- Luôn truyền biến bằng địa chỉ (ví dụ `fmt.Scan(&name, &age)`).
- Kiểm tra giá trị trả về `(n, err) := fmt.Scanf(...);` để xử lý input không hợp lệ.

4) Một vài lưu ý thực tế
- Dùng `%+v` để debug struct khi muốn thấy tên trường.
- Dùng `%#v` để xem biểu diễn Go-syntax (hữu ích khi muốn tái tạo giá trị).
- Tránh dùng `fmt.Scanf` với input không đáng tin cậy nếu không kiểm tra lỗi kỹ càng.
- Khi format số học, luôn cân nhắc precision (`.2f` cho 2 chữ số thập phân).

---

**Sự khác biệt chính giữa `bufio` và các hàm `Scan` của package `fmt` trong Go**

Cả hai đều dùng để đọc input từ console (stdin), nhưng chúng được thiết kế cho các mục đích khác nhau.

| Đặc điểm                          | `fmt.Scan / Scanf / Scanln`                                      | `bufio` (thường là `bufio.Scanner` hoặc `bufio.Reader`)                  |
|-----------------------------------|------------------------------------------------------------------|-------------------------------------------------------------------------|
| **Cách đọc dữ liệu**              | Đọc theo **token** (các từ cách nhau bởi khoảng trắng hoặc newline). Dừng khi đủ số lượng biến hoặc gặp space/newline. | `Scanner`: Đọc **từng dòng đầy đủ** (bao gồm khoảng trắng, đến khi gặp newline).<br>`Reader`: Đọc low-level hơn (theo byte, đến delimiter tùy chỉnh). |
| **Parse kiểu dữ liệu**            | Parse **tự động** trực tiếp vào biến (int, float, string,...).   | Chỉ đọc về **string** (hoặc []byte). Bạn phải tự parse sau bằng `strconv.Atoi`, `strconv.ParseFloat`,... |
| **Xử lý khoảng trắng/space**      | **Không** giữ space: "hello world" → chỉ đọc được "hello" (phần còn lại để lại cho lần đọc sau). | Giữ nguyên toàn bộ dòng, bao gồm space: "hello world" → nhận đầy đủ "hello world". |
| **Tiện lợi cho input đơn giản**   | Rất cao: ít code, parse ngay lập tức. Phù hợp nhập số nguyên, nhiều giá trị cùng lúc. | Phải viết thêm code để parse, nhưng linh hoạt hơn. |
| **Tiện lợi cho input phức tạp**   | Thấp: khó xử lý dòng có space, input lớn, hoặc định dạng tùy chỉnh. | Cao: dễ đọc cả dòng, xử lý input lớn, nhiều dòng, hoặc delimiter đặc biệt. |
| **Hiệu suất**                     | Tốt cho input nhỏ.                                               | Tốt hơn cho input lớn (buffered I/O).                                   |
| **Trường hợp dùng phổ biến**      | Chương trình console đơn giản, nhập số nhanh (ví dụ: bài tập cơ bản). | Chương trình thực tế: đọc file lớn, xử lý dòng có space, input nhiều dòng (hầu hết production code dùng bufio). |
| **Xử lý lỗi**                     | Trả về số lượng item đọc được và error.                          | `Scanner.Scan()` trả về bool (true nếu còn dòng), lỗi lấy bằng `Scanner.Err()`. |

### Ví dụ minh họa

**Dùng `fmt.Scanln` (không giữ space):**

```go
var name string
var age int
_, err := fmt.Scanln(&name, &age)  // Người dùng nhập: "Nguyen Van A 25"
if err != nil { ... }
// Kết quả: name = "Nguyen" (chỉ lấy token đầu), age không đọc được đúng (vì còn "Van", "A" ở giữa)
```

→ Không phù hợp nếu tên có space.

**Dùng `bufio.Scanner` (giữ nguyên dòng – khuyến khích):**

```go
import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

scanner := bufio.NewScanner(os.Stdin)
if scanner.Scan() {  // đọc một dòng đầy đủ
    line := scanner.Text()  // line = "Nguyen Van A 25" (giữ nguyên space)
    parts := strings.Fields(line)  // tách thành []string nếu cần
    age, _ := strconv.Atoi(parts[len(parts)-1])  // parse tuổi
    name := strings.Join(parts[:len(parts)-1], " ")  // ghép lại tên
}
```

→ Linh hoạt, xử lý tốt tên có space, nhiều trường hợp thực tế.

**Dùng `bufio.Reader` (nếu cần delimiter tùy chỉnh):**

```go
reader := bufio.NewReader(os.Stdin)
line, err := reader.ReadString('\n')  // đọc đến newline, giữ space
line = strings.TrimSpace(line)  // bỏ \n
```

### Kết luận

- **`fmt.Scan*`**: Nhanh gọn cho **input đơn giản, không có space**, phù hợp bài tập học sinh/sinh viên.
- **`bufio`**: **Linh hoạt, mạnh mẽ hơn**, được dùng hầu hết trong code thực tế (đặc biệt `bufio.Scanner` là cách idiomatic nhất để đọc từng dòng từ console).

Nếu bạn đang viết chương trình đọc input người dùng có thể chứa space (tên, câu văn,...), hãy dùng `bufio.Scanner` – nó an toàn và dễ kiểm soát hơn rất nhiều!
