package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

/*
Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

*/

const (
	timeFormat = "15:04:05"
	ntpServer  = "0.beevik-ntp.pool.ntp.org"
)

func main() {
	time, err := getTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(time)
}

func getTime() (string, error) {
	resp, err := ntp.Query(ntpServer)
	if err != nil {
		return "", err
	}
	if err := resp.Validate(); err != nil {
		return "", err
	}
	return fmt.Sprintf("Time in your time Zone is %v\nTime UTC is %v", resp.Time.Local().Format(timeFormat), resp.Time.UTC().Format(timeFormat)), nil
}
