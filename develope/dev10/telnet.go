package main

import (
	"bufio"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type TelnetClient struct {
	address    string
	connection net.Conn
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	done       chan struct{}
}

// NewTelnetClient returns a telnet client instance
func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) *TelnetClient {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})
	go func() {
		<-signals
		close(done)
	}()
	return &TelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		done:    done,
	}
}

// Connect — connects to the host.
func (t *TelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return err
	}
	t.connection = conn
	return nil
}

// Send — sends message to the host.
func (t *TelnetClient) Send() error {
	buffer, err := bufio.NewReader(t.in).ReadBytes('\n')
	switch {
	case endOfTransmissionCheck(err):
		tryToCloseChannel(t.done)
		return nil
	case err != nil:
		return err
	default:
	}
	_, err = t.connection.Write(buffer)
	return err
}

// Receive — receives messages from the host.
func (t *TelnetClient) Receive() error {
	buffer, err := bufio.NewReader(t.connection).ReadBytes('\n')
	switch {
	case endOfTransmissionCheck(err):
		tryToCloseChannel(t.done)
		return nil
	case err != nil:
		return err
	default:
	}
	_, err = t.out.Write(buffer)
	return err
}

// Close — closes the connection.
func (t *TelnetClient) Close() error {
	return t.connection.Close()
}

// Done — signals that we are done.
func (t *TelnetClient) Done() <-chan struct{} {
	return t.done
}
