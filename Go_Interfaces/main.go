package main

import (
	"fmt"
)

/*
Go language interfaces are different from other languages. In Go language, the interface is a custom type that is used to specify a set of one
or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface.
But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods
the interface requires. Or in other words, the interface is a collection of methods as well as it is a custom type.
*/
type prints interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s Price: %.2f \n", d.Name, d.Price) //.2 formats floa to 2 decimal places
}

func (d Book) printInfo() {
	fmt.Printf("Title: %s Price: %.2f \n", d.Title, d.Price) //.2 formats floa to 2 decimal places
}

func main() {
	HarryPotter := Book{
		Title: "Harry Potter 1",
		Price: 234,
	}

	Tea := Drink{
		Name:  "Tea",
		Price: 10.5,
	}

	HarryPotter.printInfo()
	Tea.printInfo()

	info := []prints{HarryPotter, Tea}

	info[0].printInfo()
	info[1].printInfo()
}
