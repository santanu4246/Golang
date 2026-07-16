package main

import "fmt"

func main (){
	age:=18

	if age >= 18{
		fmt.Println("voter")
	}else{
		fmt.Println("not a voter")
	}


	if name:="Santanu"; name == "santanu"{
		fmt.Println("true")
	}else if name == "san"{
		fmt.Println("false")
	}else{
		fmt.Println("true")
	}
}