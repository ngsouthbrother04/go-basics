```go
package main

import "fmt"

func deferExample() {
 fmt.Println("Bat dau deferExample")

 defer fmt.Println("defer 1")
 defer fmt.Println("defer 2")

 fmt.Println("Ket thuc deferExample")
}

```
