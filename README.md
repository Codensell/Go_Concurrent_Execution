## Project "Concurrent Execution"

This project explores concurrency in Go through a series of small, focused exercises.
Each task builds practical intuition about goroutines, channels, timing, synchronization, and generics.

#### Purpose

Understand Go’s concurrency model — how goroutines communicate, synchronize, and terminate safely.
Practice time-based control with timers, tickers, and deadlines.
Learn graceful shutdown and coordination with context and sync.WaitGroup.
Improve awareness of memory safety and performance under concurrent workloads.
Apply Go generics to design efficient, reusable data structures.

#### Task 1. Stopwatch for Asynchronous Tasks

Goal: Measure the execution time of multiple concurrent tasks.
Focus: goroutines, channels, synchronization, timing, and cancellation.
Expected Result:
- The program starts several asynchronous jobs.
- Once all finish (or timeout expires), total and per-task durations are printed.
- Demonstrates clean termination and proper use of WaitGroup or context.

#### Task 2. Square Number Generator

Goal: Build a concurrent number generator pipeline that computes squares.
Focus: producer–consumer patterns, channel communication, and graceful closure.
Expected Result:
- The generator emits numbers continuously or up to a limit.
- A consumer reads and prints the squares of those numbers.
- No goroutines remain hanging after completion.

#### Task 3. Ticker

Goal: Implement periodic actions using Go’s time.Ticker.
Focus: scheduling, timed loops, and resource cleanup.
Expected Result:
- The program performs a recurring action (e.g., logging a timestamp) every N seconds.
- It stops gracefully when a signal or timeout occurs.
- Demonstrates clean ticker stopping (ticker.Stop()).

#### Task 4. LRU Cache Using Generics

Goal: Implement a Least Recently Used (LRU) cache using Go generics.
Focus: synchronization primitives (sync.Mutex, sync.Map), generic types, concurrent access safety.
Expected Result:
- Cache automatically evicts the least recently used entries after reaching capacity.
- Concurrent read/write operations are handled safely.
- The structure is reusable for any key/value types via generics.