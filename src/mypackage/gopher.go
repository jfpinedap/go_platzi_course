package mypackage

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Gopher class example
type Gopher struct {
	Name     string
	AgeYears int
	idCardNumber int
}

// PrintMSG: Print a strings params
func PrintMSG(msg ...string) {
	var _msg string
	if len(msg) > 0 {
		_msg = strings.Join(msg, " ")
	} else {
		_msg = "Default message"
	}
	fmt.Println(_msg)
}

// OtherFunc: A stupid function that print a stupid message
func OtherFunc() {
	fmt.Println("A stupid function...")
}

func (g *Gopher) SetIdCradNumber(idCardNumber int) {
	g.idCardNumber = idCardNumber
}

func (g *Gopher) GetIdCradNumber() (idCardNumber int) {
	return g.idCardNumber
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err == nil {
		size += 4
		var n int
		n, err = w.Write([]byte(g.Name))
		size += int64(n)
		if err == nil {
			err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
			if err == nil {
				size += 4
			}
			return
		}
		return
	}
	return
}
