package mypackage

import "fmt"

type PC struct {
	Ram   int
	Disk  int
	Brand string
}

func (pc PC) String() string {
	return fmt.Sprintf(
		"Is a %s with with %d RAM and %d SSD",
		pc.Brand,
		pc.Ram,
		pc.Disk,
	)
}

func (pc *PC) DuplicateRam() {
	pc.Ram *= 2
}

func (pc *PC) DuplicateDisk() {
	pc.Disk *= 2
}

func StructsPointers() {
	a := 50
	b := &a
	fmt.Println(a, b, *b)
	*b = 5
	fmt.Println(a, b, *b)
}
