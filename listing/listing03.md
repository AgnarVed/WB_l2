Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Вывод: nil false
Если делаем обёртку для ошибки, то сравнивая значения с nil, можно словить ошибку, потому что ошибка внутри может хранить действительно nil, но сам интерфейс
не nil, потому что интерфейс хранит значение - ссылку на тип объекта внутри

```
