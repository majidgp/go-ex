package main

import "fmt"

func main()  {

	number := 7

	switch number {
	case 1,2:
		fmt.Println("low number")
	case 3,4,5:
		fmt.Println("normal number")
	case 6,7,8,9:
		fmt.Println("high number")
	default:
		fmt.Println("very high number")

	}
}
