# Goroutines & Concurrency trong Go

Tài liệu này tổng hợp các ý chính trong thư mục `goroutines`, trình bày lại theo cấu trúc dễ học, dễ tra cứu.

---

## Mục lục

1. [Goroutine](#1-goroutine)
2. [Defer](#2-defer)
3. [WaitGroup](#3-waitgroup)
4. [Channel](#4-channel)
5. [Select](#5-select)
6. [Context](#6-context)
7. [Quản lý CPU trong Go](#7-quản-lý-cpu-trong-go)
8. [Checklist thực hành](#8-checklist-thực-hành)

---

## 1) Goroutine

### Goroutine là gì?

Goroutine là **lightweight thread** do Go runtime quản lý, dùng để chạy tác vụ đồng thời (concurrent).

### Khởi tạo goroutine

```go
go task(1)
go func() {
    fmt.Println("Hello from goroutine")
}()
```

### Đặc điểm quan trọng

- `go f()` **không chặn** luồng hiện tại.
- Nếu `main` kết thúc trước, các goroutine còn lại sẽ bị dừng.
- Goroutine nhẹ (stack ban đầu nhỏ), có thể chạy số lượng rất lớn.
- Thứ tự chạy/hoàn thành là **không xác định** (non-deterministic).

### Ví dụ (rút gọn theo `goroutine.md`)

```go
func task(id int) {
    fmt.Printf("Task %d bat dau\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Task %d ket thuc\n", id)
}

func lyThuyetGoroutines() {
    start := time.Now()

    for i := 1; i <= 5; i++ {
        go task(i)
    }

    time.Sleep(2 * time.Second) // Demo đơn giản, không nên lạm dụng
    fmt.Println("Tong thoi gian:", time.Since(start))
}
```

> Khuyến nghị: thay `time.Sleep` bằng `sync.WaitGroup` hoặc channel để đồng bộ đúng cách.

---

## 2) Defer

`defer` trì hoãn gọi hàm cho đến khi hàm bao quanh kết thúc (return/panic).

### Quy tắc

- Chạy theo thứ tự **LIFO** (vào sau ra trước).
- Đối số của `defer` được đánh giá ngay tại dòng `defer`.
- Vẫn chạy khi có `panic` (trừ khi `os.Exit`).

### Ví dụ

```go
func deferExample() {
    fmt.Println("Bat dau deferExample")

    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")

    fmt.Println("Ket thuc deferExample")
}

// Output:
// Bat dau deferExample
// Ket thuc deferExample
// defer 2
// defer 1
```

---

## 3) WaitGroup

`sync.WaitGroup` giúp đợi một nhóm goroutine hoàn thành.

### Cách dùng chuẩn

1. `wg.Add(n)` trước khi launch goroutine.
2. Mỗi goroutine gọi `defer wg.Done()`.
3. Goroutine chính gọi `wg.Wait()`.

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
    wg.Wait()

    fmt.Println("Tong thoi gian:", time.Since(start))
}
```

### Lưu ý

- Tránh `wg.Add()` bên trong goroutine.
- Không `Add` sau khi `Wait` đã bắt đầu.
- `WaitGroup` chỉ để đồng bộ hoàn thành, không truyền dữ liệu.

---

## 4) Channel

Channel là cơ chế giao tiếp an toàn giữa các goroutine.

```go
ch1 := make(chan int)    // unbuffered
ch2 := make(chan int, 3) // buffered
```

### 4.1 Unbuffered vs Buffered

- **Unbuffered**: sender và receiver phải gặp nhau tại cùng thời điểm.
- **Buffered**: sender gửi được khi buffer còn chỗ.

### 4.2 Tránh channel `nil`

```go
var ch chan int
ch <- 1      // block vĩnh viễn
fmt.Println(<-ch) // block vĩnh viễn
```

Luôn `make` trước khi dùng:

```go
ch := make(chan int)
```

### 4.3 Quy tắc block/deadlock

- Một goroutine bị block không sao nếu còn goroutine khác chạy.
- Nếu tất cả cùng block, Go báo `fatal error: all goroutines are asleep - deadlock!`.
- `main` kết thúc thì toàn bộ chương trình dừng.

### 4.4 Đóng channel (`close`)

Quy tắc:

- Chỉ phía sender đóng channel.
- Chỉ đóng khi chắc chắn không gửi nữa.
- Gửi vào channel đã đóng sẽ panic.

### 4.5 `for range` trên channel

`for v := range ch` chỉ dừng khi channel được đóng.

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

> Điểm quan trọng (bổ sung từ code mẫu): nếu dùng `for range` để nhận mà **không có `close(ch)`**, goroutine nhận có thể chờ mãi.

### 4.6 Directional channel

- `chan<- T`: send-only
- `<-chan T`: receive-only

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

### 4.7 Pattern nhiều goroutine + gom kết quả qua channel

Ví dụ theo `waitGroup+Channel.md`:

```go
func channelTask(id int, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Task %d bat dau\n", id)
    time.Sleep(1 * time.Second)
    ch <- fmt.Sprintf("Task %d ket thuc\n", id)
}

func waitGroupChannel() {
    var wg sync.WaitGroup
    ch := make(chan string)

    for i := 1; i <= 4; i++ {
        wg.Add(1)
        go channelTask(i, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for v := range ch {
        fmt.Print(v)
    }
}
```

---

## 5) Select

`select` chờ nhiều thao tác channel cùng lúc.

### Quy tắc

- Nếu nhiều case sẵn sàng, Go chọn ngẫu nhiên một case.
- Không case nào sẵn sàng và không có `default` -> block.
- Có `default` -> non-blocking select.

### Ví dụ nhận từ nhiều channel (theo `select.md`)

```go
func selectChannel() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(3 * time.Second)
        ch1 <- "Hello from channel 1"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Hello from channel 2"
    }()

    for i := 1; i <= 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}
```

### Timeout với `time.After`

```go
select {
case msg := <-ch:
    fmt.Println("Nhan:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
}
```

---

## 6) Context

`context` giúp:

- hủy công việc đang chạy,
- giới hạn thời gian,
- truyền metadata xuyên suốt call chain.

### 6.1 `context.Background()`

Context gốc, dùng ở điểm khởi đầu.

```go
ctx := context.Background()
```

### 6.2 `context.WithCancel()`

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go worker(ctx)
cancel() // gửi tín hiệu dừng
```

### 6.3 `context.WithTimeout()`

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

result, err := callAPI(ctx)
if err != nil {
    fmt.Println(err) // context deadline exceeded
}
```

### 6.4 `context.WithDeadline()`

```go
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### 6.5 `context.WithValue()`

Chỉ nên dùng cho metadata (request ID, trace ID, auth info nhẹ), không dùng để chở business data.

```go
type keyType string
const requestIDKey keyType = "requestID"

ctx := context.WithValue(context.Background(), requestIDKey, "REQ-12345")
```

### Quy tắc vàng với context

1. Luôn là tham số đầu tiên: `func Do(ctx context.Context, ...)`
2. Không nhúng context vào struct.
3. Luôn gọi `cancel()` sau `WithCancel/WithTimeout/WithDeadline`.
4. Không truyền `nil` context.

---

## 7) Quản lý CPU trong Go

Khi học goroutine, cần tách rõ 2 ý:

- **Concurrency**: tổ chức nhiều công việc cùng tiến triển.
- **Parallelism**: nhiều công việc thực sự chạy cùng lúc trên nhiều core CPU.

### 7.1 `runtime.NumCPU()` và `runtime.GOMAXPROCS()`

- `runtime.NumCPU()`: trả về số logical CPU của máy.
- `runtime.GOMAXPROCS(n)`: đặt số OS thread có thể chạy Go code đồng thời.

```go
numCPU := runtime.NumCPU()
fmt.Println("CPU cores:", numCPU)

old := runtime.GOMAXPROCS(numCPU)
fmt.Println("Old GOMAXPROCS:", old)
```

> Từ Go 1.5+, mặc định `GOMAXPROCS` thường đã bằng số CPU khả dụng, nhưng vẫn nên hiểu để tuning khi benchmark.

### 7.2 Mô hình G-M-P (nên nhớ ngắn gọn)

- **G (Goroutine)**: tác vụ bạn tạo bằng `go`.
- **M (Machine)**: OS thread thật.
- **P (Processor)**: tài nguyên scheduler để chạy goroutine.

Go scheduler gán nhiều `G` lên một nhóm `M/P`, giúp chạy hiệu quả với chi phí thấp hơn thread truyền thống.

### 7.3 Khi nào cần quan tâm CPU tuning?

- **I/O-bound** (HTTP, DB, file): thường không cần chỉnh nhiều, tập trung vào concurrency + timeout/context.
- **CPU-bound** (tính toán nặng): đo đạc với các mức `GOMAXPROCS` khác nhau để tìm điểm tối ưu.

### 7.4 Ghi chú thực tế từ bài test `heavyTask`

- Số goroutine nên dùng vòng lặp `for i := 0; i < n; i++` để tránh chạy dư 1 task.
- Với phép cộng lớn, nên dùng `int64` để an toàn hơn trên mọi kiến trúc.
- Nên giữ kết quả tính toán (hoặc cộng dồn) để tránh bị tối ưu bỏ công việc khi benchmark.

Ví dụ khung đo thời gian:

```go
start := time.Now()
runtime.GOMAXPROCS(runtime.NumCPU())

var wg sync.WaitGroup
for i := 0; i < 20; i++ {
    wg.Add(1)
    go heavyTask(&wg)
}
wg.Wait()

fmt.Println("Time taken:", time.Since(start))
```

---

## 8) Checklist thực hành

- Ưu tiên `WaitGroup`/channel thay cho `time.Sleep` để đồng bộ.
- Dùng `close(ch)` đúng phía sender, đúng thời điểm.
- Khi đọc bằng `for range ch`, đảm bảo có nơi đóng channel.
- Dùng directional channel để giới hạn quyền send/receive.
- Dùng `select + time.After` để chống treo khi chờ dữ liệu.
- Truyền `context` xuyên suốt các hàm có I/O hoặc tác vụ dài.
