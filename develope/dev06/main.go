package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	core := NewCore()
	flag.IntVar(&core.Fields, "f", 0, "'fields' - выбрать поля (колонки)")
	flag.StringVar(&core.Delimiter, "d", "\t", "'delimiter' - использовать другой разделитель")
	flag.BoolVar(&core.Separated, "s", false, "'separated' - только строки с разделителем")
	flag.Parse()
	args := flag.Args()
	// байт в условии нет, а значит нужен  f>0, стандартный cut тоже не взлетит
	if core.Fields == 0 {
		log.Fatalln("you must use -f with some value > 0")
	}
	if len(args) == 0 {
		terminalReading(core)
	}
	fileName := args[len(args)-1]
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	splitString := strings.Split(string(file), "\n")
	for _, str := range splitString {
		if res, ok := Cut(str, core); ok {
			fmt.Println(res)
		}
	}
}
func terminalReading(c *Core) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		res, _ := Cut(text, c)
		fmt.Println(res)
	}
}
func Cut(str string, c *Core) (string, bool) {
	if c.Separated && !strings.Contains(str, c.Delimiter) {
		return "", false
	}
	splitStr := strings.Split(str, c.Delimiter)
	if c.Fields <= len(splitStr) {
		return splitStr[c.Fields-1], true
	}
	return "", false
}
