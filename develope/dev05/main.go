package main

import (
	"awesomeProject2/develope/dev05/pkg"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
Реализовать утилиту фильтрации (man grep)

Поддержать флаги:

-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	core := pkg.NewCore()
	flag.IntVar(&core.After, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&core.Before, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&core.Context, "C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&core.Count, "c", false, "'count' (количество строк)")
	flag.BoolVar(&core.IgnoreCase, "i", false, "'ignore-case' (игнорировать регистр)")
	flag.BoolVar(&core.Invert, "v", false, "'invert' (вместо совпадения, исключать)")
	flag.BoolVar(&core.Fixed, "F", false, "'fixed', точное совпадение со строкой")
	flag.BoolVar(&core.LineNum, "n", false, "'line num', печатать номер строки")
	flag.Parse()
	core.SyncOutLength()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("usage: [flags] [pattern or string] [file]")
	}
	// выделяем из запроса фразу, которую будем искать
	slicePhrase := args[:len(args)-1]
	core.Phrase = strings.Join(slicePhrase, " ")
	// читаем данные из файла
	file, err := ioutil.ReadFile(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}
	splitString := strings.Split(string(file), "\n")
	result := Grep(splitString, core)
	printRes(core, result)
}

// Grep функция поиска фразы или строки в файле с применением доп.условий
func Grep(text []string, c *pkg.Core) []*pkg.Grep {
	var result []*pkg.Grep
	var condition bool // условие сравнения
	for index, str := range text {
		// если применен -i, убираем регистр
		if c.IgnoreCase {
			str = strings.ToLower(str)
			c.PhraseToLower()
		}
		// меняем условие в зависимсоти от переданных флагов
		if c.Fixed {
			condition = c.Phrase == str // полное совпадение строки
		} else {
			condition = strings.Contains(str, c.Phrase) // совпадение подстроки
		}
		if c.Invert {
			condition = !condition
		}
		match := pkg.NewGrep()
		// если условие выполняется
		if condition {
			c.AddMatch()
			var upRange, downRange = 0, len(text) - 1
			if d := index - c.Before; d > upRange {
				upRange = d
			}
			if d := index + c.After; d < downRange {
				downRange = d
			}
			for i := upRange; i <= downRange; i++ {
				match.Result = append(match.Result, pkg.Node{
					Key:   i + 1,
					Value: text[i],
				})
			}
			result = append(result, match)
		}
	}
	return result
}
func printRes(c *pkg.Core, res []*pkg.Grep) {
	if c.Count {
		fmt.Printf("Колиество совпадений: %d\n", c.CountMatch)
	}
	for _, match := range res {
		match.Print(c.LineNum)
		fmt.Println("--------------------------------------------")
	}
}
