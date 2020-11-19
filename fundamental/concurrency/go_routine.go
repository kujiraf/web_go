package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// Race()
	// Mutex()
	Atomic()
}

/*
Atomic is sample func of go routine which data is consistent.

go run go_routine.goでも実行結果が100で安定する
*/
func Atomic() {
	fmt.Println("CPUs:", runtime.NumCPU())             // 1つのCPUで稼働している
	fmt.Println("Goroutines:", runtime.NumGoroutine()) // 1つのgo routineで稼働している

	var counter int64

	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs) // いくつのgo routineを待って関数を完了するかを設定する

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done() // go routineが完了したことを通知する
		}()
		// fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

/*
Mutex is sample func of go routine which data is consistent.

go run go_routine.goでも実行結果が100で安定する
*/
func Mutex() {
	fmt.Println("CPUs:", runtime.NumCPU())             // 1つのCPUで稼働している
	fmt.Println("Goroutines:", runtime.NumGoroutine()) // 1つのgo routineで稼働している

	counter := 0

	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs) // いくつのgo routineを待って関数を完了するかを設定する

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // ロックする
			// v := counter
			runtime.Gosched() // 他のgo routineの呼び出しを許可する。time.Sleepでも次の呼び出しをすることができるが、こっちの方が効率的
			// v++
			// counter = v
			counter++
			mu.Unlock() // ロック解除
			wg.Done()   // go routineが完了したことを通知する
		}()
		// fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

/*
Race is sample func of go routine which data is unconsistent.

go run go_routine.goで実行すると実行結果が安定しない

go run -race go_routine.goで実行すると、実行結果が100で安定する
*/
func Race() {
	fmt.Println("CPUs:", runtime.NumCPU())             // 1つのCPUで稼働している
	fmt.Println("Goroutines:", runtime.NumGoroutine()) // 1つのgo routineで稼働している

	counter := 0

	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs) // いくつのgo routineを待って関数を完了するかを設定する

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched() // 他のgo routineの呼び出しを許可する。time.Sleepでも次の呼び出しをすることができるが、こっちの方が効率的
			v++
			counter = v
			wg.Done() // go routineが完了したことを通知する
		}()
		// fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
