```go
package main

import (
 "context"
 "fmt"
 "time"
)

func worker(ctx context.Context) {
 for {
  select {
  case <-ctx.Done():
   fmt.Println("Worker is done")
   return
  default:
   priority := ctx.Value("priority")
   fmt.Println("Worker is running with priority:", priority)
   time.Sleep(100 * time.Millisecond)
  }
 }
}

func lyThuyetContextPart1() {
 ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 defer cancel()

 ctx = context.WithValue(ctx, "priority", "high")

 go worker(ctx)

 time.Sleep(3 * time.Second)
}

func nauPho(ctx context.Context, phoCh chan<- string) {
 fmt.Println("Bat dau nau Pho")
 select {
 case <-ctx.Done():
  fmt.Println("Pho bi huy")
  return
 case <-time.After(1 * time.Second):
  phoCh <- "Pho da nau xong"
 }
}

func nauPizza(ctx context.Context, pizzaCh chan<- string) {
 fmt.Println("Bat dau nau Pizza")
 select {
 case <-ctx.Done():
  fmt.Println("Pizza bi huy")
  return
 case <-time.After(2 * time.Second):
  pizzaCh <- "Pizza da nau xong"
 }
}

func lyThuyetContextPart2() {
 ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
 defer cancel()

 phoCh := make(chan string)
 pizzaCh := make(chan string)

 go nauPho(ctx, phoCh)
 go nauPizza(ctx, pizzaCh)

 for i := 1; i <= 2; i++ {
  select {
  case msgPho := <-phoCh:
   fmt.Println(msgPho)
  case msgPizza := <-pizzaCh:
   fmt.Println(msgPizza)
  case <-ctx.Done():
   fmt.Println("Timeout")
   return
  }
 }
}

func main() {
 lyThuyetContextPart2()
}
```
