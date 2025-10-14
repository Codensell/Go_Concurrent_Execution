/*
The program reads the parameter K from the command-line arguments.
The parameter K is passed through the arguments when launching the program.
K defines the ticker interval in seconds and must be of type uint.
The program prints to stdout the message Tick <i> since <time>, where <i> is the tick number and <time> is the time in seconds since the ticker started.
The program runs until the user sends a SIGTERM or SIGINT signal.
Upon receiving one of these signals, the program stops the ticker and prints the message Termination.
The ticker must operate asynchronously. It is forbidden to use functions from the time package such as time.After or time.Ticker.
You are allowed to use constants from the time package and the Sleep function.
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func parseK(args []string) (uint, error) {
	if len(args) < 2 {
		return 0, fmt.Errorf("missing K")
	}
	v, err := strconv.ParseUint(args[1], 10, strconv.IntSize)
	if err != nil || v == 0 {
		return 0, fmt.Errorf("K must be a positive (seconds)")
	}
	return uint(v), nil
}

func main() {
	k, err := parseK(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	interval := time.Duration(k) * time.Second
	start := time.Now()

	done := make(chan struct{})

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		i := 0
		for {
			deadline := time.Now().Add(interval)
			for {
				select {
				case <-done:
					return
				default:
					if !time.Now().Before(deadline) {
						goto TICK
					}
					time.Sleep(50 * time.Millisecond)
				}
			}
		TICK:
			i++
			sec := int(time.Since(start) / time.Second)
			fmt.Printf("Tick %d since %d\n", i, sec)
		}
	}()

	<-sigc
	close(done)
	fmt.Println("Termination")
}