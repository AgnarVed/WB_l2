package main

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
func getRepeatSym(n int, sym rune) []rune {
	var res []rune
	for i := 0; i < n; i++ {
		res = append(res, sym)
	}
	return res
}
func Extract(v string) string {
	switch len(v) {
	case 0:
		return ""
	case 1:
		return v
	default:
		var result []rune
		letters := []rune(v)
		for len(letters) > 0 {
			if len(letters) == 1 {
				result = append(result, letters[0])
				break
			}
			a, b := letters[0], letters[1]
			if isNumber(a) && isNumber(b) {
				return "некорректная строка"
			}
			if isNumber(b) {
				result = append(result, getRepeatSym(int(b-'0'), a)...)
				letters = letters[2:]
			} else {
				result = append(result, a)
				letters = letters[1:]
			}
		}
		return string(result)
	}
}
