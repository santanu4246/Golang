package main

import (
	"fmt"
	"time"
)

func task(){

	for i:=0; i<=10; i++{

		go func (i int){
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}

func main(){
	task()
}