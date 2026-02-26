```go
package main

import (
 "fmt"
 "time"
)

/*
Đối với unbuffered channel, phải dùng goroutine để gửi và nhận dữ liệu.
Nếu không sẽ bị deadlock vì cả sender và receiver đều chờ nhau để hoàn thành công việc của mình.
*/
func unbufferedChannel() {
 ch := make(chan int)

 go func() {
  // defer close(ch)
  ch <- 1
  ch <- 2
  ch <- 3
  fmt.Println("Sent")
 }()

 for v := range ch {
  fmt.Println(v)
 }

 fmt.Println("Received")

 time.Sleep(1 * time.Second)
}

/*
Panic và deadlock khi:
  - Sender gửi vào channel đã đầy.
  - Receiver nhận từ channel đã rỗng.

Đối với buffered channel, có thể close ngay sau khi gửi mà không cần defer
*/
func bufferedChannel() {
 ch := make(chan int, 3)

 ch <- 1
 ch <- 2
 ch <- 3
 close(ch)

 fmt.Println("Sent")

 for i := 1; i <= 3; i++ {
  fmt.Println(<-ch)
 }

 fmt.Println("Received")

 time.Sleep(1 * time.Second)
}

func lyThuyetChannel() {
 // unbufferedChannel()
 // bufferedChannel()
 // waitGroupChannel()
 selectChannel()
}

```
