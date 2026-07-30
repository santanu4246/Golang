package main

import "fmt"

type User struct {
	Name string
	age  int
}

func structEx() {

	user:= make(map[int]User)

	fmt.Println((user))

	user[1] = User{
		Name: "Santanu",
		age: 21,
	}

	user[2] = User{
		Name: "Roy",
		age: 22,
	}

	for id, u := range user {
		fmt.Println(id, u.Name, u.age)
	}
}
