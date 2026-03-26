package main

import (
	"AlgorithmPractice/algorithms"
	"fmt"
	"sync"
	"time"
)

func main() {
	elems := make([]int, 0, 100_000_000)
	for i := range cap(elems) - 1 {
		elems = append(elems, i)
	}

	wg := sync.WaitGroup{}

	mtx := sync.RWMutex{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		begin := time.Now()
		mtx.RLock()
		fmt.Println(algorithms.BinarySearch(elems, 900000))
		mtx.RUnlock()
		fmt.Println("BS: ", time.Since(begin))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		begin := time.Now()
		mtx.RLock()
		fmt.Println(algorithms.StupidSearch(elems, 900000))
		mtx.RUnlock()
		fmt.Println("SS: ", time.Since(begin))
	}()

	wg.Wait()
}
