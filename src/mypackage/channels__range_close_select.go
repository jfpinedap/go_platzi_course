package mypackage

import (
	"fmt"
	"strconv"
	"strings"
)

func ConcurrenceChannels() {
	for i := 0; i < 30; i++ {
		c := make(chan string, 3)
		fmt.Print(
			strings.Repeat(".-", int(i/3)) + strings.Repeat("*-", int(i/3)) + strconv.Itoa(i),
		)
		c <- "Hello"
		c <- "Mr."
		c <- "Frank"

		close(c)

		for e := range c {
			fmt.Println(e)
		}
	}

}
