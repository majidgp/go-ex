package main

import "fmt"

func main()  {

	v := 70

	if v <= 0 {
		fmt.Println("Value of v <=0")
	}else if v == 50  {
		fmt.Println("Value of v == 50")
	}else {
		fmt.Println("Value of v > 50")
	}
}
