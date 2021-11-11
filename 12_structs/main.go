package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Value Receiver - Not changing anything!
func (p Person) greet() string {
	return "Hello, My name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Pointer Receiver - Not changing anything!
func (p *Person) growUp() {
	p.age++
}
func main(){
	person1 :=  Person{firstName: "Nikhil",lastName: "S",city: "Zurich",gender: "Smantha",age:12}
	fmt.Println(person1)
	fmt.Println(person1.age)
	person1.age++;
	fmt.Println(person1.age)
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.growUp()
	fmt.Println(person1.age)
}