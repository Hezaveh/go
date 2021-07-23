package main

import ("fmt"
		"time"
		"sync")

/*func foo(){
	defer fmt.Println("Done3!")
	defer fmt.Println("Done2!")
	fmt.Println("Done1!")



	for i:= 0; i <5; i++ {
		defer fmt.Println(i)
		
	}
}*/

var wg sync.WaitGroup

func cleanup(){
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
}

func say(s string){ 
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println(s)
		if i == 2 {
			panic("oh its 2!")
		}
	}
}


func main(){
	//foo()

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
}