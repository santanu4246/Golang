package main

import "fmt"
// import "time"
func main(){

	// var i int

	// fmt.Println("Enter the number");
	// fmt.Scan(&i)
	// switch i{
	// case 1:
	 
	// 	fmt.Println("one")
	
	// case 2:
	// 	fmt.Println("two")

	// case 3:
	// 	fmt.Println("three")
	
	// case 4:
	// 	fmt.Println("four")

	// default:
	// 	fmt.Println("Default")
	// }

	// switch time.Now().Weekday(){

	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	//  default:
	// 	fmt.Println("Weekday")
	// }

	// type switch

	whoami := func(i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("Integar")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Others", t)
		}
	}
	whoami(6)
}