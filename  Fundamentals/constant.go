package main

import "fmt"

const (
	inte  int64      = 10
	floo  float64   = 17.5
	cmp   complex64 = 7
	fname string    = "majid"
	cbool bool      = true
)

func main()  {

	fmt.Printf("The value of inte is : %d\n", inte)

	fmt.Printf("The value of floo is : %f\n", floo)

	fmt.Printf("The value of cmp is : %g\n", cmp)

	fmt.Printf("The value of fname is : %s\n", fname)

	fmt.Printf("The value of cbool is : %t\n", cbool )
}