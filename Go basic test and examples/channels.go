//channels are used with go routines to send and receive values between them
package main

import "fmt"

func foo(c chan int, someValue int) {	//c chan int receives a channel
    c <- someValue * 5
}

func main() {
    fooVal := make(chan int)		//making channel, fooval is channel of int type all are fed into through this
    go foo(fooVal, 5)
    go foo(fooVal, 3)
    //these statement is blocking so we dont need to do th esynchronization
    v1 := <-fooval 					//receives the response from responses
    v2 := <-fooVal
    fmt.Println(v1, v2)
}