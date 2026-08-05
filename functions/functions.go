package main

import (
	"fmt"
	"errors"
)

func greet(){
	fmt.Println("hello!")
}

func add(a,b int)( int, int){
	return a+b, a-b;
}

func divied(a,b float64) (float64, error){
	if b == 0 {
		return 0 , errors.New("cannot dived by zero");
	}

	return a/b, nil;
}

func sum(nums ...int) int{

	total:=0

	for _, val := range nums{
		total+=val;
	}

	return total;
}

func main(){
	// greet()
	// sum, sub:= add(6,5)

	// fmt.Println(sum,sub)

	res, err := divied(10,2)

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	//Anonymous Function
	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	greet("Santanu")

	//variadic function
	nums:= []int{12,15}
	fmt.Println(sum(nums...))

}