package main

import "fmt"

type person struct {
	FirstName string
}

func main() {

	p1 := person{
		FirstName: "Sinthorn",
	}
	fmt.Printf("%T/n", p1)
	fmt.Println(p1)
}
