package main

import (
	"fmt"
	"sync"
	
)

func task(wg *sync.WaitGroup){

	for i:=0; i<=10; i++{
		wg.Add(1)
		go func (i int){
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

}

func main(){
	var wg sync.WaitGroup;

	task(&wg)

	wg.Wait()
}