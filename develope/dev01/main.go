package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

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
