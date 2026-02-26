```go
package main

import (
 "fmt"
 "sync"
)

func main() {
 token := 0
 var wg sync.WaitGroup
 var mutex sync.Mutex

 for i := 1; i <= 1000; i++ {
  wg.Add(1)
  go func() {
   mutex.Lock() //Khoá để đảm bảo chỉ một goroutine có thể truy cập và thay đổi giá trị của token tại một thời điểm
   token++

   mutex.Unlock() //Mở khoá để cho phép các goroutine khác có thể truy cập và thay đổi giá trị của token
   wg.Done()
  }()
 }

 wg.Wait()
 fmt.Println(token)
}

```
