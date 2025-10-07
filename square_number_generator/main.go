/*
The program reads two arguments from the command line: K and N.
The parameters K and N are passed via command-line arguments.
The program launches two functions: a generator and a squaring function.
The parameters N and M are of type int.
The generator function starts a goroutine and returns channel 1. Inside the goroutine, numbers from K to N (inclusive) are generated and sent into channel 1.
The squaring function starts a goroutine and returns channel 2. Inside the goroutine, numbers are read from channel 1, squared, and the result is sent to channel 2.
The main program (main) reads numbers from channel 2 and prints them to the console.
The squaring function must accept channel 1 as a read-only channel, which is returned by the generator function.
Both the squaring and generator functions must run concurrently.
Squaring must occur sequentially. After reading a number from channel 1, it must immediately be squared and sent to the next channel, and only then should the next number be processed.
Hint:
Channels must be created inside the functions, and returned with read/write restrictions applied.
Channels must be closed once the function finishes its work.
The generator and squaring functions must operate concurrently.
*/

package main

import (
	"flag"
	"fmt"
)

func toGenerate(k, n int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for i := k; i <= n; i++ {
			res <- i
		}
	}()
	return res
}

func toSquare(in <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for v := range in {
			res <- v * v
		}
	}()
	return res
}

func main() {
	k := flag.Int("k", -1, "start number")
	n := flag.Int("n", -1, "end number")
	flag.Parse()

	in := toGenerate(*k, *n)
	out := toSquare(in)

	for v := range out {
		fmt.Println(v)
	}
}
