package main

import "fmt"

type person struct {
	firstName string
	lastName string
}
func main() {
	p1 := person{
        firstName: "Divya",
        lastName: "Mahajan",
    }
	fmt.Println(p1)
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	fmt.Printf("%+v\n", p1)

	
}
