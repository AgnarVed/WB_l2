Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Вывод: 2 1
В функции test после выполнения тела функции, до выхода из нее, 
переменная "X" инкрементируется и затем происходит возврать именованной переменной из функции

В другой функции возврат осуществляется до выполнения блока defer, поэтому возвращаемая переменная не инкрементируется

```
