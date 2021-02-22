package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type concurrentSlice struct {
	sync.RWMutex
	items []int
}

func (cs *concurrentSlice) Append(item int) {
	cs.Lock()
	defer cs.Unlock()
	cs.items = append(cs.items, item)
}

var evenNum = &concurrentSlice{}
var oddNum = &concurrentSlice{}

func EvenOrOdd(n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			go func(x int) {
				evenNum.Append(x)
			}(i)
		} else {
			go func(x int) {
				oddNum.Append(x)
			}(i)
		}
	}
}

func main() {
	EvenOrOdd(100)
	time.Sleep(time.Millisecond)
	sort.Ints(evenNum.items)
	sort.Ints(oddNum.items)
	fmt.Println(evenNum.items, len(evenNum.items))
	fmt.Println(oddNum.items, len(oddNum.items))
}
