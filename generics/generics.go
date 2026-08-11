package main

import (
	"fmt"
)

type values interface{
	int | string
}

func print [T values] (values []T){

	for _, value := range values{
		fmt.Println(value)
	}

}
func main(){

	values:= []int{1,2,4};
	values2 := []string{"san", "roy"};
	print(values)
	print(values2)
}