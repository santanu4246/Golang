package main

import "fmt"

type User struct {
	Name string
	age  int
	Address
}

// nested struct
type Address struct {
	City    string
	Country string
}

func (u User) Greet(){

	fmt.Println("hello", u.Name)
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

	users := []User{
		{
			Name: "Santanu",
			age:  21,
		},

		{
			Name: "chudShubh",
			age:  26,
		},

		{
			Name: "roy",
			age:  22,
		},
	}

	fmt.Println(users)

	for _, user := range users{
		user.Greet()
	}



}
