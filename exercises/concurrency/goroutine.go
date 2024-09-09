package concurrency

import (
	"fmt"
	"sync"
)

var Global int = 0
var mutex sync.Mutex

// Tạo n rountine, mỗi routine chạy 1000 lần, làm sao kết quả 1000*n (consistency)

func _mutex(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		// fmt.Println(index)
		Global++
		mutex.Unlock()
	}
}

func mutexSolution(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	// Mutex
	for i := 0; i < n; i++ {
		go _mutex(&wg, i)
	}
	wg.Wait()
	fmt.Println(Global)
}

func channel(ch chan int, wg *sync.WaitGroup, index int) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		Global = <-ch
		Global++
		ch <- Global
		fmt.Println(index)
	}
}

func channelSolution(n int) {
	var wg sync.WaitGroup
	var ch chan int = make(chan int, 1)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go channel(ch, &wg, i)
	}
	ch <- Global

	wg.Wait()
	fmt.Println(Global)

}

func Concurrency(n int) {
	// mutexSolution(n)
	channelSolution(n)
}
