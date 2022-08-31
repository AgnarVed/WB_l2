Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error
Если делаем обёртку для ошибки, то сравнивая значения с nil, можно словить ошибку, потому что ошибка внутри может хранить действительно nil, но сам интерфейс
не nil, потому что интерфейс хранит значение - ссылку на тип объекта внутри

```
