package main

import ("fmt"
"sync")

var wg sync.WaitGroup


func foo(c chan int, someValue int) {
	defer wg.Done()
    c <- someValue * 5
}

func main() {
    fooVal := make(chan int, 10)			//(, 10) was the buffer size of the channel can be greater than 10 here but not less

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go foo(fooVal, i)

    }

    wg.Wait()				//if we dont add synchronization then it closes the channel before it should be like before in which it was exiting without executing
    close(fooVal)			//to close the channel

    for item := range fooVal {		//using range we just iterate directly
        fmt.Println(item)
    }
}