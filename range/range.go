package main

import "fmt"


func main(){

	nums := []int{1,2,4};

	for index, value := range nums{
		fmt.Println(index,value)
	}

	//ignore index
	for _, value := range nums{
		fmt.Println(value)
	}

	//ignore value
	for index := range nums{
		fmt.Println(index)
	}

	//maps

	m := map[string]string{
		"name":"santanu",
	}

	for key, value := range m{
		fmt.Println(key,value)
	}
}