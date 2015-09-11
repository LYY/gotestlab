package main

import (
	"fmt"
	"sync"
	"time"
)

var mt sync.Mutex
var rwmt sync.RWMutex
var wg sync.WaitGroup

func MutexExec(i int) {
	if i == 4 {
		rwmt.Lock()
	} else {
		rwmt.RLock()
	}
	fmt.Println("run:", i)
	time.Sleep(1000 * time.Millisecond)
	if i == 4 {
		rwmt.Unlock()
	} else {
		rwmt.RUnlock()
	}
	wg.Done()
}

func main() {
	for a := 0; a < 10; a++ {
		wg.Add(1)
		go MutexExec(a)
	}
	wg.Wait()
}
