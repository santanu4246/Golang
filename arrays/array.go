package main

import "fmt"

func main(){
	// numbers := [5]int{1,2,3,4,5}

	// fmt.Println(numbers[0])
	
	//2D array
	numbers := [2][2] int{{}}

	fmt.Println("Enter array elements:-")
	for i := 0; i < len(numbers); i++{
		for j:=0; j < len(numbers); j++{
			fmt.Scan(&numbers[i][j])
		}
		
	}
	fmt.Println(numbers)
}