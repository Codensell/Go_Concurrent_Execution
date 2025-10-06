/*
The program reads two launch arguments: N and M.
Parameters N and M are passed as arguments when launching the program.
The program launches N goroutines, each of which sleeps (time.Sleep) for a random duration of up to M milliseconds.
The program waits for all goroutines to finish.
The program prints a list to the console consisting of pairs <goroutine number, sleep time>, sorted in descending order of sleep time.
The goroutine number is the iteration index of the loop in which the goroutine was launched.
The sleep time is the number of milliseconds the goroutine slept.
The use of channels is not allowed.

Hint: Use the sync package. Wait for all goroutines to complete before starting the output. Use the flag package for parsing arguments.
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type result struct {
	id      int
	sleepMS int
}

func main() {
	n := flag.Int("n", -1, "number of goroutines")
	m := flag.Int("m", -1, "max sleep in milliseconds")
	flag.Parse()

	if *n < 0 {
		fmt.Println("n must be >= 0")
		return
	}
	if *m < 0 {
		fmt.Println("m must be >= 0")
		return
	}

	results := make([]result, *n)

	var wg sync.WaitGroup

	wg.Add(*n)
	for i := 0; i < *n; i++ {
		go func(i int) {
			defer wg.Done()

			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))

			ms := r.Intn(*m + 1)		

			time.Sleep(time.Duration(ms) * time.Millisecond)

			results[i] = result{id: i, sleepMS: ms}
		}(i)
	}

	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		return results[i].sleepMS > results[j].sleepMS
	})

	for _, r := range results {
		fmt.Printf("<%d, %d>\n", r.id, r.sleepMS)
	}
}
