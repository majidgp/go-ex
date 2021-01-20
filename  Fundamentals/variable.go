package main

import "fmt"

var (
	namei string  = "majid"
	agei  int64   = 30
	idi   uint64  = 571
	fl    float64 = 12.8

)

func main()  {

	fmt.Printf("The value of namei is : %s\n", namei)
	fmt.Printf("The type of namei is : %T\n", namei)

	fmt.Printf("The value of agei is : %d\n", agei)
	fmt.Printf("The type of agei is : %T\n", agei)

	fmt.Printf("The value of idi is : %d\n", idi)
	fmt.Printf("The type of idi is : %T\n", idi)

	fmt.Printf("The value of fl is : %f\n", fl)
	fmt.Printf("The type of fl is : %T\n", fl)
}