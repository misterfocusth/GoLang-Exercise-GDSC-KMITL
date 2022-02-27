package exercise

import (
	"fmt"
	"sync"
)

var sharedMap = make(map[string]int)
var mutex sync.Mutex
var wg sync.WaitGroup

func readMap() {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		_ = sharedMap["SHARED"]
		mutex.Unlock()
	}
	wg.Done()
}

func writeMap() {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		sharedMap["SHARED"] += 1
		mutex.Unlock()
	}
	wg.Done()
	fmt.Println(sharedMap["SHARED"])
}

func Exercise6() {
	wg.Add(2)
	go readMap()
	go writeMap()
	wg.Wait()
}
