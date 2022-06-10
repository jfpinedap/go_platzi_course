package mypackage

import (
	"fmt"
)


func Maps() {
	var names = make(map[int]string)
	names[1] = "Peter"
	names[2] = "Peter"
	names[1] = "Petera"
	names[3] = "Peter"
	v, ok := names[1]
	fmt.Println(v, ok)
	for k, v := range names{
		fmt.Println(k, v)
	}
}