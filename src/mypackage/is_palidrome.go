package mypackage

import (
	"fmt"
	"reflect"
	"strings"
)

type Runes []rune

func (str Runes) ReverseString() (rev_str Runes) {
	l := len(str)
	rev_str = make(Runes, l)
	for i := 0; i <= l/2; i++ {
		rev_str[i], rev_str[l-1-i] = str[l-1-i], str[i]
	}
	return rev_str
}

func (str Runes) String() string {
	return string(str)
}

func Is_palindrome(input_str string) {
	str_rune := Runes(input_str)
	rev_str := str_rune.ReverseString()
	if strings.ToLower(input_str) == strings.ToLower(rev_str.String()) {
		fmt.Println("Is palindrome")
	} else {
		fmt.Println("Is NOT palindrome")
	}
	fmt.Println(input_str, reflect.TypeOf(input_str))
	fmt.Println(rev_str, reflect.TypeOf(rev_str))
}
