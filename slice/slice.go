package main

import "fmt"

func main(){

	var n int;

	
	fmt.Println("Enter the size:-");
	fmt.Scan(&n);

	numbers := make([]int,n)
	
	fmt.Println("Enter the elements:-");
	for i := 0; i < n; i++{
		fmt.Scan(&numbers[i]);
	}

	fmt.Println(numbers)
}