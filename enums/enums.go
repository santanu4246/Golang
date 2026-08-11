package main

import "fmt"

type orderStaus int

const (
	Recevied = iota
	Prepared
	Confirmed 
	Delivered
)

func CheckSts(staus orderStaus){

	fmt.Println("Staus", staus)
}

func main(){

	CheckSts(Delivered)

}