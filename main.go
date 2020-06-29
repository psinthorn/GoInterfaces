package main

import "fmt"

type person struct {
	FirstName string
}

func (p person) speak() {
	fmt.Println("if this person have speak method then this person is also human type, and name is", p.FirstName)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a secret agent and my name is", sa.FirstName)
}

type human interface {
	speak()
}

func main() {

	p1 := person{
		FirstName: "Sinthorn",
	}
	fmt.Printf("%T\n", p1)

	sa1 := secretAgent{
		person: person{
			FirstName: "Jame",
		},
		ltk: true,
	}

	var x, y human
	x = p1
	y = sa1
	x.speak()
	y.speak()

	fmt.Printf("%T\n", x)
	fmt.Println(p1)
}
