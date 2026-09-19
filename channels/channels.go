package main

import (
	"math/rand"
	"fmt"
)

func processNum (numChan chan int){

	for num := range numChan{
		fmt.Println("number is", num)
	}
	
}
func main(){

	numChan := make(chan int)

	go processNum(numChan)

	for{
		numChan <- rand.Intn(100)
	}
	// numChan <- 5

	// messageChan:= make(chan string)

	// messageChan <- "ping"

	// msg := <-messageChan

	// fmt.Println(msg)
}