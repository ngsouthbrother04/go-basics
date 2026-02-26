# Sequential, Concurrent và Parallelism trong Golang

## 1) Sequential (Tuần tự)

**Sequential** là cách chạy công việc theo thứ tự từng bước: việc trước xong thì việc sau mới bắt đầu.

- 1 luồng xử lý chính.
- Dễ đọc, dễ debug.
- Phù hợp với tác vụ đơn giản hoặc có phụ thuộc chặt chẽ.

Ví dụ:

```go
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	fmt.Println("start", name)
	time.Sleep(1 * time.Second)
	fmt.Println("done ", name)
}

func main() {
	task("A")
	task("B")
}
```

Kết quả: `A` luôn chạy xong rồi mới đến `B`.

---

## 2) Concurrent (Đồng thời)

**Concurrency** là thiết kế chương trình để nhiều công việc có thể **tiến triển cùng lúc** (đan xen nhau), không nhất thiết cùng thời điểm vật lý trên CPU.

Trong Go, concurrency thường được xây bằng:

- `goroutine` (chạy hàm nhẹ)
- `channel` (trao đổi dữ liệu)
- `sync.WaitGroup`, `mutex`, `context`...

Ví dụ:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start", name)
	time.Sleep(1 * time.Second)
	fmt.Println("done ", name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go task("A", &wg)
	go task("B", &wg)

	wg.Wait()
}
```

Kết quả: thứ tự `A/B` có thể khác nhau giữa các lần chạy.

> Lưu ý: Concurrent **không đồng nghĩa** với Parallel. Chương trình vẫn có thể concurrent dù chỉ dùng 1 core.

---

## 3) Parallelism (Song song)

**Parallelism** là nhiều công việc thực sự chạy cùng thời điểm trên nhiều core CPU.

Trong Go:

- Goroutine cho phép concurrency.
- Việc có chạy song song thật hay không còn phụ thuộc vào số core và `GOMAXPROCS`.

Ví dụ cấu hình:

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func cpuWork(n int) {
	count := 0
	for i := 0; i < 200_000_000; i++ {
		count += i % (n + 1)
	}
	fmt.Println("done", n, count)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() { defer wg.Done(); cpuWork(1) }()
	go func() { defer wg.Done(); cpuWork(2) }()

	wg.Wait()
}
```

Nếu máy có nhiều core, 2 tác vụ CPU-bound ở trên có thể chạy song song thật.

---

## 4) So sánh nhanh

| Khái niệm | Ý nghĩa chính | Có thể trên 1 core? | Cần nhiều core để hiệu quả? |
|---|---|---|---|
| Sequential | Chạy từng việc nối tiếp | Có | Không |
| Concurrent | Nhiều việc cùng tiến triển (đan xen) | Có | Không bắt buộc |
| Parallelism | Nhiều việc chạy cùng lúc thật sự | Không (về bản chất) | Có |

---

## 5) Trong Golang nên hiểu như sau

1. **Goroutine = công cụ concurrency** (mô hình tổ chức công việc).
2. **Parallelism = cách thực thi** phụ thuộc phần cứng và runtime.
3. Mục tiêu thực tế:
   - I/O-bound (HTTP, DB, file): ưu tiên concurrency để tận dụng thời gian chờ.
   - CPU-bound (tính toán nặng): kết hợp concurrency + parallelism để giảm thời gian xử lý.

---

## 6) Kết luận ngắn

- **Sequential**: đơn giản, dễ kiểm soát.
- **Concurrent**: xử lý nhiều việc hiệu quả hơn về mặt tổ chức.
- **Parallelism**: tăng tốc thật sự cho tác vụ nặng khi có nhiều core.

Trong Go, bạn thường thiết kế theo hướng **concurrent first**, sau đó để runtime tận dụng **parallelism** khi có điều kiện phần cứng phù hợp.
