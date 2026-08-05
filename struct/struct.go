package main

import "fmt"

type User struct {
	Name    string
	age     int
	Address Address
}

// nested struct
type Address struct {
	City    string
	Country string
}

func main() {

	user := User{
		Name: "Santanu",
		age:  22,
		Address: Address{
			City: "bankura",
			Country: "India",
		},
	}

	fmt.Println(user)

}
