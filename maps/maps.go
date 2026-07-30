package main

import "fmt"

func main(){

	m:= make(map[string]string)

	m["age"]= "10"
	m["name"]= "Santanu"
	fmt.Println(m["age"])

	//update
	m["age"] = "21"
	fmt.Println(m["age"])

	//delete
	delete(m, "age")
	
	//check value exist
	value, exist := m["name"];

	fmt.Println(value);
	fmt.Println(exist)

	//loop through map
	for key, value := range m{
		fmt.Println(key, value)
	}

	structEx()
}
