package mypackage

import (
	"fmt"
	"sort"
)

func DataType() {
	x := 58
	y := 3
	var result int
	result = x / y
	fmt.Println(result)

	result = x - y
	fmt.Println(result)

	result = x + y
	fmt.Println(result)

	result = x % y
	fmt.Println(result)

	i := 5.2
	i++
	fmt.Println(i)

	bl := true
	fmt.Println(bl)

	var bl2 bool
	bl2 = false
	fmt.Println(bl2)

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	nums := []int{4, 2, 3, 1, 7, 20, -78}
	strings := []string{"4", "2", "3", "1", "7", "aR", "20", "-78", "ax", "b", "11", "10", "ar", "AA", "aa"}
	sort.Ints(nums)
	sort.Strings(strings)
	fmt.Println(nums)
	fmt.Println(strings)
	fmt.Println("A:", 'A', "  --  ", "a:", 'a')
	fmt.Println("1:", '1', "  --  ", "2:", '2')
	fmt.Println("3:", '3', "  --  ", "4:", '4')
	fmt.Println("5:", '5', "  --  ", "6:", '6')
	fmt.Println("7:", '7', "  --  ", "0:", '0')
}
