package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sg := make(chan os.Signal)

	signal.Notify(sg,
		syscall.Signal(0x01),
		syscall.Signal(0x02),
		syscall.Signal(0x03),
		syscall.Signal(0x04),
		syscall.Signal(0x05),
		syscall.Signal(0x06),
		syscall.Signal(0x07),
		syscall.Signal(0x08),
		syscall.Signal(0x09),
		syscall.Signal(0x0a),
		syscall.Signal(0x0b),
		syscall.Signal(0x0c),
		syscall.Signal(0x0e),
		syscall.Signal(0x0f),
		syscall.Signal(0x10),
		syscall.Signal(0x11),
		syscall.Signal(0x12),
		syscall.Signal(0x13),
	)

	go func(c chan os.Signal) {
		for {
			s := <-c
			fmt.Printf("read signal %s\n", s.String())
		}
	}(sg)

	for {
		time.Sleep(100000000 * time.Minute)
	}
}
