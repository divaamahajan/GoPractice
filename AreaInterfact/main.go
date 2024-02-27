package main

import "fmt"

// ------------------Area

type shape interface {
	getArea() float64
}

type square struct{
	side float64
}

type triangle struct{
	base float64
	hieght float64
}
type rectangle struct{
	length float64
	width float64
}
func main () {
	r := rectangle{length: 10, width: 2.5,}
	s := square{side:4,}
	t := triangle{base: 20,hieght: 10,}

	printArea(r)
	printArea(s)
	printArea(t)


}

func (s square) getArea() float64{	return s.side*s.side}
func (r rectangle) getArea() float64{	return r.length*r.width}
func (t triangle) getArea() float64{	return 0.5*t.base*t.hieght}
func printArea(s shape){ fmt.Println(s.getArea())}