package main

import (
	"log"
	"sync"
	"time"
)

func test_sync() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			// WaitGroupカウンターを1つ減らす
			defer wg.Done()
			do_test1(n)
		}(i)
	}
	// WaitGroupカウンターが0になるまでブロックを待機
	wg.Wait()
}

func do_test1(n int) {
	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)
}
