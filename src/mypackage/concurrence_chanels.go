package mypackage

import (
	"fmt"
	"strconv"
	"strings"
)

func print(msg string) {
	fmt.Println(msg)
}

func printCHMessage(msg string, c chan<- string) {
	c <- msg
}

func ConcurrenceCHDemo() {
	c := make(chan string, 1)
	for i := 0; i < 30; i++ {
		fmt.Print(
			strings.Repeat(".", int(i/3)) + strings.Repeat(" ", int(i/3)) + strconv.Itoa(i),
		)
		go printCHMessage("Hello", c)
		go printCHMessage("Mr.", c)
		go printCHMessage("Frank", c)
		fmt.Println(<-c)
	}
}
