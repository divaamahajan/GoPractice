package main

import "fmt"

// ------------------Pointers-->
type contactInfo struct {
	email string
	zip int
}


type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	p1 := person{
        firstName: "Divya",
        lastName: "Mahajan",
    }
	p1.contactInfo.email = "dmahajan@gmail.com"
	p1.contactInfo.zip = 12345
	// fmt.Println(p1)
	// fmt.Printf("%+v\n", p1)
	p1.print()
	p1.updateName("Diva", "Mmono")
	p1.print()
}

func (p person) print(){
	fmt.Println("----")
	fmt.Println("First Name:", p.firstName)
	fmt.Println("Last Name:",p.lastName)
	fmt.Println("Contact Info:")
	fmt.Println("\tEmail Id:",p.contactInfo.email)
	fmt.Println("\tZip Code:",p.contactInfo.zip)
	fmt.Println("----")
}

func (p *person) updateName(firstName string, lastName string){
	p.firstName = firstName
    p.lastName = lastName
}