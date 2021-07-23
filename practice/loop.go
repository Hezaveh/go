package main

import "fmt"

func main(){
	for i:=0; i<10; i++{
		fmt.Println(i)
	}

	x := 5
	for {
		fmt.Println("Do stuff", x)
		x+=3
		if x > 35{
			break
		}
	}

	for x:=5; x<35; x+=3{
		fmt.Println("Do stuff", x)
	}

	a:=3
	for x:=5; a<35; x+=3{
		fmt.Println("Do stuff", x)
		a+=4
	}

}