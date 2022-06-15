package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// alex.firstName = "Alexander"
	// fmt.Printf("%+v", alex)
	jim := person {
		firstName: "James",
		lastName: "Adams",
		contact: contactInfo{
			email: "jamesadams@gmail.com",
			zipCode: 12345,
		},
	}
	//pointers are needed in Go
	// jimPointer := &jim (var is not needed - Go can understand it from jim)
	jim.updateName("Jim")
	jim.printPerson()
	
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}
// * is asking for the memory address
func (pointerToPerson *person) updateName(newFirstName string) {
     (*pointerToPerson).firstName = newFirstName
}