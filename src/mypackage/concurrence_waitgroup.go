package mypackage

import (
	"fmt"
	"strconv"
	"sync"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	fmt.Println(msg)
	wg.Done()
}

func ConcurrenceWGDemo() {
	var wg sync.WaitGroup
	for i := 0; i < 300; i++ {
		go printMessage("Hello", &wg)
		go printMessage("Mr.", &wg)
		go printMessage("Frank", &wg)
		go func(msg string, wg *sync.WaitGroup) {
			fmt.Println(msg)
			wg.Done()
		}(strconv.Itoa(i), &wg)
		wg.Add(4)
	}
	wg.Wait()
}
