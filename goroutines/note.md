# Goroutines

## Goroutine

### Goroutine là gì?
Goroutine là một lightweight thread (luồng nhẹ) được quản lý bởi Go runtime. Đây là một trong những tính năng mạnh mẽ nhất của Go, cho phép thực thi các tác vụ đồng thời (concurrent) một cách hiệu quả.

### Cách khởi tạo Goroutine
Để khởi tạo goroutine, bạn chỉ cần sử dụng từ khóa `go` trước một hàm hoặc một biểu thức hàm.

```go
go task(1)           // Gọi hàm với goroutine
go func() {          // Hàm ẩn danh với goroutine
    fmt.Println("Hello from goroutine")
}()
```

### Đặc điểm của Goroutine

#### 1. Chạy đồng thời (Concurrent)
Khi bạn gọi một hàm với `go`, nó sẽ chạy song song với các goroutine khác và không chặn luồng chính của chương trình.

```go
go task(1)    // Không chặn, tiếp tục chạy code bên dưới ngay lập tức
go task(2)
// Code tiếp tục chạy mà không đợi task(1) và task(2) hoàn thành
```

#### 2. Main Goroutine và Program Termination
Nếu hàm `main` xong trước khi các goroutine hoàn thành công việc của chúng, chương trình sẽ kết thúc và các goroutine sẽ bị dừng lại.

**Vấn đề:**
```go
func main() {
    go task(1)
    go task(2)
    // main kết thúc ngay lập tức, các goroutine bị hủy
}
```

**Giải pháp tạm thời:**
```go
func main() {
    go task(1)
    go task(2)
    time.Sleep(2 * time.Second)  // Đợi goroutines hoàn thành
}
```

**Giải pháp tốt hơn:** Sử dụng `sync.WaitGroup` hoặc channel để đồng bộ.

#### 3. Lightweight và hiệu quả
- Goroutine rất nhẹ, chỉ tốn khoảng 2KB stack ban đầu (so với thread thông thường tốn ~1-2MB)
- Có thể chạy hàng nghìn, thậm chí hàng triệu goroutine đồng thời
- Go runtime quản lý scheduling goroutine lên các OS thread

### Ví dụ thực tế
Trong code mẫu, chúng ta chạy 5 task đồng thời:

```go
for i := 1; i <= 5; i++ {
    go task(i)
}
```

**Kết quả quan sát:**
- Nếu chạy tuần tự: 5 task × 1 giây = 5 giây
- Với goroutine: ~2 giây (vì chạy đồng thời)
- Thứ tự hoàn thành có thể không theo thứ tự (non-deterministic)

### Lưu ý quan trọng

1. **Không đoán trước được thứ tự thực thi:** Goroutine có thể chạy theo bất kỳ thứ tự nào
2. **Race condition:** Cần cẩn thận khi nhiều goroutine truy cập cùng một dữ liệu
3. **Synchronization:** Sử dụng channel, mutex hoặc WaitGroup để đồng bộ
4. **Main function:** Luôn đảm bảo `main` không kết thúc trước khi goroutine quan trọng hoàn thành

### So sánh: Concurrency vs Parallelism

- **Concurrency (đồng thời):** Nhiều task được quản lý cùng lúc (dealing with multiple things at once)
- **Parallelism (song song):** Nhiều task thực thi cùng lúc (doing multiple things at once)

Goroutine hỗ trợ cả hai, tùy thuộc vào số CPU cores và cách Go scheduler phân bổ.

### Defer trong Go
`defer` dùng để trì hoãn việc gọi một hàm cho đến khi hàm bao quanh kết thúc (return hoặc panic). Điều này hữu ích để giải phóng tài nguyên (đóng file, unlock mutex, v.v.).

#### Đặc điểm chính
- Lệnh `defer` được lưu theo cơ chế LIFO (vào sau, ra trước)
- Giá trị đối số của `defer` được đánh giá ngay tại thời điểm gặp `defer`
- `defer` vẫn chạy khi có `panic`, trừ khi chương trình bị `os.Exit`

#### Ví dụ
```go
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close() // Luôn đóng file khi hàm kết thúc

    // Xử lý đọc file ở đây
    return nil
}
```

#### Ví dụ LIFO
```go
func main() {
    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")
    fmt.Println("main")
}
// Output:
// main
// defer 2
// defer 1
```

## WaitGroup

`sync.WaitGroup` dùng để đợi một nhóm goroutine hoàn thành. Đây là cách đồng bộ phổ biến khi bạn muốn chắc chắn tất cả goroutine đã xong trước khi tiếp tục.

### Cách dùng cơ bản
1. Gọi `wg.Add(n)` để tăng bộ đếm lên `n`
2. Mỗi goroutine khi hoàn thành gọi `wg.Done()` (giảm bộ đếm)
3. Goroutine chính gọi `wg.Wait()` để đợi bộ đếm về 0
4. Nếu `WaitGroup` được truyền vào hàm thì nên dùng con trỏ

### Ví dụ
```go
func wgTask(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Printf("Task %d ket thuc\n", id)

    fmt.Printf("Task %d bat dau\n", id)
    time.Sleep(1 * time.Second)
}

func waitGroup() {
    start := time.Now()

    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go wgTask(i, &wg)
    }
    wg.Wait() // Đợi tất cả goroutine hoàn thành

    fmt.Println("Tong thoi gian:", time.Since(start))
}
```

### Lưu ý quan trọng

- Không gọi `wg.Add()` bên trong goroutine nếu có thể, vì dễ gây race
- Tránh gọi `wg.Add()` sau khi đã gọi `wg.Wait()`
- `WaitGroup` chỉ dùng để đợi, không truyền dữ liệu (dùng channel khi cần giao tiếp)

## Channel

Channel là cơ chế giao tiếp an toàn giữa các goroutine. Có 2 loại: unbuffered và buffered.

### 1. Unbuffered vs Buffered

- Unbuffered: sender và receiver phải bắt tay cùng lúc
- Buffered: sender có thể gửi trước nếu còn dung lượng buffer

```go
ch1 := make(chan int)     // unbuffered
ch2 := make(chan int, 3)  // buffered, dung lượng 3
```

### 1.1. Tránh dùng channel nil

Channel có giá trị `nil` sẽ **block vĩnh viễn** khi gửi hoặc nhận. Vì vậy, không nên khởi tạo channel bằng `var ch chan int` rồi dùng ngay mà chưa `make`.

**Ví dụ gây block:**
```go
var ch chan int
ch <- 1        // block vĩnh viễn
fmt.Println(<-ch) // block vĩnh viễn
```

**Cách đúng:**
```go
ch := make(chan int)
ch <- 1
fmt.Println(<-ch)
```

**Lưu ý:** Channel `nil` đôi khi được dùng có chủ đích để **vô hiệu hóa một case** trong `select` (bật/tắt luồng xử lý), nhưng đây là kỹ thuật nâng cao.

### 2. Quy tắc block (để tránh deadlock)

- Bên sẽ block phải được launch trước bên còn lại
- Không để cả sender và receiver block trên cùng một goroutine

**Đúng (receiver sẵn sàng trước):**
```go
ch := make(chan int)

go func() {
    fmt.Println(<-ch)
}()

ch <- 1
```

**Sai (deadlock):**
```go
ch := make(chan int)

ch <- 1        // main block ngay tại đây
go func() {
    fmt.Println(<-ch)
}()
```

### 3. Block có gây crash không?

- Một goroutine block không crash nếu còn goroutine khác chạy
- Tất cả goroutine đều block thì Go báo lỗi deadlock
- `main` kết thúc thì toàn bộ chương trình dừng lại

### 4. Đóng channel (close)

Quy tắc:
- Chỉ sender mới được `close(ch)`
- Close sau khi đã gửi xong dữ liệu
- Gửi vào channel đã đóng sẽ panic

**Pattern nhiều sender + WaitGroup:**
```go
var wg sync.WaitGroup
ch := make(chan string, 10)

for i := 1; i <= 4; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        ch <- fmt.Sprintf("Task %d", id)
    }(i)
}

go func() {
    wg.Wait()
    close(ch)
}()

for v := range ch {
    fmt.Println(v)
}
```

### 5. Directional channel (`chan<-` và `<-chan`)

- `chan<-` : chỉ được gửi
- `<-chan` : chỉ được nhận

```go
func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for v := range ch {
        fmt.Println(v)
    }
}
```

### 6. `for range` trên channel

```go
for v := range ch {
    fmt.Println(v)
}
```

Tương đương:
```go
for {
    v, ok := <-ch
    if !ok {
        break
    }
    fmt.Println(v)
}
```

### 7. `select` trong channel

`select` cho phép chờ và xử lý nhiều thao tác channel cùng lúc. Nó giống `switch`, nhưng mỗi `case` là một thao tác gửi/nhận trên channel.

#### Quy tắc quan trọng
- Nếu nhiều `case` sẵn sàng, Go sẽ chọn ngẫu nhiên một case để tránh starvation.
- Nếu không có `case` nào sẵn sàng và **không có** `default`, `select` sẽ block.
- Nếu có `default`, `select` sẽ chạy ngay `default` khi không có case sẵn sàng.

#### Ví dụ nhận từ nhiều channel
```go
select {
case msg1 := <-ch1:
    fmt.Println("ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("ch2:", msg2)
}
```

#### Ví dụ timeout với `time.After`
```go
select {
case msg := <-ch:
    fmt.Println("Nhan:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
}
```

#### Ví dụ non-blocking với `default`
```go
select {
case msg := <-ch:
    fmt.Println("Nhan:", msg)
default:
    fmt.Println("Khong co du lieu, tiep tuc lam viec")
}
```

### 8. Checklist nhanh

- Unbuffered cần đồng bộ chặt
- Buffered vẫn cần close đúng lúc
- Close: chỉ sender, chỉ 1 lần
- Dùng `for range` để đọc đến khi channel đóng
- Tránh goroutine leak bằng done channel hoặc `context`
