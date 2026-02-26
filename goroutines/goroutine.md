```go
package main

import (
 "fmt"
 "time"
)

func task(id int) {
 fmt.Printf("Task %d bat dau\n", id)
 time.Sleep(1 * time.Second)
 fmt.Printf("Task %d ket thuc\n", id)
}

func lyThuyetGoroutines() {
 //* Khởi tạo thời gian bắt đầu.
 start := time.Now()

 //? Để khởi tạo goroutine, bạn chỉ cần sử dụng từ khóa "go" trước một hàm hoặc một biểu thức hàm.
 //? Khi bạn gọi một hàm với "go", nó sẽ chạy song song với các goroutine khác và không chặn luồng chính của chương trình
 //? Nếu hàm main xong trước khi các goroutine hoàn thành công việc của chúng, chương trình sẽ kết thúc và các goroutine sẽ bị dừng lại (tránh bị memory leak).
 for i := 1; i <= 5; i++ {
  go task(i)
 }

 time.Sleep(2 * time.Second)

 fmt.Println("Tong thoi gian: ", time.Since(start))
}
```
