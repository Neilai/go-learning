package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	//对于指针接收者，则只能传入指针类型，否则会报未实现接口的错误。
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//对于值接收者，传如值或者指针都可以正常调用；
	measure(&r)
	measure(r)
	//对于指针接收者，则只能传入指针类型，否则会报未实现接口的错误。
	measure(&c)

	var g geometry
	g = &circle{radius: 10}
	if t, ok := g.(*circle); ok {
		fmt.Println("circle  type", *t)
	}
	// switch t := g.(type) {
	// case *circle:
	// 	fmt.Println("circle type", *t)
	// case rect:
	// 	fmt.Println("rect type", t)
	// }
}
