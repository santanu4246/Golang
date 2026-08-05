package main

import "fmt"


func main(){

	x:=10;

	fmt.Println(x,&x);

	var p *int = &x;

	fmt.Println(*p);

}