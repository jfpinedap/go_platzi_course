package mypackage

import (
	"fmt"
	"math"
)

type figure2D interface {
	area() float64
}

type Rectangle struct {
	Area   float64
	Base   float64
	Height float64
}

type Circle struct {
	Area   float64
	Radius float64
}

func (r Rectangle) String() string {
	return fmt.Sprintf(
		"Rectangle with Base=%.1f Height=%.1f Area=%.1f",
		r.Base,
		r.Height,
		r.Area,
	)
}

func (r Circle) String() string {
	return fmt.Sprintf(
		"Circle with Radius=%.1f Area=%.1f",
		r.Radius,
		r.Area,
	)
}

func (r *Rectangle) area() float64 {
	r.Area = r.Base * r.Height
	return r.Area
}

func (c *Circle) area() float64 {
	c.Area = math.Pi * math.Pow(c.Radius, 2)
	return c.Area
}

func CalculateArea(f figure2D) {
	fmt.Println("The area is:", f.area())
}

func InferfacesDemo() {
	rectangle := Rectangle{
		Base:   45,
		Height: 58,
	}
	circle := Circle{
		Radius: 1,
	}
	CalculateArea(&rectangle)
	CalculateArea(&circle)
	fmt.Println(rectangle)
	fmt.Println(circle)

	interfaceList := []interface{}{"sdf", 5, circle, rectangle}
	fmt.Println(interfaceList)
}
