package mypackage

import "fmt"

type MyClass struct {
	attr_1 int
	attr_2 bool
	attr_3 string
}

func ClasesInGolang() {
	mc := MyClass{}

	fmt.Printf("%v\n", mc)
	fmt.Println(mc)
}
