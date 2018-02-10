//defer got it
//panic is to create out a exit status and recover is for handling if we can recover or something


package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup				//basically adds 1 everytime a new routine called

func say(s string) {

	defer wg.Done()					//defer evaluates the sentence and do the operation only when either that set finishes or panicks out

	for i:=0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}	
	//wg.Done()						//gives a signal that the worl is done
									//here one problem can be if at any stage before this an occur occurs then it will wait for over so defer used
}

func main() {
	wg.Add(1)						//adds 1 go routine to waitgroup
	go say("Hey")
	wg.Add(1)						//adds 1 go routine to waitgroup
	go say("there")
	wg.Wait()						//Waits for return of done for all the added routines
/*
	wg.Add(1)						//adds 1 go routine to waitgroup
	go say("Hiii")
	wg.Wait()						//here it the above wait run different concurrent than this one
	*/
}