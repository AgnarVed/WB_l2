package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

// Check whether we end transmission or not.
func endOfTransmissionCheck(err error) bool {
	return err == io.EOF
}
func tryToCloseChannel(c chan struct{}) {
	select {
	case <-c:
	default:
		close(c)
	}
}

// india.colorado.edu 13
func main() {
	timeout := flag.String("timeout", "10s", "timeout for a connection")
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
		return
	}
	hostPort := flag.Arg(0)
	timeoutDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		usage()
		panic(err.Error())
	}
	telnetClient := NewTelnetClient(hostPort, timeoutDuration, os.Stdin, os.Stdout)
	err = connectAndAttach(telnetClient)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	<-telnetClient.Done()
	telnetClient.Close()
}
func usage() {
	fmt.Printf("Usage: %s [--timeout 10s] hostname:port\n", os.Args[0])
}
func connectAndAttach(tc *TelnetClient) error {
	if err := tc.Connect(); err != nil {
		return err
	}
	// Receive messages in goroutine.
	go func() {
		for {
			select {
			case <-tc.Done():
				return
			default:
			}
			tc.Receive()
		}
	}()
	// Send messages in goroutine.
	go func() {
		for {
			select {
			case <-tc.Done():
				return
			default:
			}
			tc.Send()
		}
	}()
	return nil
}
