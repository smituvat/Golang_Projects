package main

import "fmt"

const name = "smita"

var lastname = "vatgal"

func main() {
	fmt.Printf("My name: %s %s\n", name, lastname)
	fmt.Println("Enter your age:")

	// Read input from user
	var age int
	fmt.Scan(&age)
	fmt.Printf("Your age is %d\n", age)

	// Defining array
	// var hobby [3]string
	// hobby[0] = "read"

	var hobby = [3]string{"read", "write", "code"}
	fmt.Printf("Hobbies are: %s,%s,%s", hobby[0], hobby[1], hobby[2])

	//Defining slice
	// var favColors = make([]string, 0)

	// Struct
	/*
		type Address struct {
		    name, street, city, state string
		    Pincode int
		}

		or

		var a Address
		var a = Address{"Akshay", "PremNagar", "Dehradun", "Uttarakhand", 252636}

	*/

	//Maps

	/*
		        map[Key_Type]Value_Type{}
				var map_1 map[int]int
				map_2 := map[int]string{

					90: "Dog",
					91: "Cat",
					92: "Cow",
					93: "Bird",
					94: "Rabbit",
			}
	*/
}
